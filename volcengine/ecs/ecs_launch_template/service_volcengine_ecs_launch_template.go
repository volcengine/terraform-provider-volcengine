package ecs_launch_template

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEcsLaunchTemplateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func (s *VolcengineEcsLaunchTemplateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEcsLaunchTemplateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp         *map[string]interface{}
		versionsData map[string]interface{}
		results      interface{}
		ok           bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		client := s.Client.UniversalClient
		action := "DescribeLaunchTemplates"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = client.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = client.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.ReqFormat, action, condition, *resp)
		results, err = ve.ObtainSdkValue("Result.LaunchTemplates", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.LaunchTemplates is not Slice")
		}
		logger.Debug(logger.ReqFormat, action, condition, data)
		for index, ele := range data {
			action = "DescribeLaunchTemplateVersions"
			template := ele.(map[string]interface{})
			query := map[string]interface{}{
				"launch_template_id":         template["LaunchTemplateId"],
				"launch_template_versions.1": template["LatestVersionNumber"],
			}
			resp, err = client.DoCall(getUniversalInfo(action), &query)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.ReqFormat, action, query, *resp)
			versions, callErr := ve.ObtainSdkValue("Result.LaunchTemplateVersions", *resp)
			if callErr != nil {
				return data, callErr
			}
			_, ok = versions.([]interface{})
			if !ok {
				return data, errors.New("Result.LaunchTemplateVersions is not Slice")
			}
			for _, version := range versions.([]interface{}) {
				if _, ok = version.(map[string]interface{}); !ok {
					return data, errors.New("Result.LaunchTemplateVersion is not Map")
				}
				data[index].(map[string]interface{})["VersionDescription"] = version.(map[string]interface{})["VersionDescription"]
				if versionsData, ok = version.(map[string]interface{})["LaunchTemplateVersionData"].(map[string]interface{}); !ok {
					return data, errors.New("Result.LaunchTemplateVersions.LaunchTemplateVersionData is not Map")
				}
				for k, v := range versionsData {
					data[index].(map[string]interface{})[k] = v
				}
			}
		}
		return data, err
	})
}

func (s *VolcengineEcsLaunchTemplateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"launch_template_ids.1": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, fmt.Errorf("Value is not map ")
		}
	}
	logger.Debug(logger.ReqFormat, "ReadResource", data)
	if len(data) == 0 {
		return data, fmt.Errorf("LaunchTemplate %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineEcsLaunchTemplateService) RefreshResourceState(data *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineEcsLaunchTemplateService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"Eip.Bandwidth": {
				TargetField: "eip_bandwidth",
			},
			"Eip.ISP": {
				TargetField: "eip_isp",
			},
			"Eip.BillingType": {
				TargetField: "eip_billing_type",
				Convert:     billingTypeResponseConvert,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEcsLaunchTemplateService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateLaunchTemplate",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"volumes": {
					ConvertType: ve.ConvertListN,
				},
				"eip_bandwidth": {
					TargetField: "Eip.Bandwidth",
				},
				"eip_isp": {
					TargetField: "Eip.ISP",
				},
				"eip_billing_type": {
					TargetField: "Eip.BillingType",
					Convert:     billingTypeRequestConvert,
				},
				"network_interfaces": {
					ConvertType: ve.ConvertListN,
					NextLevelConvert: map[string]ve.RequestConvert{
						"security_group_ids": {
							ConvertType: ve.ConvertWithN,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.LaunchTemplateId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsLaunchTemplateService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 查询默认版本，避免错删默认版本
	launchTemplate, err := s.ReadResource(data, data.Id())
	if err != nil {
		return nil
	}
	defaultVersion := launchTemplate["DefaultVersionNumber"].(float64)
	req := map[string]interface{}{
		"launch_template_id": data.Id(),
	}
	versions, err := s.getLaunchTemplateVersions(req)
	if err != nil {
		return nil
	}
	if len(versions) > 29 {
		var oldestVersion float64
		for _, version := range versions {
			// 删除非默认版本的最老版本
			if versionMap, ok := version.(map[string]interface{}); !ok {
				return nil
			} else if versionMap["VersionNumber"].(float64) != defaultVersion &&
				(oldestVersion == 0 || versionMap["VersionNumber"].(float64) < oldestVersion) {
				oldestVersion = versionMap["VersionNumber"].(float64)
			}
		}
		_, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DeleteLaunchTemplateVersion"),
			&map[string]interface{}{
				"LaunchTemplateId": data.Id(),
				"DeleteVersions.1": oldestVersion},
		)
		if err != nil {
			return nil
		}
	}
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateLaunchTemplateVersion",
			ConvertMode: ve.RequestConvertInConvert,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if value, ok := d.GetOk("launch_template_name"); ok {
					(*call.SdkParam)["LaunchTemplateName"] = value
				}
				if value, ok := d.GetOk("instance_type_id"); ok {
					(*call.SdkParam)["InstanceTypeId"] = value
				}
				if value, ok := d.GetOk("version_description"); ok {
					(*call.SdkParam)["VersionDescription"] = value
				}
				if value, ok := d.GetOk("image_id"); ok {
					(*call.SdkParam)["ImageId"] = value
				}
				if value, ok := d.GetOk("instance_name"); ok {
					(*call.SdkParam)["InstanceName"] = value
				}
				if value, ok := d.GetOk("description"); ok {
					(*call.SdkParam)["Description"] = value
				}
				if value, ok := d.GetOk("host_name"); ok {
					(*call.SdkParam)["HostName"] = value
				}
				if value, ok := d.GetOk("hpc_cluster_id"); ok {
					(*call.SdkParam)["HpcClusterId"] = value
				}
				if value, ok := d.GetOk("instance_charge_type"); ok {
					(*call.SdkParam)["InstanceChargeType"] = value
				}
				if value, ok := d.GetOk("eip_bandwidth"); ok {
					(*call.SdkParam)["Eip.Bandwidth"] = value
				}
				if value, ok := d.GetOk("eip_isp"); ok {
					(*call.SdkParam)["Eip.ISP"] = value
				}
				if value, ok := d.GetOk("eip_billing_type"); ok {
					(*call.SdkParam)["Eip.BillingType"] = billingTypeRequestConvert(data, value)
				}
				if value, ok := d.GetOk("user_data"); ok {
					(*call.SdkParam)["UserData"] = value
				}
				if value, ok := d.GetOk("vpc_id"); ok {
					(*call.SdkParam)["VpcId"] = value
				}
				if value, ok := d.GetOk("key_pair_name"); ok {
					(*call.SdkParam)["KeyPairName"] = value
				}
				if value, ok := d.GetOk("security_enhancement_strategy"); ok {
					(*call.SdkParam)["SecurityEnhancementStrategy"] = value
				}
				if value, ok := d.GetOk("unique_suffix"); ok {
					(*call.SdkParam)["UniqueSuffix"] = value
				}
				if value, ok := d.GetOk("suffix_index"); ok {
					(*call.SdkParam)["SuffixIndex"] = value
				}
				if value, ok := d.GetOk("zone_id"); ok {
					(*call.SdkParam)["ZoneId"] = value
				}
				if value, ok := d.GetOk("volumes"); ok {
					for index, v := range value.([]interface{}) {
						if vMap, ok := v.(map[string]interface{}); ok {
							for k, v := range vMap {
								(*call.SdkParam)["Volumes."+strconv.Itoa(index+1)+"."+ve.DownLineToHump(k)] = v
							}
						}
					}
				}
				if value, ok := d.GetOk("network_interfaces"); ok {
					if len(value.([]interface{})) != 0 {
						for index, v := range value.([]interface{}) {
							if vMap, ok := v.(map[string]interface{}); ok {
								for k, v := range vMap {
									if k == "security_group_ids" {
										if len(v.([]interface{})) == 0 {
											continue
										}
										for sgIndex, sgValue := range v.([]interface{}) {
											(*call.SdkParam)["NetworkInterfaces."+strconv.Itoa(index+1)+"."+"SecurityGroupIds."+strconv.Itoa(sgIndex+1)] = sgValue
										}
									} else {
										(*call.SdkParam)["NetworkInterfaces."+strconv.Itoa(index+1)+"."+ve.DownLineToHump(k)] = v
									}
								}
							}
						}
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				delete(*call.SdkParam, "NetworkInterfaces")
				delete(*call.SdkParam, "Volumes")
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsLaunchTemplateService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteLaunchTemplate",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"LaunchTemplateId": data.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading ScalingConfiguration on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.NonRetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsLaunchTemplateService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "LaunchTemplateIds",
				ConvertType: ve.ConvertWithN,
			},
			"launch_template_names": {
				TargetField: "LaunchTemplateNames",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "LaunchTemplateName",
		IdField:      "LaunchTemplateId",
		CollectField: "launch_templates",
		ResponseConverts: map[string]ve.ResponseConvert{
			"LaunchTemplateId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"Eip.Bandwidth": {
				TargetField: "eip_bandwidth",
			},
			"Eip.ISP": {
				TargetField: "eip_isp",
			},
			"Eip.BillingType": {
				TargetField: "eip_billing_type",
				Convert:     billingTypeResponseConvert,
			},
		},
	}
}

func (s *VolcengineEcsLaunchTemplateService) ReadResourceId(id string) string {
	return id
}

func NewEcsLaunchTemplateService(client *ve.SdkClient) *VolcengineEcsLaunchTemplateService {
	return &VolcengineEcsLaunchTemplateService{
		Client:     client,
		Dispatcher: &ve.Dispatcher{},
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

var billingTypeResponseConvert = func(i interface{}) interface{} {
	var ty string
	switch i.(float64) {
	case 1:
		ty = "PrePaid"
	case 2:
		ty = "PostPaidByBandwidth"
	case 3:
		ty = "PostPaidByTraffic"
	default:
		return nil
	}
	return ty
}

var billingTypeRequestConvert = func(data *schema.ResourceData, old interface{}) interface{} {
	ty := 0
	switch old.(string) {
	case "PrePaid":
		ty = 1
	case "PostPaidByBandwidth":
		ty = 2
	case "PostPaidByTraffic":
		ty = 3
	}
	return ty
}

func (s *VolcengineEcsLaunchTemplateService) getLaunchTemplateVersions(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 30, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		client := s.Client.UniversalClient
		action := "DescribeLaunchTemplateVersions"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = client.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = client.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.ReqFormat, action, condition, *resp)
		results, err = ve.ObtainSdkValue("Result.LaunchTemplateVersions", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.LaunchTemplateVersions is not Slice")
		}
		return data, err
	})
}

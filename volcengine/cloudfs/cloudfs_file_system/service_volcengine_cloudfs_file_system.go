package cloudfs_file_system

import (
	"errors"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCloudfsFileSystemService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineCloudfsFileSystemService {
	return &VolcengineCloudfsFileSystemService{
		Client: c,
	}
}

func (s *VolcengineCloudfsFileSystemService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudfsFileSystemService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListFs"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineCloudfsFileSystemService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"FsName": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("File System %s not exist ", id)
	}

	// 针对数据湖的字段单独进行处理，无法读取出来，只能使用用户原先的参数进行替换
	if data["Mode"] == "ACC_MODE" {
		if tosBucket := resourceData.Get("tos_bucket"); len(tosBucket.(string)) > 0 {
			data["TosBucket"] = tosBucket
		}
		if tosPrefix := resourceData.Get("tos_prefix"); len(tosPrefix.(string)) > 0 {
			data["TosPrefix"] = tosPrefix
		}
		if tosAk := resourceData.Get("tos_ak"); len(tosAk.(string)) > 0 {
			data["TosAk"] = tosAk
		}
		if tosSk := resourceData.Get("tos_sk"); len(tosSk.(string)) > 0 {
			data["TosSk"] = tosSk
		}
		if tosAccountId := resourceData.Get("tos_account_id"); len(tosAccountId.(string)) > 0 {
			data["TosAccountId"] = tosAccountId
		}
	}

	// get cache access id
	defaultAccess, err := s.getDefaultAccess(id)
	if err != nil {
		return nil, err
	}
	if defaultAccess != nil {
		data["AccessId"] = defaultAccess["AccessId"]
		data["VpcRouteEnabled"] = defaultAccess["VpcRouteEnabled"]
	}
	return data, err
}

func (s *VolcengineCloudfsFileSystemService) getDefaultAccess(fsName interface{}) (map[string]interface{}, error) {
	result, err := s.Client.UniversalClient.DoCall(getUniversalInfo("ListAccess"), &map[string]interface{}{
		"FsName": fsName,
	})
	if err != nil {
		return nil, err
	}
	logger.Debug(logger.ReqFormat, "ListAccess", *result)
	accesses := (*result)["Result"].(map[string]interface{})["Items"].([]interface{})
	for _, access := range accesses {
		accessMap := access.(map[string]interface{})
		if v, ok := accessMap["IsDefault"]; ok {
			if v.(bool) {
				return accessMap, nil
			}
		}
	}
	return nil, nil
}

func (s *VolcengineCloudfsFileSystemService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "ERROR", "DOWN", "CREATE_FAILED", "ENABLE_CACHE_FAILED",
				"SCALE_UP_CACHE_FAILED", "UPDATE_FAILED", "DELETE_FAILED")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("File System  status  error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return d, status.(string), err
		},
	}

}

func (VolcengineCloudfsFileSystemService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return data, map[string]ve.ResponseConvert{
			"CacheCapacityTiB": {
				TargetField: "cache_capacity_tib",
			},
			"Name": {
				TargetField: "fs_name",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudfsFileSystemService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateFs",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cache_capacity_tib": {
					TargetField: "CacheCapacityTiB",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if v, ok := (*call.SdkParam)["SubnetId"]; ok {
					subnetAttr, err := client.VpcClient.DescribeSubnetAttributes(&vpc.DescribeSubnetAttributesInput{
						SubnetId: volcengine.String(v.(string))})
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["VpcId"] = *subnetAttr.VpcId
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(d.Get("fs_name").(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"UP"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineCloudfsFileSystemService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateFs",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"fs_name": {
					ForceGet: true,
				},
				"cache_plan": {
					ForceGet: true,
				},
				"cache_capacity_tib": {
					TargetField: "CacheCapacityTiB",
				},
				"subnet_id": {
					TargetField: "SubnetId",
				},
				"security_group_id": {
					TargetField: "SecurityGroupId",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if resourceData.HasChange("cache_plan") {
					o, n := resourceData.GetChange("cache_plan")
					if strings.Contains(n.(string), "T") && strings.Contains(o.(string), "T") {
						return false, fmt.Errorf("cache plan should remain the same")
					}

					if n.(string) == "DISABLED" {
						return false, fmt.Errorf("cannot disable cache")
					}

					if o.(string) == "DISABLED" {
						if _, ok := (*call.SdkParam)["SubnetId"]; !ok {
							return false, fmt.Errorf("need subnet id")
						}
						if _, ok := (*call.SdkParam)["SecurityGroupId"]; !ok {
							return false, fmt.Errorf("need security group id")
						}
						if _, ok := (*call.SdkParam)["CacheCapacityTiB"]; !ok {
							return false, fmt.Errorf("need cache capacity tib")
						}
					}
				}

				if v, ok := (*call.SdkParam)["SubnetId"]; ok {
					subnetAttr, err := client.VpcClient.DescribeSubnetAttributes(&vpc.DescribeSubnetAttributesInput{
						SubnetId: volcengine.String(v.(string))})
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["VpcId"] = *subnetAttr.VpcId
				}

				if len(*call.SdkParam) <= 2 {
					return false, nil
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"UP"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	if resourceData.Get("cache_plan") != "DISABLED" && resourceData.HasChange("vpc_route_enabled") {
		_, n := resourceData.GetChange("vpc_route_enabled")
		action := "DisableVpcRoute"
		if n.(bool) {
			action = "EnableVpcRoute"
		}
		enableCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      action,
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					defaultAccess, err := s.getDefaultAccess(d.Get("fs_name"))
					if err != nil {
						return false, err
					}
					if defaultAccess == nil {
						return false, fmt.Errorf("cannot find default access: %s", d.Id())
					}
					(*call.SdkParam)["AccessId"] = defaultAccess["AccessId"]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"UP"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, enableCallback)
	}
	return callbacks
}

func (s *VolcengineCloudfsFileSystemService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteFs",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"FsName": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading vpc on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudfsFileSystemService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "file_systems",
		ResponseConverts: map[string]ve.ResponseConvert{
			"CacheCapacityTiB": {
				TargetField: "cache_capacity_tib",
			},
		},
	}
}

func (s *VolcengineCloudfsFileSystemService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cfs",
		Version:     "2022-02-02",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cfs",
		Version:     "2022-02-02",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

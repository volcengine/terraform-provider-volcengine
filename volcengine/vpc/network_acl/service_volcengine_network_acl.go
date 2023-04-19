package network_acl

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"strconv"
	"time"
)

type VolcengineNetworkAclService struct {
	Client *ve.SdkClient
}

func NewNetworkAclService(c *ve.SdkClient) *VolcengineNetworkAclService {
	return &VolcengineNetworkAclService{
		Client: c,
	}
}

func (s *VolcengineNetworkAclService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNetworkAclService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeNetworkAcls"
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
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.NetworkAcls", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.NetworkAcls is not Slice")
		}

		return data, err
	})
}

func (s *VolcengineNetworkAclService) ReadResource(resourceData *schema.ResourceData, networkAclId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if networkAclId == "" {
		networkAclId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"NetworkAclIds.1": networkAclId,
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
		return data, fmt.Errorf("network acl %s is not exist ", networkAclId)
	}

	// 删除默认创建的拒绝规则
	if ingressAclEntries, ok := data["IngressAclEntries"]; ok {
		for index, entry := range ingressAclEntries.([]interface{}) {
			if priority, ok := entry.(map[string]interface{})["Priority"]; ok && priority.(float64) > 100 {
				ingressAclEntries = append(ingressAclEntries.([]interface{})[:index], ingressAclEntries.([]interface{})[index+1:]...)
				data["IngressAclEntries"] = ingressAclEntries
			}
		}
	}
	if egressAclEntries, ok := data["EgressAclEntries"]; ok {
		for index, entry := range egressAclEntries.([]interface{}) {
			if priority, ok := entry.(map[string]interface{})["Priority"]; ok && priority.(float64) > 100 {
				egressAclEntries = append(egressAclEntries.([]interface{})[:index], egressAclEntries.([]interface{})[index+1:]...)
				data["EgressAclEntries"] = egressAclEntries
			}
		}
	}

	return data, err
}

func (s *VolcengineNetworkAclService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status interface{}
			)
			//no failed status.
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			return demo, status.(string), err
		},
	}
}

func (VolcengineNetworkAclService) WithResourceResponseHandlers(acl map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return acl, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineNetworkAclService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNetworkAcl",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"ingress_acl_entries": {
					Ignore: true,
				},
				"egress_acl_entries": {
					Ignore: true,
				},
				"resources": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.NetworkAclId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 规则创建
	entryCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateNetworkAclEntries",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"ingress_acl_entries": {
					ConvertType: ve.ConvertListN,
				},
				"egress_acl_entries": {
					ConvertType: ve.ConvertListN,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["NetworkAclId"] = d.Id()
					(*call.SdkParam)["ClientToken"] = uuid.New().String()
					if _, ok := d.GetOk("ingress_acl_entries"); ok {
						(*call.SdkParam)["UpdateIngressAclEntries"] = true
					}
					if _, ok := d.GetOk("egress_acl_entries"); ok {
						(*call.SdkParam)["UpdateEgressAclEntries"] = true
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, entryCallback)

	return callbacks
}

func (s *VolcengineNetworkAclService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyNetworkAclAttributes",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"ingress_acl_entries": {
					Ignore: true,
				},
				"egress_acl_entries": {
					Ignore: true,
				},
				"resources": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["NetworkAclId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 规则修改
	if resourceData.HasChange("ingress_acl_entries") {
		ingressUpdateCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateNetworkAclEntries",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"ingress_acl_entries": {
						ConvertType: ve.ConvertListN,
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["NetworkAclId"] = d.Id()
						(*call.SdkParam)["ClientToken"] = uuid.New().String()
						(*call.SdkParam)["UpdateIngressAclEntries"] = true
						for index, entry := range d.Get("ingress_acl_entries").([]interface{}) {
							(*call.SdkParam)["IngressAclEntries."+strconv.Itoa(index+1)+".NetworkAclEntryId"] = entry.(map[string]interface{})["network_acl_entry_id"].(string)
						}
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Available"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, ingressUpdateCallback)
	}
	if resourceData.HasChange("egress_acl_entries") {
		ingressUpdateCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateNetworkAclEntries",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"egress_acl_entries": {
						ConvertType: ve.ConvertListN,
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["NetworkAclId"] = d.Id()
						(*call.SdkParam)["ClientToken"] = uuid.New().String()
						(*call.SdkParam)["UpdateEgressAclEntries"] = true
						for index, entry := range d.Get("egress_acl_entries").([]interface{}) {
							(*call.SdkParam)["EgressAclEntries."+strconv.Itoa(index+1)+".NetworkAclEntryId"] = entry.(map[string]interface{})["network_acl_entry_id"].(string)
						}
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Available"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, ingressUpdateCallback)
	}

	return callbacks
}

func (s *VolcengineNetworkAclService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNetworkAcl",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"NetworkAclId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading network acl on delete %q, %w", d.Id(), callErr))
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
	callbacks = append(callbacks, removeCallback)

	return callbacks
}

func (s *VolcengineNetworkAclService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "NetworkAclIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "NetworkAclName",
		IdField:      "NetworkAclId",
		CollectField: "network_acls",
		ResponseConverts: map[string]ve.ResponseConvert{
			"NetworkAclId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineNetworkAclService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func (s *VolcengineNetworkAclService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "vpc",
		ResourceType:         "networkacl",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

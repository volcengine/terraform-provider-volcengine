package network_acl_associate

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_acl"
	"strings"
	"time"
)

type VolcengineNetworkAclAssociateService struct {
	Client     *ve.SdkClient
}

func NewNetworkAclAssociateService(c *ve.SdkClient) *VolcengineNetworkAclAssociateService {
	return &VolcengineNetworkAclAssociateService{
		Client:     c,
	}
}

func (s *VolcengineNetworkAclAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNetworkAclAssociateService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		err     error
	)

	return ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeNetworkAcls"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return []interface{}{}, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return []interface{}{}, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.NetworkAcls", *resp)
		if err != nil {
			return []interface{}{}, err
		}
		if _, ok = results.([]interface{}); !ok {
			return []interface{}{}, errors.New("Result.NetworkAcls is not Slice")
		}
		return results.([]interface{}), err
	})
}

func (s *VolcengineNetworkAclAssociateService) ReadResource(resourceData *schema.ResourceData, associateId string) (data map[string]interface{}, err error) {
	if associateId == "" {
		associateId = resourceData.Id()
	}

	ids := strings.Split(associateId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid acl associateId: %s", associateId)
	}

	networkAclId := ids[0]
	resourceId := ids[1]
	req := map[string]interface{}{
		"NetworkAclIds.1": networkAclId,
	}

	networkAcls, err := s.ReadResources(req)
	if err != nil {
		return nil, err
	}
	if len(networkAcls) == 0 {
		return map[string]interface{}{}, fmt.Errorf("network acl %s not exist ", networkAclId)
	}
	for _, v := range networkAcls {
		if _, ok := v.(map[string]interface{}); !ok {
			return map[string]interface{}{}, errors.New("Value is not map ")
		}
	}

	aclResources := networkAcls[0].(map[string]interface{})["Resources"]
	if len(aclResources.([]interface{})) == 0 {
		return map[string]interface{}{}, fmt.Errorf("network acl resource %s:%s not exist ", networkAclId, resourceId)
	}
	for _, v := range aclResources.([]interface{}) {
		if _, ok := v.(map[string]interface{}); !ok {
			return map[string]interface{}{}, errors.New("Value is not map ")
		}
		if v.(map[string]interface{})["ResourceId"] == resourceId {
			data = v.(map[string]interface{})
		}
	}

	return data, err
}

func (s *VolcengineNetworkAclAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineNetworkAclAssociateService) WithResourceResponseHandlers(aclEntry map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return aclEntry, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNetworkAclAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateNetworkAcl",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["NetworkAclId"] = d.Get("network_acl_id")
				(*call.SdkParam)["Resource.1.ResourceId"] = d.Get("resource_id")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// ResourceData中，network_acl_associate的Id形式为'network_acl_id:resource_id'
				id := fmt.Sprintf("%s:%s", d.Get("network_acl_id"), d.Get("resource_id"))
				d.SetId(id)
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("network_acl_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				network_acl.NewNetworkAclService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("network_acl_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNetworkAclAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineNetworkAclAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateNetworkAcl",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				aclAssociateId := d.Id()
				ids := strings.Split(aclAssociateId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("error network acl associate id: %s", aclAssociateId)
				}
				(*call.SdkParam)["NetworkAclId"] = ids[0]
				(*call.SdkParam)["Resource.1.ResourceId"] = ids[1]
				return true, nil
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading acl entry on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("network_acl_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				network_acl.NewNetworkAclService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("network_acl_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNetworkAclAssociateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineNetworkAclAssociateService) ReadResourceId(id string) string {
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
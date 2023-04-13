package security_group

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint"
)

type VolcenginePrivateLinkSecurityGroupService struct {
	Client *ve.SdkClient
}

func (v *VolcenginePrivateLinkSecurityGroupService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcenginePrivateLinkSecurityGroupService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (v *VolcenginePrivateLinkSecurityGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = resourceData.Id()
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, errors.New("Invalid security group id ")
	}
	endpointId := ids[0]
	securityGroupId := ids[1]
	action := "DescribeVpcEndpointSecurityGroups"
	req := map[string]interface{}{
		"EndpointId": endpointId,
	}
	resp, err := v.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	if resp == nil {
		return data, fmt.Errorf("Security group %s not exists ", id)
	}
	securityGroupIds, err := ve.ObtainSdkValue("Result.SecurityGroupIds", *resp)
	if err != nil {
		return data, err
	}
	for _, s := range securityGroupIds.([]interface{}) {
		if _, ok := s.(string); !ok {
			return data, errors.New("security group id is not string")
		} else {
			if securityGroupId == s.(string) {
				data = map[string]interface{}{
					"SecurityGroupId": securityGroupId,
					"EndpointId":      endpointId,
				}
				break
			}
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Security group %s not exists ", id)
	}
	return data, nil
}

func (v *VolcenginePrivateLinkSecurityGroupService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return nil
}

func (v *VolcenginePrivateLinkSecurityGroupService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcenginePrivateLinkSecurityGroupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	endpointId := resourceData.Get("endpoint_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachSecurityGroupToVpcEndpoint",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := fmt.Sprintf("%s:%s", endpointId, d.Get("security_group_id"))
				d.SetId(id)
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vpc_endpoint.NewVpcEndpointService(v.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: endpointId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return endpointId
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcenginePrivateLinkSecurityGroupService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcenginePrivateLinkSecurityGroupService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachSecurityGroupFromVpcEndpoint",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"EndpointId":      resourceData.Get("endpoint_id"),
				"SecurityGroupId": resourceData.Get("security_group_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcenginePrivateLinkSecurityGroupService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (v *VolcenginePrivateLinkSecurityGroupService) ReadResourceId(id string) string {
	return id
}

func NewPrivateLinkSecurityGroupService(c *ve.SdkClient) *VolcenginePrivateLinkSecurityGroupService {
	return &VolcenginePrivateLinkSecurityGroupService{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "privatelink",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
		ContentType: ve.Default,
	}
}

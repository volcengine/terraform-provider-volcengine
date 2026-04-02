package nlb_server_group_attribute

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNlbServerGroupAttributeService struct {
	Client *ve.SdkClient
}

func NewNlbServerGroupAttributeService(c *ve.SdkClient) *VolcengineNlbServerGroupAttributeService {
	return &VolcengineNlbServerGroupAttributeService{
		Client: c,
	}
}

func (s *VolcengineNlbServerGroupAttributeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbServerGroupAttributeService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	action := "DescribeNLBServerGroupAttributes"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return nil, err
	}
	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, m, string(respBytes))

	if resp != nil {
		// Result is a single object, wrap it in slice
		if result, ok := (*resp)["Result"]; ok {
			data = append(data, result)
		}
	}
	return data, nil
}

func (s *VolcengineNlbServerGroupAttributeService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineNlbServerGroupAttributeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineNlbServerGroupAttributeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (VolcengineNlbServerGroupAttributeService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineNlbServerGroupAttributeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbServerGroupAttributeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbServerGroupAttributeService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"server_group_id": {
				TargetField: "ServerGroupId",
			},
		},
		IdField:      "ServerGroupId",
		CollectField: "server_group_attributes",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ServerGroupId": {
				TargetField: "server_group_id",
			},
			"ServerGroupName": {
				TargetField: "server_group_name",
			},
			"Description": {
				TargetField: "description",
			},
			"AccountId": {
				TargetField: "account_id",
			},
			"Protocol": {
				TargetField: "protocol",
			},
			"Type": {
				TargetField: "type",
			},
			"IpAddressVersion": {
				TargetField: "ip_address_version",
			},
			"Scheduler": {
				TargetField: "scheduler",
			},
			"VpcId": {
				TargetField: "vpc_id",
			},
			"ServerCount": {
				TargetField: "server_count",
			},
			"Status": {
				TargetField: "status",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"UpdateTime": {
				TargetField: "update_time",
			},
			"ProjectName": {
				TargetField: "project_name",
			},
			"BypassSecurityGroupEnabled": {
				TargetField: "bypass_security_group_enabled",
			},
			"ProxyProtocolType": {
				TargetField: "proxy_protocol_type",
			},
			"AnyPortEnabled": {
				TargetField: "any_port_enabled",
			},
			"ConnectionDrainEnabled": {
				TargetField: "connection_drain_enabled",
			},
			"ConnectionDrainTimeout": {
				TargetField: "connection_drain_timeout",
			},
			"PreserveClientIpEnabled": {
				TargetField: "preserve_client_ip_enabled",
			},
			"SessionPersistenceEnabled": {
				TargetField: "session_persistence_enabled",
			},
			"SessionPersistenceTimeout": {
				TargetField: "session_persistence_timeout",
			},
			"TimestampRemoveEnabled": {
				TargetField: "timestamp_remove_enabled",
			},
			"RelatedLoadBalancerIds": {
				TargetField: "related_load_balancer_ids",
			},
			"HealthCheck": {
				TargetField: "health_check",
			},
			"Servers": {
				TargetField: "servers",
			},
			"Tags": {
				TargetField: "tags",
				Convert:     ve.TransTagsToResponse,
			},
		},
	}
}

func (s *VolcengineNlbServerGroupAttributeService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

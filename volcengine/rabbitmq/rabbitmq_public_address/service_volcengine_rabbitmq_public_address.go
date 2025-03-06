package rabbitmq_public_address

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rabbitmq/rabbitmq_instance"
)

type VolcengineRabbitmqPublicAddressService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRabbitmqPublicAddressService(c *ve.SdkClient) *VolcengineRabbitmqPublicAddressService {
	return &VolcengineRabbitmqPublicAddressService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRabbitmqPublicAddressService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRabbitmqPublicAddressService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRabbitmqPublicAddressService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("Invalid rabbitmq public address id: %v ", id)
	}

	action := "DescribeInstanceDetail"
	req := map[string]interface{}{
		"InstanceId": ids[0],
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)
	basicInfo, err := ve.ObtainSdkValue("Result.BasicInstanceInfo", *resp)
	if err != nil {
		return data, err
	}
	data, ok = basicInfo.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf("DescribeInstanceDetail Result.BasicInstanceInfo is not mqp")
	}

	if eipId, exist := data["EipId"]; !exist || eipId == "" {
		return data, fmt.Errorf("rabbitmq_public_address %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRabbitmqPublicAddressService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineRabbitmqPublicAddressService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRabbitmqPublicAddressService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreatePublicAddress",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				eipId := d.Get("eip_id").(string)
				d.SetId(instanceId + ":" + eipId)
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rabbitmq_instance.NewRabbitmqInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRabbitmqPublicAddressService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRabbitmqPublicAddressService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeletePublicAddress",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("Invalid rabbitmq public address id: %v ", d.Id())
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rabbitmq_instance.NewRabbitmqInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRabbitmqPublicAddressService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRabbitmqPublicAddressService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "RabbitMQ",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

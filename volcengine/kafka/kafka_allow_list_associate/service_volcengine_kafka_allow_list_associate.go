package kafka_allow_list_associate

import (
	"errors"
	"fmt"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_instance"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKafkaAllowListAssociateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKafkaAllowListAssociateService(c *ve.SdkClient) *VolcengineKafkaAllowListAssociateService {
	return &VolcengineKafkaAllowListAssociateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKafkaAllowListAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKafkaAllowListAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		return data, err
	})
}

func (s *VolcengineKafkaAllowListAssociateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results     interface{}
		resultsMap  map[string]interface{}
		instanceMap map[string]interface{}
		instances   []interface{}
		ok          bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid kafka_allow_list_associate id: %v", id)
	}
	req := map[string]interface{}{
		"AllowListId": ids[1],
	}
	action := "DescribeAllowListDetail"
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if resultsMap, ok = results.(map[string]interface{}); !ok {
		return resultsMap, errors.New("Value is not map ")
	}
	if len(resultsMap) == 0 {
		return resultsMap, fmt.Errorf("Kafka allowlist %s not exist ", ids[1])
	}
	logger.Debug(logger.ReqFormat, action, resultsMap)
	instances, ok = resultsMap["AssociatedInstances"].([]interface{})
	if !ok {
		return data, errors.New("Value is not slice ")
	}
	logger.Debug(logger.ReqFormat, action, instances)
	for _, instance := range instances {
		if instanceMap, ok = instance.(map[string]interface{}); !ok {
			return data, errors.New("instance is not map ")
		}
		if len(instanceMap) == 0 {
			continue
		}
		if instanceMap["InstanceId"].(string) == ids[0] {
			data = resultsMap
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Kafka allowlist associate %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineKafkaAllowListAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
	}
}

func (s *VolcengineKafkaAllowListAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	allowListId := resourceData.Get("allow_list_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateAllowList",
			ConvertMode: ve.RequestConvertIgnore,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceIds"] = []string{instanceId}
				(*call.SdkParam)["AllowListIds"] = []string{allowListId}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%s", instanceId, allowListId))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				kafka_instance.NewKafkaInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKafkaAllowListAssociateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKafkaAllowListAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineKafkaAllowListAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	allowListId := resourceData.Get("allow_list_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateAllowList",
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceIds":  []string{instanceId},
				"AllowListIds": []string{allowListId},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				err := ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
				return err
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				kafka_instance.NewKafkaInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKafkaAllowListAssociateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineKafkaAllowListAssociateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Kafka",
		Version:     "2022-05-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

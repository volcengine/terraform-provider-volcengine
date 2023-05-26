package endpoint

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/instance"
)

type VolcengineRedisEndpointService struct {
	Client *ve.SdkClient
}

const (
	ActionCreateDBEndpointPublicAddress = "CreateDBEndpointPublicAddress"
	ActionDeleteDBEndpointPublicAddress = "DeleteDBEndpointPublicAddress"
	ActionDescribeDBInstanceDetail      = "DescribeDBInstanceDetail"
)

func NewRedisEndpointService(c *ve.SdkClient) *VolcengineRedisEndpointService {
	return &VolcengineRedisEndpointService{
		Client: c,
	}
}

func (s *VolcengineRedisEndpointService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRedisEndpointService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRedisEndpointService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		ids        []string
		instanceId string
		req        map[string]interface{}
		output     *map[string]interface{}
		results    interface{}
		ok         bool
	)
	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}
	ids = strings.Split(tmpId, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid redis endpoint id: %v", tmpId)
	}
	instanceId = ids[0]
	req = map[string]interface{}{
		"InstanceId": instanceId,
	}

	logger.Debug(logger.ReqFormat, ActionDescribeDBInstanceDetail, req)
	output, err = s.Client.UniversalClient.DoCall(getUniversalInfo(ActionDescribeDBInstanceDetail), &req)
	logger.Debug(logger.RespFormat, ActionDescribeDBInstanceDetail, req, *output)

	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue("Result", *output)
	if err != nil {
		return data, err
	}
	if data, ok = results.(map[string]interface{}); !ok {
		return data, errors.New("value is not map")
	}

	if _, exist := data["VisitAddrs"]; !exist {
		return nil, fmt.Errorf("not associated instance and eip. %s", tmpId)
	}

	attached := false
	for _, address := range data["VisitAddrs"].([]interface{}) {
		addr := address.(map[string]interface{})
		if addr["AddrType"].(string) == "Public" && addr["EipId"].(string) == ids[1] {
			attached = true
			break
		}
	}
	if !attached {
		return nil, fmt.Errorf("not associated instance and eip. %s", tmpId)
	}

	return map[string]interface{}{
		"InstanceId": ids[0],
		"EipId":      ids[1],
	}, nil
}

func (s *VolcengineRedisEndpointService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Delay:      time.Second,
		Pending:    []string{},
		Target:     target,
		Timeout:    timeout,
		MinTimeout: time.Second,

		Refresh: nil,
	}
}

func (s *VolcengineRedisEndpointService) WithResourceResponseHandlers(endpoint map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineRedisEndpointService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionCreateDBEndpointPublicAddress,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				output, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				return output, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["InstanceId"], ":", (*call.SdkParam)["EipId"]))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Attached"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("eip_id").(string),
				},
				instance.NewRedisDbInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRedisEndpointService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRedisEndpointService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionDeleteDBEndpointPublicAddress,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id := s.ReadResourceId(d.Id())
				ids := strings.Split(id, ":")
				instanceId := ids[0]
				eipId := ids[1]
				(*call.SdkParam)["InstanceId"] = instanceId
				(*call.SdkParam)["EipId"] = eipId
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("eip_id").(string),
				},
				instance.NewRedisDbInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRedisEndpointService) DatasourceResources(data *schema.ResourceData, resource2 *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType: ve.ContentTypeJson,
	}
}

func (s *VolcengineRedisEndpointService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Redis",
		Version:     "2020-12-07",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

package ssl_state

import (
	"errors"
	"fmt"
	mongodbInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineMongoDBSSLStateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewMongoDBSSLStateService(c *ve.SdkClient) *VolcengineMongoDBSSLStateService {
	return &VolcengineMongoDBSSLStateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineMongoDBSSLStateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineMongoDBSSLStateService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "ssl_state",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"SSLEnable": {
				TargetField: "ssl_enable",
			},
			"SSLExpiredTime": {
				TargetField: "ssl_expired_time",
			},
		},
	}
}

func (s *VolcengineMongoDBSSLStateService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	action := "DescribeDBInstanceSSL"
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

	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		logger.DebugInfo("ve.ObtainSdkValue return :%v", err)
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	return []interface{}{results}, nil
}

func (s *VolcengineMongoDBSSLStateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	resourceId := resourceData.Id()
	parts := strings.Split(resourceId, ":")
	instanceId := parts[1]

	req := map[string]interface{}{
		"InstanceId": instanceId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("SSLState %s is not exist", id)
	}
	return data, err
}

func (s *VolcengineMongoDBSSLStateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineMongoDBSSLStateService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongoDBSSLStateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceSSL",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Get("instance_id"),
				"SSLAction":  "Open",
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				d.SetId(fmt.Sprintf("ssl:%s", d.Get("instance_id")))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				mongodbInstance.NewMongoDBInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				logger.Debug("lock instance id:%s", instanceId, "")
				return instanceId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBSSLStateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceSSL",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Get("instance_id"),
				"SSLAction":  "Update",
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				d.SetId(fmt.Sprintf("ssl:%s", d.Get("instance_id")))
				d.Set("ssl_action", "noupdate")
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				mongodbInstance.NewMongoDBInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				logger.Debug("lock instance id:%s", instanceId, "")
				return instanceId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBSSLStateService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceSSL",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Get("instance_id"),
				"SSLAction":  "Close",
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				d.SetId(fmt.Sprintf("ssl:%s", d.Get("instance_id")))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				mongodbInstance.NewMongoDBInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				logger.Debug("lock instance id:%s", instanceId, "")
				return instanceId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBSSLStateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "mongodb",
		Action:      actionName,
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}

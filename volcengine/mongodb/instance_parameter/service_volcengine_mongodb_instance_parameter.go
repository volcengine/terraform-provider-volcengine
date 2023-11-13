package instance_parameter

import (
	"errors"
	"fmt"
	"strings"
	"time"

	mongodbInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineMongoDBInstanceParameterService struct {
	Client *ve.SdkClient
}

func NewMongoDBInstanceParameterService(c *ve.SdkClient) *VolcengineMongoDBInstanceParameterService {
	return &VolcengineMongoDBInstanceParameterService{
		Client: c,
	}
}

func (s *VolcengineMongoDBInstanceParameterService) GetClient() *ve.SdkClient {
	return s.Client
}

// ReadAll 保证data兼容性
func (s *VolcengineMongoDBInstanceParameterService) ReadAll(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	action := "DescribeDBInstanceParameters"

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
	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = map[string]interface{}{}
	}
	data = []interface{}{results}
	return data, err
}

func (s *VolcengineMongoDBInstanceParameterService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "DescribeDBInstanceParameters"

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
	results, err = ve.ObtainSdkValue("Result.InstanceParameters", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.InstanceParameters is not slice")
	}
	return data, err
}

func (s *VolcengineMongoDBInstanceParameterService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 3 {
		return data, fmt.Errorf("the format of import id must be 'param:instanceId:parameterName'")
	}
	req := map[string]interface{}{
		"InstanceId":     parts[1],
		"ParameterNames": parts[2],
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
		return data, fmt.Errorf("instance parameters %s is not exist", id)
	}
	if _, ok = data["ParameterNames"]; ok {
		data["ParameterName"] = data["ParameterNames"]
	}
	return data, nil
}

func (s *VolcengineMongoDBInstanceParameterService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineMongoDBInstanceParameterService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongoDBInstanceParameterService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceParameters",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"parameter_name": {
					Ignore: true,
				},
				"parameter_role": {
					Ignore: true,
				},
				"parameter_value": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ParametersObject"] = map[string]interface{}{
					"ParameterName":  d.Get("parameter_name"),
					"ParameterRole":  d.Get("parameter_role"),
					"ParameterValue": d.Get("parameter_value"),
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				parameterName := d.Get("parameter_name").(string)
				id := fmt.Sprintf("%v:%v:%v", "param", instanceId, parameterName)
				d.SetId(id)
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				mongodbInstance.NewMongoDBInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return resourceData.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBInstanceParameterService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	id := s.ReadResourceId(resourceData.Id())
	parts := strings.Split(id, ":")
	instanceId := parts[1]
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceParameters",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = instanceId
				(*call.SdkParam)["ParametersObject"] = map[string]interface{}{
					"ParameterName":  parts[2],
					"ParameterRole":  d.Get("parameter_role"),
					"ParameterValue": d.Get("parameter_value"),
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				mongodbInstance.NewMongoDBInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBInstanceParameterService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoDBInstanceParameterService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "instance_parameters",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"ParameterNames": {
				TargetField: "parameter_name",
			},
		},
	}
}

func (s *VolcengineMongoDBInstanceParameterService) ReadResourceId(id string) string {
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

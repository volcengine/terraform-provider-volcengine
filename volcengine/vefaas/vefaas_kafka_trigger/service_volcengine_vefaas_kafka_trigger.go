package vefaas_kafka_trigger

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVefaasKafkaTriggerService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVefaasKafkaTriggerService(c *ve.SdkClient) *VolcengineVefaasKafkaTriggerService {
	return &VolcengineVefaasKafkaTriggerService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVefaasKafkaTriggerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVefaasKafkaTriggerService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		item    map[string]interface{}
		result  []interface{}
		data    []interface{}
		err     error
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListTriggers"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return result, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return result, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return result, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return result, errors.New("Result.Items is not Slice")
		}

		for _, v := range data {
			if item, ok = v.(map[string]interface{}); !ok {
				return result, errors.New("Value is not map ")
			}

			if item["Type"] == "kafka" {
				result = append(result, item)
			}
		}

		return result, err
	})
}

func (s *VolcengineVefaasKafkaTriggerService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("format of vefaas kafka trigger is invalid,%s", id)
	}
	functionId := parts[0]
	kafkaTriggerId := parts[1]

	// TODO: add request info
	req := map[string]interface{}{
		"Id":         kafkaTriggerId,
		"FunctionId": functionId,
	}
	action := "GetKafkaTrigger"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	queryRelease, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	data, ok := queryRelease.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("kafka trigger %s not exist", id)
	}

	return data, err
}

func (s *VolcengineVefaasKafkaTriggerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d      map[string]interface{}
				status interface{}
			)
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)

			if err != nil {
				return nil, "", err
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVefaasKafkaTriggerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateKafkaTrigger",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"kafka_credentials": {
					ConvertType: ve.ConvertJsonObject,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				functionId, _ := ve.ObtainSdkValue("Result.FunctionId", *resp)
				kafkaTriggerId, _ := ve.ObtainSdkValue("Result.Id", *resp)

				d.SetId(fmt.Sprintf("%s:%s", functionId, kafkaTriggerId))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"ready", "failed"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVefaasKafkaTriggerService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVefaasKafkaTriggerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateKafkaTrigger",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("format of vefaas kafka trigger id is invalid,%s", d.Id())
				}
				functionId := parts[0]
				kafkaTriggerId := parts[1]
				(*call.SdkParam)["Id"] = kafkaTriggerId
				(*call.SdkParam)["FunctionId"] = functionId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVefaasKafkaTriggerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteKafkaTrigger",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("format of vefaas kafka trigger id is invalid,%s", d.Id())
				}
				functionId := parts[0]
				kafkaTriggerId := parts[1]
				(*call.SdkParam)["Id"] = kafkaTriggerId
				(*call.SdkParam)["FunctionId"] = functionId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading kafka trigger on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVefaasKafkaTriggerService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "items",
	}
}

func (s *VolcengineVefaasKafkaTriggerService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vefaas",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

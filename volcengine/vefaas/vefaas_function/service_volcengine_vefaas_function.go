package vefaas_function

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVefaasFunctionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVefaasFunctionService(c *ve.SdkClient) *VolcengineVefaasFunctionService {
	return &VolcengineVefaasFunctionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVefaasFunctionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVefaasFunctionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListFunctions"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVefaasFunctionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Id": id,
	}
	action := "GetFunction"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	queryFunction, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	data, ok := queryFunction.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("function %s not exist", id)
	}

	if resourceData.Get("tos_mount_config.0.credentials.0.access_key_id").(string) != "" && resourceData.Get("tos_mount_config.0.credentials.0.secret_access_key").(string) != "" {
		logger.Debug(logger.RespFormat, "tos_mount_config.0.credentials.0.access_key_id data is", resourceData.Get("tos_mount_config.0.credentials.0.access_key_id").(string))
		if tosMountConfig, tosMountConfigExist := data["TosMountConfig"]; tosMountConfigExist {
			logger.Debug(logger.RespFormat, "TosMountConfig data is", data["TosMountConfig"])
			tosMountConfigMap, ok := tosMountConfig.(map[string]interface{})
			if ok {
				tosMountConfigMap["Credentials"] = []interface{}{
					map[string]interface{}{
						"AccessKeyId":     resourceData.Get("tos_mount_config.0.credentials.0.access_key_id").(string),
						"SecretAccessKey": resourceData.Get("tos_mount_config.0.credentials.0.secret_access_key").(string),
					},
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineVefaasFunctionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineVefaasFunctionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateFunction",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"memory_mb": {
					TargetField: "MemoryMB",
				},
				"envs": {
					TargetField: "Envs",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"vpc_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"subnet_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"security_group_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
				"tls_config": {
					ConvertType: ve.ConvertJsonObject,
				},
				"source_access_config": {
					ConvertType: ve.ConvertJsonObject,
				},
				"nas_storage": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"nas_configs": {
							ConvertType: ve.ConvertJsonObjectArray,
						},
					},
				},
				"tos_mount_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"credentials": {
							ConvertType: ve.ConvertJsonObject,
						},
						"mount_points": {
							ConvertType: ve.ConvertJsonObjectArray,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Id", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVefaasFunctionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"MemoryMB": {
				TargetField: "memory_mb",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVefaasFunctionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateFunction",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"MemoryMB": {
					TargetField: "memory_mb",
				},
				"envs": {
					TargetField: "Envs",
					ConvertType: ve.ConvertJsonObjectArray,
					ForceGet:    true,
				},
				"vpc_config": {
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"subnet_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"security_group_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
				"tls_config": {
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
				},
				"source_access_config": {
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
				},
				"nas_storage": {
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"nas_configs": {
							ConvertType: ve.ConvertJsonObjectArray,
						},
					},
				},
				"tos_mount_config": {
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"credentials": {
							ConvertType: ve.ConvertJsonObject,
						},
						"mount_points": {
							ConvertType: ve.ConvertJsonObjectArray,
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
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

func (s *VolcengineVefaasFunctionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteFunction",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Id": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading function on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVefaasFunctionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		NameField:       "Name",
		IdField:         "Id",
		CollectField:    "items",
		ResponseConverts: map[string]ve.ResponseConvert{
			"MemoryMB": {
				TargetField: "memory_mb",
			},
		},
	}
}

func (s *VolcengineVefaasFunctionService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vefaas",
		Version:     "2024-06-06",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

//func (s *VolcengineVefaasFunctionService) ProjectTrn() *ve.ProjectTrn {
//	return &ve.ProjectTrn{
//		ServiceName:          "vefaas",
//		ResourceType:         "function",
//		ProjectResponseField: "ProjectName",
//		ProjectSchemaField:   "project_name",
//	}
//}

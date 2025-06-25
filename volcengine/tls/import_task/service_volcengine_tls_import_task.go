package import_task

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

type VolcengineImportTaskService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewImportTaskService(c *ve.SdkClient) *VolcengineImportTaskService {
	return &VolcengineImportTaskService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineImportTaskService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineImportTaskService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeImportTasks"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("RESPONSE.TaskInfo", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("RESPONSE.TaskInfo is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineImportTaskService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"TaskId": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("import_task %s not exist ", id)
	}
	if targetInfo, targetInfoExist := data["TargetInfo"]; targetInfoExist {
		if extractRule, extractRuleExist := targetInfo.(map[string]interface{})["ExtractRule"]; extractRuleExist {
			targetInfo.(map[string]interface{})["ExtractRule"] = []interface{}{
				extractRule,
			}
		}
	}
	if importSourceInfo, importSourceInfoExist := data["ImportSourceInfo"]; importSourceInfoExist {
		if tosSourceInfo, tosSourceInfoExist := importSourceInfo.(map[string]interface{})["TosSourceInfo"]; tosSourceInfoExist {
			importSourceInfo.(map[string]interface{})["TosSourceInfo"] = []interface{}{
				tosSourceInfo,
			}
		}
	}
	if importSourceInfo, importSourceInfoExist := data["ImportSourceInfo"]; importSourceInfoExist {
		if kafkaSourceInfo, kafkaSourceInfoExist := importSourceInfo.(map[string]interface{})["KafkaSourceInfo"]; kafkaSourceInfoExist {
			importSourceInfo.(map[string]interface{})["KafkaSourceInfo"] = []interface{}{
				kafkaSourceInfo,
			}
		}
	}
	return data, err
}

func (s *VolcengineImportTaskService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineImportTaskService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateImportTask",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"import_source_info": {
					TargetField: "ImportSourceInfo",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"tos_source_info": {
							TargetField: "TosSourceInfo",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"bucket": {
									TargetField: "bucket",
								},
								"prefix": {
									TargetField: "prefix",
								},
								"region": {
									TargetField: "region",
								},
								"compress_type": {
									TargetField: "compress_type",
								},
							},
						},
						"kafka_source_info": {
							TargetField: "KafkaSourceInfo",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"host": {
									TargetField: "host",
								},
								"group": {
									TargetField: "group",
								},
								"topic": {
									TargetField: "topic",
								},
								"encode": {
									TargetField: "encode",
								},
								"password": {
									TargetField: "password",
								},
								"protocol": {
									TargetField: "protocol",
								},
								"username": {
									TargetField: "username",
								},
								"mechanism": {
									TargetField: "mechanism",
								},
								"instance_id": {
									TargetField: "instance_id",
								},
								"is_need_auth": {
									TargetField: "is_need_auth",
								},
								"initial_offset": {
									TargetField: "initial_offset",
								},
								"time_source_default": {
									TargetField: "time_source_default",
								},
							},
						},
					},
				},
				"project_id": {
					TargetField: "ProjectID",
				},
				"target_info": {
					TargetField: "TargetInfo",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"extract_rule": {
							TargetField: "ExtractRule",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"keys": {
									TargetField: "Keys",
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
					},
				},
				"topic_id": {
					TargetField: "TopicID",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.TaskId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineImportTaskService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"status": {
				TargetField: "status",
			},
			"bucket": {
				TargetField: "bucket",
			},
			"prefix": {
				TargetField: "prefix",
			},
			"region": {
				TargetField: "region",
			},
			"compress_type": {
				TargetField: "compress_type",
			},
			"host": {
				TargetField: "host",
			},
			"group": {
				TargetField: "group",
			},
			"topic": {
				TargetField: "topic",
			},
			"encode": {
				TargetField: "encode",
			},
			"password": {
				TargetField: "password",
			},
			"protocol": {
				TargetField: "protocol",
			},
			"username": {
				TargetField: "username",
			},
			"mechanism": {
				TargetField: "mechanism",
			},
			"instance_id": {
				TargetField: "instance_id",
			},
			"is_need_auth": {
				TargetField: "is_need_auth",
			},
			"initial_offset": {
				TargetField: "initial_offset",
			},
			"time_source_default": {
				TargetField: "time_source_default",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineImportTaskService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyImportTask",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"description": {
					TargetField: "Description",
					ForceGet:    true,
				},
				"import_source_info": {
					TargetField: "ImportSourceInfo",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"tos_source_info": {
							TargetField: "TosSourceInfo",
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"bucket": {
									TargetField: "bucket",
									ForceGet:    true,
								},
								"prefix": {
									TargetField: "prefix",
									ForceGet:    true,
								},
								"region": {
									TargetField: "region",
									ForceGet:    true,
								},
								"compress_type": {
									TargetField: "compress_type",
									ForceGet:    true,
								},
							},
						},
						"kafka_source_info": {
							TargetField: "KafkaSourceInfo",
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"host": {
									TargetField: "host",
									ForceGet:    true,
								},
								"group": {
									TargetField: "group",
									ForceGet:    true,
								},
								"topic": {
									TargetField: "topic",
									ForceGet:    true,
								},
								"encode": {
									TargetField: "encode",
									ForceGet:    true,
								},
								"password": {
									TargetField: "password",
									ForceGet:    true,
								},
								"protocol": {
									TargetField: "protocol",
									ForceGet:    true,
								},
								"username": {
									TargetField: "username",
									ForceGet:    true,
								},
								"mechanism": {
									TargetField: "mechanism",
									ForceGet:    true,
								},
								"instance_id": {
									TargetField: "instance_id",
									ForceGet:    true,
								},
								"is_need_auth": {
									TargetField: "is_need_auth",
									ForceGet:    true,
								},
								"initial_offset": {
									TargetField: "initial_offset",
									ForceGet:    true,
								},
								"time_source_default": {
									TargetField: "time_source_default",
									ForceGet:    true,
								},
							},
						},
					},
				},
				"project_id": {
					TargetField: "ProjectID",
					ForceGet:    true,
				},
				"source_type": {
					TargetField: "SourceType",
					ForceGet:    true,
				},
				"target_info": {
					TargetField: "TargetInfo",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"extract_rule": {
							TargetField: "ExtractRule",
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"keys": {
									TargetField: "Keys",
									ForceGet:    true,
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
					},
				},
				"task_name": {
					TargetField: "TaskName",
					ForceGet:    true,
				},
				"topic_id": {
					TargetField: "TopicID",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["TaskId"] = d.Id()
				(*call.SdkParam)["Status"] = d.Get("status")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineImportTaskService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteImportTask",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"TaskId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tls import task on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineImportTaskService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		NameField:       "TaskName",
		IdField:         "TaskId",
		CollectField:    "task_info",
		ResponseConverts: map[string]ve.ResponseConvert{
			"status": {
				TargetField: "status",
			},
			"bucket": {
				TargetField: "bucket",
			},
			"prefix": {
				TargetField: "prefix",
			},
			"region": {
				TargetField: "region",
			},
			"compress_type": {
				TargetField: "compress_type",
			},
			"host": {
				TargetField: "host",
			},
			"group": {
				TargetField: "group",
			},
			"topic": {
				TargetField: "topic",
			},
			"encode": {
				TargetField: "encode",
			},
			"password": {
				TargetField: "password",
			},
			"protocol": {
				TargetField: "protocol",
			},
			"username": {
				TargetField: "username",
			},
			"mechanism": {
				TargetField: "mechanism",
			},
			"instance_id": {
				TargetField: "instance_id",
			},
			"is_need_auth": {
				TargetField: "is_need_auth",
			},
			"initial_offset": {
				TargetField: "initial_offset",
			},
			"time_source_default": {
				TargetField: "time_source_default",
			},
		},
	}
}

func (s *VolcengineImportTaskService) ReadResourceId(id string) string {
	return id
}

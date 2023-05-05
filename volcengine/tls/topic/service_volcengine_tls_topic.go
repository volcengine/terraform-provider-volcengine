package topic

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsTopicService struct {
	Client *ve.SdkClient
}

func NewTlsTopicService(c *ve.SdkClient) *VolcengineTlsTopicService {
	return &VolcengineTlsTopicService{
		Client: c,
	}
}

func (s *VolcengineTlsTopicService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTlsTopicService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp     *map[string]interface{}
		results  interface{}
		tlsTopic map[string]interface{}
		ok       bool
	)
	data, err = ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTopics"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		logger.Debug(logger.RespFormat, action, condition, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.Topics", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("RESPONSE.Topic is not Slice")
		}

		return data, err
	})
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		if tlsTopic, ok = v.(map[string]interface{}); !ok {
			return data, fmt.Errorf(" Topic Value is not map ")
		} else {
			action := "DescribeIndex"
			req := map[string]interface{}{
				"TopicId": tlsTopic["TopicId"],
			}
			logger.Debug(logger.ReqFormat, action, req)
			resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
				ContentType: ve.Default,
				HttpMethod:  ve.GET,
				Path:        []string{action},
				Client:      s.Client.BypassSvcClient.NewTlsClient(),
			}, &req)
			logger.Debug(logger.RespFormat, action, req, *resp)
			if err != nil {
				logger.DebugInfo("DescribeIndex error", "err", err, "TopicId", tlsTopic["TopicId"])
				if ve.ResourceNotFoundError(err) {
					err = nil
					continue
				} else {
					return data, err
				}
			}

			index, err := ve.ObtainSdkValue("RESPONSE", *resp)
			if err != nil {
				logger.DebugInfo("ObtainSdkValue RESPONSE error", err)
				continue
			}
			indexMap, ok := index.(map[string]interface{})
			if !ok {
				logger.Info("Index is not map")
				continue
			}
			tlsTopic["IndexCreateTime"] = indexMap["CreateTime"]
			tlsTopic["IndexModifyTime"] = indexMap["ModifyTime"]
			tlsTopic["FullText"] = indexMap["FullText"]
			tlsTopic["KeyValue"] = indexMap["KeyValue"]

			//keyValueArray, ok := tlsTopic["KeyValue"].([]interface{})
			//if !ok {
			//	logger.Info("Index KeyValues is not slice")
			//	continue
			//}
			//for _, keyValue := range keyValueArray {
			//	keyValueMap, ok := keyValue.(map[string]interface{})
			//	if !ok {
			//		logger.Info("Index KeyValue is not map")
			//		continue
			//	}
			//	valueMap, ok := keyValueMap["Value"].(map[string]interface{})
			//	if !ok {
			//		logger.Info("Index KeyValue Value is not map")
			//		continue
			//	}
			//	jsonKeys := valueMap["JsonKeys"]
			//	if jsonKeys == nil {
			//		continue
			//	}
			//	temp, err := json.Marshal(jsonKeys)
			//	if err != nil {
			//		logger.Debug("Marshal JsonKeys error", "err", err, "JsonKeys", jsonKeys)
			//		delete(valueMap, "JsonKeys")
			//		continue
			//	}
			//	valueMap["JsonKeys"] = string(temp)
			//}
		}
	}

	return data, err
}

func (s *VolcengineTlsTopicService) ReadResource(resourceData *schema.ResourceData, topicId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if topicId == "" {
		topicId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ProjectId": resourceData.Get("project_id"),
		"TopicId":   topicId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		topicMap := make(map[string]interface{})
		if topicMap, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if topicMap["TopicId"].(string) == topicId {
			data = topicMap
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tls topic %s is not exist ", topicId)
	}

	//keyValueArray, ok := data["KeyValue"].([]interface{})
	//if !ok {
	//	return data, fmt.Errorf(" Index KeyValues is not slice ")
	//}
	//for _, keyValue := range keyValueArray {
	//	keyValueMap, ok := keyValue.(map[string]interface{})
	//	if !ok {
	//		return data, fmt.Errorf(" Index KeyValue is not map ")
	//	}
	//	valueMap, ok := keyValueMap["Value"].(map[string]interface{})
	//	if !ok {
	//		return data, fmt.Errorf(" Index KeyValue Value is not map ")
	//	}
	//	keyValueMap["Value"] = []interface{}{valueMap}
	//}

	return data, err
}

func (s *VolcengineTlsTopicService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineTlsTopicService) WithResourceResponseHandlers(topic map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		if _, existIndex := topic["KeyValue"]; !existIndex {
			return topic, nil, nil
		}
		keyValueArray, ok := topic["KeyValue"].([]interface{})
		if !ok {
			return topic, nil, fmt.Errorf(" Index KeyValues is not slice ")
		}
		for _, keyValue := range keyValueArray {
			valueArray := make([]interface{}, 0)
			keyValueMap, ok := keyValue.(map[string]interface{})
			if !ok {
				return topic, nil, fmt.Errorf(" Index KeyValue is not map ")
			}
			valueMap, ok := keyValueMap["Value"].(map[string]interface{})
			if !ok {
				return topic, nil, fmt.Errorf(" Index KeyValue Value is not map ")
			}

			if jsonKeys, ok := valueMap["JsonKeys"]; ok {
				subKeyValueArray, ok := jsonKeys.([]interface{})
				if !ok {
					return topic, nil, fmt.Errorf(" Index JsonKeys KeyValues is not slice ")
				}
				for _, subKeyValue := range subKeyValueArray {
					subValueArray := make([]interface{}, 0)
					subKeyValueMap, ok := subKeyValue.(map[string]interface{})
					if !ok {
						return topic, nil, fmt.Errorf(" Index JsonKeys KeyValue is not map ")
					}
					subValueMap, ok := subKeyValueMap["Value"].(map[string]interface{})
					if !ok {
						return topic, nil, fmt.Errorf(" Index JsonKeys KeyValue Value is not map ")
					}
					subValueArray = append(subValueArray, subValueMap)
					subKeyValueMap["Value"] = subValueArray
				}
			}

			valueArray = append(valueArray, valueMap)
			keyValueMap["Value"] = valueArray
		}
		return topic, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineTlsTopicService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	topicCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateTopic",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"full_text": {
					Ignore: true,
				},
				"key_value": {
					Ignore: true,
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
				id, _ := ve.ObtainSdkValue("RESPONSE.TopicId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	callbacks = append(callbacks, topicCallback)

	indexCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateIndex",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"full_text": {
					ConvertType: ve.ConvertJsonObject,
				},
				"key_value": {
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"value": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"json_keys": {
									ConvertType: ve.ConvertJsonObjectArray,
									NextLevelConvert: map[string]ve.RequestConvert{
										"value": {
											ConvertType: ve.ConvertJsonObject,
										},
									},
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["TopicId"] = d.Id()
					return true, nil
				}
				return false, nil
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
		},
	}
	callbacks = append(callbacks, indexCallback)

	return callbacks
}

func (s *VolcengineTlsTopicService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	topicCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyTopic",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"topic_name": {
					TargetField: "TopicName",
				},
				"ttl": {
					TargetField: "Ttl",
				},
				"auto_split": {
					TargetField: "AutoSplit",
				},
				"max_split_shard": {
					TargetField: "MaxSplitShard",
				},
				"enable_tracking": {
					TargetField: "EnableTracking",
				},
				"time_key": {
					TargetField: "TimeKey",
				},
				"time_format": {
					TargetField: "TimeFormat",
				},
				"description": {
					TargetField: "Description",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["TopicId"] = d.Id()
					return true, nil
				}
				return false, nil
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
	callbacks = append(callbacks, topicCallback)

	logger.DebugInfo("testModifyIndex", resourceData.HasChange("full_text"), resourceData.HasChange("key_value"))
	if resourceData.HasChanges("full_text", "key_value") {
		var (
			createIndex bool
			deleteIndex bool
		)
		indexCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyIndex",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"full_text": {
						ConvertType: ve.ConvertJsonObject,
						ForceGet:    true,
					},
					"key_value": {
						ConvertType: ve.ConvertJsonObjectArray,
						ForceGet:    true,
						//NextLevelConvert: map[string]ve.RequestConvert{
						//	"value": {
						//		ConvertType: ve.ConvertJsonObject,
						//		NextLevelConvert: map[string]ve.RequestConvert{
						//			"json_keys": {
						//				ConvertType: ve.ConvertJsonObjectArray,
						//				NextLevelConvert: map[string]ve.RequestConvert{
						//					"value": {
						//						ConvertType: ve.ConvertJsonObject,
						//						NextLevelConvert: map[string]ve.RequestConvert{
						//							"case_sensitive": {
						//								Ignore: true,
						//							},
						//							"delimiter": {
						//								Ignore: true,
						//							},
						//							"include_chinese": {
						//								Ignore: true,
						//							},
						//							"sql_flag": {
						//								Ignore: true,
						//							},
						//						},
						//					},
						//				},
						//			},
						//		},
						//	},
						//},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					oldFullText, _ := d.GetChange("full_text")
					oldKeyValue, _ := d.GetChange("key_value")
					logger.DebugInfo("testCreateIndex", oldFullText, oldFullText, len(oldFullText.([]interface{})) == 0, len(oldKeyValue.([]interface{})) == 0)
					if len(oldFullText.([]interface{})) == 0 && len(oldKeyValue.([]interface{})) == 0 {
						// 当 Index 从无到有时，调用 CreateIndex 来新建该 Topic 的索引
						createIndex = true
					}
					if len(*call.SdkParam) == 0 {
						// 当 Index 相关的参数被清空时，调用 DeleteIndex 来删除该 Topic 的索引
						deleteIndex = true
					}
					(*call.SdkParam)["TopicId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					if createIndex {
						call.Action = "CreateIndex"
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
							ContentType: ve.ApplicationJSON,
							HttpMethod:  ve.POST,
							Path:        []string{call.Action},
							Client:      s.Client.BypassSvcClient.NewTlsClient(),
						}, call.SdkParam)
					} else if deleteIndex {
						call.Action = "DeleteIndex"
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
							ContentType: ve.ApplicationJSON,
							HttpMethod:  ve.DELETE,
							Path:        []string{call.Action},
							Client:      s.Client.BypassSvcClient.NewTlsClient(),
						}, call.SdkParam)
					} else {
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
							ContentType: ve.ApplicationJSON,
							HttpMethod:  ve.PUT,
							Path:        []string{call.Action},
							Client:      s.Client.BypassSvcClient.NewTlsClient(),
						}, call.SdkParam)
					}
				},
			},
		}
		callbacks = append(callbacks, indexCallback)
	}

	return callbacks
}

func (s *VolcengineTlsTopicService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteTopic",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TopicId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading tls topic on delete %q, %w", d.Id(), callErr))
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
	callbacks = append(callbacks, removeCallback)

	return callbacks
}

func (s *VolcengineTlsTopicService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "TopicName",
		IdField:      "TopicId",
		CollectField: "tls_topics",
		ResponseConverts: map[string]ve.ResponseConvert{
			"TopicId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineTlsTopicService) ReadResourceId(id string) string {
	return id
}

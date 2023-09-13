package index

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsIndexService struct {
	Client *ve.SdkClient
}

func NewTlsIndexService(c *ve.SdkClient) *VolcengineTlsIndexService {
	return &VolcengineTlsIndexService{
		Client: c,
	}
}

func (s *VolcengineTlsIndexService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTlsIndexService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp  *map[string]interface{}
		index interface{}
		ok    bool
	)

	topicIds, exist := condition["TopicIds"]
	if !exist {
		return data, err
	}
	if _, ok = topicIds.([]interface{}); !ok {
		return data, fmt.Errorf(" topic ids is not slice ")
	}
	for _, topicId := range topicIds.([]interface{}) {
		action := "DescribeIndex"
		req := map[string]interface{}{
			"TopicId": topicId,
		}
		logger.DebugInfo(logger.ReqFormat, action, req)
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &req)
		logger.Debug(logger.RespFormat, action, req, *resp)
		if err != nil {
			return data, err
		}
		index, err = ve.ObtainSdkValue("RESPONSE", *resp)
		if err != nil {
			return data, err
		}

		indexMap, ok := index.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf(" Index is not map ")
		}
		keyValue, exist := indexMap["KeyValue"]
		if !exist || keyValue == nil {
			continue
		}
		indexMap["KeyValue"], err = transKeyValueToResponse(keyValue)
		userInnerKeyValue, ok := indexMap["UserInnerKeyValue"]
		if !ok || userInnerKeyValue == nil {
			continue
		}
		indexMap["UserInnerKeyValue"], err = transKeyValueToResponse(userInnerKeyValue)
		if err != nil {
			return data, err
		}
		data = append(data, index)
	}

	return data, err
}

func (s *VolcengineTlsIndexService) ReadResource(resourceData *schema.ResourceData, indexId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if indexId == "" {
		indexId = s.ReadResourceId(resourceData.Id())
	}
	items := strings.Split(indexId, ":")
	if len(items) != 2 {
		return data, fmt.Errorf(" invalid index id: %s", indexId)
	}
	req := map[string]interface{}{
		"TopicIds": []interface{}{items[1]},
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
		return data, fmt.Errorf("tls index %s is not exist ", indexId)
	}

	return data, err
}

func (s *VolcengineTlsIndexService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineTlsIndexService) WithResourceResponseHandlers(index map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return index, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineTlsIndexService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	indexCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateIndex",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"full_text": {
					ConvertType: ve.ConvertJsonObject,
				},
				"key_value": {
					// 框架层对于set套set的类型转换有bug，手动处理
					Ignore: true,
				},
				"user_inner_key_value": {
					// 框架层对于set套set的类型转换有bug，手动处理
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					if keyValueSet, ok := d.GetOk("key_value"); ok {
						keyValues, err := transKeyValueToRequest(keyValueSet)
						if err != nil {
							return false, err
						}
						(*call.SdkParam)["KeyValue"] = keyValues
					}
					if userInnerKeyValueSet, ok := d.GetOk("user_inner_key_value"); ok {
						keyValues, err := transKeyValueToRequest(userInnerKeyValueSet)
						if err != nil {
							return false, err
						}
						(*call.SdkParam)["UserInnerKeyValue"] = keyValues
					}
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
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%v:%v", "index", d.Get("topic_id")))
				return nil
			},
		},
	}
	callbacks = append(callbacks, indexCallback)

	return callbacks
}

func (s *VolcengineTlsIndexService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

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
					// 框架层对于set套set的类型转换有bug，手动处理
					Ignore: true,
				},
				"user_inner_key_value": {
					// 框架层对于set套set的类型转换有bug，手动处理
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["TopicId"] = d.Get("topic_id")
				if keyValueSet, ok := d.GetOk("key_value"); ok {
					keyValues, err := transKeyValueToRequest(keyValueSet)
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["KeyValue"] = keyValues
				}
				if userInnerKeyValueSet, ok := d.GetOk("user_inner_key_value"); ok {
					keyValues, err := transKeyValueToRequest(userInnerKeyValueSet)
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["UserInnerKeyValue"] = keyValues
				}
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
	callbacks = append(callbacks, indexCallback)

	return callbacks
}

func (s *VolcengineTlsIndexService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteIndex",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TopicId": resourceData.Get("topic_id"),
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading tls index on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineTlsIndexService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "TopicIds",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		IdField:      "IndexId",
		CollectField: "tls_indexes",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"TopicId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
		ExtraData: func(i []interface{}) ([]interface{}, error) {
			for index, ele := range i {
				element := ele.(map[string]interface{})
				i[index].(map[string]interface{})["IndexId"] = fmt.Sprintf("%s-%d", "index", element["TopicId"])
			}
			return i, nil
		},
	}
}

func (s *VolcengineTlsIndexService) ReadResourceId(id string) string {
	return id
}

func transKeyValueToRequest(keyValueSet interface{}) ([]interface{}, error) {
	keyValues := make([]interface{}, 0)
	for _, k := range keyValueSet.(*schema.Set).List() {
		keyValue := make(map[string]interface{})
		valueMap := make(map[string]interface{})
		kMap, ok := k.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("key value struct error")
		}
		keyValue["Key"] = kMap["key"]
		valueMap["ValueType"] = kMap["value_type"]
		if v, ok := kMap["case_sensitive"]; ok {
			valueMap["CaseSensitive"] = v
		}
		if v, ok := kMap["include_chinese"]; ok {
			valueMap["IncludeChinese"] = v
		}
		if v, ok := kMap["delimiter"]; ok {
			valueMap["Delimiter"] = v
		}
		if v, ok := kMap["sql_flag"]; ok {
			valueMap["SqlFlag"] = v
		}
		if v, ok := kMap["json_keys"]; ok {
			jsonKeys := make([]interface{}, 0)
			jsonList, ok := v.(*schema.Set)
			if !ok {
				return nil, fmt.Errorf("json keys struct error")
			}
			for _, key := range jsonList.List() {
				jsonKey := make(map[string]interface{})
				keyMap, ok := key.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("json key struct error")
				}
				if v, ok = keyMap["key"]; ok {
					jsonKey["Key"] = v
				}
				if v, ok = keyMap["value_type"]; ok {
					jsonValue := make(map[string]interface{})
					jsonValue["ValueType"] = v
					jsonKey["Value"] = jsonValue
				}
				jsonKeys = append(jsonKeys, jsonKey)
			}
			valueMap["JsonKeys"] = jsonKeys
		}
		keyValue["Value"] = valueMap
		keyValues = append(keyValues, keyValue)
	}
	return keyValues, nil
}

func transKeyValueToResponse(keyValue interface{}) ([]interface{}, error) {
	keyValueArray, ok := keyValue.([]interface{})
	if !ok {
		return []interface{}{}, fmt.Errorf(" Index KeyValues is not slice ")
	}
	for _, v := range keyValueArray {
		keyValueMap, ok := v.(map[string]interface{})
		if !ok {
			return []interface{}{}, fmt.Errorf(" Index KeyValue is not map ")
		}
		valueMap, ok := keyValueMap["Value"].(map[string]interface{})
		if !ok {
			return []interface{}{}, fmt.Errorf(" Index KeyValue Value is not map ")
		}
		for k1, v1 := range valueMap {
			if k1 == "JsonKeys" && v1 != nil {
				jsonArr, ok := v1.([]interface{})
				if !ok {
					return []interface{}{}, fmt.Errorf(" Index KeyValues JsonKeys is not slice ")
				}
				for _, v2 := range jsonArr {
					jsonMap, ok := v2.(map[string]interface{})
					if !ok {
						return []interface{}{}, fmt.Errorf(" Index KeyValue JsonKeys is not map ")
					}
					jsonValueMap, ok := jsonMap["Value"].(map[string]interface{})
					if !ok {
						return []interface{}{}, fmt.Errorf(" Index KeyValue JsonKeys Value is not map ")
					}
					for k3, v3 := range jsonValueMap {
						jsonMap[k3] = v3
						delete(jsonMap, "Value")
					}
				}
			}
			keyValueMap[k1] = v1
			delete(keyValueMap, "Value")
		}
	}
	return keyValueArray, nil
}

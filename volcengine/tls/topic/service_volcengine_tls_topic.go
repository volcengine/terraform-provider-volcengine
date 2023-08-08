package topic

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
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	if tags, exist := condition["Tags"]; exist {
		tagsArr, ok := tags.([]interface{})
		if !ok {
			return data, fmt.Errorf(" Tags in condition is not slice ")
		}
		tagsBytes, err := json.Marshal(tagsArr)
		if err != nil {
			return data, fmt.Errorf(" json marshal tags error: %v", err)
		}
		condition["Tags"] = string(tagsBytes)
	}

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
	projectId, exist := resourceData.GetOkExists("project_id")
	if !exist {
		// import topic 时需要先查询 ProjectId
		action := "DescribeTopic"
		condition := map[string]interface{}{
			"TopicId": topicId,
		}
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		logger.Debug(logger.RespFormat, action, condition, resp)
		if err != nil {
			return data, fmt.Errorf(" DescribeTopic Error: %v", err)
		}
		projectId, err = ve.ObtainSdkValue("RESPONSE.ProjectId", *resp)
		if err != nil || projectId == "" {
			return data, fmt.Errorf(" ObtainSdkValue RESPONSE.ProjectId Error")
		}
	}
	req := map[string]interface{}{
		"ProjectId": projectId,
		"TopicId":   topicId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		var topicMap map[string]interface{}
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

	return data, err
}

func (s *VolcengineTlsTopicService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineTlsTopicService) WithResourceResponseHandlers(topic map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
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
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
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

	// 更新Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

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
		RequestConverts: map[string]ve.RequestConvert{
			"tags": {
				TargetField: "Tags",
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		NameField:    "TopicName",
		IdField:      "TopicId",
		CollectField: "tls_topics",
		ContentType:  ve.ContentTypeJson,
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

func (s *VolcengineTlsTopicService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveTagsFromResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceType"] = "topic"
					(*call.SdkParam)["ResourcesList"] = []string{resourceData.Id()}
					(*call.SdkParam)["TagKeyList"] = make([]string, 0)
					for _, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeyList"] = append((*call.SdkParam)["TagKeyList"].([]string), tag.(map[string]interface{})["key"].(string))
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
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddTagsToResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceType"] = "topic"
					(*call.SdkParam)["ResourcesList"] = []string{resourceData.Id()}
					(*call.SdkParam)["Tags"] = make([]map[string]interface{}, 0)
					for _, tag := range addedTags.List() {
						(*call.SdkParam)["Tags"] = append((*call.SdkParam)["Tags"].([]map[string]interface{}), map[string]interface{}{
							"Key":   tag.(map[string]interface{})["key"],
							"Value": tag.(map[string]interface{})["value"],
						})
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
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

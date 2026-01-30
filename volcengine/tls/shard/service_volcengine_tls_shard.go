package shard

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type Service struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *Service {
	return &Service{
		Client: c,
	}
}

func (s *Service) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *Service) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeShards"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.Shards", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Shards is not Slice")
		}

		return data, err
	})
}

func (s *Service) ReadResource(resourceData *schema.ResourceData, vpcId string) (data map[string]interface{}, err error) {
	topicId := resourceData.Get("topic_id").(string)
	shardId := resourceData.Get("shard_id").(int)
	if topicId == "" {
		return nil, errors.New("topic_id is empty")
	}

	action := "DescribeShards"
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
	logger.Debug(logger.RespFormat, action, resp)
	if err != nil {
		return nil, err
	}

	results, err := ve.ObtainSdkValue("RESPONSE.Shards", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return nil, fmt.Errorf("shards not found for topic %s", topicId)
	}

	shards, ok := results.([]interface{})
	if !ok {
		return nil, errors.New("Result.Shards is not Slice")
	}

	for _, v := range shards {
		shardMap := v.(map[string]interface{})
		if int(shardMap["ShardId"].(float64)) == shardId {
			data = make(map[string]interface{})
			data["topic_id"] = topicId
			data["shard_id"] = shardId
			// We don't really have other properties to set as this resource is an action
			return data, nil
		}
	}

	return nil, fmt.Errorf("shard %d not found in topic %s", shardId, topicId)
}

func (s *Service) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil

}

func (Service) WithResourceResponseHandlers(vpc map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return vpc, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *Service) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ManualShardSplit",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"topic_id": {
					TargetField: "TopicId",
				},
				"shard_id": {
					TargetField: "ShardId",
				},
				"shard_count": {
					TargetField: "Number",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%d", d.Get("topic_id").(string), d.Get("shard_id").(int)))

				results, err := ve.ObtainSdkValue("RESPONSE.Shards", *resp)
				if err != nil {
					return err
				}
				if results == nil {
					return nil
				}
				shards, ok := results.([]interface{})
				if !ok {
					return errors.New("Result.Shards is not Slice")
				}

				var shardList []interface{}
				for _, v := range shards {
					shardMap := v.(map[string]interface{})
					shardList = append(shardList, map[string]interface{}{
						"topic_id":            shardMap["TopicId"],
						"shard_id":            shardMap["ShardId"],
						"inclusive_begin_key": shardMap["InclusiveBeginKey"],
						"exclusive_end_key":   shardMap["ExclusiveEndKey"],
						"status":              shardMap["Status"],
						"modify_time":         shardMap["ModifyTime"],
						"stop_write_time":     shardMap["StopWriteTime"],
					})
				}
				return d.Set("shards", shardList)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *Service) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *Service) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *Service) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "TlsShardId",
		CollectField: "shards",
		ExtraData: func(i []interface{}) ([]interface{}, error) {
			for index, ele := range i {
				element := ele.(map[string]interface{})
				i[index].(map[string]interface{})["TlsShardId"] = fmt.Sprintf("%s-%d", element["TopicId"], element["ShardId"])
			}
			return i, nil
		},
	}
}

func (s *Service) ReadResourceId(id string) string {
	return id
}

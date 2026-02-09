package check_point

import (
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
	return nil, nil
}

func (s *Service) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *Service) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *Service) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *Service) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *Service) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *Service) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *Service) DatasourceResources(d *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *Service) ReadResourceId(id string) string {
	return id
}

func (s *Service) DescribeCheckPoint(m map[string]interface{}) (data []interface{}, err error) {
	action := "DescribeCheckPoint"
	urlParam := map[string]string{
		"ProjectId": fmt.Sprintf("%v", m["ProjectId"]),
		"TopicId":   fmt.Sprintf("%v", m["TopicId"]),
		"ShardId":   fmt.Sprintf("%v", m["ShardId"]),
	}
	body := map[string]interface{}{}
	if v, ok := m["ConsumerGroupName"]; ok && v != "" {
		body["ConsumerGroupName"] = v
	}
	logger.Debug(logger.ReqFormat, action, body)

	resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.ApplicationJSON,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		UrlParam:    urlParam,
		Client:      s.Client.BypassSvcClient.NewTlsClient(),
	}, &body)
	if err != nil {
		return data, err
	}

	checkpoint, err := ve.ObtainSdkValue("RESPONSE.Checkpoint", *resp)
	if err != nil {
		return data, err
	}
	shardID, _ := ve.ObtainSdkValue("RESPONSE.ShardID", *resp)
	var finalShardID interface{}
	if shardID != nil {
		if f, ok := shardID.(float64); ok {
			finalShardID = int(f)
		} else {
			finalShardID = shardID
		}
	}

	data = []interface{}{
		map[string]interface{}{
			"checkpoint": fmt.Sprintf("%v", checkpoint),
			"shard_id":   finalShardID,
		},
	}
	logger.Debug(logger.RespFormat, action, body, data)
	return data, nil
}

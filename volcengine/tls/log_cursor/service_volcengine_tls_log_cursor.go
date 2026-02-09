package log_cursor

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

func (s *Service) DescribeCursor(m map[string]interface{}) (data []interface{}, err error) {
	action := "DescribeCursor"
	urlParam := map[string]string{
		"TopicId": fmt.Sprintf("%v", m["TopicId"]),
		"ShardId": fmt.Sprintf("%v", m["ShardId"]),
	}
	body := map[string]interface{}{
		"From": m["From"],
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
	logger.Debug(logger.RespFormat, action, resp)

	cursor, err := ve.ObtainSdkValue("RESPONSE.Cursor", *resp)
	if err != nil {
		return data, err
	}

	data = []interface{}{
		map[string]interface{}{
			"cursor":   cursor,
			"topic_id": urlParam["TopicId"],
			"shard_id": m["ShardId"],
			"from":     body["From"],
		},
	}
	return data, nil
}

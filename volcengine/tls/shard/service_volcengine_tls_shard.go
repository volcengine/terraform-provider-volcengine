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
		}, &m)
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
	return data, err
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
	return []ve.Callback{}

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

package host

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
		action := "DescribeHosts"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.HostInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.HostInfos is not Slice")
		}

		var res []interface{}
		for _, ele := range data {
			ele.(map[string]interface{})["HostGroupId"] = condition["HostGroupId"] // required field
			res = append(res, ele)
		}
		return res, nil
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

func (s *Service) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	var converter ve.RequestConvert
	val, ok := d.GetOkExists("heartbeat_status")
	if ok {
		converter = ve.RequestConvert{
			Convert: func(data *schema.ResourceData, i interface{}) interface{} {
				return val
			},
		}
	} else {
		converter = ve.RequestConvert{
			Ignore: true,
		}
	}

	return ve.DataSourceInfo{
		IdField:      "HostId",
		CollectField: "host_infos",
		RequestConverts: map[string]ve.RequestConvert{
			"heartbeat_status": converter,
		},
		ExtraData: func(i []interface{}) ([]interface{}, error) {
			for index, ele := range i {
				element := ele.(map[string]interface{})
				i[index].(map[string]interface{})["HostId"] = fmt.Sprintf("%s-%d", element["HostGroupId"], element["Ip"])
			}
			return i, nil
		},
	}
}

func (s *Service) ReadResourceId(id string) string {
	return id
}

package host

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

func (s *Service) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")

	req := map[string]interface{}{
		"HostGroupId": ids[0],
	}
	if len(ids) == 2 && ids[1] != "" {
		req["Ip"] = ids[1]
		results, err = s.ReadResources(req)
		if err != nil {
			return data, err
		}
		if len(results) == 0 {
			// Host not found, return mock data to indicate resource exists (as a deletion task)
			return map[string]interface{}{
				"host_group_id": ids[0],
				"ip":            ids[1],
			}, nil
		} else {
			// Host found, return nil to force recreation (which triggers deletion)
			return nil, nil
		}
	} else {
		req["HeartbeatStatus"] = 0
		results, err = s.ReadResources(req)
		if err != nil {
			return data, err
		}
		if len(results) == 0 {
			// No abnormal hosts, task completed
			return map[string]interface{}{
				"host_group_id": ids[0],
			}, nil
		} else {
			// Abnormal hosts found, trigger recreation (deletion)
			return nil, nil
		}
	}
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
			Action:      "DeleteHost",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"HostGroupId": resourceData.Get("host_group_id"),
				"Ip":          resourceData.Get("ip"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				ip := d.Get("ip").(string)
				if ip == "" {
					call.Action = "DeleteAbnormalHosts"
					delete(*call.SdkParam, "Ip")
				} else {
					call.Action = "DeleteHost"
				}

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				if d.Get("ip").(string) == "" {
					d.SetId(fmt.Sprintf("%s", d.Get("host_group_id").(string)))
				} else {
					d.SetId(fmt.Sprintf("%s:%s", d.Get("host_group_id").(string), d.Get("ip").(string)))
				}
				return nil
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

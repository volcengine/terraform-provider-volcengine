package escloud_ip_white_list

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/escloud_v2/escloud_instance_v2"
)

type VolcengineEscloudIpWhiteListService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEscloudIpWhiteListService(c *ve.SdkClient) *VolcengineEscloudIpWhiteListService {
	return &VolcengineEscloudIpWhiteListService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEscloudIpWhiteListService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEscloudIpWhiteListService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeInstances"

		// 重新组织 Filter 的格式
		if filter, filterExist := condition["Filters"]; filterExist {
			newFilter := make([]interface{}, 0)
			for k, v := range filter.(map[string]interface{}) {
				newFilter = append(newFilter, map[string]interface{}{
					"Name":   k,
					"Values": v,
				})
			}
			condition["Filters"] = newFilter
		}

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		results, err = ve.ObtainSdkValue("Result.Instances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Instances is not Slice")
		}

		return data, err
	})
}

func (s *VolcengineEscloudIpWhiteListService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(id, ":")
	if len(ids) != 3 {
		return data, fmt.Errorf("Invalid ip white list id: %s ", id)
	}

	req := map[string]interface{}{
		"Filters": map[string]interface{}{
			"InstanceId": []string{ids[0]},
		},
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
		return data, fmt.Errorf("escloud_instance_v2 %s not exist ", id)
	}

	if ids[1] == "public" && ids[2] == "es" {
		data["IpList"] = strings.Split(data["ESPublicIpWhitelist"].(string), ",")
	} else if ids[1] == "public" && ids[2] == "kibana" {
		data["IpList"] = strings.Split(data["KibanaPublicIpWhitelist"].(string), ",")
	} else if ids[1] == "private" && ids[2] == "es" {
		data["IpList"] = strings.Split(data["ESPrivateIpWhitelist"].(string), ",")
	} else if ids[1] == "private" && ids[2] == "kibana" {
		data["IpList"] = strings.Split(data["KibanaPrivateIpWhitelist"].(string), ",")
	}

	return data, err
}

func (s *VolcengineEscloudIpWhiteListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineEscloudIpWhiteListService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEscloudIpWhiteListService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyIpWhitelist",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"ip_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ipArr := d.Get("ip_list").(*schema.Set).List()
				ipList := make([]string, 0)
				for _, id := range ipArr {
					ipList = append(ipList, id.(string))
				}
				(*call.SdkParam)["IpList"] = strings.Join(ipList, ",")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				ipType := d.Get("type").(string)
				component := d.Get("component").(string)
				d.SetId(instanceId + ":" + ipType + ":" + component)
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				escloud_instance_v2.NewEscloudInstanceV2Service(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEscloudIpWhiteListService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyIpWhitelist",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_id": {
					TargetField: "InstanceId",
					ForceGet:    true,
				},
				"type": {
					TargetField: "Type",
					ForceGet:    true,
				},
				"component": {
					TargetField: "Component",
					ForceGet:    true,
				},
				"ip_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ipArr := d.Get("ip_list").(*schema.Set).List()
				ipList := make([]string, 0)
				for _, id := range ipArr {
					ipList = append(ipList, id.(string))
				}
				(*call.SdkParam)["IpList"] = strings.Join(ipList, ",")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				escloud_instance_v2.NewEscloudInstanceV2Service(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEscloudIpWhiteListService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyIpWhitelist",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_id": {
					TargetField: "InstanceId",
					ForceGet:    true,
				},
				"type": {
					TargetField: "Type",
					ForceGet:    true,
				},
				"component": {
					TargetField: "Component",
					ForceGet:    true,
				},
				"ip_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["IpList"] = ""
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				escloud_instance_v2.NewEscloudInstanceV2Service(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEscloudIpWhiteListService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineEscloudIpWhiteListService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ESCloud",
		Version:     "2023-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

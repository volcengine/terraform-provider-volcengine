package cloud_server

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

type VolcengineCloudServerService struct {
	Client     *ve.SdkClient
}

func NewCloudServerService(c *ve.SdkClient) *VolcengineCloudServerService {
	return &VolcengineCloudServerService{
		Client:     c,
	}
}

func (s *VolcengineCloudServerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudServerService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	if idsObject, ok := m["ids"]; !ok {
		return ve.WithPageNumberQuery(m, "limit", "page", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
			action := "ListCloudServers"
			bytes, _ := json.Marshal(condition)
			logger.Debug(logger.ReqFormat, action, string(bytes))

			if condition == nil {
				resp, err = s.Client.UniversalClient.DoCall(universalGet(action), nil)
				if err != nil {
					return data, err
				}
			} else {
				resp, err = s.Client.UniversalClient.DoCall(universalGet(action), &condition)
				if err != nil {
					return data, err
				}
			}
			respBytes, _ := json.Marshal(resp)
			logger.Debug(logger.RespFormat, action, condition, string(respBytes))
			results, err = ve.ObtainSdkValue("Result.cloud_servers", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			if data, ok = results.([]interface{}); !ok {
				return data, errors.New(" Result.cloud_servers is not Slice")
			}
			return data, err
		})
	} else {
		ids := idsObject.([]interface{})
		for _, id := range ids {
			con := map[string]interface{}{
				"cloud_server_id": id,
			}
			logger.Debug(logger.ReqFormat, "GetCloudServer", con)
			resp, err = s.Client.UniversalClient.DoCall(universalGet("GetCloudServer"), &con)
			if err != nil {
				return data, err
			}

			bytes, _ := json.Marshal(resp)
			logger.Debug(logger.RespFormat, "GetCloudServer", con, string(bytes))
			results, err = ve.ObtainSdkValue("Result.cloud_server", *resp)
			if err != nil {
				return data, err
			}
			data = append(data, results)
		}
	}
	return data, err
}

func (s *VolcengineCloudServerService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ids": []interface{}{id},
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
		return data, fmt.Errorf("Cloud Server %s not exist ", id)
	}

	ins, err := s.getInstances(map[string]interface{}{
		"cloud_server_identities": id,
	})
	if err != nil {
		return nil, err
	}
	if len(ins) == 0 {
		return data, err
	}
	insIds := make([]interface{}, 0)
	for _, ele := range ins {
		insIds = append(insIds, ele.(map[string]interface{})["instance_identity"])
	}
	data["default_instance_id"] = insIds[0]
	return data, err
}

func (s *VolcengineCloudServerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,

		Refresh: func() (result interface{}, state string, err error) {
			var (
				resp   map[string]interface{}
				status interface{}
			)
			resp, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			logger.Debug("Refresh Cloud Server status resp:%v", "ReadResource", resp)
			status, err = ve.ObtainSdkValue("instance_status", resp)
			if err != nil {
				return nil, "", err
			}
			insStatues := status.([]interface{})
			for _, ele := range insStatues {
				status := ele.(map[string]interface{})
				if status["status"] != "running" && status["instance_count"].(float64) > 0 {
					return resp, "Creating", nil
				}
			}
			// 所有的实例都处在运行中才可以
			return resp, "Running", nil
		},
	}
}

func (s *VolcengineCloudServerService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		if cfg, ok := d["secret_config"]; ok {
			secretType := cfg.(map[string]interface{})["secret_type"].(float64)
			switch secretType {
			case 3:
				d["secret_type"] = "KeyPair"
			default:
				d["secret_type"] = "Password"
			}
		}

		return d, map[string]ve.ResponseConvert{
			"name": {
				TargetField: "cloudserver_name",
			},
			"spec": {
				TargetField: "spec_name",
			},
			"image": {
				TargetField: "image_id",
				Convert: func(i interface{}) interface{} {
					return i.(map[string]interface{})["image_identity"]
				},
			},
			"schedule_strategy_configs": {
				TargetField: "schedule_strategy",
				Convert: func(i interface{}) interface{} {
					return []interface{}{i}
				},
			},
			"network": {
				TargetField: "network_config",
				Convert: func(i interface{}) interface{} {
					return []interface{}{i}
				},
			},
			"storage": {
				TargetField: "storage_config",
				Convert: func(i interface{}) interface{} {
					res := map[string]interface{}{}
					ele := i.(map[string]interface{})

					if sd, ok := ele["system_disk"]; ok {
						res["system_disk"] = []interface{}{sd}
					}
					if ddl, ok := ele["data_disk_list"]; ok {
						res["data_disk_list"] = ddl
					}
					return res
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

var diskSpecConvert = map[string]ve.RequestConvert{
	"storage_type": {
		TargetField: "storage_type",
	},
	"capacity": {
		TargetField: "capacity",
	},
}

var createResourceConvert = map[string]ve.RequestConvert{
	"cloudserver_name": {
		TargetField: "cloudserver_name",
	},
	"image_id": {
		TargetField: "image_id",
	},
	"spec_name": {
		TargetField: "spec_name",
	},
	"server_area_level": {
		TargetField: "server_area_level",
	},
	"secret_type": {
		TargetField: "secret_config.secret_type",
		Convert: func(data *schema.ResourceData, old interface{}) interface{} {
			ty := 0
			switch old.(string) {
			case "KeyPair":
				ty = 3
			case "Password":
				ty = 2
			}
			return ty
		},
	},
	"secret_data": {
		TargetField: "secret_config.secret_data",
	},
	"custom_data": {
		TargetField: "custom_data",
		ConvertType: ve.ConvertJsonObject,
		NextLevelConvert: map[string]ve.RequestConvert{
			"data": {
				TargetField: "data",
			},
		},
	},
	"billing_config": {
		TargetField: "billing_config",
		ConvertType: ve.ConvertJsonObject,
		NextLevelConvert: map[string]ve.RequestConvert{
			"computing_billing_method": {
				TargetField: "computing_billing_method",
			},
			"bandwidth_billing_method": {
				TargetField: "bandwidth_billing_method",
			},
		},
	},
	"schedule_strategy": {
		TargetField: "schedule_strategy",
		ConvertType: ve.ConvertJsonObject,
		NextLevelConvert: map[string]ve.RequestConvert{
			"schedule_strategy": {
				TargetField: "schedule_strategy",
			},
			"price_strategy": {
				TargetField: "price_strategy",
			},
			"network_strategy": {
				TargetField: "network_strategy",
			},
		},
	},
	"network_config": {
		TargetField: "network_config",
		ConvertType: ve.ConvertJsonObject,
		NextLevelConvert: map[string]ve.RequestConvert{
			"bandwidth_peak": {
				TargetField: "bandwidth_peak",
			},
			"internal_bandwidth_peak": {
				TargetField: "internal_bandwidth_peak",
			},
			"enable_ipv6": {
				TargetField: "enable_ipv6",
			},
			"custom_internal_interface_name": {
				TargetField: "custom_internal_interface_name",
			},
			"custom_external_interface_name": {
				TargetField: "custom_external_interface_name",
			},
		},
	},
	"storage_config": {
		TargetField: "storage_config",
		ConvertType: ve.ConvertJsonObject,
		NextLevelConvert: map[string]ve.RequestConvert{
			"system_disk": {
				TargetField:      "system_disk",
				ConvertType:      ve.ConvertJsonObject,
				NextLevelConvert: diskSpecConvert,
			},
			"data_disk": {
				TargetField:      "data_disk",
				ConvertType:      ve.ConvertJsonObject,
				NextLevelConvert: diskSpecConvert,
			},
			"data_disk_list": {
				TargetField:      "data_disk_list",
				ConvertType:      ve.ConvertJsonObjectArray,
				NextLevelConvert: diskSpecConvert,
			},
		},
	},

	"default_area_name": {
		TargetField: "instance_area_config.area_name",
	},
	"default_isp": {
		TargetField: "instance_area_config.isp",
	},
	"default_cluster_name": {
		TargetField: "instance_area_config.cluster_name",
	},
}

func (s *VolcengineCloudServerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateCloudServer",
			ContentType: ve.ContentTypeJson,
			Convert:     createResourceConvert,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				config := (*call.SdkParam)["instance_area_config"].(map[string]interface{})
				config["num"] = 1 // 默认创建一条

				(*call.SdkParam)["instance_area_nums"] = []interface{}{
					config,
				}

				bytes, _ := json.Marshal(call.SdkParam)
				logger.Debug(logger.ReqFormat, call.Action, string(bytes))
				return s.Client.UniversalClient.DoCall(universalPost(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.cloud_server_identity", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudServerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	return callbacks
}

func (s *VolcengineCloudServerService) getInstances(condition map[string]interface{}) ([]interface{}, error) {
	return ve.WithPageNumberQuery(condition, "limit", "page", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		logger.Debug(logger.ReqFormat, "ListInstances", condition)
		resp, err := s.Client.UniversalClient.DoCall(universalGet("ListInstances"), &condition)
		if err != nil {
			return nil, err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, "ListInstances", condition, string(respBytes))
		results, err := ve.ObtainSdkValue("Result.instances", *resp)
		if err != nil {
			return nil, err
		}
		if results == nil {
			results = []interface{}{}
		}
		data, ok := results.([]interface{})
		if !ok {
			return nil, errors.New(" Result.instances is not Slice")
		}
		return data, nil
	})
}

func (s *VolcengineCloudServerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteCloudServer",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["cloud_server_id"] = resourceData.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(universalPost(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudServerService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType: ve.ContentTypeJson,
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ids",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		IdField:      "cloud_server_identity",
		CollectField: "cloud_servers",
		NameField:    "name",
		ResponseConverts: map[string]ve.ResponseConvert{
			"cloud_server_identity": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineCloudServerService) ReadResourceId(id string) string {
	return id
}

func universalPost(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "veenedge",
		Version:     "2021-04-30",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func universalGet(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "veenedge",
		Version:     "2021-04-30",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
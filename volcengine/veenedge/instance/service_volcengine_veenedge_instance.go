package instance

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
)

type VolcengineInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewInstanceService(c *ve.SdkClient) *VolcengineInstanceService {
	return &VolcengineInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func interfaceSlice2String(ele []interface{}) []string {
	var res []string
	for _, i := range ele {
		res = append(res, i.(string))
	}
	return res
}

func (s *VolcengineInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "limit", "page", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListInstances"
		// 列表查询接口全是 , 拼接形式
		if ids, ok := condition["instance_identities"]; ok {
			condition["instance_identities"] = strings.Join(interfaceSlice2String(ids.([]interface{})), ",")
		}
		if status, ok := condition["status"]; ok {
			condition["status"] = strings.Join(interfaceSlice2String(status.([]interface{})), ",")
		}
		if ids, ok := condition["cloud_server_identities"]; ok {
			condition["cloud_server_identities"] = strings.Join(interfaceSlice2String(ids.([]interface{})), ",")
		}
		if names, ok := condition["instance_names"]; ok {
			condition["instance_names"] = strings.Join(interfaceSlice2String(names.([]interface{})), ",")
		}

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
		results, err = ve.ObtainSdkValue("Result.instances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New(" Result.instances is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	results, err = s.ReadResources(map[string]interface{}{
		"instance_identities": []interface{}{id},
	})
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Instance %s not exist ", id)
	}
	data["instance_id"] = id // 填充import来的字段
	return data, err
}

func (s *VolcengineInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      3 * time.Second,
		MinTimeout: 3 * time.Second,
		Target:     target,
		Timeout:    timeout,

		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "error")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("instance status error,status %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineInstanceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		if ct, ok := d["cluster"]; ok {
			cluster := ct.(map[string]interface{})
			// instance里返回的 region 就是 area_name
			d["area_name"] = cluster["region"]
			d["isp"] = cluster["isp"]
			d["cluster_name"] = cluster["cluster_name"]
		}

		return d, map[string]ve.ResponseConvert{
			"instance_name": {
				TargetField: "name",
			},
			"cloud_server_identity": {
				TargetField: "cloudserver_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	if id, ok := resourceData.GetOk("instance_id"); ok {
		return []ve.Callback{{
			Call: ve.SdkCall{
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return nil, nil
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					d.SetId(id.(string))
					return nil
				},
			},
		}}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateInstance",
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_id": {
					TargetField: "instance_id",
				},
				"cloudserver_id": {
					TargetField: "cloud_server_identity",
				},
				"area_name": {
					TargetField: "instance_area_config.area_name",
				},
				"isp": {
					TargetField: "instance_area_config.isp",
				},
				"default_isp": {
					TargetField: "instance_area_config.default_isp",
				},
				"cluster_name": {
					TargetField: "instance_area_config.cluster_name",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				con := map[string]interface{}{
					"cloud_server_id": (*call.SdkParam)["cloud_server_identity"],
				}
				logger.Debug(logger.ReqFormat, "GetCloudServer", con)
				resp, err := s.Client.UniversalClient.DoCall(universalGet("GetCloudServer"), &con)
				if err != nil {
					return false, err
				}

				bytes, _ := json.Marshal(resp)
				logger.Debug(logger.RespFormat, "GetCloudServer", con, string(bytes))
				res, err := ve.ObtainSdkValue("Result.cloud_server", *resp)
				if err != nil {
					return false, err
				}
				if serverAreaLevel, ok := res.(map[string]interface{})["server_area_level"]; ok {
					(*call.SdkParam)["server_area_level"] = serverAreaLevel
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				if _, ok := (*call.SdkParam)["instance_area_config"]; ok {
					config := (*call.SdkParam)["instance_area_config"].(map[string]interface{})
					config["num"] = 1 // 默认创建一条

					(*call.SdkParam)["instance_area_nums"] = []interface{}{
						config,
					}
				}

				bytes, _ := json.Marshal(call.SdkParam)
				logger.Debug(logger.ReqFormat, call.Action, string(bytes))
				return s.Client.UniversalClient.DoCall(universalPost(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ids, _ := ve.ObtainSdkValue("Result.instance_id_list", *resp)
				instanceIds, ok := ids.([]interface{})
				if !ok {
					return fmt.Errorf("instance_id_list result is not slice")
				}
				if len(instanceIds) == 0 {
					return fmt.Errorf("create instance error")
				}
				d.SetId(instanceIds[0].(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("name") {
		id := resourceData.Id()
		name := resourceData.Get("name")

		renameCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SetInstanceName",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["instance_identity"] = id
					(*call.SdkParam)["instance_name"] = name
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(universalPost(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, renameCallback)
	}

	if resourceData.HasChange("secret_type") || resourceData.HasChange("secret_data") {
		id := resourceData.Id()

		ty := 0
		secretType := resourceData.Get("secret_type")
		switch secretType {
		case "KeyPair":
			ty = 3
		case "Password":
			ty = 2
		}

		resetAdminPasswdCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ResetLoginCredential",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["instance_identity"] = id
					(*call.SdkParam)["secret_type"] = ty
					(*call.SdkParam)["secret_data"] = resourceData.Get("secret_data")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(universalPost(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, resetAdminPasswdCallback)
	}
	return callbacks
}

func (s *VolcengineInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	stopCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "StopInstances",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["instance_identities"] = []interface{}{resourceData.Id()}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(universalPost("StopInstances"), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"stop"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "OfflineInstances",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["instance_identities"] = []interface{}{resourceData.Id()}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(universalPost("OfflineInstances"), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading instance on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{stopCallback, callback}
}

func (s *VolcengineInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "instance_identities",
				ConvertType: ve.ConvertJsonArray,
			},
			"names": {
				TargetField: "instance_names",
				ConvertType: ve.ConvertJsonArray,
			},
			"cloud_server_ids": {
				TargetField: "cloud_server_identities",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "status",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		IdField:      "instance_identity",
		CollectField: "instances",
		ResponseConverts: map[string]ve.ResponseConvert{
			"instance_identity": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineInstanceService) ReadResourceId(id string) string {
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

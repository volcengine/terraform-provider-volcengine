package addon

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

var addonSupportUpdate []string

func init() {
	addonSupportUpdate = []string{
		"cr-credential-controller",
		"apmplus-opentelemetry-collector",
		"cluster-autoscaler",
		"ingress-nginx",
		"p2p-accelerator",
		"nvidia-device-plugin",
		"prometheus-agent",
		"scheduler-plugin",
		"mgpu",
	}
}

type VolcengineVkeAddonService struct {
	Client *ve.SdkClient
}

func NewVkeAddonService(c *ve.SdkClient) *VolcengineVkeAddonService {
	return &VolcengineVkeAddonService{
		Client: c,
	}
}

func (s *VolcengineVkeAddonService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVkeAddonService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	data, err = ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "ListAddons"
		logger.Debug(logger.ReqFormat, action, condition)
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
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
	return data, err
}

func (s *VolcengineVkeAddonService) ReadResource(resourceData *schema.ResourceData, resourceId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if resourceId == "" {
		resourceId = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(resourceId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid addon id")
	}

	clusterId := ids[0]
	name := ids[1]

	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"ClusterIds": []string{clusterId},
			"Names":      []string{name},
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
		return data, fmt.Errorf("Vke Addon %s:%s not exist ", clusterId, name)
	}
	data["CompleteConfig"] = data["Config"]
	if cfg, ok := resourceData.GetOkExists("config"); ok {
		// 返回的 config 可能会添加默认参数，这里始终使用创建的
		data["Config"] = cfg
	}
	return data, err
}

func (s *VolcengineVkeAddonService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	failStateTimes := 0 // 处于Failed状态的组件可能经过短暂的时间（自愈）状态变成 Running
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status.Phase", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					failStateTimes++
					if failStateTimes > 10 { // 硬编码：检查10次还是错误状态
						return nil, "", fmt.Errorf(" Vke addon status error, status:%s", status.(string))
					}
					return demo, "", nil
				}
			}
			return demo, status.(string), err
		},
	}

}

func (VolcengineVkeAddonService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineVkeAddonService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAddon",
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				//check addon exist
				clusterId := (*call.SdkParam)["ClusterId"]
				name := (*call.SdkParam)["Name"]
				data, err := s.ReadResource(d, fmt.Sprintf("%s:%s", clusterId, name))
				if err != nil {
					if !strings.Contains(err.Error(), "not exist") { // 其他类型的错误
						return nil, err
					}
				}
				//addon exist
				if len(data) > 0 {
					version := data["Version"]

					hclVersion, ok := d.GetOk("version")
					if ok && hclVersion != version { // 如果用户在tf中定义了 version，并且与已有的不相等
						params := map[string]interface{}{
							"ClusterId": clusterId,
							"Name":      name,
							"Version":   hclVersion,
						}
						logger.Debug(logger.ReqFormat, "UpdateAddonVersion", params)
						_, err = s.Client.UniversalClient.DoCall(getUniversalInfo("UpdateAddonVersion"), &params)
						if err != nil {
							return nil, err
						}
						// wait status
						_, err = s.RefreshResourceState(d, []string{"Running"}, resourceData.Timeout(schema.TimeoutCreate), fmt.Sprintf("%s:%s", clusterId, name)).WaitForState()
						if err != nil {
							return nil, err
						}
					}

					// 接下来检查 Config
					return s.updateConfig(name, data, client, call)
				}
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("cluster_id"), d.Get("name"))
				d.SetId(id)
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

func (s *VolcengineVkeAddonService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	versionCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateAddonVersion",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"version": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) == 0 {
					return false, nil
				}

				databaseId := d.Id()
				ids := strings.Split(databaseId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid addon id")
				}
				(*call.SdkParam)["ClusterId"] = ids[0]
				(*call.SdkParam)["Name"] = ids[1]

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				req, err := json.Marshal(*call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.ReqFormat, call.Action, string(req))
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateAddonConfig",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"config": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) == 0 {
					return false, nil
				}

				databaseId := d.Id()
				ids := strings.Split(databaseId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid addon id")
				}
				(*call.SdkParam)["ClusterId"] = ids[0]
				(*call.SdkParam)["Name"] = ids[1]

				//if ids[1] == "ingress-nginx" {
				//	return false, fmt.Errorf("ingress-nginx addon prohibits updating config")
				//}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				req, err := json.Marshal(*call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.ReqFormat, call.Action, string(req))
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}

	return []ve.Callback{versionCallback, callback}
}

func (s *VolcengineVkeAddonService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAddon",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				databaseId := d.Id()
				ids := strings.Split(databaseId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid addon id")
				}
				(*call.SdkParam)["ClusterId"] = ids[0]
				(*call.SdkParam)["Name"] = ids[1]
				(*call.SdkParam)["CascadingDeleteResources"] = []string{"Crd"}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				if strings.Contains(baseErr.Error(), "forbidden to delete required addon") { // 一些组件禁止删除，直接返回
					msg := fmt.Sprintf("error: %s. msg: %s",
						baseErr.Error(),
						"If you want to remove it form terraform state, "+
							"please use `terraform state rm volcengine_vke_addon.resource_name` command ")
					return fmt.Errorf(msg)
				}
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading addon on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{callback}
}

func (s *VolcengineVkeAddonService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"cluster_ids": {
				TargetField: "Filter.ClusterIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"names": {
				TargetField: "Filter.Names",
				ConvertType: ve.ConvertJsonArray,
			},
			"deploy_mode": {
				TargetField: "Filter.DeployModes",
				ConvertType: ve.ConvertJsonArray,
			},
			"deploy_node_types": {
				TargetField: "Filter.DeployNodeTypes",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "Filter.Statuses",
				ConvertType: ve.ConvertJsonObjectArray,
				NextLevelConvert: map[string]ve.RequestConvert{
					"phase": {
						TargetField: "Phase",
					},
					"conditions_type": {
						TargetField: "ConditionsType",
					},
				},
			},
			"create_client_token": {
				TargetField: "Filter.CreateClientToken",
			},
			"update_client_token": {
				TargetField: "Filter.UpdateClientToken",
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		CollectField: "addons",
	}
}

func (s *VolcengineVkeAddonService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineVkeAddonService) updateConfig(name interface{}, data map[string]interface{}, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
	var (
		source    map[string]interface{}
		target    map[string]interface{}
		needCheck bool
	)
	if config, ok := data["CompleteConfig"].(string); ok {
		err := json.Unmarshal([]byte(config), &source)
		if err != nil {
			return nil, err
		}
		needCheck = true
	} else {
		needCheck = false
	}

	if c, ok := (*call.SdkParam)["Config"]; ok {
		if config, ok1 := c.(string); ok1 {
			err := json.Unmarshal([]byte(config), &target)
			if err != nil {
				return nil, err
			}
			needCheck = true
		} else {
			needCheck = false
		}
	} else {
		needCheck = false
	}
	if needCheck && !reflect.DeepEqual(source, target) && checkSupportUpdate(fmt.Sprintf("%s", name)) {
		//update config
		logger.Debug(logger.ReqFormat, "UpdateAddonConfig", *call.SdkParam)
		return s.Client.UniversalClient.DoCall(getUniversalInfo("UpdateAddonConfig"), call.SdkParam)
	}
	return nil, nil
}

func checkSupportUpdate(name string) bool {
	for _, addon := range addonSupportUpdate {
		if name == addon {
			return true
		}
	}
	return false
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vke",
		Version:     "2022-05-12",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

package cr_registry

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCrRegistryService struct {
	Client *ve.SdkClient
}

func NewCrRegistryService(c *ve.SdkClient) *VolcengineCrRegistryService {
	return &VolcengineCrRegistryService{
		Client: c,
	}
}

func (s *VolcengineCrRegistryService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCrRegistryService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	pageCall := func(condition map[string]interface{}) ([]interface{}, error) {
		// Get registry
		action := "ListRegistries"
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
			return data, errors.New("Results.Items is not slice")
		}

		for i, v := range data {
			ins := v.(map[string]interface{})
			condition := &map[string]interface{}{
				"Registry": ins["Name"],
			}

			status, err := ve.ObtainSdkValue("Status.Phase", ins)
			if err != nil {
				return data, err
			}
			if status.(string) == "Creating" || status.(string) == "Deleting" || status.(string) == "Failed" {
				logger.DebugInfo("registry status is Creating/Deleting/Failed,skip GetUser and ListDomains%s", "")
				continue
			}

			//get user
			action = "GetUser"
			logger.Debug(logger.ReqFormat, action, condition)
			resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), condition)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, condition, *resp)
			username, err := ve.ObtainSdkValue("Result.Username", *resp)
			if err != nil {
				return data, err
			}
			userStatus, err := ve.ObtainSdkValue("Result.Status", *resp)
			if err != nil {
				return data, err
			}

			data[i].(map[string]interface{})["Username"] = username
			data[i].(map[string]interface{})["UserStatus"] = userStatus

			//get domains
			action = "ListDomains"
			logger.Debug(logger.ReqFormat, action, condition)
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), condition)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, condition, *resp)
			results, err = ve.ObtainSdkValue("Result.Items", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			data[i].(map[string]interface{})["Domains"] = results
		}

		return data, err
	}

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, pageCall)
}

func (s *VolcengineCrRegistryService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Names": []string{id},
		},
	}

	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}

	for _, v := range results {
		data, ok = v.(map[string]interface{})
		if !ok {
			return data, errors.New("value is not a map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("CrRegistry %s is not exist", id)
	}
	return data, err
}

func (s *VolcengineCrRegistryService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, name string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,

		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				userStatus interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			demo, err = s.ReadResource(resourceData, name)
			if err != nil {
				return nil, "", err
			}
			logger.Debug("Refresh CrRegistry status resp:%v", "ReadResource", demo)

			status, err = ve.ObtainSdkValue("Status.Phase", demo)
			if err != nil {
				return nil, "", err
			}

			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("CrRegistry status error,status %s", status.(string))
				}
			}

			//must wait user status
			if len(target) > 0 && target[0] == "Active" {
				userStatus, err = ve.ObtainSdkValue("UserStatus", demo)
				if userStatus != "Active" {
					status = "InActive"
				} else {
					if status == "Running" {
						status = userStatus
					}
				}
			}
			return demo, status.(string), err
		},
	}
}

func (s *VolcengineCrRegistryService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, map[string]ve.ResponseConvert{
			"SkipSSLVerify": {
				TargetField: "skip_ssl_verify",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCrRegistryService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRegistry",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"project": {
					TargetField: "Project",
				},
				"resource_tags": {
					TargetField: "ResourceTags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"type": {
					TargetField: "Type",
				},
				"proxy_cache_enabled": {
					TargetField: "ProxyCacheEnabled",
				},
				"proxy_cache": {
					TargetField: "ProxyCache",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"skip_ssl_verify": {
							TargetField: "SkipSSLVerify",
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Name"] = resourceData.Get("name")
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := d.Get("name").(string)
				d.SetId(id)
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)
	if password, ok := resourceData.GetOkExists("password"); ok {
		action := "SetUser"
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      action,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["Registry"] = resourceData.Get("name")
					(*call.SdkParam)["Password"] = base64.StdEncoding.EncodeToString([]byte(password.(string)))
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Active"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, callback)
	}
	return callbacks
}

func (s *VolcengineCrRegistryService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	if resourceData.HasChange("password") {
		action := "SetUser"
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      action,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					bytes := []byte(resourceData.Get("password").(string))
					(*call.SdkParam)["Registry"] = resourceData.Get("name")
					(*call.SdkParam)["Password"] = base64.StdEncoding.EncodeToString(bytes)
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Active"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, callback)
	}
	return callbacks
}

func (s *VolcengineCrRegistryService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRegistry",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Name"] = d.Id()
				(*call.SdkParam)["DeleteImmediately"] = d.Get("delete_immediately")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				deleteImmediately := d.Get("delete_immediately").(bool)
				// 如选择立即销毁，则进行removed检查
				if deleteImmediately {
					return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCrRegistryService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		IdField:      "Name",
		CollectField: "registries",
		RequestConverts: map[string]ve.RequestConvert{
			"names": {
				TargetField: "Filter.Names",
				ConvertType: ve.ConvertJsonArray,
			},
			"types": {
				TargetField: "Filter.Types",
				ConvertType: ve.ConvertJsonArray,
			},
			"projects": {
				TargetField: "Filter.Projects",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "Filter.Statuses",
				ConvertType: ve.ConvertJsonObjectArray,
			},
			"resource_tags": {
				TargetField: "ResourceTagFilters",
				ConvertType: ve.ConvertJsonObjectArray,
				NextLevelConvert: map[string]ve.RequestConvert{
					"key": {
						TargetField: "Key",
					},
					"values": {
						TargetField: "Values",
						ConvertType: ve.ConvertJsonArray,
					},
				},
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"SkipSSLVerify": {
				TargetField: "skip_ssl_verify",
			},
		},
	}
}

func (s *VolcengineCrRegistryService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineCrRegistryService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "cr",
		ResourceType:         "instance",
		ProjectResponseField: "Project",
		ProjectSchemaField:   "project",
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cr",
		Version:     "2022-05-12",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

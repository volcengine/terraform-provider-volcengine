package cr_instance

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCrRegistryService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCrRegistryService(c *ve.SdkClient) *VolcengineCrRegistryService {
	return &VolcengineCrRegistryService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
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

		logger.Debug(logger.RespFormat, action, condition, resp)
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

		//get user
		for i, v := range data {
			ins := v.(map[string]interface{})
			cond := &map[string]interface{}{
				"Registry": ins["Registry"],
			}

			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("GetUser"), cond)
			if err != nil {
				return data, err
			}

			results, err = ve.ObtainSdkValue("Result", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				return data, fmt.Errorf("getUser return an empty result")
			}

			status, err := ve.ObtainSdkValue("Result.Status", *resp)
			if err != nil {
				return data, err
			}
			username, err := ve.ObtainSdkValue("Result.Username", *resp)
			if err != nil {
				return data, err
			}

			user := map[string]interface{}{
				"Status":   status,
				"Username": username,
			}
			data[i].(map[string]interface{})["User"] = user
		}

		return data, err
	}

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, pageCall)
}

func (s *VolcengineCrRegistryService) ReadResource(resourceData *schema.ResourceData, name string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if name == "" {
		name = s.ReadResourceId(resourceData.Id())
	} else {
		resourceData.Set("registry", name)
	}

	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Names": []string{name},
		},
	}

	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}

	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not a map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("CrRegistry %s is not exist", name)
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
			return demo, status.(string), err
		},
	}
}

func (s *VolcengineCrRegistryService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCrRegistryService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRegistry",
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Name"] = resourceData.Get("name")
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
	return []ve.Callback{callback}
}

func (s *VolcengineCrRegistryService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCrRegistryService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRegistry",
			ContentType: ve.ContentTypeJson,
			//ConvertMode: ve.RequestConvertAll,
			//Convert: map[string]ve.RequestConvert{
			//	"Name": {
			//		TargetField: "name",
			//	},
			//	"DeleteImmediately": {
			//		TargetField: "delete_immediately",
			//	},
			//},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Name"] = resourceData.Id()
				delete_imme := resourceData.Get("delete_immediately")
				logger.DebugInfo("delete_imme: %v", delete_imme)
				(*call.SdkParam)["DeleteImmediately"] = resourceData.Get("delete_immediately").(bool)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
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
			"statuses": {
				TargetField: "Filter.Statuses",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"Name": {
				TargetField: "registry",
			},
		},
	}
}

func (s *VolcengineCrRegistryService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cr_pre",
		Version:     "2022-05-12",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

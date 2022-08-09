package user

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCrUserService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCrUserService(c *ve.SdkClient) *VolcengineCrUserService {
	return &VolcengineCrUserService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCrUserService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCrUserService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)

	pageCall := func(condition map[string]interface{}) ([]interface{}, error) {
		action := "GetUser"

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

		logger.Debug(logger.RespFormat, action, resp)
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

		return []interface{}{user}, err
	}

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, pageCall)
}

func (s *VolcengineCrUserService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)

	if id == "" {
		id = resourceData.Get("registry").(string)
	} else {
		resourceData.Set("registry", id) // for import,writeback registry
	}
	req := map[string]interface{}{
		"Registry": id,
	}

	results, err = s.ReadResources(req)

	if err != nil {
		return data, err
	}

	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("GetUser value is not a map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cr user %s is not exist", id)
	}
	return data, err
}

func (s *VolcengineCrUserService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,

		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status interface{}
			)

			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			logger.DebugInfo("Refresh CrUser status resp:%v", demo)

			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}

			return demo, status.(string), err
		},
	}
}

func (VolcengineCrUserService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCrUserService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	action := "SetUser"
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				registry := resourceData.Get("registry").(string)
				password := resourceData.Get("password").(string)

				bytes := []byte(password)
				passwdBase64 := base64.StdEncoding.EncodeToString(bytes)

				(*call.SdkParam)["Registry"] = registry
				(*call.SdkParam)["password"] = passwdBase64
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := d.Get("registry").(string)
				logger.DebugInfo("After create instance state,registry name:%s", id)
				d.SetId(id)
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active", "Inactive"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineCrUserService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return s.CreateResource(resourceData, resource)
}

func (s *VolcengineCrUserService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCrUserService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		CollectField: "users",
	}
}

func (s *VolcengineCrUserService) ReadResourceId(id string) string {
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

package cr_repository

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

type VolcengineCrRepositoryService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCrRepositoryService(c *ve.SdkClient) *VolcengineCrRepositoryService {
	return &VolcengineCrRepositoryService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCrRepositoryService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCrRepositoryService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	pageCall := func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRepositories"

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
		return data, err
	}

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, pageCall)
}

func (s *VolcengineCrRepositoryService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	parts := strings.Split(id, ":")
	if len(parts) != 3 {
		return data, fmt.Errorf("the id format must be 'registry:namespace:name'")
	}
	registry := parts[0]
	namespace := parts[1]
	name := parts[2]

	req := map[string]interface{}{
		"Registry": registry,
		"Filter": map[string]interface{}{
			"Namespaces": []string{namespace},
			"Names":      []string{name},
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
		return data, fmt.Errorf("CrRepository %s is not exist", name)
	}
	return data, err
}

func (s *VolcengineCrRepositoryService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, name string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineCrRepositoryService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCrRepositoryService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRepository",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				registry := d.Get("registry").(string)
				namespace := d.Get("namespace").(string)
				name := d.Get("name").(string)
				id := registry + ":" + namespace + ":" + name
				d.SetId(id)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCrRepositoryService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRepository",
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Registry"] = resourceData.Get("registry").(string)
				(*call.SdkParam)["Namespace"] = resourceData.Get("namespace").(string)
				(*call.SdkParam)["Name"] = resourceData.Get("name").(string)
				if resourceData.HasChange("description") {
					(*call.SdkParam)["Description"] = resourceData.Get("description").(string)
				}
				if resourceData.HasChange("access_level") {
					(*call.SdkParam)["AccessLevel"] = resourceData.Get("access_level").(string)
				}
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

func (s *VolcengineCrRepositoryService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRepository",
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id := resourceData.Id()
				parts := strings.Split(id, ":")
				if len(parts) != 3 {
					return false, fmt.Errorf("the id format must be 'registry:namespace:name'")
				}
				(*call.SdkParam)["Registry"] = parts[0]
				(*call.SdkParam)["Namespace"] = parts[1]
				(*call.SdkParam)["Name"] = parts[2]
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

func (s *VolcengineCrRepositoryService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		IdField:      "Name",
		CollectField: "repositories",
		RequestConverts: map[string]ve.RequestConvert{
			"namespaces": {
				TargetField: "Filter.Namespaces",
				ConvertType: ve.ConvertJsonArray,
			},
			"names": {
				TargetField: "Filter.Names",
				ConvertType: ve.ConvertJsonArray,
			},
			"access_levels": {
				TargetField: "Filter.AccessLevels",
				ConvertType: ve.ConvertJsonArray,
			},
		},
	}
}

func (s *VolcengineCrRepositoryService) ReadResourceId(id string) string {
	return id
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

package cr_tag

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

type VolcengineCrTagService struct {
	Client *ve.SdkClient
}

func NewCrTagService(c *ve.SdkClient) *VolcengineCrTagService {
	return &VolcengineCrTagService{
		Client: c,
	}
}

func (s *VolcengineCrTagService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCrTagService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	pageCall := func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListTags"

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
		return data, err
	}

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, pageCall)
}

func (s *VolcengineCrTagService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	logger.DebugInfo("read resource id :%s", id)
	parts := strings.Split(id, ":")
	if len(parts) != 4 {
		return data, fmt.Errorf("the format of import id must be 'registry:namespace:repository:tag...'")
	}

	registry := parts[0]
	namespace := parts[1]
	repository := parts[2]
	name := parts[3]

	req := map[string]interface{}{
		"Registry":   registry,
		"Namespace":  namespace,
		"Repository": repository,
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
		return data, fmt.Errorf("cr tag %s not exist", name)
	}
	return data, err
}

func (s *VolcengineCrTagService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, name string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineCrTagService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	if _, ok := instance["ImageAttributes"]; !ok {
		instance["ImageAttributes"] = []interface{}{}
	}
	if _, ok := instance["ChartAttribute"]; !ok {
		instance["ChartAttribute"] = map[string]interface{}{}
	}
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineCrTagService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCrTagService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCrTagService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteTags",
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id := resourceData.Id()
				parts := strings.Split(id, ":")
				if len(parts) != 4 {
					return false, fmt.Errorf("the id format must be 'registry:namespace:repository:tag'")
				}
				(*call.SdkParam)["Registry"] = parts[0]
				(*call.SdkParam)["Namespace"] = parts[1]
				(*call.SdkParam)["Repository"] = parts[2]
				(*call.SdkParam)["Names"] = []string{parts[3]}
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

func (s *VolcengineCrTagService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		IdField:      "Name",
		CollectField: "tags",
		RequestConverts: map[string]ve.RequestConvert{
			"names": {
				TargetField: "Filter.Names",
				ConvertType: ve.ConvertJsonArray,
			},
			"types": {
				TargetField: "Filter.Types",
				ConvertType: ve.ConvertJsonArray,
			},
		},
	}
}

func (s *VolcengineCrTagService) ReadResourceId(id string) string {
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

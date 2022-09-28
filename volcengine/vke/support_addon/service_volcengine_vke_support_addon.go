package support_addon

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVkeSupportAddonService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVkeSupportAddonService(c *ve.SdkClient) *VolcengineVkeSupportAddonService {
	return &VolcengineVkeSupportAddonService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVkeSupportAddonService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVkeSupportAddonService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	action := "ListSupportedAddons"
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
}

func (s *VolcengineVkeSupportAddonService) ReadResource(resourceData *schema.ResourceData, clusterId string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineVkeSupportAddonService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineVkeSupportAddonService) WithResourceResponseHandlers(cluster map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineVkeSupportAddonService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}

}

func (s *VolcengineVkeSupportAddonService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineVkeSupportAddonService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineVkeSupportAddonService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"name": {
				TargetField: "Filter.Name",
			},
			"pod_network_modes": {
				TargetField: "Filter.PodNetworkModes",
				ConvertType: ve.ConvertJsonArray,
			},
			"deploy_modes": {
				TargetField: "Filter.DeployModes",
				ConvertType: ve.ConvertJsonArray,
			},
			"deploy_node_types": {
				TargetField: "Filter.DeployNodeTypes",
				ConvertType: ve.ConvertJsonArray,
			},
			"necessaries": {
				TargetField: "Filter.Necessaries",
				ConvertType: ve.ConvertJsonArray,
			},
			"categories": {
				TargetField: "Filter.Categories",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		CollectField: "addons",
	}
}

func (s *VolcengineVkeSupportAddonService) ReadResourceId(id string) string {
	return id
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

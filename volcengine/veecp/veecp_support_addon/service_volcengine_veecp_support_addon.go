package veecp_support_addon

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVeecpSupportAddonService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVeecpSupportAddonService(c *ve.SdkClient) *VolcengineVeecpSupportAddonService {
	return &VolcengineVeecpSupportAddonService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVeecpSupportAddonService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVeecpSupportAddonService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	if _, ok := condition["Filter"]; ok {
		if kubernetesVersions, ok := condition["Filter"].(map[string]interface{})["KubernetesVersions"]; ok {
			condition["Filter"].(map[string]interface{})["Versions.Compatibilities.KubernetesVersions"] = kubernetesVersions
			delete(condition["Filter"].(map[string]interface{}), "KubernetesVersions")
		}
	}

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

func (s *VolcengineVeecpSupportAddonService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineVeecpSupportAddonService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineVeecpSupportAddonService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (VolcengineVeecpSupportAddonService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVeecpSupportAddonService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineVeecpSupportAddonService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineVeecpSupportAddonService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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
			"kubernetes_versions": {
				TargetField: "Filter.KubernetesVersions",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		CollectField: "addons",
	}
}

func (s *VolcengineVeecpSupportAddonService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "veecp_openapi",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

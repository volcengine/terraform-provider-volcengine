package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackVkeNodeService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVestackVkeNodeService(c *ve.SdkClient) *VestackVkeNodeService {
	return &VestackVkeNodeService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackVkeNodeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackVkeNodeService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListNodes"
		// 单独适配VKE Conditions.Type 字段，该字段 API 表示不规范
		if filter, filterExist := condition["Filter"]; filterExist {
			if statuses, exist := filter.(map[string]interface{})["Statuses"]; exist {
				for index, status := range statuses.([]interface{}) {
					if ty, ex := status.(map[string]interface{})["ConditionsType"]; ex {
						condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{})["Conditions.Type"] = ty
						delete(condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{}), "ConditionsType")
					}
				}
			}
		}

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.VkeClient.ListNodesCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.VkeClient.ListNodesCommon(&condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
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
}

func (s *VestackVkeNodeService) ReadResource(resourceData *schema.ResourceData, nodeId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if nodeId == "" {
		nodeId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []string{nodeId},
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
		return data, fmt.Errorf("Node %s not exist ", nodeId)
	}
	return data, err
}

func (s *VestackVkeNodeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VestackVkeNodeService) WithResourceResponseHandlers(vpc map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return vpc, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackVkeNodeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}

}

func (s *VestackVkeNodeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VestackVkeNodeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VestackVkeNodeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"cluster_ids": {
				TargetField: "Filter.ClusterIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"name": {
				TargetField: "Filter.Name",
			},
			"create_client_token": {
				TargetField: "Filter.CreateClientToken",
			},
			"node_pool_ids": {
				TargetField: "Filter.NodePoolIds",
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
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "nodes",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Id": {
				TargetField: "id",
				KeepDefault: true,
			},
			"Status.Phase": {
				TargetField: "phase",
			},
			"Status.Conditions": {
				TargetField: "condition_types",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						for _, _data := range dd {
							results = append(results, _data.(map[string]interface{})["Type"])
						}
					}
					return results
				},
			},
		},
	}
}

func (s *VestackVkeNodeService) ReadResourceId(id string) string {
	return id
}

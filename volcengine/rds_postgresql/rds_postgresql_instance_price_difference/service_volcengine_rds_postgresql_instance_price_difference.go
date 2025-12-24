package rds_postgresql_instance_price_difference

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlInstancePriceDifferenceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstancePriceDifferenceService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstancePriceDifferenceService {
	return &VolcengineRdsPostgresqlInstancePriceDifferenceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp   *map[string]interface{}
		result interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstancePriceDifference"

		// ModifyType is Temporary, RollbackTime is required
		temporarySet := false
		if mt, ok := condition["ModifyType"].(string); ok && mt == "Temporary" {
			temporarySet = true
			if rt, ok2 := condition["RollbackTime"].(string); !ok2 || rt == "" {
				return data, errors.New("rollback_time is required when modify_types is Temporary")
			}
		}
		if v, ok := condition["NodeInfo"]; ok && v != nil {
			nodes, ok2 := v.([]interface{})
			if !ok2 {
				return data, errors.New("node_info is not Slice")
			}
			var (
				primarySpec    string
				secondarySpec  string
				primaryCount   int
				secondaryCount int
				sanitized      []interface{}
			)
			for _, n := range nodes {
				node, ok3 := n.(map[string]interface{})
				if !ok3 {
					continue
				}
				nodeType, _ := node["NodeType"].(string)
				zoneId, _ := node["ZoneId"].(string)
				nodeSpec, _ := node["NodeSpec"].(string)
				nodeId, _ := node["NodeId"].(string)
				operate, _ := node["NodeOperateType"].(string)
				if nodeType == "Primary" {
					primaryCount++
					primarySpec = nodeSpec
				} else if nodeType == "Secondary" {
					secondaryCount++
					secondarySpec = nodeSpec
				}
				nodeNew := make(map[string]interface{})
				nodeNew["NodeType"] = nodeType
				nodeNew["ZoneId"] = zoneId
				nodeNew["NodeSpec"] = nodeSpec
				if nodeId == "" && temporarySet {
					return data, errors.New("node_id is required when modify_type is Temporary")
				} else {
					nodeNew["NodeId"] = nodeId
				}
				if operate != "" {
					nodeNew["NodeOperateType"] = operate
				}
				sanitized = append(sanitized, nodeNew)
			}
			if primaryCount != 1 || secondaryCount != 1 {
				return data, errors.New("must provide exactly one Primary and one Secondary in NodeInfo")
			}
			if primarySpec == "" || secondarySpec == "" || primarySpec != secondarySpec {
				return data, errors.New("the Primary and Secondary node_spec must be identical")
			}
			condition["NodeInfo"] = sanitized
		} else {
			return data, errors.New("NodeInfo is required")
		}

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return data, err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		result, err = ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if result == nil {
			return []interface{}{}, nil
		}
		if m, okm := result.(map[string]interface{}); okm {
			// normalize empty ChargeItemPrices to empty list
			if cip, ok3 := m["ChargeItemPrices"]; !ok3 || cip == nil {
				m["ChargeItemPrices"] = []interface{}{}
			}
			return []interface{}{m}, nil
		}
		return data, errors.New("the Result is not Map")
	})
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return nil, "", err
		},
	}
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlInstancePriceDifferenceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"node_info": {
				TargetField: "NodeInfo",
				ConvertType: ve.ConvertJsonObjectArray,
			},
			"charge_info": {
				TargetField: "ChargeInfo",
				ConvertType: ve.ConvertJsonObject,
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"ChargeItemPrices": {
				TargetField: "charge_item_prices",
				KeepDefault: true,
			},
		},
		NameField:    "PayablePrice",
		IdField:      "PayablePrice",
		CollectField: "instances_price",
		ContentType:  ve.ContentTypeJson,
	}
}

func (s *VolcengineRdsPostgresqlInstancePriceDifferenceService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

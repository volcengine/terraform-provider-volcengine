package rds_postgresql_instance_price_detail

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlInstancePriceDetailService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstancePriceDetailService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstancePriceDetailService {
	return &VolcengineRdsPostgresqlInstancePriceDetailService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstancePriceDetail"
		// sanitize NodeInfo: must contain one Primary and one Secondary with same NodeSpec; default NodeOperateType to Create; ignore NodeId
		if condition != nil {
			if v, ok1 := condition["NodeInfo"]; ok1 && v != nil {
				nodes, ok2 := v.([]interface{})
				if !ok2 {
					return data, errors.New("NodeInfo is not Slice")
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
					operate, _ := node["NodeOperateType"].(string)
					// count and record specs
					if nodeType == "Primary" {
						primaryCount++
						primarySpec = nodeSpec
					} else if nodeType == "Secondary" {
						secondaryCount++
						secondarySpec = nodeSpec
					}
					// build minimal node map，符合API要求
					nodeNew := make(map[string]interface{})
					nodeNew["NodeType"] = nodeType
					if zoneId != "" {
						nodeNew["ZoneId"] = zoneId
					}
					if nodeSpec != "" {
						nodeNew["NodeSpec"] = nodeSpec
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
		}

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		// extract single result object and wrap as slice
		results, err = ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			return []interface{}{}, nil
		}
		if m, okm := results.(map[string]interface{}); okm {
			// normalize empty ChargeItemPrices to empty list
			if cip, ok3 := m["ChargeItemPrices"]; !ok3 || cip == nil {
				m["ChargeItemPrices"] = []interface{}{}
			}
			return []interface{}{m}, nil
		}
		return data, errors.New("the Result is not Map")
	})
}

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlInstancePriceDetailService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
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
			"storage_type": {
				TargetField: "StorageType",
				ConvertType: ve.ConvertDefault,
			},
			"storage_space": {
				TargetField: "StorageSpace",
				ConvertType: ve.ConvertDefault,
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

func (s *VolcengineRdsPostgresqlInstancePriceDetailService) ReadResourceId(id string) string {
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

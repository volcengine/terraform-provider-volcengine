package rds_postgresql_engine_version_parameter

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlEngineVersionParameterService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlEngineVersionParameterService(c *ve.SdkClient) *VolcengineRdsPostgresqlEngineVersionParameterService {
	return &VolcengineRdsPostgresqlEngineVersionParameterService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBEngineVersionParameters"

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
		version, _ := ve.ObtainSdkValue("Result.DBEngineVersion", *resp)
		count, _ := ve.ObtainSdkValue("Result.ParameterCount", *resp)
		results, err = ve.ObtainSdkValue("Result.Parameters", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if arr, ok := results.([]interface{}); ok {
			wrapped := map[string]interface{}{
				"db_engine_version": version,
				"parameter_count":   count,
				"parameters":        convertEngineParams(arr),
			}
			return []interface{}{wrapped}, nil
		}
		return data, errors.New("Result.Parameters is not Slice")
	})
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	req := map[string]interface{}{}
	if v, ok := resourceData.GetOk("db_engine"); ok {
		req["DBEngine"] = v
	}
	if v, ok := resourceData.GetOk("db_engine_version"); ok {
		req["DBEngineVersion"] = v
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); ok {
			break
		}
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return nil, "", err
		},
	}
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlEngineVersionParameterService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"db_engine": {
				TargetField: "DBEngine",
			},
			"db_engine_version": {
				TargetField: "DBEngineVersion",
			},
		},
		CollectField: "db_engine_version_parameters",
		ResponseConverts: map[string]ve.ResponseConvert{
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
		},
	}
}

func (s *VolcengineRdsPostgresqlEngineVersionParameterService) ReadResourceId(id string) string {
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

func convertEngineParams(arr []interface{}) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(arr))
	for _, it := range arr {
		if m, ok := it.(map[string]interface{}); ok {
			item := map[string]interface{}{}
			if v, ok2 := m["Name"]; ok2 {
				item["name"] = v
			}
			if v, ok2 := m["Type"]; ok2 {
				item["type"] = v
			}
			if v, ok2 := m["DefaultValue"]; ok2 {
				item["default_value"] = v
			}
			if v, ok2 := m["ForceRestart"]; ok2 {
				item["force_restart"] = v
			}
			if v, ok2 := m["CheckingCode"]; ok2 {
				item["checking_code"] = v
			}
			out = append(out, item)
		}
	}
	return out
}

package rds_postgresql_instance_parameter

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlInstanceParameterService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstanceParameterService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceParameterService {
	return &VolcengineRdsPostgresqlInstanceParameterService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstanceParameters"

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

		result := make(map[string]interface{})
		if v, e := ve.ObtainSdkValue("Result.InstanceId", *resp); e == nil && v != nil {
			result["instance_id"] = v
		}
		if v, e := ve.ObtainSdkValue("Result.DBEngineVersion", *resp); e == nil && v != nil {
			result["db_engine_version"] = v
		}
		if v, e := ve.ObtainSdkValue("Result.ParameterCount", *resp); e == nil && v != nil {
			result["parameter_count"] = v
		}

		if v, e := ve.ObtainSdkValue("Result.Parameters", *resp); e == nil && v != nil {
			if arr, ok := v.([]interface{}); ok {
				result["parameters"] = convertParamArray(arr)
			}
		}
		if v, e := ve.ObtainSdkValue("Result.NoneKernelParameters", *resp); e == nil && v != nil {
			if arr, ok := v.([]interface{}); ok {
				result["none_kernel_parameters"] = convertParamArray(arr)
			}
		}

		return []interface{}{result}, nil

	})
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlInstanceParameterService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlInstanceParameterService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id": {
				TargetField: "InstanceId",
			},
			"parameter_name": {
				TargetField: "ParameterName",
			},
		},
		CollectField: "instance_parameters",
		ContentType:  ve.ContentTypeJson,
	}
}

func (s *VolcengineRdsPostgresqlInstanceParameterService) ReadResourceId(id string) string {
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

func convertParamArray(arr []interface{}) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(arr))
	for _, it := range arr {
		if m, ok := it.(map[string]interface{}); ok {
			item := map[string]interface{}{}
			if v, ok2 := m["CheckingCode"]; ok2 {
				item["checking_code"] = v
			}
			if v, ok2 := m["DefaultValue"]; ok2 {
				item["default_value"] = v
			}
			if v, ok2 := m["Description"]; ok2 {
				item["description"] = v
			}
			if v, ok2 := m["DescriptionZH"]; ok2 {
				item["description_zh"] = v
			}
			if v, ok2 := m["ForceRestart"]; ok2 {
				item["force_restart"] = v
			}
			if v, ok2 := m["Name"]; ok2 {
				item["name"] = v
			}
			if v, ok2 := m["Type"]; ok2 {
				item["type"] = v
			}
			if v, ok2 := m["Value"]; ok2 {
				item["value"] = v
			}
			out = append(out, item)
		}
	}
	return out
}

package rds_postgresql_instance_spec

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlInstanceSpecService struct {
	Client *ve.SdkClient
}

func NewRdsPostgresqlInstanceSpecService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceSpecService {
	return &VolcengineRdsPostgresqlInstanceSpecService{
		Client: c,
	}
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstanceSpecs"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return nil, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return nil, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.InstanceSpecs", *resp)
		if err != nil {
			return nil, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return nil, errors.New("Result.InstanceSpecs is not Slice")
		}
		return data, nil
	})
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) ReadResource(*schema.ResourceData, string) (map[string]interface{}, error) {
	return nil, nil
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlInstanceSpecService) CreateResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlInstanceSpecService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) ModifyResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) RemoveResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"zone_id": {
				TargetField: "ZoneId",
			},
			"db_engine_version": {
				TargetField: "DBEngineVersion",
			},
			"spec_code": {
				TargetField: "SpecCode",
			},
			"storage_type": {
				TargetField: "StorageType",
			},
		},
		CollectField: "instance_specs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Connection": {
				TargetField: "connection",
			},
			"SpecCode": {
				TargetField: "spec_code",
			},
			"VCPU": {
				TargetField: "v_cpu",
			},
			"Memory": {
				TargetField: "memory",
			},
			"StorageType": {
				TargetField: "storage_type",
			},
			"SpecStatus": {
				TargetField: "spec_status",
			},
			"RegionId": {
				TargetField: "region_id",
			},
			// 条件返回字段：仅当请求包含时，后端才会返回
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"ZoneId": {
				TargetField: "zone_id",
			},
		},
		ContentType: ve.ContentTypeJson,
	}
}

func (s *VolcengineRdsPostgresqlInstanceSpecService) ReadResourceId(id string) string {
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

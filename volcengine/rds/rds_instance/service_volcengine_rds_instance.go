package rds_instance

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsInstanceService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func NewRdsInstanceService(c *volc.SdkClient) *VolcengineRdsInstanceService {
	return &VolcengineRdsInstanceService{
		Client:     c,
		Dispatcher: &volc.Dispatcher{},
	}
}

func (s *VolcengineRdsInstanceService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		ok          bool
		rdsInstance map[string]interface{}
	)
	data, err = volc.WithPageOffsetQuery(m, "Limit", "Offset", 20, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListDBInstances"
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

		results, err = volc.ObtainSdkValue("Result.Datas", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Datas is not Slice")
		}
		return data, err
	})
	if err != nil {
		return data, err
	}

	for _, v := range data {
		if rdsInstance, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else {
			// query rds connection info
			connResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeDBInstanceConnection"), &map[string]interface{}{
				"InstanceId": rdsInstance["InstanceId"],
			})
			if err != nil {
				logger.Info("DescribeDBInstanceConnection error:", err)
				continue
			}
			connInfo, err := volc.ObtainSdkValue("Result.ConnectionInfo", *connResp)
			if err != nil {
				logger.Info("ObtainSdkValue Result.ConnectionInfo error:", err)
				continue
			}
			if connInfo != nil {
				rdsInstance["ConnectionInfo"] = connInfo
			}
		}
	}

	return data, err
}

func (s *VolcengineRdsInstanceService) ReadResource(resourceData *schema.ResourceData, rdsInstanceId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if rdsInstanceId == "" {
		rdsInstanceId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InstanceId": rdsInstanceId,
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
		return data, fmt.Errorf("Rds instance %s not exist ", rdsInstanceId)
	}

	return data, err
}

func (s *VolcengineRdsInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = volc.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("Rds instance status error, status:%s ", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineRdsInstanceService) WithResourceResponseHandlers(rdsInstance map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		instanceSpecName, err := volc.ObtainSdkValue("InstanceSpec.SpecName", rdsInstance)
		if err != nil {
			return nil, nil, err
		}
		rdsInstance["InstanceSpecName"] = instanceSpecName
		return rdsInstance, map[string]volc.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBEngine": {
				TargetField: "db_engine",
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"StorageSpaceGB": {
				TargetField: "storage_space_gb",
			},
			"VpcID": {
				TargetField: "vpc_id",
			},
		}, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineRdsInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateDBInstance",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertAll,
			Convert: map[string]volc.RequestConvert{
				"db_engine": {
					TargetField: "DBEngine",
				},
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"storage_space_gb": {
					TargetField: "StorageSpaceGB",
				},
				"vpc_id": {
					TargetField: "VpcID",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建rdsInstance
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := volc.ObtainSdkValue("Result.InstanceId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &volc.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []volc.Callback{callback}

}

func (s *VolcengineRdsInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	return []volc.Callback{}
}

func (s *VolcengineRdsInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteDBInstance",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除RdsInstance
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if volc.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rds instance on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "InstanceName",
		IdField:      "InstanceId",
		CollectField: "rds_instances",
		ResponseConverts: map[string]volc.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBEngine": {
				TargetField: "db_engine",
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"StorageSpaceGB": {
				TargetField: "storage_space_gb",
			},
			"VpcID": {
				TargetField: "vpc_id",
			},
		},
	}
}

func (s *VolcengineRdsInstanceService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2018-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}

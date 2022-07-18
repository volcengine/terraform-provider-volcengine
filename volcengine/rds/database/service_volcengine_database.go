package database

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineDatabaseService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func NewDatabaseService(c *volc.SdkClient) *VolcengineDatabaseService {
	return &VolcengineDatabaseService{
		Client:     c,
		Dispatcher: &volc.Dispatcher{},
	}
}

func (s *VolcengineDatabaseService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineDatabaseService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	list, err := volc.WithPageOffsetQuery(m, "Limit", "Offset", 20, 0, func(condition map[string]interface{}) (data []interface{}, err error) {
		var (
			resp    *map[string]interface{}
			results interface{}
			ok      bool
		)
		rdsClient := s.Client.RdsClient
		action := "ListDatabases"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = rdsClient.ListDatabasesCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = rdsClient.ListDatabasesCommon(&condition)
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
		return list, err
	}

	targetDBName := m["DBName"]

	// 拼接id
	res := make([]interface{}, 0)
	for _, d := range list {
		db, ok := d.(map[string]interface{})
		if !ok {
			continue
		}

		if targetDBName != nil && targetDBName.(string) != db["DBName"].(string) {
			// ListDatabases接口不支持根据dbName过滤，这里手动过滤下
			continue
		}

		db["Id"] = fmt.Sprintf("%s:%s", m["InstanceId"], db["DBName"])
		res = append(res, db)
	}
	return res, nil
}

func (s *VolcengineDatabaseService) ReadResource(resourceData *schema.ResourceData, DatabaseId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if DatabaseId == "" {
		DatabaseId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(DatabaseId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid database id")
	}

	instanceId := ids[0]
	dbName := ids[1]

	req := map[string]interface{}{
		"InstanceId": instanceId,
		"DBName":     dbName,
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
		return data, fmt.Errorf("Database %s not exist ", DatabaseId)
	}

	return data, err
}

func (s *VolcengineDatabaseService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status interface{}
			)
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = volc.ObtainSdkValue("DBStatus", demo)
			if err != nil {
				return nil, "", err
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (VolcengineDatabaseService) WithResourceResponseHandlers(database map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return database, nil, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineDatabaseService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateDatabase",
			ConvertMode: volc.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建Database
				return s.Client.RdsClient.CreateDatabaseCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("instance_id"), d.Get("db_name"))
				d.SetId(id)
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

func (s *VolcengineDatabaseService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	return []volc.Callback{}
}

func (s *VolcengineDatabaseService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteDatabase",
			ConvertMode: volc.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				databaseId := d.Id()
				ids := strings.Split(databaseId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid rds account id")
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["DBName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除Database
				return s.Client.RdsClient.DeleteDatabaseCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if volc.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading RDS account on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineDatabaseService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		RequestConverts: map[string]volc.RequestConvert{
			"db_status": {
				TargetField: "DBStatus",
			},
		},
		NameField:    "DBName",
		CollectField: "databases",
		ResponseConverts: map[string]volc.ResponseConvert{
			"DBPrivileges": {
				TargetField: "db_privileges",
			},
			"DBStatus": {
				TargetField: "db_status",
			},
			"DBName": {
				TargetField: "db_name",
			},
		},
	}
}

func (s *VolcengineDatabaseService) ReadResourceId(id string) string {
	return id
}

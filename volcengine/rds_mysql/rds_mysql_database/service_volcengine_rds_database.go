package rds_mysql_database

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/common"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsMysqlDatabaseService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func NewRdsMysqlDatabaseService(c *volc.SdkClient) *VolcengineRdsMysqlDatabaseService {
	return &VolcengineRdsMysqlDatabaseService{
		Client:     c,
		Dispatcher: &volc.Dispatcher{},
	}
}

func (s *VolcengineRdsMysqlDatabaseService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlDatabaseService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	list, err := volc.WithPageOffsetQuery(m, "PageSize", "PageNumber", 20, 0, func(condition map[string]interface{}) (data []interface{}, err error) {
		var (
			resp    *map[string]interface{}
			results interface{}
			ok      bool
		)
		action := "DescribeDatabases"
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
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		results, err = volc.ObtainSdkValue("Result.Databases", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Databases is not Slice")
		}
		return data, err
	})
	if err != nil {
		return list, err
	}

	dbName := m["DBName"]
	res := make([]interface{}, 0)
	for _, d := range list {
		db, ok := d.(map[string]interface{})
		if !ok {
			continue
		}
		// DBName是模糊搜索，需要过滤一下
		if dbName != "" && dbName != db["DBName"].(string) {
			continue
		}
		// 拼接id
		db["Id"] = fmt.Sprintf("%s:%s", m["InstanceId"], db["DBName"])
		res = append(res, db)
	}
	return res, nil
}

func (s *VolcengineRdsMysqlDatabaseService) ReadResource(resourceData *schema.ResourceData, rdsDatabaseId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if rdsDatabaseId == "" {
		rdsDatabaseId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(rdsDatabaseId, ":")
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
		return data, fmt.Errorf("RDS database %s not exist ", rdsDatabaseId)
	}

	return data, err
}

func (s *VolcengineRdsMysqlDatabaseService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineRdsMysqlDatabaseService) WithResourceResponseHandlers(database map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return database, map[string]volc.ResponseConvert{
			"DBName": {
				TargetField: "db_name",
			},
		}, nil
	}
	return []volc.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlDatabaseService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateDatabase",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertAll,
			Convert: map[string]volc.RequestConvert{
				"database_privileges": {
					TargetField: "DatabasePrivileges",
					ConvertType: volc.ConvertJsonObjectArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建Database
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("instance_id"), d.Get("db_name"))
				d.SetId(id)
				return nil
			},
		},
	}
	return []volc.Callback{callback}
}

func hasAccountNameInSet(account string, set *schema.Set) bool {
	for _, item := range set.List() {
		if m, ok := item.(map[string]interface{}); ok {
			if v, ok := m["account_name"]; ok && v.(string) == account {
				return true
			}
		}
	}
	return false
}

func (s *VolcengineRdsMysqlDatabaseService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callbacks := make([]volc.Callback, 0)
	if resourceData.HasChange("database_privileges") {
		addPrivis, removePrivs, _, _ := common.GetSetDifference("database_privileges", resourceData, RdsMysqlDatabasePrivilegeHash, false)
		if addPrivis.Len() != 0 {
			callback := volc.Callback{
				Call: volc.SdkCall{
					Action:      "GrantDatabasePrivilege",
					ConvertMode: volc.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
						(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
						(*call.SdkParam)["DBName"] = d.Get("db_name")
						dbPrivileges := make([]map[string]interface{}, 0)
						for _, item := range addPrivis.List() {
							m, _ := item.(map[string]interface{})
							privi := make(map[string]interface{})
							if v, ok := m["account_name"]; ok {
								privi["AccountName"] = v
							}
							if v, ok := m["account_privilege"]; ok {
								privi["AccountPrivilege"] = v
							}
							if v, ok := m["account_privilege_detail"]; ok {
								privi["AccountPrivilegeDetail"] = v
							}
							dbPrivileges = append(dbPrivileges, privi)
						}
						(*call.SdkParam)["DatabasePrivileges"] = dbPrivileges
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}

			callbacks = append(callbacks, callback)

		}
		if removePrivs.Len() != 0 {
			callback := volc.Callback{
				Call: volc.SdkCall{
					Action:      "RevokeDatabasePrivilege",
					ConvertMode: volc.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
						(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
						(*call.SdkParam)["DBName"] = d.Get("db_name")
						accountNames := make([]string, 0)
						for _, item := range addPrivis.List() {
							m, ok := item.(map[string]interface{})
							if !ok {
								continue
							}
							account := m["account_name"].(string)
							// 过滤掉有Grant操作的account,Grant权限方式为覆盖，先取消原有权限，再赋新权限，此处无需再取消一次。
							if addPrivis != nil && addPrivis.Len() != 0 && hasAccountNameInSet(account, addPrivis) {
								continue
							}
							accountNames = append(accountNames, account)
						}
						(*call.SdkParam)["AccountNames"] = accountNames
						if len(accountNames) == 0 {
							return false, nil
						}
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}
			callbacks = append(callbacks, callback)
		}
	}

	return callbacks
}

func (s *VolcengineRdsMysqlDatabaseService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteDatabase",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				databaseId := d.Id()
				ids := strings.Split(databaseId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid rds database id")
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["DBName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//删除Database
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
							return resource.NonRetryableError(fmt.Errorf("error on reading RDS database on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRdsMysqlDatabaseService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType: volc.ContentTypeJson,
		RequestConverts: map[string]volc.RequestConvert{
			"db_name": {
				TargetField: "DBName",
			},
		},
		NameField:    "DBName",
		CollectField: "databases",
		ResponseConverts: map[string]volc.ResponseConvert{
			"DBName": {
				TargetField: "db_name",
			},
		},
	}
}

func (s *VolcengineRdsMysqlDatabaseService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2022-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}

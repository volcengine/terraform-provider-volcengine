package vedb_mysql_account

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vedb_mysql/vedb_mysql_database"
)

type VolcengineVedbMysqlAccountService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVedbMysqlAccountService(c *ve.SdkClient) *VolcengineVedbMysqlAccountService {
	return &VolcengineVedbMysqlAccountService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVedbMysqlAccountService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVedbMysqlAccountService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeDBAccounts"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = universalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Accounts", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Accounts is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVedbMysqlAccountService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		account map[string]interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid veDB mysql account id")
	}

	instanceId := ids[0]
	accountName := ids[1]

	req := map[string]interface{}{
		"InstanceId":  instanceId,
		"AccountName": accountName,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}

	for _, r := range results {
		account, ok = r.(map[string]interface{})
		if !ok {
			return data, errors.New("Value is not map ")
		}
		if accountName == account["AccountName"].(string) {
			data = account
			break
		}
	}

	if len(data) == 0 {
		return data, fmt.Errorf("veDB account %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineVedbMysqlAccountService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("vedb_mysql_account status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVedbMysqlAccountService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBAccount",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"account_privileges": {
					TargetField: "AccountPrivileges",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"db_name": {
							TargetField: "DBName",
						},
					},
				},
			},
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				privileges := d.Get("account_privileges")
				if privileges == nil || privileges.(*schema.Set).Len() == 0 {
					return true, nil
				}
				for _, privilege := range privileges.(*schema.Set).List() {
					privilegeMap, ok := privilege.(map[string]interface{})
					if !ok {
						return false, fmt.Errorf("account_privilege is not map")
					}
					id := fmt.Sprintf("%s:%s", d.Get("instance_id"), privilegeMap["db_name"])
					_, err := vedb_mysql_database.NewVedbMysqlDatabaseService(client).ReadResource(d, id)
					if err != nil {
						return false, err
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("instance_id"), d.Get("account_name"))
				d.SetId(id)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVedbMysqlAccountService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"DBName": {
				TargetField: "db_name",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVedbMysqlAccountService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if resourceData.HasChange("account_password") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ResetDBAccount",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam: &map[string]interface{}{
					"InstanceId":      resourceData.Get("instance_id"),
					"AccountName":     resourceData.Get("account_name"),
					"AccountPassword": resourceData.Get("account_password"),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)
	}
	if resourceData.HasChange("account_privileges") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "GrantDBAccountPrivilege",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"account_privileges": {
						TargetField: "AccountPrivileges",
						ForceGet:    true,
						NextLevelConvert: map[string]ve.RequestConvert{
							"db_name": {
								TargetField: "DBName",
							},
						},
						ConvertType: ve.ConvertJsonObjectArray,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					(*call.SdkParam)["AccountName"] = d.Get("account_name")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)

		add, remove, _, _ := ve.GetSetDifference("account_privileges", resourceData, veDBMysqlAccountPrivilegeHash, false)
		if remove != nil && remove.Len() > 0 {
			removeCallback := ve.Callback{
				Call: ve.SdkCall{
					Action:      "RevokeDBAccountPrivilege",
					ConvertMode: ve.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
						(*call.SdkParam)["AccountName"] = d.Get("account_name")
						dbNames := make([]string, 0)
						for _, item := range remove.List() {
							m, ok := item.(map[string]interface{})
							if !ok {
								continue
							}
							removeDbName := m["db_name"].(string)
							// 过滤掉有Grant操作的db_name,Grant权限方式为覆盖，先取消原有权限，再赋新权限，此处无需再取消一次。
							if add != nil && add.Len() > 0 && hasDbNameInSet(removeDbName, add) {
								continue
							}
							dbNames = append(dbNames, removeDbName)
						}
						if len(dbNames) == 0 {
							return false, nil
						}
						(*call.SdkParam)["DBNames"] = strings.Join(dbNames, ",")
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}
			callbacks = append(callbacks, removeCallback)
		}
	}
	return callbacks
}

func hasDbNameInSet(dbName string, set *schema.Set) bool {
	for _, item := range set.List() {
		if m, ok := item.(map[string]interface{}); ok {
			if v, ok := m["db_name"]; ok && v.(string) == dbName {
				if detail, ok := m["account_privilege_detail"].(string); ok {
					if len(detail) == 0 {
						return false
					}
				} else {
					return false
				}
				return true
			}
		}
	}
	return false
}

func (s *VolcengineVedbMysqlAccountService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBAccount",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				rdsAccountId := d.Id()
				ids := strings.Split(rdsAccountId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid veDB mysql account id")
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["AccountName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//删除RdsMysqlAccount
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading veDB mysql account on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{callback}
}

func (s *VolcengineVedbMysqlAccountService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "AccountName",
		CollectField: "accounts",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"DBName": {
				TargetField: "db_name",
			},
		},
	}
}

func (s *VolcengineVedbMysqlAccountService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vedbm",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

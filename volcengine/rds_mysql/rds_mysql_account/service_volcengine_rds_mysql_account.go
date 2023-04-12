package rds_mysql_account

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	database "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_database"
)

type VolcengineRdsMysqlAccountService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func NewRdsMysqlAccountService(c *volc.SdkClient) *VolcengineRdsMysqlAccountService {
	return &VolcengineRdsMysqlAccountService{
		Client:     c,
		Dispatcher: &volc.Dispatcher{},
	}
}

func (s *VolcengineRdsMysqlAccountService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlAccountService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	return volc.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 0, func(condition map[string]interface{}) (data []interface{}, err error) {
		var (
			resp    *map[string]interface{}
			results interface{}
			ok      bool
		)
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
		results, err = volc.ObtainSdkValue("Result.Accounts", *resp)
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

func (s *VolcengineRdsMysqlAccountService) ReadResource(resourceData *schema.ResourceData, RdsMysqlAccountId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		account = make(map[string]interface{})
		ok      bool
	)
	if RdsMysqlAccountId == "" {
		RdsMysqlAccountId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(RdsMysqlAccountId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid rds mysql account id")
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
		return data, fmt.Errorf("RDS account %s not exist ", RdsMysqlAccountId)
	}

	return data, err
}

func (s *VolcengineRdsMysqlAccountService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineRdsMysqlAccountService) WithResourceResponseHandlers(rdsAccount map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return rdsAccount, map[string]volc.ResponseConvert{
			"DBName": {
				TargetField: "db_name",
			},
		}, nil
	}
	return []volc.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlAccountService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateDBAccount",
			ConvertMode: volc.RequestConvertAll,
			Convert: map[string]volc.RequestConvert{
				"account_privileges": {
					TargetField: "AccountPrivileges",
					ConvertType: volc.ConvertJsonObjectArray,
					NextLevelConvert: map[string]volc.RequestConvert{
						"db_name": {
							TargetField: "DBName",
						},
					},
				},
			},
			ContentType: volc.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
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
					_, err := database.NewRdsMysqlDatabaseService(client).ReadResource(d, id)
					if err != nil {
						return false, err
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建RdsMysqlAccount
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("instance_id"), d.Get("account_name"))
				d.SetId(id)
				return nil
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAccountService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callbacks := []volc.Callback{}
	if resourceData.HasChange("account_password") {
		callback := volc.Callback{
			Call: volc.SdkCall{
				Action:      "ResetDBAccount",
				ConvertMode: volc.RequestConvertIgnore,
				SdkParam: &map[string]interface{}{
					"InstanceId":      resourceData.Get("instance_id"),
					"AccountName":     resourceData.Get("account_name"),
					"AccountPassword": resourceData.Get("account_password"),
				},
				ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)
	}
	if resourceData.HasChange("account_privileges") {
		callback := volc.Callback{
			Call: volc.SdkCall{
				Action:      "GrantDBAccountPrivilege",
				ConvertMode: volc.RequestConvertInConvert,
				ContentType: volc.ContentTypeJson,
				Convert: map[string]volc.RequestConvert{
					"account_privileges": {
						TargetField: "AccountPrivileges",
						ForceGet:    true,
						NextLevelConvert: map[string]volc.RequestConvert{
							"db_name": {
								TargetField: "DBName",
							},
						},
						ConvertType: volc.ConvertJsonObjectArray,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					(*call.SdkParam)["AccountName"] = d.Get("account_name")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)

		add, remove, _, _ := volc.GetSetDifference("account_privileges", resourceData, RdsMysqlAccountPrivilegeHash, false)
		if remove != nil && remove.Len() > 0 {
			removeCallback := volc.Callback{
				Call: volc.SdkCall{
					Action:      "RevokeDBAccountPrivilege",
					ConvertMode: volc.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
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
					ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
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

func (s *VolcengineRdsMysqlAccountService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteDBAccount",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				rdsAccountId := d.Id()
				ids := strings.Split(rdsAccountId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid rds account id")
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["AccountName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//删除RdsMysqlAccount
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
							return resource.NonRetryableError(fmt.Errorf("error on reading RDS account on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRdsMysqlAccountService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "AccountName",
		CollectField: "accounts",
		ResponseConverts: map[string]volc.ResponseConvert{
			"DBName": {
				TargetField: "db_name",
			},
		},
	}
}

func (s *VolcengineRdsMysqlAccountService) ReadResourceId(id string) string {
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

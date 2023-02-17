package rds_mysql_account

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
	list, err := volc.WithPageOffsetQuery(m, "PageSize", "PageNumber", 20, 0, func(condition map[string]interface{}) (data []interface{}, err error) {
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
	if err != nil {
		return list, err
	}

	accountName := m["AccountName"]
	res := make([]interface{}, 0)
	for _, a := range list {
		account, ok := a.(map[string]interface{})
		if !ok {
			continue
		}
		// accountName是模糊搜索，需要过滤一下
		if accountName != "" && accountName != account["AccountName"].(string) {
			continue
		}
		// 拼接id
		account["Id"] = fmt.Sprintf("%s:%s", m["InstanceId"], account["AccountName"])
		res = append(res, account)
	}
	return res, nil
}

func (s *VolcengineRdsMysqlAccountService) ReadResource(resourceData *schema.ResourceData, RdsMysqlAccountId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
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
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
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
				},
			},
			ContentType: volc.ContentTypeJson,
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

func hasDbNameInSet(dbName string, set *schema.Set) bool {
	for _, item := range set.List() {
		if m, ok := item.(map[string]interface{}); ok {
			if v, ok := m["db_name"]; ok && v.(string) == dbName {
				return true
			}
		}
	}
	return false
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
		addPrivis, removePrivs, _, _ := common.GetSetDifference("account_privileges", resourceData, RdsMysqlAccountPrivilegeHash, false)
		logger.DebugInfo("account_privileges %v", resourceData.Get("account_privileges"))
		logger.DebugInfo("addPrivis %v len %d", addPrivis, addPrivis.Len())
		logger.DebugInfo("removePrivs %v len %d", removePrivs, addPrivis.Len())

		if addPrivis != nil && addPrivis.Len() != 0 {
			callback := volc.Callback{
				Call: volc.SdkCall{
					Action:      "GrantDBAccountPrivilege",
					ConvertMode: volc.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
						(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
						(*call.SdkParam)["AccountName"] = d.Get("account_name")
						accountPrivileges := make([]map[string]interface{}, 0)
						for _, item := range addPrivis.List() {
							m, _ := item.(map[string]interface{})
							privi := make(map[string]interface{})
							if v, ok := m["db_name"]; ok {
								privi["DBName"] = v
							}
							if v, ok := m["account_privilege"]; ok {
								privi["AccountPrivilege"] = v
							}
							if v, ok := m["account_privilege_detail"]; ok {
								privi["AccountPrivilegeDetail"] = v
							}
							accountPrivileges = append(accountPrivileges, privi)
						}
						(*call.SdkParam)["AccountPrivileges"] = accountPrivileges
						logger.DebugInfo("accountPrivileges %v", accountPrivileges)
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
		if removePrivs != nil && removePrivs.Len() != 0 {
			callback := volc.Callback{
				Call: volc.SdkCall{
					Action:      "RevokeDBAccountPrivilege",
					ConvertMode: volc.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
						(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
						(*call.SdkParam)["AccountName"] = d.Get("account_name")
						dbNames := make([]string, 0)
						for _, item := range addPrivis.List() {
							m, ok := item.(map[string]interface{})
							if !ok {
								continue
							}
							dbName := m["db_name"].(string)
							// 过滤掉有Grant操作的db_name,Grant权限方式为覆盖，先取消原有权限，再赋新权限，此处无需再取消一次。
							if addPrivis != nil && addPrivis.Len() != 0 && hasDbNameInSet(dbName, addPrivis) {
								continue
							}
							dbNames = append(dbNames, dbName)
						}
						if len(dbNames) == 0 {
							return false, nil
						}
						(*call.SdkParam)["DBNames"] = dbNames
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

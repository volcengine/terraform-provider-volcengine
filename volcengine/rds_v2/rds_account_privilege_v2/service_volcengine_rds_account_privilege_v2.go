package rds_account_privilege_v2

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

type VolcengineRdsAccountPrivilegeService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func NewRdsAccountPrivilegeService(c *volc.SdkClient) *VolcengineRdsAccountPrivilegeService {
	return &VolcengineRdsAccountPrivilegeService{
		Client:     c,
		Dispatcher: &volc.Dispatcher{},
	}
}

func (s *VolcengineRdsAccountPrivilegeService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsAccountPrivilegeService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	list, err := volc.WithPageOffsetQuery(m, "Limit", "Offset", 20, 0, func(condition map[string]interface{}) (data []interface{}, err error) {
		var (
			resp    *map[string]interface{}
			results interface{}
			ok      bool
		)
		universalClient := s.Client.UniversalClient
		action := "ListAccounts"
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

	// 拼接id
	res := make([]interface{}, 0)
	for _, a := range list {
		account, ok := a.(map[string]interface{})
		if !ok {
			continue
		}
		account["Id"] = fmt.Sprintf("%s:%s", m["InstanceId"], account["AccountName"])
		res = append(res, account)
	}
	return res, nil
}

func (s *VolcengineRdsAccountPrivilegeService) ReadResource(resourceData *schema.ResourceData, RdsAccountPrivilegeId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if RdsAccountPrivilegeId == "" {
		RdsAccountPrivilegeId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(RdsAccountPrivilegeId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid rds account id")
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
		return data, fmt.Errorf("RDS account %s not exist ", RdsAccountPrivilegeId)
	}

	return data, err
}

func (s *VolcengineRdsAccountPrivilegeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineRdsAccountPrivilegeService) WithResourceResponseHandlers(rdsAccountPrivilege map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return rdsAccountPrivilege, map[string]volc.ResponseConvert{
			"DBPrivileges": {
				TargetField: "db_privileges",
				Convert: func(v interface{}) interface{} {
					dbPrivileges, ok := v.([]interface{})
					if !ok {
						return v
					}

					for i := range dbPrivileges {
						dbPrivilege, ok := dbPrivileges[i].(map[string]interface{})
						if !ok {
							continue
						}
						dbPrivilege["DbName"] = dbPrivilege["DBName"]
					}

					return dbPrivileges
				},
			},
		}, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineRdsAccountPrivilegeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	grantAccountPrivilegeCallbacks := make([]volc.Callback, 0)

	dbPrivileges := resourceData.Get("db_privileges").(*schema.Set).List()
	for _, dbPrivilege := range dbPrivileges {
		localPrivilege := dbPrivilege
		callback := volc.Callback{
			Call: volc.SdkCall{
				Action:      "GrantAccountPrivilege",
				ConvertMode: volc.RequestConvertIgnore,
				ContentType: volc.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					(*call.SdkParam)["AccountName"] = d.Get("account_name")
					if privilege, ok := localPrivilege.(map[string]interface{}); ok {
						(*call.SdkParam)["DBName"] = privilege["db_name"]
						(*call.SdkParam)["AccountPrivilege"] = privilege["account_privilege"]
						(*call.SdkParam)["AccountPrivilegeStr"] = privilege["account_privilege_str"]
					} else {
						return false, errors.New("db_privileges' element is not map")
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					//创建RdsAccountPrivilege
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
					id := fmt.Sprintf("%s:%s", d.Get("instance_id"), d.Get("account_name"))
					d.SetId(id)
					return nil
				},
			},
		}
		grantAccountPrivilegeCallbacks = append(grantAccountPrivilegeCallbacks, callback)
	}

	return grantAccountPrivilegeCallbacks
}

func (s *VolcengineRdsAccountPrivilegeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	if !resourceData.HasChange("db_privileges") {
		return []volc.Callback{}
	}

	callbacks := make([]volc.Callback, 0)

	// 1. compute add and remove
	add, remove, _, _ := volc.GetSetDifference("db_privileges", resourceData, RdsAccountPrivilegeHash, false)
	logger.Info("[ModifyRdsAccountPrivilege] add: %+v, remove: %+v", add, remove)

	// 2. remove callback
	if len(remove.List()) > 0 {
		removeCallback := volc.Callback{
			Call: volc.SdkCall{
				Action:      "RevokeAccountPrivilege",
				ConvertMode: volc.RequestConvertIgnore,
				ContentType: volc.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					(*call.SdkParam)["AccountName"] = d.Get("account_name")
					removeDBNames := make([]string, 0)
					for _, removePrivilege := range remove.List() {
						removeDBNames = append(removeDBNames, removePrivilege.(map[string]interface{})["db_name"].(string))
					}
					(*call.SdkParam)["DBNames"] = strings.Join(removeDBNames, ",")
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

	// 3. add callbacks
	if len(add.List()) > 0 {
		grantAccountPrivilegeCallbacks := make([]volc.Callback, 0)

		for _, addPrivilege := range add.List() {
			localPrivilege := addPrivilege
			callback := volc.Callback{
				Call: volc.SdkCall{
					Action:      "GrantAccountPrivilege",
					ConvertMode: volc.RequestConvertIgnore,
					ContentType: volc.ContentTypeJson,
					BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
						(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
						(*call.SdkParam)["AccountName"] = d.Get("account_name")
						if privilege, ok := localPrivilege.(map[string]interface{}); ok {
							(*call.SdkParam)["DBName"] = privilege["db_name"]
							(*call.SdkParam)["AccountPrivilege"] = privilege["account_privilege"]
							(*call.SdkParam)["AccountPrivilegeStr"] = privilege["account_privilege_str"]
						} else {
							return false, errors.New("db_privileges' element is not map")
						}
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						//创建RdsAccountPrivilege
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}
			grantAccountPrivilegeCallbacks = append(grantAccountPrivilegeCallbacks, callback)
		}
		callbacks = append(callbacks, grantAccountPrivilegeCallbacks...)
	}

	return callbacks
}

func (s *VolcengineRdsAccountPrivilegeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	dbPrivileges := resourceData.Get("db_privileges").(*schema.Set).List()
	if len(dbPrivileges) == 0 {
		return []volc.Callback{}
	}

	dbNames := make([]string, 0)
	for _, dbPrivilege := range dbPrivileges {
		if privilege, ok := dbPrivilege.(map[string]interface{}); ok {
			dbNames = append(dbNames, privilege["db_name"].(string))
		}
	}

	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "RevokeAccountPrivilege",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
				(*call.SdkParam)["AccountName"] = d.Get("account_name")
				(*call.SdkParam)["DBNames"] = strings.Join(dbNames, ",")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//删除RdsAccountPrivilege
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
							return resource.NonRetryableError(fmt.Errorf("error on reading RDS account privilege on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRdsAccountPrivilegeService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{}
}

func (s *VolcengineRdsAccountPrivilegeService) ReadResourceId(id string) string {
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

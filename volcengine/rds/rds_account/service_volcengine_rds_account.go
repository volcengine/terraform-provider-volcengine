package rds_account

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

type VolcengineRdsAccountService struct {
	Client *volc.SdkClient
}

func NewRdsAccountService(c *volc.SdkClient) *VolcengineRdsAccountService {
	return &VolcengineRdsAccountService{
		Client: c,
	}
}

func (s *VolcengineRdsAccountService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsAccountService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
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

func (s *VolcengineRdsAccountService) ReadResource(resourceData *schema.ResourceData, RdsAccountId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if RdsAccountId == "" {
		RdsAccountId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(RdsAccountId, ":")
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
		return data, fmt.Errorf("RDS account %s not exist ", RdsAccountId)
	}

	return data, err
}

func (s *VolcengineRdsAccountService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineRdsAccountService) WithResourceResponseHandlers(rdsAccount map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return rdsAccount, nil, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineRdsAccountService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateAccount",
			ConvertMode: volc.RequestConvertAll,
			ContentType: volc.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建RdsAccount
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

func (s *VolcengineRdsAccountService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	// TODO support reset password, grant and revoke privileges
	return []volc.Callback{}
}

func (s *VolcengineRdsAccountService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteAccount",
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
				//删除RdsAccount
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

func (s *VolcengineRdsAccountService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "AccountName",
		CollectField: "rds_accounts",
		ResponseConverts: map[string]volc.ResponseConvert{
			"DBPrivileges": {
				TargetField: "db_privileges",
				Convert: func(i interface{}) interface{} {
					dbPrivileges, ok := i.([]interface{})
					if !ok {
						return i
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
		},
	}
}

func (s *VolcengineRdsAccountService) ReadResourceId(id string) string {
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

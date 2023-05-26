package account

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

type VolcengineRedisAccountService struct {
	Client *volc.SdkClient
}

func NewAccountService(c *volc.SdkClient) *VolcengineRedisAccountService {
	return &VolcengineRedisAccountService{
		Client: c,
	}
}

func (s *VolcengineRedisAccountService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRedisAccountService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	universalClient := s.Client.UniversalClient
	action := "ListDBAccount"
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
}

func (s *VolcengineRedisAccountService) ReadResource(resourceData *schema.ResourceData, RedisAccountId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if RedisAccountId == "" {
		RedisAccountId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(RedisAccountId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid redis account id")
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
		return data, fmt.Errorf("Redis account %s not exist ", RedisAccountId)
	}

	return data, err
}

func (s *VolcengineRedisAccountService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineRedisAccountService) WithResourceResponseHandlers(account map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return account, nil, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineRedisAccountService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateDBAccount",
			ConvertMode: volc.RequestConvertAll,
			ContentType: volc.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
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

func (s *VolcengineRedisAccountService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "ModifyDBAccount",
			ConvertMode: volc.RequestConvertAll,
			ContentType: volc.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				accountId := d.Id()
				ids := strings.Split(accountId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid redis account id")
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["AccountName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRedisAccountService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteDBAccount",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				redisAccountId := d.Id()
				ids := strings.Split(redisAccountId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid redis account id")
				}

				if ids[1] == "default" {
					return false, fmt.Errorf("can not delete `default` account of redis instance")
				}

				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["AccountName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall, baseErr error) error {
				// 不能删除 default 账号
				if strings.Contains(baseErr.Error(), "can not delete `default` account of redis instance") {
					msg := fmt.Sprintf("error: %s. msg: %s",
						baseErr.Error(),
						"If you want to remove it form terraform state, "+
							"please use `terraform state rm volcengine_redis_account.resource_name` command ")
					return fmt.Errorf(msg)
				}
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if volc.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading redis account on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRedisAccountService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "AccountName",
		CollectField: "accounts",
	}
}

func (s *VolcengineRedisAccountService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "redis",
		Version:     "2020-12-07",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}

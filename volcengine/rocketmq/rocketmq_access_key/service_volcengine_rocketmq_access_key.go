package rocketmq_access_key

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
)

type VolcengineRocketmqAccessKeyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRocketmqAccessKeyService(c *ve.SdkClient) *VolcengineRocketmqAccessKeyService {
	return &VolcengineRocketmqAccessKeyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRocketmqAccessKeyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRocketmqAccessKeyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	data, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAccessKeys"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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

		results, err = ve.ObtainSdkValue("Result.AccessKeysInfo", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AccessKeysInfo is not Slice")
		}
		return data, err
	})
	if err != nil {
		return data, err
	}

	for _, v := range data {
		mqKey, ok := v.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("The value of Result.AccessKeysInfo is not map ")
		}
		mqKey["InstanceId"] = m["InstanceId"]
		req := map[string]interface{}{
			"InstanceId": mqKey["InstanceId"],
			"AccessKey":  mqKey["AccessKey"],
		}

		// query access key detail
		accessAction := "DescribeAccessKeyDetail"
		logger.Debug(logger.ReqFormat, accessAction, req)
		accessResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(accessAction), &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, accessAction, *accessResp)

		permissionInfo, err := ve.ObtainSdkValue("Result.TopicPermissions", *accessResp)
		if err != nil {
			return data, err
		}
		if permissionInfo == nil {
			permissionInfo = []interface{}{}
		}
		permissionArr, ok := permissionInfo.([]interface{})
		if !ok {
			return data, fmt.Errorf("DescribeAccessKeyDetail Result.TopicPermissions is not slice")
		}
		mqKey["TopicPermissions"] = permissionArr

		// query sk
		skAction := "DescribeSecretKey"
		logger.Debug(logger.ReqFormat, skAction, req)
		groupResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(skAction), &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, skAction, *groupResp)

		sk, err := ve.ObtainSdkValue("Result.SecretKey", *groupResp)
		if err != nil {
			return data, err
		}
		mqKey["SecretKey"] = sk
	}

	return data, err
}

func (s *VolcengineRocketmqAccessKeyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("Invalid rocketmq key id: %v ", id)
	}

	req := map[string]interface{}{
		"InstanceId": ids[0],
		"AccessKey":  ids[1],
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
		return data, fmt.Errorf("rocketmq_access_key %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRocketmqAccessKeyService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineRocketmqAccessKeyService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRocketmqAccessKeyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAccessKey",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// query access_key by description
				instanceId := d.Get("instance_id").(string)
				description := d.Get("description").(string)
				req := map[string]interface{}{
					"InstanceId": instanceId,
				}
				results, err := s.ReadResources(req)
				if err != nil {
					return fmt.Errorf("CreateAccessKey AfterCall Error: %v", err)
				}

				var accessKey string
				for _, v := range results {
					result, ok := v.(map[string]interface{})
					if !ok {
						return fmt.Errorf("CreateAccessKey AfterCall Error: The value of query result is not map")
					}
					if result["Description"].(string) == description {
						accessKey = result["AccessKey"].(string)
						break
					}
				}
				if accessKey == "" {
					return fmt.Errorf("CreateAccessKey AfterCall Error: Cannot query AccessKey by Description")
				}

				d.SetId(instanceId + ":" + accessKey)
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRocketmqAccessKeyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyAccessKeyAllAuthority",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"all_authority": {
					TargetField: "AllAuthority",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("Invalid rocketmq key id: %v ", d.Id())
					}
					(*call.SdkParam)["InstanceId"] = ids[0]
					(*call.SdkParam)["AccessKey"] = ids[1]
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRocketmqAccessKeyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAccessKey",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("Invalid rocketmq key id: %v ", d.Id())
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["AccessKey"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRocketmqAccessKeyService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "access_keys",
		ContentType:  ve.ContentTypeJson,
	}
}

func (s *VolcengineRocketmqAccessKeyService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "RocketMQ",
		Version:     "2023-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

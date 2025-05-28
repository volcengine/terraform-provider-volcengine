package kms_key_schedule_deletion

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsKeyScheduleDeletionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsKeyScheduleDeletionService(c *ve.SdkClient) *VolcengineKmsKeyScheduleDeletionService {
	return &VolcengineKmsKeyScheduleDeletionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsKeyScheduleDeletionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsKeyScheduleDeletionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineKmsKeyScheduleDeletionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		ok          bool
		resp        *map[string]interface{}
		keyId       string
		keyName     string
		keyringName string
		req         map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	parts := strings.Split(id, ":")
	switch len(parts) {
	case 1:
		keyId = parts[0]
	case 2:
		keyName = parts[0]
		keyringName = parts[1]
	default:
		return data, fmt.Errorf("format of kms key schedule deletion id is invalid,%s", id)
	}
	if keyId != "" {
		req = map[string]interface{}{
			"KeyID": keyId,
		}
	} else {
		req = map[string]interface{}{
			"KeyringName": keyringName,
			"KeyName":     keyName,
		}
	}
	action := "DescribeKey"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	keyring, err := ve.ObtainSdkValue("Result.Key", *resp)
	if err != nil {
		return data, err
	}
	data, ok = keyring.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("kms key schedule deletion %s not exist", id)
	}

	data["KeyID"] = data["ID"]

	return data, err
}

func (s *VolcengineKmsKeyScheduleDeletionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsKeyScheduleDeletionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ScheduleKeyDeletion",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"key_id": {
					TargetField: "KeyID",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				if resourceData.Get("key_id") != "" {
					d.SetId(resourceData.Get("key_id").(string))
				} else {
					// set id to key_name:keyring_name
					d.SetId(fmt.Sprintf("%s:%s", resourceData.Get("key_name").(string), resourceData.Get("keyring_name").(string)))
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsKeyScheduleDeletionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"KeyID": {
				TargetField: "key_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsKeyScheduleDeletionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsKeyScheduleDeletionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CancelKeyDeletion",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"key_id": {
					TargetField: "KeyID",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var (
					keyId       string
					keyName     string
					keyringName string
				)
				parts := strings.Split(d.Id(), ":")
				switch len(parts) {
				case 1:
					keyId = parts[0]
				case 2:
					keyName = parts[0]
					keyringName = parts[1]
				default:
					return false, fmt.Errorf("format of kms key schedule deletion id is invalid,%s", d.Id())
				}
				if keyId != "" {
					(*call.SdkParam)["KeyID"] = keyId
				} else {
					(*call.SdkParam)["KeyringName"] = keyringName
					(*call.SdkParam)["KeyName"] = keyName
				}
				return true, nil
			},

			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return s.checkResourceUtilRemoved(d, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKmsKeyScheduleDeletionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineKmsKeyScheduleDeletionService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func (s *VolcengineKmsKeyScheduleDeletionService) checkResourceUtilRemoved(d *schema.ResourceData, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		keyStatus, _ := s.ReadResource(d, d.Id())
		// 能查询成功代表还在删除中，重试
		if keyStatus["KeyState"] == "PendingDelete" {
			return resource.RetryableError(fmt.Errorf("resource still in removing status "))
		} else {
			if keyStatus["KeyState"] == "Enable" {
				return nil
			} else {
				return resource.NonRetryableError(fmt.Errorf("kms key schedule deletion status is not Enable "))
			}
		}
	})
}

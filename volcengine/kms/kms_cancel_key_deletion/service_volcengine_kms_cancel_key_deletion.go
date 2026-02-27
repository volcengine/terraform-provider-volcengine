package kms_cancel_key_deletion

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsCancelKeyDeletionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsCancelKeyDeletionService(c *ve.SdkClient) *VolcengineKmsCancelKeyDeletionService {
	return &VolcengineKmsCancelKeyDeletionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsCancelKeyDeletionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsCancelKeyDeletionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineKmsCancelKeyDeletionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("format of kms cancel key deletion id is invalid,%s", id)
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

	key, err := ve.ObtainSdkValue("Result.Key", *resp)
	if err != nil {
		return data, err
	}
	data, ok = key.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("kms cancel key deletion %s not exist", id)
	}

	data["KeyID"] = data["ID"]

	return data, err
}

func (s *VolcengineKmsCancelKeyDeletionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsCancelKeyDeletionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CancelKeyDeletion",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"key_id": {
					TargetField: "KeyID",
				},
				"key_name": {
					TargetField: "KeyName",
				},
				"keyring_name": {
					TargetField: "KeyringName",
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
					d.SetId(fmt.Sprintf("%s:%s", resourceData.Get("key_name").(string), resourceData.Get("keyring_name").(string)))
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsCancelKeyDeletionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"KeyID": {
				TargetField: "key_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsCancelKeyDeletionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsCancelKeyDeletionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsCancelKeyDeletionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineKmsCancelKeyDeletionService) ReadResourceId(id string) string {
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

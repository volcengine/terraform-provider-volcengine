package kms_key_primary_region

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsPrimaryRegionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsKeyPrimaryRegionService(c *ve.SdkClient) *VolcengineKmsPrimaryRegionService {
	return &VolcengineKmsPrimaryRegionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsPrimaryRegionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsPrimaryRegionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineKmsPrimaryRegionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp        *map[string]interface{}
		ok          bool
		keyId       string
		keyName     string
		keyringName string
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
		return data, fmt.Errorf("format of kms primary region id is invalid,%s", id)
	}

	req := make(map[string]interface{})
	if keyId != "" {
		req["KeyID"] = keyId
	} else {
		req["KeyringName"] = keyringName
		req["KeyName"] = keyName
	}

	action := "DescribeKey"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	keyRaw, err := ve.ObtainSdkValue("Result.Key", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = keyRaw.(map[string]interface{}); !ok {
		return data, fmt.Errorf("Result.Key is not Map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("kms primary region %s not exist", id)
	}

	primaryRegion, _ := ve.ObtainSdkValue("MultiRegionConfiguration.PrimaryKey.Region", data)
	if primaryRegion != nil {
		data["PrimaryRegion"] = primaryRegion
	}
	data["KeyID"] = data["ID"]

	return data, nil
}

func (s *VolcengineKmsPrimaryRegionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsPrimaryRegionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdatePrimaryRegion",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"key_name": {
					TargetField: "KeyName",
				},
				"keyring_name": {
					TargetField: "KeyringName",
				},
				"key_id": {
					TargetField: "KeyID",
				},
				"primary_region": {
					TargetField: "PrimaryRegion",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalPostInfo(call.Action), call.SdkParam)
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

func (VolcengineKmsPrimaryRegionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"KeyID": {
				TargetField: "key_id",
			},
			"PrimaryRegion": {
				TargetField: "primary_region",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsPrimaryRegionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsPrimaryRegionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsPrimaryRegionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineKmsPrimaryRegionService) ReadResourceId(id string) string {
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

func getUniversalPostInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

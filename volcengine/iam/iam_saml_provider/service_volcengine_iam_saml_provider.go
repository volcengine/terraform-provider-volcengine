package iam_saml_provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamSamlProviderService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamSamlProviderService(c *ve.SdkClient) *VolcengineIamSamlProviderService {
	return &VolcengineIamSamlProviderService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineIamSamlProviderService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamSamlProviderService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListSAMLProviders"

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
		results, err = ve.ObtainSdkValue("Result.SAMLProviders", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.SAMLProviders is not Slice")
		}
		for index, ele := range data {
			provider := ele.(map[string]interface{})
			query := map[string]interface{}{
				// 跟API文档不对应，文档写的返回字段为SAMLProviderName，实际为ProviderName
				"SAMLProviderName": provider["ProviderName"],
			}
			action = "GetSAMLProvider"
			logger.Debug(logger.ReqFormat, action, query)
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &query)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, query, *resp)
			document, err := ve.ObtainSdkValue("Result.EncodedSAMLMetadataDocument", *resp)
			if err != nil {
				return data, err
			}
			data[index].(map[string]interface{})["EncodedSAMLMetadataDocument"] = document
		}
		return data, err
	})
}

func (s *VolcengineIamSamlProviderService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		temp    map[string]interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if temp, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else if temp["ProviderName"].(string) == id {
			data = temp
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("iam_saml_provider %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineIamSamlProviderService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineIamSamlProviderService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateSAMLProvider",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"encoded_saml_metadata_document": {
					TargetField: "EncodedSAMLMetadataDocument",
				},
				"saml_provider_name": {
					TargetField: "SAMLProviderName",
				},
				"sso_type": {
					TargetField: "SSOType",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfoPost(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := d.Get("saml_provider_name")
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineIamSamlProviderService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamSamlProviderService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateSAMLProvider",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"encoded_saml_metadata_document": {
					TargetField: "NewEncodedSAMLMetadataDocument",
				},
				"description": {
					TargetField: "Description",
				},
				"status": {
					TargetField: "Status",
				},
				"sso_type": {
					TargetField: "SSOType",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["SAMLProviderName"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfoPost(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamSamlProviderService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteSAMLProvider",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"SAMLProviderName": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamSamlProviderService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "SAMLProviderName",
		IdField:      "Trn",
		CollectField: "providers",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ProviderName": {
				TargetField: "saml_provider_name",
			},
			"EncodedSAMLMetadataDocument": {
				TargetField: "encoded_saml_metadata_document",
			},
			"SSOType": {
				TargetField: "sso_type",
			},
		},
	}
}

func (s *VolcengineIamSamlProviderService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return getReqInfo(actionName, ve.GET)
}

func getUniversalInfoPost(actionName string) ve.UniversalInfo {
	return getReqInfo(actionName, ve.POST)
}

func getReqInfo(actionName string, httMethod ve.HttpMethod) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  httMethod,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

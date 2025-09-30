package iam_oidc_provider

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

type VolcengineIamOidcProviderService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamOidcProviderService(c *ve.SdkClient) *VolcengineIamOidcProviderService {
	return &VolcengineIamOidcProviderService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineIamOidcProviderService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamOidcProviderService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListOIDCProviders"

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
		results, err = ve.ObtainSdkValue("Result.OIDCProviders", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.OIDCProviders is not Slice")
		}
		for _, ele := range data {
			provider, ok := ele.(map[string]interface{})
			if !ok {
				continue
			}
			query := map[string]interface{}{
				"OIDCProviderName": provider["ProviderName"],
			}
			action = "GetOIDCProvider"
			logger.Debug(logger.ReqFormat, action, query)
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &query)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, query, *resp)

			issuerUrl, err := ve.ObtainSdkValue("Result.IssuerURL", *resp)
			if err != nil {
				return data, err
			}
			provider["IssuerURL"] = issuerUrl

			issuanceLimitTime, err := ve.ObtainSdkValue("Result.IssuanceLimitTime", *resp)
			if err != nil {
				return data, err
			}
			provider["IssuanceLimitTime"] = issuanceLimitTime

			clientIds, err := ve.ObtainSdkValue("Result.ClientIDs", *resp)
			if err != nil {
				return data, err
			}
			provider["ClientIDs"] = clientIds

			thumbprints, err := ve.ObtainSdkValue("Result.Thumbprints", *resp)
			if err != nil {
				return data, err
			}
			provider["Thumbprints"] = thumbprints
		}
		return data, err
	})
}

func (s *VolcengineIamOidcProviderService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("iam_oidc_provider %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineIamOidcProviderService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamOidcProviderService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"OIDCProviderName": {
				TargetField: "oidc_provider_name",
			},
			"IssuerURL": {
				TargetField: "issuer_url",
			},
			"ClientIDs": {
				TargetField: "client_ids",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamOidcProviderService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateOIDCProvider",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"oidc_provider_name": {
					TargetField: "OIDCProviderName",
				},
				"issuer_url": {
					TargetField: "IssuerURL",
				},
				"client_ids": {
					TargetField: "ClientIDs",
					ConvertType: ve.ConvertWithN,
				},
				"thumbprints": {
					TargetField: "Thumbprints",
					ConvertType: ve.ConvertWithN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.OIDCProviderName", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamOidcProviderService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateOIDCProvider",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"oidc_provider_name": {
					TargetField: "OIDCProviderName",
					ForceGet:    true,
				},
				"description": {
					TargetField: "Description",
				},
				"issuance_limit_time": {
					TargetField: "IssuanceLimitTime",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	callbacks = append(callbacks, callback)

	// client_ids
	if resourceData.HasChange("client_ids") {
		addClientIds, removeClientIds, _, _ := ve.GetSetDifference("client_ids", resourceData, schema.HashString, false)
		for _, element := range addClientIds.List() {
			callbacks = append(callbacks, s.updateOidcProviderCallback(resourceData, "AddClientIDToOIDCProvider", "ClientID", element))
		}
		for _, element := range removeClientIds.List() {
			callbacks = append(callbacks, s.updateOidcProviderCallback(resourceData, "RemoveClientIDFromOIDCProvider", "ClientID", element))
		}
	}

	// thumbprints
	if resourceData.HasChange("thumbprints") {
		addThumbprints, removeThumbprints, _, _ := ve.GetSetDifference("thumbprints", resourceData, schema.HashString, false)
		for _, element := range addThumbprints.List() {
			callbacks = append(callbacks, s.updateOidcProviderCallback(resourceData, "AddThumbprintToOIDCProvider", "Thumbprint", element))
		}
		for _, element := range removeThumbprints.List() {
			callbacks = append(callbacks, s.updateOidcProviderCallback(resourceData, "RemoveThumbprintFromOIDCProvider", "Thumbprint", element))
		}
	}

	return callbacks
}

func (s *VolcengineIamOidcProviderService) updateOidcProviderCallback(resourceData *schema.ResourceData, action, field string, element interface{}) ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["OIDCProviderName"] = d.Id()
				(*call.SdkParam)[field] = element.(string)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Id()
			},
		},
	}
}

func (s *VolcengineIamOidcProviderService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteOIDCProvider",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"OIDCProviderName": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading iam oidc provider on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{callback}
}

func (s *VolcengineIamOidcProviderService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "OIDCProviderName",
		CollectField: "oidc_providers",
		ResponseConverts: map[string]ve.ResponseConvert{
			"IssuerURL": {
				TargetField: "issuer_url",
			},
			"ClientIDs": {
				TargetField: "client_ids",
			},
		},
	}
}

func (s *VolcengineIamOidcProviderService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

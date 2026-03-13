package iam_oauth_provider

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamOAuthProviderService struct {
	Client *ve.SdkClient
}

func NewIamOAuthProviderService(c *ve.SdkClient) *VolcengineIamOAuthProviderService {
	return &VolcengineIamOAuthProviderService{
		Client: c,
	}
}

func (s *VolcengineIamOAuthProviderService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamOAuthProviderService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "CreateOAuthProvider",
				ConvertMode: ve.RequestConvertIgnore,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param, err := ve.ResourceDateToRequest(d, resource, false, s.createRequestConvert(), ve.RequestConvertInConvert, ve.ContentTypeDefault)
					if err != nil {
						return nil, err
					}
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					if v, ok := d.Get("oauth_provider_name").(string); ok && v != "" {
						d.SetId(v)
					}
					return nil
				},
			},
		},
	}
}

func (s *VolcengineIamOAuthProviderService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "UpdateOAuthProvider",
				ConvertMode: ve.RequestConvertIgnore,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param, err := ve.ResourceDateToRequest(d, resource, true, s.createRequestConvert(), ve.RequestConvertInConvert, ve.ContentTypeDefault)
					if err != nil {
						return nil, err
					}
					param["OAuthProviderName"] = d.Id()
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		},
	}
}

func (s *VolcengineIamOAuthProviderService) createRequestConvert() map[string]ve.RequestConvert {
	return map[string]ve.RequestConvert{
		"oauth_provider_name": {TargetField: "OAuthProviderName"},
		"sso_type":            {TargetField: "SSOType"},
		"status":              {TargetField: "Status"},
		"description":         {TargetField: "Description"},
		"client_id":           {TargetField: "ClientId"},
		"client_secret":       {TargetField: "ClientSecret"},
		"user_info_url":       {TargetField: "UserInfoURL"},
		"token_url":           {TargetField: "TokenURL"},
		"authorize_url":       {TargetField: "AuthorizeURL"},
		"authorize_template":  {TargetField: "AuthorizeTemplate"},
		"scope":               {TargetField: "Scope"},
		"identity_map_type":   {TargetField: "IdentityMapType"},
		"idp_identity_key":    {TargetField: "IdpIdentityKey"},
	}
}

func (s *VolcengineIamOAuthProviderService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "DeleteOAuthProvider",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam: &map[string]interface{}{
					"OAuthProviderName": data.Id(),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		},
	}
}

func (s *VolcengineIamOAuthProviderService) ReadResource(d *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	action := "GetOAuthProvider"
	if id == "" {
		id = d.Id()
	}
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{
		"OAuthProviderName": id,
	})
	if err != nil {
		return nil, err
	}
	result, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	if resMap, ok := result.(map[string]interface{}); ok {
		return resMap, nil
	}
	return nil, fmt.Errorf("result is not map[string]interface{}")
}

func (s *VolcengineIamOAuthProviderService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	action := "GetOAuthProvider"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return data, err
	}

	result, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if result == nil {
		return []interface{}{}, nil
	}

	return []interface{}{result}, nil
}

func (s *VolcengineIamOAuthProviderService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ResponseConverts: map[string]ve.ResponseConvert{
			"OAuthProviderName": {TargetField: "oauth_provider_name"},
			"ProviderId":        {TargetField: "provider_id"},
			"SSOType":           {TargetField: "sso_type"},
			"Status":            {TargetField: "status"},
			"Description":       {TargetField: "description"},
			"ClientId":          {TargetField: "client_id"},
			"ClientSecret":      {TargetField: "client_secret"},
			"UserInfoURL":       {TargetField: "user_info_url"},
			"TokenURL":          {TargetField: "token_url"},
			"AuthorizeURL":      {TargetField: "authorize_url"},
			"AuthorizeTemplate": {TargetField: "authorize_template"},
			"Scope":             {TargetField: "scope"},
			"IdentityMapType":   {TargetField: "identity_map_type"},
			"IdpIdentityKey":    {TargetField: "idp_identity_key"},
			"Trn":               {TargetField: "trn"},
			"CreateDate":        {TargetField: "create_date"},
			"UpdateDate":        {TargetField: "update_date"},
		},
		CollectField: "providers",
		RequestConverts: map[string]ve.RequestConvert{
			"oauth_provider_name": {TargetField: "OAuthProviderName"},
		},
	}
}

func (s *VolcengineIamOAuthProviderService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}
func (s *VolcengineIamOAuthProviderService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}
func (s *VolcengineIamOAuthProviderService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	method := ve.GET
	if actionName == "CreateOAuthProvider" || actionName == "UpdateOAuthProvider" {
		method = ve.POST
	}
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  method,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

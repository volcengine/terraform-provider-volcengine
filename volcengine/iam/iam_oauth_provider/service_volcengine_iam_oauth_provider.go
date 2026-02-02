package iam_oauth_provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"time"
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
				SdkParam: &map[string]interface{}{
					"OAuthProviderName": data.Get("oauth_provider_name").(string),
					"SSOType":           data.Get("sso_type").(int),
					"Status":            data.Get("status").(int),
					"Description":       data.Get("description").(string),
					"ClientId":          data.Get("client_id").(string),
					"ClientSecret":      data.Get("client_secret").(string),
					"UserInfoURL":       data.Get("user_info_url").(string),
					"TokenURL":          data.Get("token_url").(string),
					"AuthorizeURL":      data.Get("authorize_url").(string),
					"AuthorizeTemplate": data.Get("authorize_template").(string),
					"Scope":             data.Get("scope").(string),
					"IdentityMapType":   data.Get("identity_map_type").(int),
					"IdpIdentityKey":    data.Get("idp_identity_key").(string),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					d.SetId(d.Get("oauth_provider_name").(string))
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
				SdkParam: &map[string]interface{}{
					"OAuthProviderName": data.Id(),
					"Status":            data.Get("status").(int),
					"Description":       data.Get("description").(string),
					"ClientId":          data.Get("client_id").(string),
					"ClientSecret":      data.Get("client_secret").(string),
					"UserInfoURL":       data.Get("user_info_url").(string),
					"TokenURL":          data.Get("token_url").(string),
					"AuthorizeURL":      data.Get("authorize_url").(string),
					"AuthorizeTemplate": data.Get("authorize_template").(string),
					"Scope":             data.Get("scope").(string),
					"IdentityMapType":   data.Get("identity_map_type").(int),
					"IdpIdentityKey":    data.Get("idp_identity_key").(string),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		},
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
	return result.(map[string]interface{}), nil
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

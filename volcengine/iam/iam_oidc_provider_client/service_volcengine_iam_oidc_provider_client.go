package iam_oidc_provider_client

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

type VolcengineIamOidcProviderClientService struct {
	Client *ve.SdkClient
}

func NewIamOidcProviderClientService(c *ve.SdkClient) *VolcengineIamOidcProviderClientService {
	return &VolcengineIamOidcProviderClientService{
		Client: c,
	}
}

func (s *VolcengineIamOidcProviderClientService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamOidcProviderClientService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "AddClientIDToOIDCProvider",
				ConvertMode: ve.RequestConvertIgnore,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param, err := ve.ResourceDateToRequest(d, resource, false, s.createRequestConvert(), ve.RequestConvertInConvert, ve.ContentTypeDefault)
					if err != nil {
						return nil, err
					}
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					pName, _ := d.Get("oidc_provider_name").(string)
					cId, _ := d.Get("client_id").(string)
					if pName != "" && cId != "" {
						d.SetId(fmt.Sprintf("%s:%s", pName, cId))
					}
					return nil
				},
			},
		},
	}
}

func (s *VolcengineIamOidcProviderClientService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "RemoveClientIDFromOIDCProvider",
				ConvertMode: ve.RequestConvertIgnore,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					parts := strings.Split(d.Id(), ":")
					if len(parts) != 2 {
						return nil, fmt.Errorf("invalid id format")
					}
					param := map[string]interface{}{
						"OIDCProviderName": parts[0],
						"ClientID":         parts[1],
					}
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		},
	}
}

func (s *VolcengineIamOidcProviderClientService) createRequestConvert() map[string]ve.RequestConvert {
	return map[string]ve.RequestConvert{
		"oidc_provider_name": {TargetField: "OIDCProviderName"},
		"client_id":          {TargetField: "ClientID"},
	}
}

func (s *VolcengineIamOidcProviderClientService) ReadResource(d *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = d.Id()
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid id format")
	}
	providerName := parts[0]
	clientId := parts[1]

	action := "GetOIDCProvider"
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{
		"OIDCProviderName": providerName,
	})
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return nil, nil
		}
		return nil, err
	}

	clientIds, err := ve.ObtainSdkValue("Result.ClientIDs", *resp)
	if err != nil {
		return nil, err
	}
	if clientIds == nil {
		return nil, nil
	}

	list, ok := clientIds.([]interface{})
	if !ok {
		return nil, fmt.Errorf("ClientIDs is not a list")
	}

	for _, v := range list {
		if vStr, ok := v.(string); ok && vStr == clientId {
			return map[string]interface{}{
				"oidc_provider_name": providerName,
				"client_id":          clientId,
				"id":                 id,
			}, nil
		}
	}

	return nil, nil
}

func (s *VolcengineIamOidcProviderClientService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineIamOidcProviderClientService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamOidcProviderClientService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamOidcProviderClientService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineIamOidcProviderClientService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineIamOidcProviderClientService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	method := ve.POST
	if actionName == "GetOIDCProvider" {
		method = ve.GET
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

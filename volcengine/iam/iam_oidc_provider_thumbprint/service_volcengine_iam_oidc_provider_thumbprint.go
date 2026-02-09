package iam_oidc_provider_thumbprint

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

type VolcengineIamOidcProviderThumbprintService struct {
	Client *ve.SdkClient
}

func NewIamOidcProviderThumbprintService(c *ve.SdkClient) *VolcengineIamOidcProviderThumbprintService {
	return &VolcengineIamOidcProviderThumbprintService{
		Client: c,
	}
}

func (s *VolcengineIamOidcProviderThumbprintService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamOidcProviderThumbprintService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "AddThumbprintToOIDCProvider",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam: &map[string]interface{}{
					"OIDCProviderName": data.Get("oidc_provider_name").(string),
					"Thumbprint":       data.Get("thumbprint").(string),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					d.SetId(fmt.Sprintf("%s:%s", d.Get("oidc_provider_name").(string), d.Get("thumbprint").(string)))
					return nil
				},
			},
		},
	}
}

func (s *VolcengineIamOidcProviderThumbprintService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "RemoveThumbprintFromOIDCProvider",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					parts := strings.Split(d.Id(), ":")
					if len(parts) != 2 {
						return false, fmt.Errorf("invalid id format")
					}
					(*call.SdkParam)["OIDCProviderName"] = parts[0]
					(*call.SdkParam)["Thumbprint"] = parts[1]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		},
	}
}

func (s *VolcengineIamOidcProviderThumbprintService) ReadResource(d *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = d.Id()
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid id format")
	}
	providerName := parts[0]
	thumbprint := parts[1]

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

	thumbprints, err := ve.ObtainSdkValue("Result.Thumbprints", *resp)
	if err != nil {
		return nil, err
	}
	if thumbprints == nil {
		return nil, nil
	}

	list, ok := thumbprints.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Thumbprints is not a list")
	}

	for _, v := range list {
		if v.(string) == thumbprint {
			return map[string]interface{}{
				"oidc_provider_name": providerName,
				"thumbprint":         thumbprint,
				"id":                 id,
			}, nil
		}
	}

	return nil, nil
}

func (s *VolcengineIamOidcProviderThumbprintService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineIamOidcProviderThumbprintService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamOidcProviderThumbprintService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamOidcProviderThumbprintService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineIamOidcProviderThumbprintService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineIamOidcProviderThumbprintService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
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

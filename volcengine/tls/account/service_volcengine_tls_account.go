package account

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsAccountService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsAccountService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsAccountService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp          *map[string]interface{}
		accountStatus interface{}
	)

	action := "GetAccountStatus"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.Default,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, nil)
	if err != nil {
		return nil, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	accountStatus, err = ve.ObtainSdkValue("RESPONSE", *resp)
	if err != nil {
		return nil, err
	}

	statusMap, ok := accountStatus.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("response is not map[string]interface{}")
	}

	// Account is a singleton, return as slice for batch query compatibility
	data = append(data, statusMap)

	return data, nil
}

func (v *VolcengineTlsAccountService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	action := "GetAccountStatus"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.Default,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, nil)
	if err != nil {
		return nil, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	accountStatus, err := ve.ObtainSdkValue("RESPONSE", *resp)
	if err != nil {
		return nil, err
	}

	statusMap, ok := accountStatus.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("response is not map[string]interface{}")
	}

	// Set resource ID as "default" since it's a singleton
	resourceData.SetId("default")

	return statusMap, nil
}

func (v *VolcengineTlsAccountService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsAccountService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsAccountService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ActiveTlsAccount",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// Set resource ID as "default" since it's a singleton
				d.SetId("default")
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineTlsAccountService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineTlsAccountService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	// No Delete API available for account
	return []ve.Callback{}
}

func (v *VolcengineTlsAccountService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:      ve.ContentTypeJson,
		CollectField:     "tls_accounts",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (v *VolcengineTlsAccountService) ReadResourceId(id string) string {
	return id
}

func NewTlsAccountService(c *ve.SdkClient) *VolcengineTlsAccountService {
	return &VolcengineTlsAccountService{
		Client: c,
	}
}

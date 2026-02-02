package iam_identity_provider

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamIdentityProviderService struct {
	Client *ve.SdkClient
}

func NewIamIdentityProviderService(c *ve.SdkClient) *VolcengineIamIdentityProviderService {
	return &VolcengineIamIdentityProviderService{
		Client: c,
	}
}

func (s *VolcengineIamIdentityProviderService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamIdentityProviderService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithPageNumberQuery(m, "Limit", "Offset", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "ListIdentityProviders"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err := universalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, resp)
		results, err := ve.ObtainSdkValue("Result.IdentityProviders", *resp)
		if err != nil {
			return nil, err
		}
		if results == nil {
			return []interface{}{}, nil
		}
		return results.([]interface{}), nil
	})
}

func (s *VolcengineIamIdentityProviderService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		CollectField:    "providers",
		ResponseConverts: map[string]ve.ResponseConvert{
			"SSOType": {
				TargetField: "sso_type",
			},
		},
	}
}

func (s *VolcengineIamIdentityProviderService) CreateResource(d *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamIdentityProviderService) ReadResource(d *schema.ResourceData, id string) (map[string]interface{}, error) {
	return nil, nil
}

func (s *VolcengineIamIdentityProviderService) ModifyResource(d *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamIdentityProviderService) RemoveResource(d *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamIdentityProviderService) RefreshResourceState(d *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamIdentityProviderService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineIamIdentityProviderService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
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

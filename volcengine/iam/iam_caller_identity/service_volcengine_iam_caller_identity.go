package iam_caller_identity

import (
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamCallerIdentityService struct {
	Client *ve.SdkClient
}

func NewIamCallerIdentityService(c *ve.SdkClient) *VolcengineIamCallerIdentityService {
	return &VolcengineIamCallerIdentityService{
		Client: c,
	}
}

func (s *VolcengineIamCallerIdentityService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamCallerIdentityService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	action := "GetCallerIdentity"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
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

func (s *VolcengineIamCallerIdentityService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ResponseConverts: map[string]ve.ResponseConvert{
			"AccountId": {
				TargetField: "account_id",
				Convert: func(i interface{}) interface{} {
					if v, ok := i.(float64); ok {
						return strconv.FormatFloat(v, 'f', 0, 64)
					}
					return i
				},
			},
			"Trn": {
				TargetField: "trn",
			},
			"IdentityType": {
				TargetField: "identity_type",
			},
			"IdentityId": {
				TargetField: "identity_id",
				Convert: func(i interface{}) interface{} {
					if v, ok := i.(float64); ok {
						return strconv.FormatFloat(v, 'f', 0, 64)
					}
					return i
				},
			},
		},
		CollectField: "caller_identities",
	}
}

func (s *VolcengineIamCallerIdentityService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}
func (s *VolcengineIamCallerIdentityService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}
func (s *VolcengineIamCallerIdentityService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}
func (s *VolcengineIamCallerIdentityService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}
func (s *VolcengineIamCallerIdentityService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}
func (s *VolcengineIamCallerIdentityService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}
func (s *VolcengineIamCallerIdentityService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "sts",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

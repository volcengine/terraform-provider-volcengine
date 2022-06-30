package iam_policy

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
	"time"
)

type VestackIamPolicyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamPolicyService(c *ve.SdkClient) *VestackIamPolicyService {
	return &VestackIamPolicyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (v *VestackIamPolicyService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VestackIamPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListPolicies"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}

		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.PolicyMetadata", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.PolicyMetadata is not Slice")
		}
		return data, err
	})
}

func (v *VestackIamPolicyService) ReadResource(data *schema.ResourceData, s string) (map[string]interface{}, error) {
	panic("implement me")
}

func (v *VestackIamPolicyService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	panic("implement me")
}

func (v *VestackIamPolicyService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	panic("implement me")
}

func (v *VestackIamPolicyService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	panic("implement me")
}

func (v *VestackIamPolicyService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	panic("implement me")
}

func (v *VestackIamPolicyService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	panic("implement me")
}

func (v *VestackIamPolicyService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ResponseConverts: map[string]ve.ResponseConvert{
			"PolicyName": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
		NameField:    "PolicyName",
		IdField:      "PolicyName",
		CollectField: "policies",
	}
}

func (v *VestackIamPolicyService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

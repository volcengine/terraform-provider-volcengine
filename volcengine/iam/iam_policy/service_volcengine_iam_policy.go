package iam_policy

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamPolicyService struct {
	Client *ve.SdkClient
}

func NewIamPolicyService(c *ve.SdkClient) *VolcengineIamPolicyService {
	return &VolcengineIamPolicyService{
		Client: c,
	}
}

func (s *VolcengineIamPolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListPolicies"
		logger.Debug(logger.ReqFormat, action, condition)
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

func (s *VolcengineIamPolicyService) ReadResource(resourceData *schema.ResourceData, policyId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if policyId == "" {
		policyId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Query": policyId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Policy %s not exist ", policyId)
	}
	return data, err
}

func (s *VolcengineIamPolicyService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamPolicyService) WithResourceResponseHandlers(policy map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		policy["Id"] = policy["PolicyName"]
		return policy, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamPolicyService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	createIamPolicyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreatePolicy",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				policyName, err := ve.ObtainSdkValue("Result.Policy.PolicyName", *resp)
				if err != nil {
					return err
				}
				d.SetId(policyName.(string))
				return nil
			},
		},
	}
	return []ve.Callback{createIamPolicyCallback}
}

func (s *VolcengineIamPolicyService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	updatePolicyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdatePolicy",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["PolicyName"] = d.Get("policy_name")
				if d.HasChange("policy_name") {
					oldPolicyName, newPolicyName := d.GetChange("policy_name")
					(*call.SdkParam)["PolicyName"] = oldPolicyName
					(*call.SdkParam)["NewPolicyName"] = newPolicyName
				}
				if d.HasChange("policy_document") {
					(*call.SdkParam)["NewPolicyDocument"] = d.Get("policy_document")
				}
				if d.HasChange("description") {
					(*call.SdkParam)["NewDescription"] = d.Get("description")
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				if d.HasChange("policy_name") {
					policyName, err := ve.ObtainSdkValue("Result.Policy.PolicyName", *resp)
					if err != nil {
						return err
					}
					d.SetId(policyName.(string))
				}
				return nil
			},
		},
	}
	return []ve.Callback{updatePolicyCallback}
}

func (s *VolcengineIamPolicyService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	deletePolicyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeletePolicy",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["PolicyName"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading iam policy on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{deletePolicyCallback}
}

func (s *VolcengineIamPolicyService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
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

func (s *VolcengineIamPolicyService) ReadResourceId(id string) string {
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

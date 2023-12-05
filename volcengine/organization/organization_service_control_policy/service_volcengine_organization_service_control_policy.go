package organization_service_control_policy

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcenginePolicyService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcenginePolicyService {
	return &VolcenginePolicyService{
		Client: c,
	}
}

func (s *VolcenginePolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcenginePolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) (data []interface{}, err error) {
		action := "ListServiceControlPolicies"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return nil, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return nil, err
			}
		}

		results, err = ve.ObtainSdkValue("Result.ServiceControlPolicies", *resp)
		if err != nil {
			return nil, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ServiceControlPolicies is not Slice")
		}

		// 获取每一个策略内容
		for _, ele := range data {
			temp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("GetServiceControlPolicy"), &map[string]interface{}{
				"PolicyID": ele.(map[string]interface{})["PolicyID"],
			})
			if err != nil {
				return nil, err
			}
			statement, err := ve.ObtainSdkValue("Result.Statement", *temp)
			if err != nil {
				return nil, err
			}
			ele.(map[string]interface{})["Statement"] = statement
		}
		return data, err
	})
}

func (s *VolcenginePolicyService) ReadResource(resourceData *schema.ResourceData, policyId string) (data map[string]interface{}, err error) {
	if policyId == "" {
		policyId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"PolicyID": policyId,
	}
	temp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("GetServiceControlPolicy"), &req)
	if err != nil {
		return nil, err
	}
	res, err := ve.ObtainSdkValue("Result", *temp)
	if err != nil {
		return nil, err
	}
	return res.(map[string]interface{}), nil
}

func (s *VolcenginePolicyService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcenginePolicyService) WithResourceResponseHandlers(policy map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcenginePolicyService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	createIamPolicyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateServiceControlPolicy",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(postUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				policyId, err := ve.ObtainSdkValue("Result.PolicyId", *resp)
				if err != nil {
					return err
				}
				d.SetId(policyId.(string))
				return nil
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-ServiceControlPolicy"
			},
		},
	}
	return []ve.Callback{createIamPolicyCallback}
}

func (s *VolcenginePolicyService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	updatePolicyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateServiceControlPolicy",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"policy_name": {
					ForceGet: true,
				},
				"description": {
					ForceGet: true,
				},
				"statement": {
					ForceGet: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["PolicyID"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(postUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return "lock-ServiceControlPolicy"
			},
		},
	}
	return []ve.Callback{updatePolicyCallback}
}

func (s *VolcenginePolicyService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	deletePolicyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteServiceControlPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["PolicyID"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(postUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return "lock-ServiceControlPolicy"
			},
		},
	}
	return []ve.Callback{deletePolicyCallback}
}

func (s *VolcenginePolicyService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ResponseConverts: map[string]ve.ResponseConvert{
			"PolicyID": {
				TargetField: "id",
			},
		},
		NameField:    "PolicyName",
		IdField:      "PolicyID",
		CollectField: "policies",
	}
}

func (s *VolcenginePolicyService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "organization",
		Version:     "2022-01-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

func postUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "organization",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		Action:      actionName,
		ContentType: ve.ApplicationJSON,
	}
}

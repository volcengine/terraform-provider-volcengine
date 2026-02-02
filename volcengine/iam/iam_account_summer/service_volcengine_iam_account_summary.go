package iam_account_summer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"time"
)

type VolcengineIamAccountSummaryService struct {
	Client *ve.SdkClient
}

func NewIamAccountSummaryService(c *ve.SdkClient) *VolcengineIamAccountSummaryService {
	return &VolcengineIamAccountSummaryService{
		Client: c,
	}
}

func (s *VolcengineIamAccountSummaryService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamAccountSummaryService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	action := "GetAccountSummary"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
	if err != nil {
		return data, err
	}

	result, err := ve.ObtainSdkValue("Result.SummaryMap", *resp)
	if err != nil {
		return data, err
	}
	if result == nil {
		return []interface{}{}, nil
	}

	return []interface{}{result}, nil
}

func (s *VolcengineIamAccountSummaryService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ResponseConverts: map[string]ve.ResponseConvert{
			"AccessKeysPerUserQuota":              {TargetField: "access_keys_per_user_quota"},
			"AccessKeysPerAccountQuota":           {TargetField: "access_keys_per_account_quota"},
			"AttachedPoliciesPerGroupQuota":       {TargetField: "attached_policies_per_group_quota"},
			"AttachedSystemPoliciesPerGroupQuota": {TargetField: "attached_system_policies_per_group_quota"},
			"AttachedPoliciesPerRoleQuota":        {TargetField: "attached_policies_per_role_quota"},
			"AttachedSystemPoliciesPerRoleQuota":  {TargetField: "attached_system_policies_per_role_quota"},
			"AttachedSystemPoliciesPerUserQuota":  {TargetField: "attached_system_policies_per_user_quota"},
			"AttachedPoliciesPerUserQuota":        {TargetField: "attached_policies_per_user_quota"},
			"GroupsQuota":                         {TargetField: "groups_quota"},
			"PoliciesQuota":                       {TargetField: "policies_quota"},
			"RolesQuota":                          {TargetField: "roles_quota"},
			"UsersQuota":                          {TargetField: "users_quota"},
			"PolicySize":                          {TargetField: "policy_size"},
			"RolesUsage":                          {TargetField: "roles_usage"},
			"UsersUsage":                          {TargetField: "users_usage"},
			"GroupsUsage":                         {TargetField: "groups_usage"},
			"PoliciesUsage":                       {TargetField: "policies_usage"},
			"GroupsPerUserQuota":                  {TargetField: "groups_per_user_quota"},
		},
		CollectField: "account_summaries",
	}
}

func (s *VolcengineIamAccountSummaryService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}
func (s *VolcengineIamAccountSummaryService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}
func (s *VolcengineIamAccountSummaryService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}
func (s *VolcengineIamAccountSummaryService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}
func (s *VolcengineIamAccountSummaryService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}
func (s *VolcengineIamAccountSummaryService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}
func (s *VolcengineIamAccountSummaryService) ReadResourceId(id string) string {
	return id
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

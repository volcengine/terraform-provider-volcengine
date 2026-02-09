package iam_entities_policy

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamEntitiesPolicyService struct {
	Client *ve.SdkClient
}

func NewIamEntitiesPolicyService(c *ve.SdkClient) *VolcengineIamEntitiesPolicyService {
	return &VolcengineIamEntitiesPolicyService{
		Client: c,
	}
}

func (s *VolcengineIamEntitiesPolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamEntitiesPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var resp *map[string]interface{}
	action := "ListEntitiesForPolicy"
	logger.Debug(logger.ReqFormat, action, m)
	if m == nil {
		m = make(map[string]interface{})
	}
	// Save input parameters before DoCall in case m is modified
	policyName := m["PolicyName"]
	policyType := m["PolicyType"]
	entityFilter := m["EntityFilter"]

	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return data, err
	}
	if resp == nil {
		return data, errors.New("response is nil")
	}

	resultMap := make(map[string]interface{})

	// Extract and map entities to snake_case for manual mapping in Read function
	users, _ := ve.ObtainSdkValue("Result.PolicyUsers", *resp)
	resultMap["users"] = convertEntities(users, map[string]string{
		"UserName":    "user_name",
		"DisplayName": "display_name",
		"Description": "description",
		"AttachDate":  "attach_date",
		"ID":          "id",
	})

	roles, _ := ve.ObtainSdkValue("Result.PolicyRoles", *resp)
	resultMap["roles"] = convertEntities(roles, map[string]string{
		"RoleName":    "role_name",
		"DisplayName": "display_name",
		"Description": "description",
		"AttachDate":  "attach_date",
		"ID":          "id",
	})

	userGroups, _ := ve.ObtainSdkValue("Result.PolicyUserGroups", *resp)
	resultMap["user_groups"] = convertEntities(userGroups, map[string]string{
		"UserGroupName": "user_group_name",
		"DisplayName":   "display_name",
		"Description":   "description",
		"AttachDate":    "attach_date",
		"ID":            "id",
	})

	total, _ := ve.ObtainSdkValue("Result.Total", *resp)
	if total == nil {
		resultMap["total_count"] = 0
	} else {
		if f, ok := total.(float64); ok {
			resultMap["total_count"] = int(f)
		} else {
			resultMap["total_count"] = total
		}
	}

	// Inject input parameters
	resultMap["policy_name"] = policyName
	resultMap["policy_type"] = policyType
	resultMap["entity_filter"] = entityFilter
	resultMap["id"] = policyName

	return []interface{}{resultMap}, nil
}

func convertEntities(entities interface{}, mapping map[string]string) []interface{} {
	if entities == nil {
		return []interface{}{}
	}
	list, ok := entities.([]interface{})
	if !ok {
		return []interface{}{}
	}
	var result []interface{}
	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			mapped := make(map[string]interface{})
			for sdkKey, tfKey := range mapping {
				if v, ok := m[sdkKey]; ok {
					// Convert float64 to int for ID fields if necessary
					if tfKey == "id" {
						if f, ok := v.(float64); ok {
							mapped[tfKey] = int(f)
							continue
						}
					}
					mapped[tfKey] = v
				}
			}
			// Handle PolicyScope
			if scopes, ok := m["PolicyScope"].([]interface{}); ok {
				var mappedScopes []interface{}
				for _, scope := range scopes {
					if sm, ok := scope.(map[string]interface{}); ok {
						mappedScopes = append(mappedScopes, map[string]interface{}{
							"policy_scope_type":    sm["PolicyScopeType"],
							"project_name":         sm["ProjectName"],
							"project_display_name": sm["ProjectDisplayName"],
							"attach_date":          sm["AttachDate"],
						})
					}
				}
				mapped["policy_scope"] = mappedScopes
			}
			result = append(result, mapped)
		}
	}
	return result
}

func (s *VolcengineIamEntitiesPolicyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, errors.New("not support")
}

func (s *VolcengineIamEntitiesPolicyService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamEntitiesPolicyService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineIamEntitiesPolicyService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIamEntitiesPolicyService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIamEntitiesPolicyService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIamEntitiesPolicyService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"policy_name": {
				TargetField: "PolicyName",
			},
			"policy_type": {
				TargetField: "PolicyType",
			},
			"entity_filter": {
				TargetField: "EntityFilter",
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"PolicyUsers": {
				TargetField: "users",
			},
			"PolicyRoles": {
				TargetField: "roles",
			},
			"PolicyUserGroups": {
				TargetField: "user_groups",
			},
			"Total": {
				TargetField: "total_count",
			},
			"PolicyName": {
				TargetField: "policy_name",
			},
			"PolicyType": {
				TargetField: "policy_type",
			},
			"EntityFilter": {
				TargetField: "entity_filter",
			},
			"ID": {
				TargetField: "id",
			},
		},
		IdField: "ID",
	}
}

func (s *VolcengineIamEntitiesPolicyService) ReadResourceId(id string) string {
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

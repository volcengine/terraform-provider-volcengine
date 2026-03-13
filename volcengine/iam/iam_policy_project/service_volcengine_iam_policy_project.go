package iam_policy_project

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

type VolcengineIamPolicyProjectService struct {
	Client *ve.SdkClient
}

func NewIamPolicyProjectService(c *ve.SdkClient) *VolcengineIamPolicyProjectService {
	return &VolcengineIamPolicyProjectService{
		Client: c,
	}
}

func (s *VolcengineIamPolicyProjectService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamPolicyProjectService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	v := data.Get("project_names")
	tagsSet, ok := v.(*schema.Set)
	if !ok {
		return []ve.Callback{{Err: errors.New("project_names is not *schema.Set")}}
	}
	projects := tagsSet.List()
	var callbacks []ve.Callback
	for i, p := range projects {
		pName, ok := p.(string)
		if !ok {
			return []ve.Callback{{Err: errors.New("project_name item is not string")}}
		}
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "AttachPolicyInProject",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if call.SdkParam == nil {
						param := make(map[string]interface{})
						call.SdkParam = &param
					}
					principalType, ok := d.Get("principal_type").(string)
					if !ok {
						return false, errors.New("principal_type is not string")
					}
					principalName, ok := d.Get("principal_name").(string)
					if !ok {
						return false, errors.New("principal_name is not string")
					}
					policyType, ok := d.Get("policy_type").(string)
					if !ok {
						return false, errors.New("policy_type is not string")
					}
					policyName, ok := d.Get("policy_name").(string)
					if !ok {
						return false, errors.New("policy_name is not string")
					}
					(*call.SdkParam)["PrincipalType"] = principalType
					(*call.SdkParam)["PrincipalName"] = principalName
					(*call.SdkParam)["PolicyType"] = policyType
					(*call.SdkParam)["PolicyName"] = policyName
					(*call.SdkParam)["ProjectName.1"] = pName
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		if i == len(projects)-1 {
			callback.Call.AfterCall = func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				var projectNames []string
				for _, p := range projects {
					if pStr, ok := p.(string); ok {
						projectNames = append(projectNames, pStr)
					}
				}
				principalType, ok := d.Get("principal_type").(string)
				if !ok {
					return errors.New("principal_type is not string")
				}
				principalName, ok := d.Get("principal_name").(string)
				if !ok {
					return errors.New("principal_name is not string")
				}
				policyType, ok := d.Get("policy_type").(string)
				if !ok {
					return errors.New("policy_type is not string")
				}
				policyName, ok := d.Get("policy_name").(string)
				if !ok {
					return errors.New("policy_name is not string")
				}
				d.SetId(fmt.Sprintf("%s:%s:%s:%s:%s",
					principalType,
					principalName,
					policyType,
					policyName,
					strings.Join(projectNames, ",")))
				return nil
			}
		}
		callbacks = append(callbacks, callback)
	}
	return callbacks
}

func (s *VolcengineIamPolicyProjectService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(data.Id(), ":")
	if len(ids) < 5 {
		return nil
	}
	principalType := ids[0]
	principalName := ids[1]
	policyType := ids[2]
	policyName := ids[3]
	projectNames := strings.Split(ids[4], ",")

	var callbacks []ve.Callback
	for _, pName := range projectNames {
		projectName := pName
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "DetachPolicyInProject",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if call.SdkParam == nil {
						param := make(map[string]interface{})
						call.SdkParam = &param
					}
					(*call.SdkParam)["PrincipalType"] = principalType
					(*call.SdkParam)["PrincipalName"] = principalName
					(*call.SdkParam)["PolicyType"] = policyType
					(*call.SdkParam)["PolicyName"] = policyName
					(*call.SdkParam)["ProjectName.1"] = projectName
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)
	}
	return callbacks
}

func (s *VolcengineIamPolicyProjectService) ReadResource(d *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = d.Id()
	}
	ids := strings.Split(id, ":")
	if len(ids) < 5 {
		return nil, fmt.Errorf("invalid id format")
	}
	principalType := ids[0]
	principalName := ids[1]
	policyType := ids[2]
	policyName := ids[3]
	projectNamesStr := ids[4]
	projectNames := strings.Split(projectNamesStr, ",")

	var action string
	var param map[string]interface{}
	switch principalType {
	case "User":
		action = "ListAttachedUserPolicies"
		param = map[string]interface{}{"UserName": principalName}
	case "Role":
		action = "ListAttachedRolePolicies"
		param = map[string]interface{}{"RoleName": principalName}
	case "UserGroup":
		action = "ListAttachedUserGroupPolicies"
		param = map[string]interface{}{"UserGroupName": principalName}
	default:
		return nil, fmt.Errorf("invalid principal type: %s", principalType)
	}

	resp, err := s.Client.UniversalClient.DoCall(ve.UniversalInfo{
		ServiceName: "iam",
		Action:      action,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
	}, &param)
	if err != nil {
		if strings.Contains(err.Error(), "NotExist") {
			return nil, nil
		}
		return nil, err
	}

	results, err := ve.ObtainSdkValue("Result.AttachedPolicyMetadata", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return nil, nil
	}

	list, ok := results.([]interface{})
	if !ok {
		return nil, fmt.Errorf("AttachedPolicyMetadata not list")
	}

	foundProjects := make(map[string]bool)
	hasSystemScope := false
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New("result item is not map")
		}
		pName, ok := m["PolicyName"].(string)
		if !ok {
			return nil, errors.New("result item PolicyName is not string")
		}
		pType, ok := m["PolicyType"].(string)
		if !ok {
			return nil, errors.New("result item PolicyType is not string")
		}
		if strings.EqualFold(pName, policyName) && strings.EqualFold(pType, policyType) {
			scopes, ok := m["PolicyScope"].([]interface{})
			if ok {
				for _, scope := range scopes {
					sm, ok := scope.(map[string]interface{})
					if !ok {
						return nil, errors.New("policy scope item is not map")
					}
					scopeType, ok := sm["PolicyScopeType"].(string)
					if !ok {
						return nil, errors.New("policy scope type is not string")
					}
					if strings.EqualFold(scopeType, "System") {
						hasSystemScope = true
					}
					pName, ok := sm["ProjectName"].(string)
					if ok && pName != "" {
						foundProjects[pName] = true
					}
				}
			}
		}
	}

	var foundList []interface{}
	for _, name := range projectNames {
		if name != "" {
			if hasSystemScope || foundProjects[name] {
				foundList = append(foundList, name)
			}
		}
	}

	if len(foundList) == 0 {
		return nil, fmt.Errorf("policy %s:%s not found in projects %v for %s %s", policyType, policyName, projectNames, principalType, principalName)
	}

	return map[string]interface{}{
		"PrincipalType": principalType,
		"PrincipalName": principalName,
		"PolicyType":    policyType,
		"PolicyName":    policyName,
		"ProjectNames":  foundList,
		"ID":            id,
	}, nil
}

func (s *VolcengineIamPolicyProjectService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineIamPolicyProjectService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamPolicyProjectService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamPolicyProjectService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineIamPolicyProjectService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineIamPolicyProjectService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2021-08-01",
		HttpMethod:  ve.POST,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

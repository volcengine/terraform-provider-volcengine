package vmp_rule

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVmpRuleService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineVmpRuleService {
	return &VolcengineVmpRuleService{
		Client: c,
	}
}

func (s *VolcengineVmpRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVmpRuleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRules"
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
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}

		return data, err
	})
}

func (s *VolcengineVmpRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineVmpRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineVmpRuleService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}

}

func (s *VolcengineVmpRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}

}

func (s *VolcengineVmpRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	return callbacks
}

func (s *VolcengineVmpRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineVmpRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"rule_file_names": {
				TargetField: "Filter.RuleFileNames",
				ConvertType: ve.ConvertJsonArray,
			},
			"rule_group_names": {
				TargetField: "Filter.RuleGroupNames",
				ConvertType: ve.ConvertJsonArray,
			},
			"kind": {
				TargetField: "Filter.Kind",
			},
			"status": {
				TargetField: "Filter.Status",
			},
			"workspace_id": {
				TargetField: "Filter.WorkspaceId",
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Name",
		CollectField: "rules",
	}
}

func (s *VolcengineVmpRuleService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vmp",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

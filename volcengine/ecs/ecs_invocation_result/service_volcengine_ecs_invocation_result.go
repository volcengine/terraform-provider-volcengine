package ecs_invocation_result

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEcsInvocationResultService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEcsInvocationResultService(c *ve.SdkClient) *VolcengineEcsInvocationResultService {
	return &VolcengineEcsInvocationResultService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEcsInvocationResultService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEcsInvocationResultService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeInvocationResults"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.InvocationResults", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.InvocationResults is not Slice")
		}

		return data, err
	})
}

func (s *VolcengineEcsInvocationResultService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineEcsInvocationResultService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineEcsInvocationResultService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return data, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEcsInvocationResultService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEcsInvocationResultService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEcsInvocationResultService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEcsInvocationResultService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"invocation_result_status": {
				TargetField: "InvocationResultStatus",
				Convert: func(data *schema.ResourceData, i interface{}) interface{} {
					var status string
					statusSet, ok := data.GetOk("invocation_result_status")
					if !ok {
						return status
					}
					statusList := statusSet.(*schema.Set).List()
					statusArr := make([]string, 0)
					for _, value := range statusList {
						statusArr = append(statusArr, value.(string))
					}
					status = strings.Join(statusArr, ",")
					return status
				},
			},
		},
		IdField:      "InvocationResultId",
		CollectField: "invocation_results",
		ResponseConverts: map[string]ve.ResponseConvert{
			"InvocationResultId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineEcsInvocationResultService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

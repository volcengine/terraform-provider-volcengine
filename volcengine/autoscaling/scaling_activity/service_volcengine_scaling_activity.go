package scaling_activity

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineScalingActivityService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewScalingActivityService(c *ve.SdkClient) *VolcengineScalingActivityService {
	return &VolcengineScalingActivityService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineScalingActivityService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalingActivityService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeScalingActivities"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = universalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.ScalingActivities", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ScalingActivities is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineScalingActivityService) ReadResource(resourceData *schema.ResourceData, activityId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if activityId == "" {
		activityId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ScalingActivityIds.1": activityId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Scaling Activity %s not exist ", activityId)
	}
	return data, err
}

func (s *VolcengineScalingActivityService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineScalingActivityService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return data, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineScalingActivityService) CreateResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineScalingActivityService) ModifyResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineScalingActivityService) RemoveResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineScalingActivityService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ScalingActivityIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		IdField:      "ScalingActivityId",
		CollectField: "activities",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ScalingActivityId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineScalingActivityService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "auto_scaling",
		Action:      actionName,
		Version:     "2020-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

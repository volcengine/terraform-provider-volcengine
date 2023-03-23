package scaling_instance

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"time"
)

type VolcengineScalingInstanceService struct {
	Client *ve.SdkClient
}

func (s *VolcengineScalingInstanceService) ReadResource(data *schema.ResourceData, s2 string) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (s *VolcengineScalingInstanceService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s2 string) *resource.StateChangeConf {
	return nil
}

func NewScalingInstanceService(c *ve.SdkClient) *VolcengineScalingInstanceService {
	return &VolcengineScalingInstanceService{
		Client: c,
	}
}

func (s *VolcengineScalingInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalingInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeScalingInstances"
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
		logger.Debug(logger.RespFormat, action, action, resp)
		results, err = ve.ObtainSdkValue("Result.ScalingInstances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ScalingInstances is not Slice")
		}
		return data, err
	})
}

func (VolcengineScalingInstanceService) WithResourceResponseHandlers(scalingInstance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingInstance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineScalingInstanceService) CreateResource(d *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingInstanceService) ModifyResource(data *schema.ResourceData, s2 *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingInstanceService) RemoveResource(data *schema.ResourceData, s2 *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingInstanceService) DatasourceResources(data *schema.ResourceData, s2 *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "InstanceIds",
				ConvertType: ve.ConvertWithN,
			},
			"status": {
				TargetField: "Status",
				ConvertType: ve.ConvertDefault,
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
		IdField:      "InstanceId",
		CollectField: "scaling_instances",
	}
}

func (s *VolcengineScalingInstanceService) ReadResourceId(id string) string {
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

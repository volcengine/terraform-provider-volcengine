package scaling_group_enable

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineScalingGroupEnableService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func (s *VolcengineScalingGroupEnableService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeScalingGroups"
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
		logger.Debug(logger.RespFormat, action, action, resp)
		results, err = ve.ObtainSdkValue("Result.ScalingGroups", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ScalingGroups is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineScalingGroupEnableService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("Invalid ScalingGroupEnable Id ")
	}
	req := map[string]interface{}{
		"ScalingGroupIds.1": ids[1],
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
		return data, fmt.Errorf("ScalingGroup %s not exist ", ids[1])
	}
	if data["LifecycleState"] != "Active" {
		return data, fmt.Errorf("ScalingGroup %s is not active", ids[1])
	}
	return data, err
}

func (s *VolcengineScalingGroupEnableService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineScalingGroupEnableService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineScalingGroupEnableService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	param := &map[string]interface{}{
		"ScalingGroupId": data.Get("scaling_group_id").(string),
	}
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableScalingGroup",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam:    param,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.ScalingGroupId", *resp)
				d.SetId(fmt.Sprintf("enable:%s", id.(string)))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineScalingGroupEnableService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingGroupEnableService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	param := &map[string]interface{}{
		"ScalingGroupId": data.Get("scaling_group_id").(string),
	}
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisableScalingGroup",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam:    param,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineScalingGroupEnableService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineScalingGroupEnableService) ReadResourceId(id string) string {
	return id
}

func NewScalingGroupEnableService(client *ve.SdkClient) *VolcengineScalingGroupEnableService {
	return &VolcengineScalingGroupEnableService{
		Client:     client,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineScalingGroupEnableService) GetClient() *ve.SdkClient {
	return s.Client
}

func getUniversalInfo(action string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "auto_scaling",
		Action:      action,
		Version:     "2020-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

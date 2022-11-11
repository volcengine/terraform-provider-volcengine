package scaling_configuration_enable

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

type VolcengineScalingConfigurationEnableService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func (s *VolcengineScalingConfigurationEnableService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1,
		func(condition map[string]interface{}) ([]interface{}, error) {
			client := s.Client.UniversalClient
			action := "DescribeScalingConfigurations"
			logger.Debug(logger.ReqFormat, action, condition)
			if condition == nil {
				resp, err = client.DoCall(getUniversalInfo(action), nil)
				if err != nil {
					return data, err
				}
			} else {
				resp, err = client.DoCall(getUniversalInfo(action), &condition)
				if err != nil {
					return data, err
				}
			}
			logger.Debug(logger.RespFormat, action, condition, *resp)
			results, err = ve.ObtainSdkValue("Result.ScalingConfigurations", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			if data, ok = results.([]interface{}); !ok {
				return data, errors.New("Result.ScalingConfigurations is not Slice")
			}
			return data, err
		})
}

func (s *VolcengineScalingConfigurationEnableService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"ScalingGroupId":            ids[0],
		"ScalingConfigurationIds.1": ids[1],
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
		return data, fmt.Errorf("The ScalingConfiguration %s bound to ScalingGroup %s does not exist ", ids[1], ids[0])
	}
	return data, err
}

func (s *VolcengineScalingConfigurationEnableService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s2 string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineScalingConfigurationEnableService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineScalingConfigurationEnableService) CreateResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	groupId := data.Get("scaling_group_id").(string)
	configId := data.Get("scaling_configuration_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableScalingConfiguration",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ScalingGroupId":         groupId,
				"ScalingConfigurationId": configId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["ScalingGroupId"], ":", (*call.SdkParam)["ScalingConfigurationId"]))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineScalingConfigurationEnableService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingConfigurationEnableService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingConfigurationEnableService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineScalingConfigurationEnableService) ReadResourceId(id string) string {
	return id
}

func NewScalingConfigurationEnableService(client *ve.SdkClient) *VolcengineScalingConfigurationEnableService {
	return &VolcengineScalingConfigurationEnableService{
		Client:     client,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineScalingConfigurationEnableService) GetClient() *ve.SdkClient {
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

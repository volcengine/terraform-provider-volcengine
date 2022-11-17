package scaling_configuration_attachment

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

type VolcengineScalingConfigurationAttachmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func (s *VolcengineScalingConfigurationAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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

func (s *VolcengineScalingConfigurationAttachmentService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, errors.New("Invalid ScalingConfigurationAttachment Id ")
	}
	req := map[string]interface{}{
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
		return data, fmt.Errorf("The ScalingConfiguration %s does not exist ", ids[1])
	}
	return data, err
}

func (s *VolcengineScalingConfigurationAttachmentService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s2 string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineScalingConfigurationAttachmentService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineScalingConfigurationAttachmentService) CreateResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var (
		readData map[string]interface{}
		err      error
		configId string
		groupId  string
	)
	configId = data.Get("scaling_configuration_id").(string)
	readData, err = s.ReadResource(data, fmt.Sprintf("enable:%s", configId))
	if err != nil {
		logger.DebugInfo("Failed to read scaling configuration resource", false)
		return []ve.Callback{}
	}
	groupId = readData["ScalingGroupId"].(string)
	logger.Debug(logger.RespFormat, "Read ScalingGroupId", configId, groupId)
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
				d.SetId(fmt.Sprintf("enable:%s", (*call.SdkParam)["ScalingConfigurationId"]))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineScalingConfigurationAttachmentService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingConfigurationAttachmentService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingConfigurationAttachmentService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineScalingConfigurationAttachmentService) ReadResourceId(id string) string {
	return id
}

func NewScalingConfigurationAttachmentService(client *ve.SdkClient) *VolcengineScalingConfigurationAttachmentService {
	return &VolcengineScalingConfigurationAttachmentService{
		Client:     client,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineScalingConfigurationAttachmentService) GetClient() *ve.SdkClient {
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

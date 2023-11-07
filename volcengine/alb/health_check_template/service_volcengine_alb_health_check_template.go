package alb_health_check_template

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineAlbHealthCheckTemplateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAlbHealthCheckTemplateService(c *ve.SdkClient) *VolcengineAlbHealthCheckTemplateService {
	return &VolcengineAlbHealthCheckTemplateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAlbHealthCheckTemplateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAlbHealthCheckTemplateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeHealthCheckTemplates"

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
		results, err = ve.ObtainSdkValue("Result.HealthCheckTemplates", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.HealthCheckTemplates is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineAlbHealthCheckTemplateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"HealthCheckTemplateIds.1": id,
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
		return data, fmt.Errorf("alb_health_check_template %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineAlbHealthCheckTemplateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineAlbHealthCheckTemplateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateHealthCheckTemplates",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"health_check_template_name": {
					TargetField: "HealthCheckTemplates.1.HealthCheckTemplateName",
				},
				"description": {
					TargetField: "HealthCheckTemplates.1.Description",
				},
				"health_check_interval": {
					TargetField: "HealthCheckTemplates.1.HealthCheckInterval",
				},
				"health_check_timeout": {
					TargetField: "HealthCheckTemplates.1.HealthCheckTimeout",
				},
				"healthy_threshold": {
					TargetField: "HealthCheckTemplates.1.HealthyThreshold",
				},
				"unhealthy_threshold": {
					TargetField: "HealthCheckTemplates.1.UnhealthyThreshold",
				},
				"health_check_method": {
					TargetField: "HealthCheckTemplates.1.HealthCheckMethod",
				},
				"health_check_domain": {
					TargetField: "HealthCheckTemplates.1.HealthCheckDomain",
				},
				"health_check_uri": {
					TargetField: "HealthCheckTemplates.1.HealthCheckURI",
				},
				"health_check_http_code": {
					TargetField: "HealthCheckTemplates.1.HealthCheckHttpCode",
				},
				"health_check_protocol": {
					TargetField: "HealthCheckTemplates.1.HealthCheckProtocol",
				},
				"health_check_http_version": {
					TargetField: "HealthCheckTemplates.1.HealthCheckHttpVersion",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ids, err := ve.ObtainSdkValue("Result.HealthCheckTemplateIDs", *resp)
				if err != nil {
					return err
				}
				idArr, ok := ids.([]interface{})
				if !ok || len(idArr) == 0 {
					return fmt.Errorf("ids is invalid")
				}
				d.SetId(idArr[0].(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineAlbHealthCheckTemplateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAlbHealthCheckTemplateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyHealthCheckTemplatesAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"health_check_template_name": {
					TargetField: "HealthCheckTemplates.1.HealthCheckTemplateName",
				},
				"description": {
					TargetField: "HealthCheckTemplates.1.Description",
				},
				"health_check_interval": {
					TargetField: "HealthCheckTemplates.1.HealthCheckInterval",
				},
				"health_check_timeout": {
					TargetField: "HealthCheckTemplates.1.HealthCheckTimeout",
				},
				"healthy_threshold": {
					TargetField: "HealthCheckTemplates.1.HealthyThreshold",
				},
				"unhealthy_threshold": {
					TargetField: "HealthCheckTemplates.1.UnhealthyThreshold",
				},
				"health_check_method": {
					TargetField: "HealthCheckTemplates.1.HealthCheckMethod",
				},
				"health_check_domain": {
					TargetField: "HealthCheckTemplates.1.HealthCheckDomain",
				},
				"health_check_uri": {
					TargetField: "HealthCheckTemplates.1.HealthCheckURI",
				},
				"health_check_http_code": {
					TargetField: "HealthCheckTemplates.1.HealthCheckHttpCode",
				},
				"health_check_protocol": {
					TargetField: "HealthCheckTemplates.1.HealthCheckProtocol",
				},
				"health_check_http_version": {
					TargetField: "HealthCheckTemplates.1.HealthCheckHttpVersion",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["HealthCheckTemplates.1.HealthCheckTemplateId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineAlbHealthCheckTemplateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteHealthCheckTemplates",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"HealthCheckTemplateIds.1": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineAlbHealthCheckTemplateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "HealthCheckTemplateIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "HealthCheckTemplateName",
		IdField:      "HealthCheckTemplateId",
		CollectField: "health_check_templates",
		ResponseConverts: map[string]ve.ResponseConvert{
			"HealthCheckTemplateId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"HealthCheckURI": {
				TargetField: "health_check_uri",
			},
		},
	}
}

func (s *VolcengineAlbHealthCheckTemplateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "alb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

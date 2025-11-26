package rds_postgresql_parameter_template

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

type VolcengineRdsPostgresqlParameterTemplateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlParameterTemplateService(c *ve.SdkClient) *VolcengineRdsPostgresqlParameterTemplateService {
	return &VolcengineRdsPostgresqlParameterTemplateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlParameterTemplateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlParameterTemplateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListParameterTemplates"

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
		results, err = ve.ObtainSdkValue("Result.TemplateInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.TemplateInfos is not Slice")
		}
		for _, value := range data {
			item, ok := value.(map[string]interface{})
			if !ok {
				return data, errors.New("Value is not map ")
			}
			action = "DescribeParameterTemplate"
			req := map[string]interface{}{
				"TemplateId": item["TemplateId"],
			}
			logger.Debug(logger.ReqFormat, action, req)
			detailResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, *detailResp)
			templateParams, err := ve.ObtainSdkValue("Result.TemplateInfo.TemplateParams", *detailResp)
			if err != nil {
				return data, err
			}
			item["TemplateParams"] = templateParams
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlParameterTemplateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	back := make(map[string]interface{})
	for _, v := range results {
		if back, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if back["TemplateId"].(string) == id {
			data = back
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("rds_postgresql_parameter_template %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlParameterTemplateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("rds_postgresql_parameter_template status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineRdsPostgresqlParameterTemplateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 如果设置了 src_template_id 则走 CloneParameterTemplate
	if v, ok := resourceData.GetOk("src_template_id"); ok && v.(string) != "" {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "CloneParameterTemplate",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["SrcTemplateId"] = d.Get("src_template_id")
					(*call.SdkParam)["TemplateName"] = d.Get("template_name")
					if desc, ok := d.GetOk("template_desc"); ok {
						(*call.SdkParam)["TemplateDesc"] = desc
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					id, _ := ve.ObtainSdkValue("Result.TemplateId", *resp)
					d.SetId(id.(string))
					_ = d.Set("template_id", id)
					return nil
				},
			},
		}
		return []ve.Callback{callback}
	}
	// 如果设置了 instance_id 则走 SaveAsParameterTemplate
	if v, ok := resourceData.GetOk("instance_id"); ok && v.(string) != "" {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SaveAsParameterTemplate",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					(*call.SdkParam)["TemplateName"] = d.Get("template_name")
					if desc, ok := d.GetOk("template_desc"); ok {
						(*call.SdkParam)["TemplateDesc"] = desc
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					id, _ := ve.ObtainSdkValue("Result.TemplateId", *resp)
					d.SetId(id.(string))
					_ = d.Set("template_id", id)
					return nil
				},
			},
		}
		return []ve.Callback{callback}
	}
	// 否则走 CreateParameterTemplate
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateParameterTemplate",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"template_name":         {TargetField: "TemplateName"},
				"template_type":         {TargetField: "TemplateType"},
				"template_type_version": {TargetField: "TemplateTypeVersion"},
				"template_desc":         {TargetField: "TemplateDesc"},
				"template_params": {ConvertType: ve.ConvertJsonObjectArray, NextLevelConvert: map[string]ve.RequestConvert{
					"name":  {TargetField: "Name", ForceGet: true},
					"value": {TargetField: "Value", ForceGet: true},
				}},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.TemplateId", *resp)
				d.SetId(id.(string))
				_ = d.Set("template_id", id)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlParameterTemplateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"TemplateId":          {TargetField: "template_id"},
			"TemplateName":        {TargetField: "template_name"},
			"TemplateType":        {TargetField: "template_type"},
			"TemplateTypeVersion": {TargetField: "template_type_version"},
			"TemplateDesc":        {TargetField: "template_desc"},
			"NeedRestart":         {TargetField: "need_restart"},
			"ParameterNum":        {TargetField: "parameter_num"},
			"TemplateCategory":    {TargetField: "template_category"},
			"TemplateSource":      {TargetField: "template_source"},
			"UpdateTime":          {TargetField: "update_time"},
			"CreateTime":          {TargetField: "create_time"},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlParameterTemplateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyParameterTemplate",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam:    &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				tid := d.Get("template_id")
				if tid == nil || tid == "" {
					tid = d.Id()
				}
				(*call.SdkParam)["TemplateId"] = tid
				(*call.SdkParam)["TemplateName"] = d.Get("template_name")
				if desc, ok := d.GetOk("template_desc"); ok {
					(*call.SdkParam)["TemplateDesc"] = desc
				}
				params := make([]map[string]interface{}, 0)
				if v, ok := d.GetOk("template_params"); ok && v != nil {
					if set, okSet := v.(*schema.Set); okSet {
						for _, it := range set.List() {
							m, _ := it.(map[string]interface{})
							name, _ := m["name"].(string)
							value, _ := m["value"].(string)
							if name == "" {
								continue
							}
							params = append(params, map[string]interface{}{
								"Name":  name,
								"Value": value,
							})
						}
					} else if arr, okArr := v.([]interface{}); okArr {
						for _, it := range arr {
							m, _ := it.(map[string]interface{})
							name, _ := m["name"].(string)
							value, _ := m["value"].(string)
							if name == "" {
								continue
							}
							params = append(params, map[string]interface{}{
								"Name":  name,
								"Value": value,
							})
						}
					}
				}
				if len(params) == 0 {
					current, err := s.ReadResource(d, d.Id())
					if err != nil {
						return false, err
					}
					if tp, ok := current["TemplateParams"].([]interface{}); ok {
						for _, it := range tp {
							m, _ := it.(map[string]interface{})
							name, _ := m["Name"].(string)
							value, _ := m["Value"].(string)
							if name == "" {
								continue
							}
							params = append(params, map[string]interface{}{
								"Name":  name,
								"Value": value,
							})
						}
					}
				}
				if len(params) > 0 {
					(*call.SdkParam)["TemplateParams"] = params
				}
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

func (s *VolcengineRdsPostgresqlParameterTemplateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteParameterTemplate",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["TemplateId"] = d.Get("template_id")
				return true, nil
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

func (s *VolcengineRdsPostgresqlParameterTemplateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"template_category":     {TargetField: "TemplateCategory"},
			"template_type":         {TargetField: "TemplateType"},
			"template_type_version": {TargetField: "TemplateTypeVersion"},
			"template_source":       {TargetField: "TemplateSource"},
		},
		NameField:    "TemplateName",
		IdField:      "TemplateId",
		CollectField: "template_infos",
		ResponseConverts: map[string]ve.ResponseConvert{
			"TemplateId":          {TargetField: "template_id", KeepDefault: true},
			"TemplateName":        {TargetField: "template_name"},
			"TemplateType":        {TargetField: "template_type"},
			"TemplateTypeVersion": {TargetField: "template_type_version"},
			"TemplateDesc":        {TargetField: "template_desc"},
			"TemplateCategory":    {TargetField: "template_category"},
			"TemplateSource":      {TargetField: "template_source"},
			"AccountId":           {TargetField: "account_id"},
			"CreateTime":          {TargetField: "create_time"},
			"UpdateTime":          {TargetField: "update_time"},
			"ParameterNum":        {TargetField: "parameter_num"},
			"NeedRestart":         {TargetField: "need_restart"},
			"TemplateParams":      {TargetField: "template_params"},
		},
	}
}

func (s *VolcengineRdsPostgresqlParameterTemplateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

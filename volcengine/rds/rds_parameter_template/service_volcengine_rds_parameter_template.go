package rds_parameter_template

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsParameterTemplateService struct {
	Client *volc.SdkClient
}

func NewRdsParameterTemplateService(c *volc.SdkClient) *VolcengineRdsParameterTemplateService {
	return &VolcengineRdsParameterTemplateService{
		Client: c,
	}
}

func (s *VolcengineRdsParameterTemplateService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsParameterTemplateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp                 *map[string]interface{}
		results              interface{}
		ok                   bool
		rdsParameterTemplate map[string]interface{}
	)
	action := "ListParameterTemplates"
	logger.Debug(logger.ReqFormat, action, m)
	if m == nil {
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
		if err != nil {
			return data, err
		}
	} else {
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
		if err != nil {
			return data, err
		}
	}

	results, err = volc.ObtainSdkValue("Result.Datas", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.Datas is not Slice")
	}

	for _, v := range data {
		if rdsParameterTemplate, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else {
			// query rds connection info
			templateInfoResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeParameterTemplate"), &map[string]interface{}{
				"TemplateId": rdsParameterTemplate["TemplateId"],
			})
			if err != nil {
				logger.Info("DescribeParameterTemplate error:", err)
				continue
			}
			templateInfo, err := volc.ObtainSdkValue("Result.TemplateInfo", *templateInfoResp)
			if err != nil {
				logger.Info("ObtainSdkValue Result.TemplateInfo error:", err)
				continue
			}
			if templateInfo != nil {
				if templateInfoMap, ok := templateInfo.(map[string]interface{}); ok {
					rdsParameterTemplate["TemplateParams"] = templateInfoMap["TemplateParams"]
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineRdsParameterTemplateService) ReadResource(resourceData *schema.ResourceData, rdsParameterTemplateId string) (data map[string]interface{}, err error) {
	var (
		ok bool
	)
	if rdsParameterTemplateId == "" {
		rdsParameterTemplateId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"TemplateId": rdsParameterTemplateId,
	}
	action := "DescribeParameterTemplate"
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return nil, err
	}
	result, err := volc.ObtainSdkValue("Result.TemplateInfo", *resp)
	if err != nil {
		return nil, err
	}

	if data, ok = result.(map[string]interface{}); !ok {
		return nil, errors.New("Result.TemplateInfo is not map")
	}

	return data, nil
}

func (s *VolcengineRdsParameterTemplateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineRdsParameterTemplateService) WithResourceResponseHandlers(rdsParameterTemplate map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return rdsParameterTemplate, nil, nil
	}
	return []volc.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsParameterTemplateService) ReadResourceByTemplateName(templateName string) (map[string]interface{}, error) {
	results, err := s.ReadResources(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	for _, v := range results {
		if rdsParameterTemplate, ok := v.(map[string]interface{}); !ok {
			return nil, errors.New("Value is not map ")
		} else {
			if rdsParameterTemplate["TemplateName"] == templateName {
				return rdsParameterTemplate, nil
			}
		}
	}

	return nil, errors.New("parameter template not found")
}

func (s *VolcengineRdsParameterTemplateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateParameterTemplate",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertAll,
			Convert: map[string]volc.RequestConvert{
				"template_params": {
					ConvertType: volc.ConvertJsonObjectArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建rdsParameterTemplate
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				// 2018 CreateParameterTemplate 接口未返回 TemplateId，需要通过唯一TemplateName查询
				parameterTemplate, err := s.ReadResourceByTemplateName(d.Get("template_name").(string))
				if err != nil {
					return err
				}
				d.SetId(parameterTemplate["TemplateId"].(string))
				return nil
			},
		},
	}
	return []volc.Callback{callback}

}

func (s *VolcengineRdsParameterTemplateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "ModifyParameterTemplate",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TemplateId": resourceData.Id(),
			},
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				(*call.SdkParam)["TemplateName"] = d.Get("template_name")
				(*call.SdkParam)["TemplateDesc"] = d.Get("template_desc")
				templateParams := d.Get("template_params")
				templateParamsReq := make([]map[string]interface{}, 0)
				if templateParamsList, ok := templateParams.([]interface{}); !ok {
					return false, errors.New("template_params is not array")
				} else {
					for _, v := range templateParamsList {
						templateParam := v.(map[string]interface{})
						templateParamsReq = append(templateParamsReq, map[string]interface{}{
							"Name":         templateParam["name"],
							"RunningValue": templateParam["running_value"],
						})
					}
				}

				(*call.SdkParam)["TemplateParams"] = templateParamsReq
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//修改rdsParameterTemplate
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsParameterTemplateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteParameterTemplate",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TemplateId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//删除RdsParameterTemplate
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if volc.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rds parameter template on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsParameterTemplateService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "TemplateName",
		IdField:      "TemplateId",
		CollectField: "rds_parameter_templates",
		ResponseConverts: map[string]volc.ResponseConvert{
			"TemplateId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineRdsParameterTemplateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2018-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}

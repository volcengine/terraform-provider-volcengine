package waf_instance_ctl

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineWafInstanceCtlService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewWafInstanceCtlService(c *ve.SdkClient) *VolcengineWafInstanceCtlService {
	return &VolcengineWafInstanceCtlService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineWafInstanceCtlService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineWafInstanceCtlService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineWafInstanceCtlService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Region":      id,
		"ProjectName": resourceData.Get("project_name"),
	}
	client := s.Client.UniversalClient
	action := "GetInstanceCtl"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = client.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}

	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = results.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}

	return data, err
}

func (s *VolcengineWafInstanceCtlService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineWafInstanceCtlService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateInstanceCtl",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Region"] = client.Region
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(client.Region)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineWafInstanceCtlService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineWafInstanceCtlService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateInstanceCtl",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"project_name": {
					TargetField: "ProjectName",
					ForceGet:    true,
				},
				"allow_enable": {
					TargetField: "AllowEnable",
					ForceGet:    true,
				},
				"block_enable": {
					TargetField: "BlockEnable",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Region"] = client.Region
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

func (s *VolcengineWafInstanceCtlService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	//callback := ve.Callback{
	//	Call: ve.SdkCall{
	//		Action:      "UpdateInstanceCtl",
	//		ConvertMode: ve.RequestConvertIgnore,
	//		ContentType: ve.ContentTypeJson,
	//		BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
	//			(*call.SdkParam)["Region"] = client.Region
	//			(*call.SdkParam)["AllowEnable"] = 0
	//			(*call.SdkParam)["BlockEnable"] = 0
	//			return true, nil
	//		},
	//		ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
	//			logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
	//			return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
	//		},
	//		AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
	//			return s.checkResourceUtilRemoved(d, 5*time.Minute)
	//		},
	//		CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
	//			//出现错误后重试
	//			return resource.Retry(5*time.Minute, func() *resource.RetryError {
	//				_, callErr := s.ReadResource(d, "")
	//				if callErr != nil {
	//					if ve.ResourceNotFoundError(callErr) {
	//						return nil
	//					} else {
	//						return resource.NonRetryableError(fmt.Errorf("error on  reading waf domain on delete %q, %w", d.Id(), callErr))
	//					}
	//				}
	//				_, callErr = call.ExecuteCall(d, client, call)
	//				if callErr == nil {
	//					return nil
	//				}
	//				return resource.RetryableError(callErr)
	//			})
	//		},
	//	},
	//}
	logger.Debug(logger.ReqFormat, "RemoveResource", "Remove only from tf management")
	return []ve.Callback{}
}

func (s *VolcengineWafInstanceCtlService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineWafInstanceCtlService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "waf",
		Version:     "2023-12-25",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

func (s *VolcengineWafInstanceCtlService) checkResourceUtilRemoved(d *schema.ResourceData, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		instanceCtl, _ := s.ReadResource(d, d.Id())
		logger.Debug(logger.RespFormat, "instanceCtl", instanceCtl)

		// 能查询成功代表还在删除中，重试
		allowEnableInt, ok := instanceCtl["AllowEnable"].(float64)
		if !ok {
			return resource.NonRetryableError(fmt.Errorf("AllowEnable is not float64"))
		}
		blockEnable, ok := instanceCtl["BlockEnable"].(float64)
		if !ok {
			return resource.NonRetryableError(fmt.Errorf("BlockEnable is not float64"))
		}
		if int(allowEnableInt) == 1 || int(blockEnable) == 1 {
			return resource.RetryableError(fmt.Errorf("resource still in removing status "))
		} else {
			if int(allowEnableInt) == 0 && int(blockEnable) == 0 {
				return nil
			} else {
				return resource.NonRetryableError(fmt.Errorf("instanceCtl status is not disable "))
			}
		}
	})
}

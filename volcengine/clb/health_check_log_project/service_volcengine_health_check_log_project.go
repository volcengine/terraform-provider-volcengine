package health_check_log_project

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineHealthCheckLogProjectService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewHealthCheckLogProjectService(c *ve.SdkClient) *VolcengineHealthCheckLogProjectService {
	return &VolcengineHealthCheckLogProjectService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineHealthCheckLogProjectService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineHealthCheckLogProjectService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	action := "DescribeHealthCheckLogProjectAttributes"

	bytes, _ := json.Marshal(m)
	logger.Debug(logger.ReqFormat, action, string(bytes))

	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
	if err != nil {
		return data, err
	}

	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, string(respBytes))

	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}

	if results == nil {
		results = []interface{}{}
	}
	// 单个结果，包装成列表返回
	if resultMap, ok := results.(map[string]interface{}); ok {
		data = []interface{}{resultMap}
	}
	return data, err
}

func (s *VolcengineHealthCheckLogProjectService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "DescribeHealthCheckLogProjectAttributes"
	logger.Debug(logger.ReqFormat, action, id)

	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
	if err != nil {
		return data, err
	}

	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, string(respBytes))

	// 从响应中提取Result
	results, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}

	if results == nil {
		return data, fmt.Errorf("health_check_log_project %s not exist", id)
	}

	// 处理Result - 它应该是一个map包含LogProjectId
	if resultMap, ok := results.(map[string]interface{}); ok {
		data = make(map[string]interface{})
		data["log_project_id"] = resultMap["LogProjectId"]
		data["id"] = resultMap["LogProjectId"]
	} else {
		return data, fmt.Errorf("invalid result format for health_check_log_project %s", id)
	}

	if data["log_project_id"] == nil || data["log_project_id"].(string) == "" {
		return data, fmt.Errorf("health_check_log_project %s not exist", id)
	}
	return data, err
}

func (s *VolcengineHealthCheckLogProjectService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateHealthCheckLogProject",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam:    &map[string]interface{}{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.LogProjectId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineHealthCheckLogProjectService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteHealthCheckLogProject",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam:    &map[string]interface{}{},
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

func (s *VolcengineHealthCheckLogProjectService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineHealthCheckLogProjectService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d map[string]interface{}
			)
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			return d, "Available", err
		},
	}
}

func (s *VolcengineHealthCheckLogProjectService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineHealthCheckLogProjectService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "LogProjectId",
		IdField:      "LogProjectId",
		CollectField: "health_check_log_projects",
		ResponseConverts: map[string]ve.ResponseConvert{
			"LogProjectId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineHealthCheckLogProjectService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

package iam_login_profile

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamLoginProfileService struct {
	Client *ve.SdkClient
}

func NewIamLoginProfileService(c *ve.SdkClient) *VolcengineIamLoginProfileService {
	return &VolcengineIamLoginProfileService{
		Client: c,
	}
}

func (s *VolcengineIamLoginProfileService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamLoginProfileService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineIamLoginProfileService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	condition := map[string]interface{}{"UserName": id}
	action := "GetLoginProfile"
	logger.Debug(logger.ReqFormat, action, condition)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)
	result, err = ve.ObtainSdkValue("Result.LoginProfile", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("login profile %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineIamLoginProfileService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamLoginProfileService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		delete(v, "Password")
		return v, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineIamLoginProfileService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateLoginProfile",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				time.Sleep(5 * time.Second)
				d.SetId(d.Get("user_name").(string))
				return nil
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineIamLoginProfileService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "UpdateLoginProfile",
			ConvertMode:    ve.RequestConvertAll,
			RequestIdField: "UserName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				time.Sleep(5 * time.Second)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamLoginProfileService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "DeleteLoginProfile",
			ConvertMode:    ve.RequestConvertIgnore,
			RequestIdField: "UserName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineIamLoginProfileService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineIamLoginProfileService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

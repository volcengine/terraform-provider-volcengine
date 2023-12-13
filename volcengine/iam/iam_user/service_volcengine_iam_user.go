package iam_user

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamUserService struct {
	Client *ve.SdkClient
}

func NewIamUserService(c *ve.SdkClient) *VolcengineIamUserService {
	return &VolcengineIamUserService{
		Client: c,
	}
}

func (s *VolcengineIamUserService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamUserService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		nameSet = make(map[string]bool)
	)
	if _, ok = m["UserNames.1"]; ok {
		i := 1
		for {
			filed := fmt.Sprintf("UserNames.%d", i)
			tmpId, ok := m[filed]
			if !ok {
				break
			}
			nameSet[tmpId.(string)] = true
			i++
			delete(m, filed)
		}
	}
	cens, err := ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "ListUsers"
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
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.UserMetadata", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.UserMetadata is not Slice")
		}
		return data, err
	})
	if err != nil || len(nameSet) == 0 {
		return cens, err
	}

	res := make([]interface{}, 0)
	for _, cen := range cens {
		if !nameSet[cen.(map[string]interface{})["UserName"].(string)] {
			continue
		}
		res = append(res, cen)
	}
	return res, nil
}

func (s *VolcengineIamUserService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"UserNames.1": id,
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
		return data, fmt.Errorf("user %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineIamUserService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamUserService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return v, map[string]ve.ResponseConvert{
			"AccountId": {
				TargetField: "account_id",
				Convert: func(i interface{}) interface{} {
					return strconv.FormatFloat(i.(float64), 'f', 0, 64)
				},
			},
			"Id": {
				TargetField: "user_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineIamUserService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateUser",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				d.SetId(d.Get("user_name").(string))
				return nil
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineIamUserService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateUser",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"user_name": {
					TargetField: "NewUserName",
					ConvertType: ve.ConvertDefault,
				},
				"display_name": {
					TargetField: "NewDisplayName",
					ConvertType: ve.ConvertDefault,
					Convert:     defaultConvert,
				},
				"mobile_phone": {
					TargetField: "NewMobilePhone",
				},
				"email": {
					TargetField: "NewEmail",
					Convert:     defaultConvert,
				},
				"description": {
					TargetField: "NewDescription",
					ConvertType: ve.ConvertDefault,
					Convert:     defaultConvert,
				},
			},
			RequestIdField: "UserName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				if d.HasChange("user_name") {
					d.SetId(d.Get("user_name").(string))
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamUserService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "DeleteUser",
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

func (s *VolcengineIamUserService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"user_names": {
				TargetField: "UserNames",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "UserName",
		IdField:      "UserName",
		CollectField: "users",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Id": {
				TargetField: "user_id",
			},
			"AccountId": {
				TargetField: "account_id",
				Convert: func(i interface{}) interface{} {
					return strconv.FormatFloat(i.(float64), 'f', 0, 64)
				},
			},
		},
	}
}

func (s *VolcengineIamUserService) ReadResourceId(id string) string {
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

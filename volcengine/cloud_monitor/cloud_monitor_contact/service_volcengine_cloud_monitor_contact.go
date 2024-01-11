package cloud_monitor_contact

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCloudMonitorContactService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineCloudMonitorContactService {
	return &VolcengineCloudMonitorContactService{
		Client: c,
	}
}

func (s *VolcengineCloudMonitorContactService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudMonitorContactService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListContactsByIds"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return nil, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return nil, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.Data", *resp)

		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Data is not Slice")
		}

		return data, err
	})
}

func (s *VolcengineCloudMonitorContactService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Ids": []string{id},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Contact %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCloudMonitorContactService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineCloudMonitorContactService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	// 防止读取不一致
	phone, exist := data["Phone"]
	if exist && phone == "-" {
		delete(data, "Phone")
	}
	email, exist := data["Email"]
	if exist && email == "-" {
		delete(data, "Email")
	}
	return []ve.ResourceResponseHandler{}

}

func (s *VolcengineCloudMonitorContactService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateContacts",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				data, err := ve.ObtainSdkValue("Result.Data", *resp)
				if err != nil {
					return err
				}
				if len(data.([]interface{})) == 0 {
					return fmt.Errorf("create contact failed")
				}
				d.SetId(data.([]interface{})[0].(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudMonitorContactService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateContacts",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				(*call.SdkParam)["Name"] = d.Get("name")
				(*call.SdkParam)["Email"] = d.Get("email")

				phone := d.Get("phone")
				if len(phone.(string)) > 0 {
					(*call.SdkParam)["Phone"] = phone
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, callback)
	return callbacks
}

func (s *VolcengineCloudMonitorContactService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteContactsByIds",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Ids": []string{resourceData.Id()},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudMonitorContactService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Ids",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "contacts",
	}
}

func (s *VolcengineCloudMonitorContactService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Volc_Observe",
		Version:     "2018-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

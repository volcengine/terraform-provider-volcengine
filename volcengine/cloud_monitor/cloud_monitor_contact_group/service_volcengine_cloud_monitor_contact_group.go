package cloud_monitor_contact_group

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

type VolcengineCloudMonitorContactGroupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudMonitorContactGroupService(c *ve.SdkClient) *VolcengineCloudMonitorContactGroupService {
	return &VolcengineCloudMonitorContactGroupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudMonitorContactGroupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudMonitorContactGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListContactGroups"

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

func (s *VolcengineCloudMonitorContactGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		temp    map[string]interface{}
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
	for _, v := range results {
		if temp, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else if temp["Id"].(string) == id {
			data = temp
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cloud_monitor_contact_group %s not exist ", id)
	}
	if _, ok = data["Contacts"]; ok {
		idList := make([]string, 0)
		for _, contact := range data["Contacts"].([]interface{}) {
			if contactMap, ok := contact.(map[string]interface{}); ok {
				idList = append(idList, contactMap["Id"].(string))
			}
		}
		data["ContactsIdList"] = idList
	}
	return data, err
}

func (s *VolcengineCloudMonitorContactGroupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineCloudMonitorContactGroupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateContactGroup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"contacts_id_list": {
					TargetField: "ContactsIdList",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				data, _ := ve.ObtainSdkValue("Result.Data", *resp)
				dataArr, ok := data.([]interface{})
				if !ok || len(dataArr) < 1 {
					return fmt.Errorf("create contact group failed, error id")
				}
				d.SetId(dataArr[0].(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineCloudMonitorContactGroupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudMonitorContactGroupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateContactGroup",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
					ForceGet:    true,
				},
				"description": {
					TargetField: "Description",
				},
				"contacts_id_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
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
	callbacks = append(callbacks, callback)
	if resourceData.HasChange("contacts_id_list") {
		contactCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateContactGroupWithContacts",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["Id"] = d.Id()
					list, ok := d.GetOk("contacts_id_list")
					if !ok {
						(*call.SdkParam)["ContactsIdList"] = []string{}
					} else {
						(*call.SdkParam)["ContactsIdList"] = list.(*schema.Set).List()
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
		callbacks = append(callbacks, contactCallback)
	}
	return callbacks
}

func (s *VolcengineCloudMonitorContactGroupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteContactGroupByIds",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Ids": []string{resourceData.Id()},
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

func (s *VolcengineCloudMonitorContactGroupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "groups",
	}
}

func (s *VolcengineCloudMonitorContactGroupService) ReadResourceId(id string) string {
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

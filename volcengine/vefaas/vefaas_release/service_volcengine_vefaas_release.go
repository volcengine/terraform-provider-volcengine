package vefaas_release

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVefaasReleaseService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVefaasReleaseService(c *ve.SdkClient) *VolcengineVefaasReleaseService {
	return &VolcengineVefaasReleaseService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVefaasReleaseService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVefaasReleaseService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListReleaseRecords"

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
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVefaasReleaseService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("format of vefaas funtion id release record id is invalid,%s", id)
	}
	functionId := parts[0]

	if err != nil {
		return data, fmt.Errorf(" revisionNumber cannot convert to int ")
	}

	req := map[string]interface{}{
		"FunctionId": functionId,
	}
	action := "GetReleaseStatus"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	queryRelease, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	data, ok := queryRelease.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("release %s not exist", id)
	}

	return data, err
}

func (s *VolcengineVefaasReleaseService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineVefaasReleaseService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "Release",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				functionId, _ := ve.ObtainSdkValue("Result.FunctionId", *resp)
				releaseRecordId, _ := ve.ObtainSdkValue("Result.ReleaseRecordId", *resp)

				d.SetId(fmt.Sprintf("%s:%s", functionId, releaseRecordId))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVefaasReleaseService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVefaasReleaseService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRelease",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("format of vefaas funtion id release record id is invalid,%s", d.Id())
				}
				functionId := parts[0]
				(*call.SdkParam)["FunctionId"] = functionId
				(*call.SdkParam)["TargetTrafficWeight"] = d.Get("target_traffic_weight")
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

func (s *VolcengineVefaasReleaseService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	canRollBack, status, functionId, releaseRecordId, err := s.checkReleaseRecordStatusCanRollBack(resourceData)
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}

	if !canRollBack {
		logger.Info("vefass function release %s status is %s, not support abort ", functionId, status)
		return []ve.Callback{}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AbortRelease",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"FunctionId": functionId,
				"Async":      true,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				stateConf := s.buildReleaseRecordStateConf(resourceData, functionId, releaseRecordId)
				_, err = stateConf.WaitForState()
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVefaasReleaseService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts:  map[string]ve.RequestConvert{},
		IdField:          "FunctionId",
		CollectField:     "items",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineVefaasReleaseService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vefaas",
		Version:     "2024-06-06",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func (s *VolcengineVefaasReleaseService) getReleaseRecordsStatus(functionId, releaseRecordId string) (data map[string]interface{}, err error) {
	var (
		items []interface{}
	)
	type Filter struct {
		Name   string   `json:"Name"`
		Values []string `json:"Values"`
	}
	filters := []Filter{
		{
			Name:   "Id",
			Values: []string{releaseRecordId},
		},
	}
	req := map[string]interface{}{
		"FunctionId": functionId,
		"Filters":    filters,
	}
	items, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, ele := range items {
		item, ok := ele.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf(" item is not Map ")
		}
		if item["Id"].(string) == releaseRecordId {
			data = item
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("vefass function release record %s not exist ", functionId)
	}
	return data, nil
}

func (s *VolcengineVefaasReleaseService) buildReleaseRecordStateConf(resourceData *schema.ResourceData, functionId, releaseRecordId string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     []string{"rollbacked"},
		Timeout:    resourceData.Timeout(schema.TimeoutDelete),
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d      map[string]interface{}
				status interface{}
			)
			d, err = s.getReleaseRecordsStatus(functionId, releaseRecordId)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)

			if err != nil {
				return nil, "", err
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVefaasReleaseService) checkReleaseRecordStatusCanRollBack(resourceData *schema.ResourceData) (bool, string, string, string, error) {
	var (
		status          interface{}
		functionId      string
		releaseRecordId string
	)
	parts := strings.Split(resourceData.Id(), ":")
	if len(parts) != 2 {
		return false, "", "", "", fmt.Errorf("format of vefaas funtion id release record id is invalid,%s", resourceData.Id())
	}
	functionId = parts[0]
	releaseRecordId = parts[1]

	statusData, err := s.getReleaseRecordsStatus(functionId, releaseRecordId)
	if err != nil {
		return false, "", "", "", err
	}
	status, err = ve.ObtainSdkValue("Status", statusData)
	if err != nil {
		return false, "", "", "", err
	}

	if status.(string) == "rolling" {
		return true, status.(string), functionId, releaseRecordId, nil
	} else {
		return false, status.(string), functionId, releaseRecordId, nil
	}

}

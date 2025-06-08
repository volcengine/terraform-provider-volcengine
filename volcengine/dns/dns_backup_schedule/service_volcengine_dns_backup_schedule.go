package dns_backup_schedule

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

type VolcengineDnsBackupScheduleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewDnsBackupScheduleService(c *ve.SdkClient) *VolcengineDnsBackupScheduleService {
	return &VolcengineDnsBackupScheduleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineDnsBackupScheduleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineDnsBackupScheduleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineDnsBackupScheduleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		zidInt  int
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	zidInt, err = strconv.Atoi(id)
	if err != nil {
		return data, fmt.Errorf(" ZID cannot convert to int ")
	}
	req := map[string]interface{}{
		"ZID": zidInt,
	}
	client := s.Client.UniversalClient
	action := "QueryBackupSchedule"
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

	if data["Schedule"].(float64) == 0 {
		return data, fmt.Errorf("dns_record_schedule %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineDnsBackupScheduleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
	}
}

func (s *VolcengineDnsBackupScheduleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateBackupSchedule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"zid": {
					TargetField: "ZID",
				},
				"schedule": {
					TargetField: "Schedule",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(strconv.Itoa((*call.SdkParam)["ZID"].(int)))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineDnsBackupScheduleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"Schedule": {
				TargetField: "schedule",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineDnsBackupScheduleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateBackupSchedule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"zid": {
					TargetField: "ZID",
				},
				"schedule": {
					TargetField: "Schedule",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				zid, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, fmt.Errorf(" ZID cannot convert to int ")
				}
				(*call.SdkParam)["ZID"] = zid
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineDnsBackupScheduleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateBackupSchedule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"zid": {
					TargetField: "ZID",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				zid, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, fmt.Errorf(" ZID cannot convert to int ")
				}
				(*call.SdkParam)["ZID"] = zid
				(*call.SdkParam)["Schedule"] = 0
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineDnsBackupScheduleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineDnsBackupScheduleService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "DNS",
		Version:     "2018-08-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "DNS",
		Version:     "2018-08-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

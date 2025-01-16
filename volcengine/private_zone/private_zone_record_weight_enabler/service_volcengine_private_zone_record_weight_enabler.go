package private_zone_record_weight_enabler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcenginePrivateZoneRecordWeightEnablerService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewPrivateZoneRecordWeightEnablerService(c *ve.SdkClient) *VolcenginePrivateZoneRecordWeightEnablerService {
	return &VolcenginePrivateZoneRecordWeightEnablerService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRecordSets"

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

		results, err = ve.ObtainSdkValue("Result.RecordSets", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RecordSets is not Slice")
		}
		return data, err
	})
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid record weight enabler Id: %s", id)
	}

	zid, err := strconv.Atoi(ids[0])
	if err != nil {
		return data, fmt.Errorf(" ZID cannot convert to int ")
	}

	req := map[string]interface{}{
		"ZID":         zid,
		"RecordSetID": ids[1],
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
		return data, fmt.Errorf("private_zone_record_set %s not exist ", id)
	}
	return data, err
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcenginePrivateZoneRecordWeightEnablerService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ID": {
				TargetField: "record_set_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRecordSet",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"record_set_id": {
					TargetField: "RecordSetID",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.ID", *resp)
				zid := d.Get("zid").(int)
				d.SetId(strconv.Itoa(zid) + ":" + id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRecordSet",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"weight_enabled": {
					TargetField: "WeightEnabled",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RecordSetID"] = d.Get("record_set_id").(string)
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

func (s *VolcenginePrivateZoneRecordWeightEnablerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcenginePrivateZoneRecordWeightEnablerService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "private_zone",
		Version:     "2022-06-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
		Action:      actionName,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "private_zone",
		Version:     "2022-06-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		RegionType:  ve.Global,
		Action:      actionName,
	}
}

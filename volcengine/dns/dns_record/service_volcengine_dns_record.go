package dns_record

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

type VolcengineDnsRecordService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewDnsRecordService(c *ve.SdkClient) *VolcengineDnsRecordService {
	return &VolcengineDnsRecordService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineDnsRecordService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineDnsRecordService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRecords"
		logger.Debug(logger.ReqFormat, action, condition)
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
		results, err = ve.ObtainSdkValue("Result.Records", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Records is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineDnsRecordService) ReadResource(resourceData *schema.ResourceData, id string) (result map[string]interface{}, err error) {
	var (
		data     map[string]interface{}
		results  []interface{}
		ok       bool
		zid      string
		recordId string
		zidInt   int
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("format of dns record resource id is invalid,%s", id)
	}
	zid = parts[0]
	recordId = parts[1]
	zidInt, err = strconv.Atoi(zid)
	if err != nil {
		return data, fmt.Errorf(" ZID cannot convert to int ")
	}
	req := map[string]interface{}{
		"ZID": zidInt,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}

		if data["RecordID"] == recordId {
			result = data
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("dns_record %s not exist ", id)
	}
	return result, err
}

func (s *VolcengineDnsRecordService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineDnsRecordService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRecord",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"zid": {
					TargetField: "ZID",
				},
				"Host": {
					TargetField: "host",
				},
				"Type": {
					TargetField: "type",
				},
				"Line": {
					TargetField: "line",
				},
				"Value": {
					TargetField: "value",
				},
				"Remark": {
					TargetField: "remark",
				},
				"Weight": {
					TargetField: "weight",
				},
				"TTL": {
					TargetField: "ttl",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.RecordID", *resp)
				//d.SetId(id.(string))
				d.SetId(fmt.Sprintf("%s:%s", strconv.Itoa(d.Get("zid").(int)), id.(string)))

				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineDnsRecordService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"RecordID": {
				TargetField: "record_id",
			},
			"PQDN": {
				TargetField: "pqdn",
			},
			"Host": {
				TargetField: "host",
			},
			"Type": {
				TargetField: "type",
			},
			"TTL": {
				TargetField: "ttl",
			},
			"Line": {
				TargetField: "line",
			},
			"Value": {
				TargetField: "value",
			},
			"Weight": {
				TargetField: "weight",
			},
			"Enable": {
				TargetField: "enable",
			},
			"RecordSetID": {
				TargetField: "record_set_id",
			},
			"Remark": {
				TargetField: "remark",
			},
			"Operators": {
				TargetField: "operators",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineDnsRecordService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	if resourceData.HasChanges("type", "project_name", "value", "remark", "ttl", "weight") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateRecord",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"record_id": {
						TargetField: "RecordID",
					},
					"host": {
						TargetField: "Host",
					},
					"line": {
						TargetField: "Line",
					},
					"type": {
						TargetField: "Type",
					},
					"value": {
						TargetField: "Value",
					},
					"remark": {
						TargetField: "Remark",
					},
					"ttl": {
						TargetField: "TTL",
					},
					"weight": {
						TargetField: "Weight",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if d.Get("record_id").(string) == "" {
						return false, errors.New("record id cannot be empty")
					}
					(*call.SdkParam)["RecordID"] = d.Get("record_id")
					(*call.SdkParam)["Host"] = d.Get("host")
					(*call.SdkParam)["Line"] = d.Get("line")

					//if resourceData.HasChange("project_name") {
					//	(*call.SdkParam)["ProjectName"] = d.Get("project_name")
					//}
					if resourceData.HasChange("type") {
						(*call.SdkParam)["Type"] = d.Get("type")
					}
					if resourceData.HasChange("value") {
						(*call.SdkParam)["Value"] = d.Get("value")
					}
					if resourceData.HasChange("remark") {
						(*call.SdkParam)["Remark"] = d.Get("remark")
					}
					if resourceData.HasChange("ttl") {
						(*call.SdkParam)["TTL"] = d.Get("ttl")
					}
					if resourceData.HasChange("weight") {
						(*call.SdkParam)["Weight"] = d.Get("weight")
					}
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
		callbacks = append(callbacks, callback)
	}

	return callbacks
}

func (s *VolcengineDnsRecordService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRecord",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"RecordID": resourceData.Get("record_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading reocrd on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{callback}
}

func (s *VolcengineDnsRecordService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"zid": {
				TargetField: "ZID",
			},
			"host": {
				TargetField: "Host",
			},
			"line": {
				TargetField: "Line",
			},
			"type": {
				TargetField: "Type",
			},
			"value": {
				TargetField: "Value",
			},
		},
		//NameField:    "PQDN",
		IdField:      "RecordID",
		CollectField: "records",
		ResponseConverts: map[string]ve.ResponseConvert{
			"RecordID": {
				TargetField: "record_id",
			},
			"PQDN": {
				TargetField: "pqdn",
			},
			"Host": {
				TargetField: "host",
			},
			"Type": {
				TargetField: "type",
			},
			"TTL": {
				TargetField: "ttl",
			},
			"Line": {
				TargetField: "line",
			},
			"Value": {
				TargetField: "value",
			},
			"Weight": {
				TargetField: "weight",
			},
			"Enable": {
				TargetField: "enable",
			},
			"RecordSetID": {
				TargetField: "record_set_id",
			},
			"Remark": {
				TargetField: "remark",
			},
			"Operators": {
				TargetField: "operators",
			},
		},
	}
}

func (s *VolcengineDnsRecordService) ReadResourceId(id string) string {
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

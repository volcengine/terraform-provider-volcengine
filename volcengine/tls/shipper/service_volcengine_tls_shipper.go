package shipper

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

type VolcengineShipperService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewShipperService(c *ve.SdkClient) *VolcengineShipperService {
	return &VolcengineShipperService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineShipperService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineShipperService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeShippers"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("RESPONSE.Shippers", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("RESPONSE.Shippers is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineShipperService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ShipperId": id,
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
		return data, fmt.Errorf("shipper %s not exist ", id)
	}
	if contentInfos, contentInfoExist := data["ContentInfo"]; contentInfoExist {
		if csvInfo, csvInfoExist := contentInfos.(map[string]interface{})["CsvInfo"]; csvInfoExist {
			contentInfos.(map[string]interface{})["CsvInfo"] = []interface{}{
				csvInfo,
			}
		}
	}
	if contentInfos, contentInfoExist := data["ContentInfo"]; contentInfoExist {
		if jsonInfo, jsonInfoExist := contentInfos.(map[string]interface{})["JsonInfo"]; jsonInfoExist {
			contentInfos.(map[string]interface{})["JsonInfo"] = []interface{}{
				jsonInfo,
			}
		}
	}
	// ShipperEndTime ShipperStartTime 不做diff
	delete(data, "ShipperEndTime")
	delete(data, "ShipperStartTime")
	return data, err
}

func (s *VolcengineShipperService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineShipperService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateShipper",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"content_info": {
					TargetField: "ContentInfo",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"csv_info": {
							TargetField: "CsvInfo",
							ConvertType: ve.ConvertJsonObject,
						},
						"json_info": {
							TargetField: "JsonInfo",
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"kafka_shipper_info": {
					TargetField: "KafkaShipperInfo",
					ConvertType: ve.ConvertJsonObject,
				},
				"tos_shipper_info": {
					TargetField: "TosShipperInfo",
					ConvertType: ve.ConvertJsonObject,
				},
				"role_trn": {
					TargetField: "RoleTrn",
					ConvertType: ve.ConvertDefault,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.ShipperId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineShipperService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineShipperService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyShipper",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"content_info": {
					TargetField: "ContentInfo",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"csv_info": {
							TargetField: "CsvInfo",
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
						},
						"format": {
							TargetField: "Format",
							ForceGet:    true,
						},
						"json_info": {
							TargetField: "JsonInfo",
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"kafka_shipper_info": {
					TargetField: "KafkaShipperInfo",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
				},
				"shipper_name": {
					TargetField: "ShipperName",
					ForceGet:    true,
				},
				"shipper_type": {
					TargetField: "ShipperType",
					ForceGet:    true,
				},
				"status": {
					TargetField: "Status",
					ForceGet:    true,
				},
				"tos_shipper_info": {
					TargetField: "TosShipperInfo",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
				},
				"role_trn": {
					TargetField: "RoleTrn",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ShipperId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineShipperService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteShipper",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"ShipperId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tls shipper on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineShipperService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		NameField:       "ShipperName",
		IdField:         "ShipperId",
		CollectField:    "shippers",
		ResponseConverts: map[string]ve.ResponseConvert{
			"role_trn": {
				TargetField: "RoleTrn",
			},
		},
	}
}

func (s *VolcengineShipperService) ReadResourceId(id string) string {
	return id
}

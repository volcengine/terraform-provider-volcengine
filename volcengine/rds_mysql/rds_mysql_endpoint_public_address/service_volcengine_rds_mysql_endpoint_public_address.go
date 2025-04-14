package rds_mysql_endpoint_public_address

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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_instance"
)

type VolcengineRdsMysqlEndpointPublicAddressService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsMysqlEndpointPublicAddressService(c *ve.SdkClient) *VolcengineRdsMysqlEndpointPublicAddressService {
	return &VolcengineRdsMysqlEndpointPublicAddressService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstanceDetail"

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
		results, err = ve.ObtainSdkValue("Result.Endpoints", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Endpoints is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results      []interface{}
		ok           bool
		temp         map[string]interface{}
		endpointData map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	// instanceId:endpointId:eipId
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"InstanceId": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if temp, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else {
			if endpointId, ok := temp["EndpointId"]; ok {
				if ids[1] == endpointId.(string) {
					endpointData = temp
				}
			}
		}
	}
	if len(endpointData) == 0 {
		return data, fmt.Errorf("rds_mysql_endpoint_public_address %s not exist ", id)
	}
	logger.Debug(logger.ReqFormat, "Endpoint Data", endpointData)
	addresses := endpointData["Addresses"]
	if addresses != nil {
		for _, addr := range addresses.([]interface{}) {
			if eipId, ok := addr.(map[string]interface{})["EipId"]; ok {
				if eipId.(string) == ids[2] {
					data = addr.(map[string]interface{})
				}
			}
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("rds_mysql_endpoint_public_address %s not exist ", id)
	}
	logger.Debug(logger.ReqFormat, "Address Data", data)
	return data, err
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data   map[string]interface{}
				status interface{}
			)
			// 资源异步且无状态，加假状态防止读取不到
			if err = resource.Retry(10*time.Minute, func() *resource.RetryError {
				data, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}
			status = "Success"
			return data, status.(string), err
		},
	}
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBEndpointPublicAddress",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"domain": {
					Ignore: true,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				endpointId := d.Get("endpoint_id").(string)
				eipId := d.Get("eip_id").(string)
				d.SetId(fmt.Sprintf("%s:%s:%s", instanceId, endpointId, eipId))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Success"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rds_mysql_instance.NewRdsMysqlInstanceService(s.Client): {
					ResourceId: resourceData.Get("instance_id").(string),
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
				},
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Attached"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("eip_id").(string),
				},
			},
		},
	}
	callbacks = append(callbacks, callback)
	if domain, ok := resourceData.GetOk("domain"); ok {
		modifyCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointAddress",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					(*call.SdkParam)["EndpointId"] = d.Get("endpoint_id")
					(*call.SdkParam)["NetworkType"] = "Public"
					arr := strings.Split(domain.(string), ".")
					if len(arr) < 2 {
						return false, fmt.Errorf("domain is not valid")
					}
					(*call.SdkParam)["DomainPrefix"] = arr[0]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("instance_id").(string)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rds_mysql_instance.NewRdsMysqlInstanceService(s.Client): {
						ResourceId: resourceData.Get("instance_id").(string),
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					},
				},
			},
		}
		callbacks = append(callbacks, modifyCallback)
	}
	return callbacks
}

func (VolcengineRdsMysqlEndpointPublicAddressService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if resourceData.HasChange("domain") {
		modifyCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointAddress",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					(*call.SdkParam)["EndpointId"] = d.Get("endpoint_id")
					(*call.SdkParam)["NetworkType"] = "Public"
					arr := strings.Split(d.Get("domain").(string), ".")
					if len(arr) < 2 {
						return false, fmt.Errorf("domain is not valid")
					}
					(*call.SdkParam)["DomainPrefix"] = arr[0]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("instance_id").(string)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rds_mysql_instance.NewRdsMysqlInstanceService(s.Client): {
						ResourceId: resourceData.Get("instance_id").(string),
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					},
				},
			},
		}
		callbacks = append(callbacks, modifyCallback)
	}
	return callbacks
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBEndpointPublicAddress",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": ids[0],
				"EndpointId": ids[1],
				"EipId":      ids[2],
				"Domain":     resourceData.Get("domain"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rds_mysql_instance.NewRdsMysqlInstanceService(s.Client): {
					ResourceId: resourceData.Get("instance_id").(string),
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
				},
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("eip_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsMysqlEndpointPublicAddressService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

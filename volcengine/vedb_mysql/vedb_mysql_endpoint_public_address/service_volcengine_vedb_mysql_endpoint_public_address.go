package vedb_mysql_endpoint_public_address

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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vedb_mysql/vedb_mysql_instance"
)

type VolcengineVedbMysqlEndpointPublicAddressService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVedbMysqlEndpointPublicAddressService(c *ve.SdkClient) *VolcengineVedbMysqlEndpointPublicAddressService {
	return &VolcengineVedbMysqlEndpointPublicAddressService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBEndpoint"

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

func (s *VolcengineVedbMysqlEndpointPublicAddressService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		tmpData map[string]interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 3 {
		return nil, errors.New("ids length is not correct")
	}
	req := map[string]interface{}{
		"InstanceId": ids[0],
		"EndpointId": ids[1],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if tmpData, ok = v.(map[string]interface{}); !ok {
			return tmpData, errors.New("Value is not map ")
		}
	}
	if len(tmpData) == 0 {
		return data, fmt.Errorf("vedb_mysql_allowlist_associate %s not exist ", id)
	}
	addresses := tmpData["Addresses"]
	if a, ok := addresses.([]interface{}); ok && len(addresses.([]interface{})) > 0 {
		for _, add := range a {
			addMap := add.(map[string]interface{})
			if addMap["EipId"] == ids[2] {
				data = addMap
			}
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("vedb_mysql_allowlist_associate %s not exist ", id)
	}
	data["EndpointId"] = ids[1]
	data["InstanceId"] = ids[0]
	return data, err
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("vedb_mysql_endpoint_public_address status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBEndpointPublicAddress",
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
				d.SetId(fmt.Sprint((*call.SdkParam)["InstanceId"], ":", (*call.SdkParam)["EndpointId"], ":", (*call.SdkParam)["EipId"]))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vedb_mysql_instance.NewVedbMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("instance_id").(string),
				},
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Attached"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("eip_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVedbMysqlEndpointPublicAddressService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBEndpointPublicAddress",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["EndpointId"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vedb_mysql_instance.NewVedbMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("instance_id").(string),
				},
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("eip_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineVedbMysqlEndpointPublicAddressService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vedbm",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

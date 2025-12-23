package rds_postgresql_endpoint_public_address

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	eip_address "github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	rdsPgInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance"
)

type VolcengineRdsPostgresqlEndpointPublicAddressService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlEndpointPublicAddressService(c *ve.SdkClient) *VolcengineRdsPostgresqlEndpointPublicAddressService {
	return &VolcengineRdsPostgresqlEndpointPublicAddressService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return []interface{}{}, nil
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return map[string]interface{}{}, nil
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBEndpointPublicAddress",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				endpointId := d.Get("endpoint_id").(string)
				eipId := d.Get("eip_id").(string)
				d.SetId(fmt.Sprintf("%s:%s:%s", instanceId, endpointId, eipId))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Success"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
					ResourceId: resourceData.Get("instance_id").(string),
					Target:     []string{"Running"}, Timeout: resourceData.Timeout(schema.TimeoutCreate)},
				eip_address.NewEipAddressService(s.Client): {
					ResourceId: resourceData.Get("eip_id").(string),
					Target:     []string{"Attached"}, Timeout: resourceData.Timeout(schema.TimeoutCreate)},
			},
			LockId: func(d *schema.ResourceData) string { return d.Get("instance_id").(string) },
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlEndpointPublicAddressService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	// 资源id 按照 instance_id:endpoint_id:eip_id 设置
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBEndpointPublicAddress",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": ids[0],
				"EndpointId": ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				eip_address.NewEipAddressService(s.Client): {
					ResourceId: resourceData.Get("eip_id").(string),
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
				},
			},
			LockId: func(d *schema.ResourceData) string { return d.Get("instance_id").(string) },
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlEndpointPublicAddressService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

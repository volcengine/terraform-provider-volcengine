package network_interface_attach

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNetworkInterfaceAttachService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewNetworkInterfaceAttachService(c *ve.SdkClient) *VolcengineNetworkInterfaceAttachService {
	return &VolcengineNetworkInterfaceAttachService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineNetworkInterfaceAttachService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNetworkInterfaceAttachService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineNetworkInterfaceAttachService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp               *map[string]interface{}
		results            interface{}
		deviceId           interface{}
		eniType            interface{}
		ok                 bool
		networkInterfaceId string
		targetInstanceId   string
		ids                []string
	)

	if id == "" {
		id = resourceData.Id()
	}

	ids = strings.Split(id, ":")
	networkInterfaceId = ids[0]
	targetInstanceId = ids[1]

	req := map[string]interface{}{
		"NetworkInterfaceId": networkInterfaceId,
	}
	vpc := s.Client.VpcClient
	action := "DescribeNetworkInterfaceAttributes"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = vpc.DescribeNetworkInterfaceAttributesCommon(&req)
	if err != nil {
		return data, err
	}

	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = results.(map[string]interface{}); !ok {
		return data, errors.New("value is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("network interface attributes %s not exist ", networkInterfaceId)
	}

	eniType, ok = data["Type"]
	if !ok {
		return data, errors.New("eni type not exist")
	}
	if eniType.(string) != "secondary" {
		return data, errors.New("only secondary eni support attach/detach")
	}

	deviceId, ok = data["DeviceId"]
	if !ok {
		return data, errors.New("device id not exist")
	}
	if len(deviceId.(string)) == 0 {
		return data, errors.New("not associate")
	}
	if deviceId.(string) != targetInstanceId {
		return data, fmt.Errorf("network interface %s does not bound target device. bound_instance_id %s, target_instance_id %s",
			networkInterfaceId, deviceId.(string), targetInstanceId)
	}
	return data, err
}

func (s *VolcengineNetworkInterfaceAttachService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil && !strings.Contains(err.Error(), "not associate") {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("network interface attach status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (VolcengineNetworkInterfaceAttachService) WithResourceResponseHandlers(networkInterface map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return networkInterface, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineNetworkInterfaceAttachService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachNetworkInterface",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.AttachNetworkInterfaceCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["NetworkInterfaceId"], ":", (*call.SdkParam)["InstanceId"]))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"InUse"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNetworkInterfaceAttachService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineNetworkInterfaceAttachService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachNetworkInterface",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"NetworkInterfaceId": ids[0],
				"InstanceId":         ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.DetachNetworkInterfaceCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading network interface on delete %q, %w", d.Id(), callErr))
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNetworkInterfaceAttachService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineNetworkInterfaceAttachService) ReadResourceId(id string) string {
	items := strings.Split(id, ":")
	return items[0]
}

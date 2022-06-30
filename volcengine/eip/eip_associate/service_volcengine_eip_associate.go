package eip_associate

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

type VolcengineEipAssociateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEipAssociateService(c *ve.SdkClient) *VolcengineEipAssociateService {
	return &VolcengineEipAssociateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEipAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEipAssociateService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineEipAssociateService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		resp             *map[string]interface{}
		results          interface{}
		ok               bool
		allocationId     string
		targetInstanceId string
		ids              []string
	)

	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}

	ids = strings.Split(tmpId, ":")
	allocationId = ids[0]
	targetInstanceId = ids[1]

	req := map[string]interface{}{
		"AllocationId": allocationId,
	}
	vpc := s.Client.VpcClient
	action := "DescribeEipAddressAttributes"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = vpc.DescribeEipAddressAttributesCommon(&req)
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
		return data, fmt.Errorf("eip address attributes %s not exist ", allocationId)
	}
	if instanceId, ok := data["InstanceId"]; !ok {
		return data, errors.New("instance id not exist")
	} else {
		if len(instanceId.(string)) == 0 {
			return data, errors.New("not associate")
		}
		if instanceId.(string) != targetInstanceId {
			return data, fmt.Errorf("eip address %s does not associate target instance. assoicate_instance_id %s, target_instance_id %s",
				allocationId, instanceId.(string), targetInstanceId)
		}
	}
	return data, err
}

func (s *VolcengineEipAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("eip address status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (VolcengineEipAssociateService) WithResourceResponseHandlers(eip map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return eip, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineEipAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateEipAddress",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.AssociateEipAddressCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["AllocationId"], ":", (*call.SdkParam)["InstanceId"]))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Attached"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEipAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineEipAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(s.ReadResourceId(resourceData.Id()), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateEipAddress",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"AllocationId": ids[0],
				"InstanceId":   ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.DisassociateEipAddressCommon(call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutDelete),
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading eip associate on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineEipAssociateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineEipAssociateService) ReadResourceId(id string) string {
	return id
}

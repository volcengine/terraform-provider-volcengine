package ha_vip_associate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineHaVipAssociateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewHaVipAssociateService(c *ve.SdkClient) *VolcengineHaVipAssociateService {
	return &VolcengineHaVipAssociateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineHaVipAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineHaVipAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeHaVips"

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

		results, err = ve.ObtainSdkValue("Result.HaVips", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.HaVips is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineHaVipAssociateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid ha vip associate Id: %s", id)
	}
	haVipId := ids[0]
	instanceId := ids[1]
	req := map[string]interface{}{
		"HaVipIds.1": haVipId,
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
		return data, fmt.Errorf("ha_vip_associate %s not exist ", id)
	}

	if instanceType, exist := data["AssociatedInstanceType"]; exist {
		data["InstanceType"] = instanceType
	}
	if instanceIds, exist := data["AssociatedInstanceIds"]; exist && instanceIds != nil {
		instanceIdArr, ok := instanceIds.([]interface{})
		if !ok {
			return data, fmt.Errorf("AssociatedInstanceIds is not slice")
		}
		for _, id := range instanceIdArr {
			if instanceId == id {
				data["InstanceId"] = id
				return data, err
			}
		}
		return data, fmt.Errorf("The instance %s is not associated to the ha vip %s ", instanceId, haVipId)
	}

	return data, err
}

func (s *VolcengineHaVipAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("ha_vip_associate status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineHaVipAssociateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineHaVipAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateHaVip",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("ha_vip_id"), d.Get("instance_id"))
				d.SetId(id)
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("ha_vip_id").(string)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"InUse"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineHaVipAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineHaVipAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateHaVip",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				associateId := d.Id()
				ids := strings.Split(associateId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("error ha vip associate id: %s", associateId)
				}
				(*call.SdkParam)["HaVipId"] = ids[0]
				(*call.SdkParam)["InstanceId"] = ids[1]
				(*call.SdkParam)["InstanceType"] = d.Get("instance_type")
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading ha vip associate on delete %q, %w", d.Id(), callErr))
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("ha_vip_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineHaVipAssociateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineHaVipAssociateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

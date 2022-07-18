package rds_ip_list

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsIpListService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func NewRdsIpListService(c *volc.SdkClient) *VolcengineRdsIpListService {
	return &VolcengineRdsIpListService{
		Client:     c,
		Dispatcher: &volc.Dispatcher{},
	}
}

func (s *VolcengineRdsIpListService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsIpListService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	rdsClient := s.Client.RdsClient
	action := "ListDBInstanceIPLists"
	logger.Debug(logger.ReqFormat, action, condition)
	if condition == nil {
		resp, err = rdsClient.ListDBInstanceIPListsCommon(nil)
		if err != nil {
			return data, err
		}
	} else {
		resp, err = rdsClient.ListDBInstanceIPListsCommon(&condition)
		if err != nil {
			return data, err
		}
	}

	results, err = volc.ObtainSdkValue("Result.Datas", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.Datas is not Slice")
	}

	targetGroupName := condition["GroupName"]

	// 生成ID
	res := make([]interface{}, 0)
	for _, a := range results.([]interface{}) {
		ipList, ok := a.(map[string]interface{})
		if !ok {
			continue
		}

		if targetGroupName != nil && targetGroupName.(string) != ipList["GroupName"].(string) {
			// ListDBInstanceIPLists接口不支持根据GroupName过滤，这里手动过滤下
			continue
		}

		ipList["Id"] = fmt.Sprintf("%s:%s", condition["InstanceId"], ipList["GroupName"])
		res = append(res, ipList)
	}

	return res, err
}

func (s *VolcengineRdsIpListService) ReadResource(resourceData *schema.ResourceData, rdsIpListId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if rdsIpListId == "" {
		rdsIpListId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(rdsIpListId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid RDS ip list id %s", rdsIpListId)
	}

	instanceId := ids[0]
	groupName := ids[1]

	req := map[string]interface{}{
		"InstanceId": instanceId,
		"GroupName":  groupName,
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
		return data, fmt.Errorf("RDS ip list %s not exist ", rdsIpListId)
	}

	return data, err
}

func (s *VolcengineRdsIpListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineRdsIpListService) WithResourceResponseHandlers(rdsIpList map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return rdsIpList, map[string]volc.ResponseConvert{
			"IPList": {
				TargetField: "ip_list",
			},
		}, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineRdsIpListService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateDBInstanceIPList",
			ConvertMode: volc.RequestConvertAll,
			Convert: map[string]volc.RequestConvert{
				"ip_list": {
					TargetField: "IPList",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建RdsIpList
				return s.Client.RdsClient.CreateDBInstanceIPListCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("instance_id"), d.Get("group_name"))
				d.SetId(id)
				return nil
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsIpListService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "ModifyDBInstanceIPList",
			ConvertMode: volc.RequestConvertAll,
			Convert: map[string]volc.RequestConvert{
				"ip_list": {
					TargetField: "IPList",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid RDS ip list id %s", d.Id())
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["GroupName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.RdsClient.ModifyDBInstanceIPListCommon(call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsIpListService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteDBInstanceIPList",
			ConvertMode: volc.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				rdsIpListId := d.Id()
				ids := strings.Split(rdsIpListId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid RDS ip list id %s", rdsIpListId)
				}
				(*call.SdkParam)["InstanceId"] = ids[0]
				(*call.SdkParam)["GroupName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除RdsIpList
				return s.Client.RdsClient.DeleteDBInstanceIPListCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if volc.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading RDS ip list on delete %q, %w", d.Id(), callErr))
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
	return []volc.Callback{callback}
}

func (s *VolcengineRdsIpListService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		NameField:    "GroupName",
		CollectField: "rds_ip_lists",
		ResponseConverts: map[string]volc.ResponseConvert{
			"IPList": {
				TargetField: "ip_list",
			},
		},
	}
}

func (s *VolcengineRdsIpListService) ReadResourceId(id string) string {
	return id
}

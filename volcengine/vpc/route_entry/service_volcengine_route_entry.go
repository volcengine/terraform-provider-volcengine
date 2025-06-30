package route_entry

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/vpc"
)

type VolcengineRouteEntryService struct {
	Client *ve.SdkClient
}

func NewRouteEntryService(c *ve.SdkClient) *VolcengineRouteEntryService {
	return &VolcengineRouteEntryService{
		Client: c,
	}
}

func (s *VolcengineRouteEntryService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRouteEntryService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		entries []interface{}
		res     = make([]interface{}, 0)
		ids     interface{}
		idsMap  = make(map[string]bool)
		ok      bool
	)
	entries, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeRouteEntryList"
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

		results, err = ve.ObtainSdkValue("Result.RouteEntries", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RouteEntries is not Slice")
		}
		return data, err
	})
	if err != nil {
		return entries, err
	}
	ids, ok = m["RouteEntryIds"]
	if !ok || ids == nil {
		return entries, nil
	}
	for _, id := range ids.(*schema.Set).List() {
		idsMap[strings.Trim(id.(string), " ")] = true
	}
	if len(idsMap) == 0 {
		return entries, nil
	}
	for _, entry := range entries {
		if _, ok = idsMap[entry.(map[string]interface{})["RouteEntryId"].(string)]; ok {
			res = append(res, entry)
		}
	}
	return res, nil
}

func (s *VolcengineRouteEntryService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(tmpId, ":")
	if len(ids) < 2 {
		return nil, fmt.Errorf("error route tmp id %s", tmpId)
	}
	req := map[string]interface{}{
		"RouteEntryId":   ids[1],
		"RouteTableId":   ids[0],
		"RouteEntryType": "Custom",
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("route entry %s not exist ", tmpId)
	}
	return data, err
}

func (s *VolcengineRouteEntryService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				failStates []string
				status     interface{}
			)
			failStates = append(failStates, "Error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("route entry error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (VolcengineRouteEntryService) WithResourceResponseHandlers(entries map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return entries, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineRouteEntryService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var vpcId string
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRouteEntry",
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				routeTableId := resourceData.Get("route_table_id").(string)
				resp, err := route_table.NewRouteTableService(s.Client).ReadResource(resourceData, routeTableId)
				if err != nil {
					return false, err
				}
				vpcId = resp["VpcId"].(string)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				id, _ := ve.ObtainSdkValue("Result.RouteEntryId", *resp)
				d.SetId(fmt.Sprint((*call.SdkParam)["RouteTableId"], ":", id))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("route_table_id").(string)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			// 外部定义vpcId无法传入ExtraRefresh中
			ExtraRefreshCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (map[ve.ResourceService]*ve.StateRefresh, error) {
				return map[ve.ResourceService]*ve.StateRefresh{
					vpc.NewVpcService(s.Client): {
						Target:     []string{"Available"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
						ResourceId: vpcId,
					},
				}, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRouteEntryService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	ids := strings.Split(s.ReadResourceId(resourceData.Id()), ":")
	var vpcId string
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyRouteEntry",
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				routeTableId := resourceData.Get("route_table_id").(string)
				resp, err := route_table.NewRouteTableService(s.Client).ReadResource(resourceData, routeTableId)
				if err != nil {
					return false, err
				}
				vpcId = resp["VpcId"].(string)

				(*call.SdkParam)["RouteEntryId"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("route_table_id").(string)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
			// 外部定义vpcId无法传入ExtraRefresh中
			ExtraRefreshCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (map[ve.ResourceService]*ve.StateRefresh, error) {
				return map[ve.ResourceService]*ve.StateRefresh{
					vpc.NewVpcService(s.Client): {
						Target:     []string{"Available"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
						ResourceId: vpcId,
					},
				}, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRouteEntryService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	var vpcId string
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRouteEntry",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"RouteEntryId": ids[1],
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				routeTableId := resourceData.Get("route_table_id").(string)
				resp, err := route_table.NewRouteTableService(s.Client).ReadResource(resourceData, routeTableId)
				if err != nil {
					return false, err
				}
				vpcId = resp["VpcId"].(string)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("route_table_id").(string)
			},
			// 外部定义vpcId无法传入ExtraRefresh中
			ExtraRefreshCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (map[ve.ResourceService]*ve.StateRefresh, error) {
				return map[ve.ResourceService]*ve.StateRefresh{
					vpc.NewVpcService(s.Client): {
						Target:     []string{"Available"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
						ResourceId: vpcId,
					},
				}, nil
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading route entry on delete %q, %w", d.Id(), callErr))
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
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 3*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRouteEntryService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "RouteEntryIds",
			},
		},
		NameField:    "RouteEntryName",
		IdField:      "RouteEntryId",
		CollectField: "route_entries",
		ResponseConverts: map[string]ve.ResponseConvert{
			"RouteEntryId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func importRouteEntry(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form RouteTableId:RouteEntryId")
	}
	err = d.Set("route_table_id", items[0])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}
	err = d.Set("route_entry_id", items[1])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	return []*schema.ResourceData{d}, nil
}

func (s *VolcengineRouteEntryService) ReadResourceId(id string) string {
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

package route_table

import (
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

type VolcengineTRRouteTableService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func (v *VolcengineTRRouteTableService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTRRouteTableService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTransitRouterRouteTables"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.ReqFormat, action, condition, *resp)
		results, err = ve.ObtainSdkValue("Result.TransitRouterRouteTables", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.TransitRouterRouteTables is not Slice")
		}
		return data, err
	})
}

func (v *VolcengineTRRouteTableService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return nil, errors.New("id err")
	}
	req := map[string]interface{}{
		"TransitRouterId":              ids[0],
		"TransitRouterRouteTableIds.1": ids[1],
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("TransitRouterRouteTable %s not exists", id)
	}
	return data, err
}

func (v *VolcengineTRRouteTableService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			demo, err = v.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, s := range failStates {
				if s == status.(string) {
					return nil, "", fmt.Errorf("TransitRouterRouteTable status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (v *VolcengineTRRouteTableService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTRRouteTableService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateTransitRouterRouteTable",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.TransitRouterRouteTableId", *resp)
				d.SetId(fmt.Sprint((*call.SdkParam)["TransitRouterId"], ":", id.(string)))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: data.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("transit_router_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTRRouteTableService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyTransitRouterRouteTableAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"description": {
					ConvertType: ve.ConvertDefault,
				},
				"transit_router_route_table_name": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(data.Id(), ":")
				(*call.SdkParam)["TransitRouterRouteTableId"] = ids[1]
				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: data.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("transit_router_id").(string)
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 更新Tags
	callbacks = v.setResourceTags(data, callbacks)

	return callbacks
}

func (v *VolcengineTRRouteTableService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(data.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteTransitRouterRouteTable",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TransitRouterRouteTableId": ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tr route table on delete %q, %w", d.Id(), callErr))
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
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 3*time.Minute)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("transit_router_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTRRouteTableService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "TransitRouterRouteTableIds",
				ConvertType: ve.ConvertWithN,
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"value": {
						TargetField: "Values.1",
					},
				},
			},
		},
		CollectField: "route_tables",
	}
}

func (v *VolcengineTRRouteTableService) ReadResourceId(s string) string {
	return s
}

func (v *VolcengineTRRouteTableService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					ids := strings.Split(resourceData.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid route table id")
					}
					(*call.SdkParam)["ResourceIds.1"] = ids[1]
					(*call.SdkParam)["ResourceType"] = "transitrouterroutetable"
					for index, v := range removedTags.List() {
						tag, ok := v.(map[string]interface{})
						if !ok {
							return false, fmt.Errorf("Tags is not map ")
						}
						(*call.SdkParam)["TagKeys."+strconv.Itoa(index+1)] = tag["key"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					ids := strings.Split(resourceData.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid route table id")
					}
					(*call.SdkParam)["ResourceIds.1"] = ids[1]
					(*call.SdkParam)["ResourceType"] = "transitrouterroutetable"
					for index, v := range addedTags.List() {
						tag, ok := v.(map[string]interface{})
						if !ok {
							return false, fmt.Errorf("Tags is not map ")
						}
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Key"] = tag["key"].(string)
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Value"] = tag["value"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func NewTRRouteTableService(c *ve.SdkClient) *VolcengineTRRouteTableService {
	return &VolcengineTRRouteTableService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "transitrouter",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

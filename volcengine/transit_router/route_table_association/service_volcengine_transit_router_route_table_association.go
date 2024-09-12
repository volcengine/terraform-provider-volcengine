package route_table_association

import (
	"errors"
	"fmt"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

type VolcengineTRRouteTableAssociationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func (v *VolcengineTRRouteTableAssociationService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTRRouteTableAssociationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTransitRouterRouteTableAssociations"
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
		results, err = ve.ObtainSdkValue("Result.TransitRouterRouteTableAssociations", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.TransitRouterRouteTableAssociations is not Slice")
		}
		return data, err
	})
}

func (v *VolcengineTRRouteTableAssociationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		"TransitRouterRouteTableId": ids[1],
		"TransitRouterAttachmentId": ids[0],
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
		return data, fmt.Errorf("TransitRouterRouteTableAssociation %s not exists", id)
	}
	return data, err
}

func (v *VolcengineTRRouteTableAssociationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("TransitRouterRouteTableAssociation status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (v *VolcengineTRRouteTableAssociationService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTRRouteTableAssociationService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateTransitRouterAttachmentToRouteTable",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["TransitRouterAttachmentId"], ":", (*call.SdkParam)["TransitRouterRouteTableId"]))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: data.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				attachmentId := d.Get("transit_router_attachment_id").(string)
				trId, err := v.describeTransitRouterId(attachmentId)
				if err != nil {
					logger.DebugInfo("LockId get transit_router_id error: ", err)
					return ""
				}
				return trId
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTRRouteTableAssociationService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcengineTRRouteTableAssociationService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DissociateTransitRouterAttachmentFromRouteTable",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TransitRouterRouteTableId": data.Get("transit_router_route_table_id"),
				"TransitRouterAttachmentId": data.Get("transit_router_attachment_id"),
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tr route table association on delete %q, %w", d.Id(), callErr))
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
				attachmentId := d.Get("transit_router_attachment_id").(string)
				trId, err := v.describeTransitRouterId(attachmentId)
				if err != nil {
					logger.DebugInfo("LockId get transit_router_id error: ", err)
					return ""
				}
				return trId
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTRRouteTableAssociationService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "associations",
	}
}

func (v *VolcengineTRRouteTableAssociationService) ReadResourceId(s string) string {
	return s
}

func (v *VolcengineTRRouteTableAssociationService) describeTransitRouterId(attachmentId string) (string, error) {
	action := "DescribeTransitRouterAttachments"
	req := map[string]interface{}{
		"TransitRouterAttachmentIds.1": attachmentId,
	}
	resp, err := v.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return "", err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)
	results, err := ve.ObtainSdkValue("Result.TransitRouterAttachments", *resp)
	if err != nil {
		return "", err
	}
	if results == nil {
		results = []interface{}{}
	}
	trAttachments, ok := results.([]interface{})
	if !ok {
		return "", errors.New("Result.TransitRouterAttachments is not Slice")
	}
	if len(trAttachments) == 0 {
		return "", fmt.Errorf("TransitRouterAttachments %s not exist", attachmentId)
	}
	trAttachment, ok := trAttachments[0].(map[string]interface{})
	if !ok {
		return "", errors.New("The value of Result.TransitRouterAttachments is not map ")
	}
	trId := trAttachment["TransitRouterId"].(string)
	return trId, nil
}

func NewTRRouteTableAssociationService(c *ve.SdkClient) *VolcengineTRRouteTableAssociationService {
	return &VolcengineTRRouteTableAssociationService{
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

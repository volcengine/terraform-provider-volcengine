package transit_router_vpc_attachment

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type Service struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewService(c *ve.SdkClient) *Service {
	return &Service{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *Service) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *Service) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTransitRouterVpcAttachments"
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
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.TransitRouterAttachments", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.TransitRouterAttachments is not Slice")
		}
		return data, err
	})
}

func (s *Service) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")

	req := map[string]interface{}{
		"TransitRouterId":              parts[0],
		"TransitRouterAttachmentIds.1": parts[1],
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
		return data, fmt.Errorf("TransitRouter Vpc Attachment %s not exist ", id)
	}
	return data, err
}

func (s *Service) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")

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

			status, err = ve.ObtainSdkValue("Status", data)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("TransitRouter Vpc Attachment status error, status:%s", status.(string))
				}
			}
			return data, status.(string), err
		},
	}

}

func (Service) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return data, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *Service) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateTransitRouterVpcAttachment",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"attach_points": {
					ConvertType: ve.ConvertListN,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.TransitRouterAttachmentId", *resp)
				d.SetId(fmt.Sprintf("%s:%s", d.Get("transit_router_id"), id))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("transit_router_id").(string)
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *Service) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyTransitRouterVpcAttachmentAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"description": {
					TargetField: "Description",
				},
				"transit_router_attachment_name": {
					TargetField: "TransitRouterAttachmentName",
				},
				"auto_publish_route_enabled": {
					TargetField: "AutoPublishRouteEnabled",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["TransitRouterAttachmentId"] = d.Get("transit_router_attachment_id")
				if d.HasChange("attach_points") {
					points := d.Get("attach_points").(*schema.Set).List()
					for index, point := range points {
						p := point.(map[string]interface{})
						(*call.SdkParam)[fmt.Sprintf("AttachPoints.%d.SubnetId", index+1)] = p["subnet_id"]
						(*call.SdkParam)[fmt.Sprintf("AttachPoints.%d.ZoneId", index+1)] = p["zone_id"]
					}
				}
				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("transit_router_id").(string)
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 更新Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *Service) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteTransitRouterVpcAttachment",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TransitRouterAttachmentId": resourceData.Get("transit_router_attachment_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("transit_router_id").(string)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading transit router vpc attachment on delete %q, %w", d.Id(), callErr))
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

func (s *Service) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"transit_router_attachment_ids": {
				TargetField: "TransitRouterAttachmentIds",
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
		NameField:        "TransitRouterAttachmentName",
		IdField:          "TransitRouterAttachmentId",
		CollectField:     "attachments",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *Service) ReadResourceId(id string) string {
	return id
}

func (s *Service) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
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
					(*call.SdkParam)["ResourceType"] = "transitrouterattachment"
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
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
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
					(*call.SdkParam)["ResourceType"] = "transitrouterattachment"
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
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
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

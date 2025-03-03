package tos_bucket_inventory

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

type VolcengineTosBucketInventoryService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketInventoryService(c *ve.SdkClient) *VolcengineTosBucketInventoryService {
	return &VolcengineTosBucketInventoryService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketInventoryService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketInventoryService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		action  string
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action = "ListBucketInventory"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     condition[ve.BypassDomain].(string),
		UrlParam: map[string]string{
			"inventory": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue(ve.BypassResponse+".InventoryConfigurations", *resp)
	if err != nil {
		return data, err
	}

	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New("InventoryConfigurations is not Slice")
	}
	return data, err
}

func (s *VolcengineTosBucketInventoryService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid tos inventory id: %s", id)
	}

	action := "GetBucketInventory"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     ids[0],
		UrlParam: map[string]string{
			"inventory": "",
			"id":        ids[1],
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketInventory Resp is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_inventory %s not exist ", id)
	}

	data["BucketName"] = ids[0]
	if destination, ok := data["Destination"].(map[string]interface{}); ok {
		if tosBucketDestination, ok := destination["TOSBucketDestination"].(map[string]interface{}); ok {
			destination["TOSBucketDestination"] = []interface{}{tosBucketDestination}
		}
	}

	return data, err
}

func (s *VolcengineTosBucketInventoryService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineTosBucketInventoryService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"TOSBucketDestination": {
				TargetField: "tos_bucket_destination",
			},
			"Id": {
				TargetField: "inventory_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketInventoryService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	//create inventory
	callback := s.createOrUpdateInventory(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketInventoryService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	//create inventory
	callback := s.createOrUpdateInventory(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketInventoryService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketInventory",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid tos inventory id: %s", d.Id())
				}
				(*call.SdkParam)["BucketName"] = ids[0]
				(*call.SdkParam)["Id"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Domain:      (*call.SdkParam)["BucketName"].(string),
					UrlParam: map[string]string{
						"inventory": "",
						"id":        (*call.SdkParam)["Id"].(string),
					},
				}, nil)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tos inventory on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketInventoryService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	bucketName := data.Get("bucket_name")
	inventoryId, ok := data.GetOk("inventory_id")
	return ve.DataSourceInfo{
		ServiceCategory: ve.ServiceBypass,
		RequestConverts: map[string]ve.RequestConvert{
			"bucket_name": {
				ConvertType: ve.ConvertDefault,
				SpecialParam: &ve.SpecialParam{
					Type: ve.DomainParam,
				},
			},
			"inventory_id": {
				Ignore: true,
			},
		},
		NameField:    "Id",
		IdField:      "InventoryId",
		CollectField: "inventory_configurations",
		ResponseConverts: map[string]ve.ResponseConvert{
			"TOSBucketDestination": {
				TargetField: "tos_bucket_destination",
			},
		},
		ExtraData: func(sourceData []interface{}) (extraData []interface{}, err error) {
			for _, v := range sourceData {
				if ok {
					if inventoryId.(string) == v.(map[string]interface{})["Id"].(string) {
						v.(map[string]interface{})["InventoryId"] = bucketName.(string) + ":" + v.(map[string]interface{})["Id"].(string)
						extraData = append(extraData, v)
						break
					} else {
						continue
					}
				} else {
					v.(map[string]interface{})["InventoryId"] = bucketName.(string) + ":" + v.(map[string]interface{})["Id"].(string)
					extraData = append(extraData, v)
				}

			}
			return extraData, err
		},
	}
}

func (s *VolcengineTosBucketInventoryService) createOrUpdateInventory(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketInventory",
			ConvertMode:     ve.RequestConvertInConvert,
			ContentType:     ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
					ForceGet: isUpdate,
				},
				"inventory_id": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Id",
					ForceGet:    isUpdate,
				},
				"is_enabled": {
					ConvertType: ve.ConvertDefault,
					TargetField: "IsEnabled",
					ForceGet:    true,
				},
				"included_object_versions": {
					ConvertType: ve.ConvertDefault,
					TargetField: "IncludedObjectVersions",
					ForceGet:    true,
				},
				"schedule": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "Schedule",
					ForceGet:    true,
				},
				"filter": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "Filter",
					ForceGet:    true,
				},
				"optional_fields": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "OptionalFields",
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"field": {
							ConvertType: ve.ConvertJsonArray,
							TargetField: "Field",
						},
					},
				},
				"destination": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "Destination",
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"tos_bucket_destination": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "TosBucketDestination",
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id := d.Get("inventory_id")
				(*call.SdkParam)["InventoryId"] = id.(string)

				var sourceParam map[string]interface{}
				sourceParam, err := ve.SortAndStartTransJson((*call.SdkParam)[ve.BypassParam].(map[string]interface{}))
				if err != nil {
					return false, err
				}
				(*call.SdkParam)[ve.BypassParam] = sourceParam

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建 Inventory
				param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod:  ve.PUT,
					ContentType: ve.ApplicationJSON,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"inventory": "",
						"id":        (*call.SdkParam)["InventoryId"].(string),
					},
				}, &param)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string) + ":" + (*call.SdkParam)["InventoryId"].(string))
				return nil
			},
		},
	}

	return callback
}

func (s *VolcengineTosBucketInventoryService) ReadResourceId(id string) string {
	return id
}

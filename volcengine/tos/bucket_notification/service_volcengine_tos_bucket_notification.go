package tos_bucket_notification

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

type VolcengineTosBucketNotificationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketNotificationService(c *ve.SdkClient) *VolcengineTosBucketNotificationService {
	return &VolcengineTosBucketNotificationService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketNotificationService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketNotificationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketNotificationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid tos notification id: %s", id)
	}

	action := "GetBucketNotificationV2"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     ids[0],
		UrlParam: map[string]string{
			"notification_v2": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketNotificationV2 Resp is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_notification %s not exist ", id)
	}

	rules, ok := data["Rules"].([]interface{})
	if !ok {
		return data, fmt.Errorf("tos_bucket_notification %s rules is not slice ", id)
	}
	if len(rules) == 0 {
		return data, fmt.Errorf("tos_bucket_notification %s rules not exist ", id)
	}

	// 根据 rule id 查找对应 rule
	rule := make(map[string]interface{})
	for _, v := range rules {
		ruleMap, ok := v.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("tos_bucket_notification %s rule is not map ", id)
		}
		if ids[1] == ruleMap["RuleId"] {
			rule = ruleMap
			break
		}
	}
	if len(rule) == 0 {
		return data, fmt.Errorf("tos_bucket_notification %s rule not exist ", id)
	}

	if destination, exist := rule["Destination"]; exist {
		if destinationMap, ok := destination.(map[string]interface{}); ok {
			rule["Destination"] = []interface{}{destinationMap}
		}
	}
	if filter, exist := rule["Filter"]; exist {
		if filterMap, ok := filter.(map[string]interface{}); ok {
			if tosKey, exist := filterMap["TOSKey"]; exist {
				if tosKeyMap, ok := tosKey.(map[string]interface{}); ok {
					filterMap["TOSKey"] = []interface{}{tosKeyMap}
				}
			}
			rule["Filter"] = []interface{}{filterMap}
		}
	}
	data["Rules"] = []interface{}{rule}
	data["BucketName"] = ids[0]

	return data, err
}

func (s *VolcengineTosBucketNotificationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineTosBucketNotificationService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"TOSKey": {
				TargetField: "tos_key",
			},
			"VeFaaS": {
				TargetField: "ve_faas",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketNotificationService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	//create inventory
	callback := s.createOrUpdateNotification(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketNotificationService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	//create inventory
	callback := s.createOrUpdateNotification(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketNotificationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	if len(ids) != 2 {
		return []ve.Callback{{
			Err: fmt.Errorf("invalid tos bucket notification id: %s", resourceData.Id()),
		}}
	}
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "PutBucketNotificationV2",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BucketName"] = ids[0]
				(*call.SdkParam)[ve.BypassParam] = make(map[string]interface{})
				//version := d.Get("version").(string)
				//(*call.SdkParam)[ve.BypassParam] = map[string]interface{}{
				//	"Version": version,
				//}
				return true, nil
			},
			AfterLocked: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) error {
				// 获取存量 rules 信息
				action := "GetBucketNotificationV2"
				logger.Debug(logger.ReqFormat, action, d.Get("bucket_name"))
				data, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod: ve.GET,
					Domain:     ids[0],
					UrlParam: map[string]string{
						"notification_v2": "",
					},
				}, nil)
				logger.Debug(logger.RespFormat, action, data, err)
				if err != nil {
					return err
				}

				v, _ := ve.ObtainSdkValue("Rules", (*data)[ve.BypassResponse])
				rules, ok := v.([]interface{})
				if !ok {
					return fmt.Errorf("tos_bucket_notification %s rules is not slice ", d.Id())
				}
				for index, v := range rules {
					ruleMap, ok := v.(map[string]interface{})
					if !ok {
						return fmt.Errorf("tos_bucket_notification %s rule is not map ", d.Id())
					}
					if ids[1] == ruleMap["RuleId"] {
						rules = append(rules[:index], rules[index+1:]...)
						break
					}
				}
				(*call.SdkParam)[ve.BypassParam].(map[string]interface{})["Rules"] = rules
				return nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod:  ve.PUT,
					ContentType: ve.ApplicationJSON,
					Domain:      (*call.SdkParam)["BucketName"].(string),
					UrlParam: map[string]string{
						"notification_v2": "",
					},
				}, &param)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket realtime log on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("bucket_name").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineTosBucketNotificationService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketNotificationService) createOrUpdateNotification(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketNotificationV2",
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
				//"version": {
				//	ConvertType: ve.ConvertDefault,
				//	TargetField: "Version",
				//	ForceGet:    isUpdate,
				//},
				"rules": {
					ConvertType: ve.ConvertJsonObjectArray,
					TargetField: "Rules",
					//ForceGet: true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"rule_id": {
							ConvertType: ve.ConvertDefault,
							TargetField: "RuleId",
							ForceGet:    isUpdate,
						},
						"events": {
							ConvertType: ve.ConvertJsonArray,
							TargetField: "Events",
							ForceGet:    isUpdate,
						},
						"destination": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Destination",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"ve_faas": {
									ConvertType: ve.ConvertJsonObjectArray,
									TargetField: "VeFaaS",
									ForceGet:    isUpdate,
									NextLevelConvert: map[string]ve.RequestConvert{
										"function_id": {
											ConvertType: ve.ConvertDefault,
											TargetField: "FunctionId",
										},
									},
								},
							},
						},
						"filter": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Filter",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"tos_key": {
									ConvertType: ve.ConvertJsonObject,
									TargetField: "TOSKey",
									ForceGet:    isUpdate,
									NextLevelConvert: map[string]ve.RequestConvert{
										"filter_rules": {
											ConvertType: ve.ConvertJsonObjectArray,
											TargetField: "FilterRules",
											ForceGet:    isUpdate,
											NextLevelConvert: map[string]ve.RequestConvert{
												"name": {
													ConvertType: ve.ConvertDefault,
													TargetField: "Name",
													ForceGet:    isUpdate,
												},
												"value": {
													ConvertType: ve.ConvertDefault,
													TargetField: "Value",
													ForceGet:    isUpdate,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			//BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
			//	id := d.Get("rules.0.rule_id")
			//	(*call.SdkParam)["RuleId"] = id.(string)
			//
			//	var sourceParam map[string]interface{}
			//	sourceParam, err := ve.SortAndStartTransJson((*call.SdkParam)[ve.BypassParam].(map[string]interface{}))
			//	if err != nil {
			//		return false, err
			//	}
			//	(*call.SdkParam)[ve.BypassParam] = sourceParam
			//
			//	return true, nil
			//},
			AfterLocked: s.beforePutBucketNotification(isUpdate),
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建 Notification
				param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod:  ve.PUT,
					ContentType: ve.ApplicationJSON,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"notification_v2": "",
					},
				}, &param)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string) + ":" + d.Get("rules.0.rule_id").(string))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("bucket_name").(string)
			},
		},
	}

	return callback
}

func (s *VolcengineTosBucketNotificationService) beforePutBucketNotification(isUpdate bool) ve.CallFunc {

	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) error {
		action := "GetBucketNotificationV2"
		logger.Debug(logger.ReqFormat, action, d.Get("bucket_name"))
		data, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			HttpMethod: ve.GET,
			Domain:     (*call.SdkParam)[ve.BypassDomain].(string),
			UrlParam: map[string]string{
				"notification_v2": "",
			},
		}, nil)
		logger.Debug(logger.RespFormat, action, data, err)
		return s.beforeTosPutNotification(d, call, data, err, isUpdate)
	}
}

func (s *VolcengineTosBucketNotificationService) beforeTosPutNotification(d *schema.ResourceData, call ve.SdkCall, data *map[string]interface{}, err error, isUpdate bool) error {
	if err != nil {
		return err
	}
	id := d.Get("rules.0.rule_id")

	var sourceAclParam map[string]interface{}
	sourceAclParam, err = ve.SortAndStartTransJson((*call.SdkParam)[ve.BypassParam].(map[string]interface{}))
	if err != nil {
		return err
	}
	v, _ := ve.ObtainSdkValue("Rules", (*data)[ve.BypassResponse])
	rules, ok := v.([]interface{})
	if !ok {
		return fmt.Errorf("tos_bucket_notification %s rules is not slice ", id)
	}
	if len(rules) == 0 && isUpdate {
		return fmt.Errorf("tos_bucket_notification %s rules is empty", id)
	}
	// 根据 rule id 查找对应 rule
	rule := make(map[string]interface{})
	for index, v := range rules {
		ruleMap, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("tos_bucket_notification %s rule is not map ", id)
		}
		if id == ruleMap["RuleId"] {
			rule = ruleMap
			rules = append(rules[:index], rules[index+1:]...)
			break
		}
	}
	if len(rule) == 0 && isUpdate {
		return fmt.Errorf("tos_bucket_notification %s rule not exist ", id)
	}
	if len(rule) > 0 && !isUpdate {
		return fmt.Errorf("tos_bucket_notification %s rule is existed ", id)
	}
	// merge rule
	rulesParam, _ := ve.ObtainSdkValue("Rules", sourceAclParam)
	if rulesParam != nil {
		_, ok := rulesParam.([]interface{})
		if !ok {
			return fmt.Errorf("tos_bucket_notification %s rules is not slice ", id)
		}
		rulesParam = append(rulesParam.([]interface{}), rules...)
	}
	sourceAclParam["Rules"] = rulesParam

	(*call.SdkParam)[ve.BypassParam] = sourceAclParam
	return nil
}

func (s *VolcengineTosBucketNotificationService) ReadResourceId(id string) string {
	return id
}

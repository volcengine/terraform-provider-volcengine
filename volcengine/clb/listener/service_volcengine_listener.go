package listener

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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/clb"
)

type VolcengineListenerService struct {
	Client *ve.SdkClient
}

func NewListenerService(c *ve.SdkClient) *VolcengineListenerService {
	return &VolcengineListenerService{
		Client: c,
	}
}

func (s *VolcengineListenerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineListenerService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeListeners"
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

		results, err = ve.ObtainSdkValue("Result.Listeners", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Listeners is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineListenerService) ReadResource(resourceData *schema.ResourceData, listenerId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if listenerId == "" {
		listenerId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ListenerIds.1": listenerId,
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
		return data, fmt.Errorf("Listener %s not exist ", listenerId)
	}

	action := "DescribeListenerAttributes"
	listenerResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{
		"ListenerId": listenerId,
	})
	if err != nil {
		return nil, err
	}

	listenerAttrMap := make(map[string]interface{})

	timeout, err := ve.ObtainSdkValue("Result.EstablishedTimeout", *listenerResp)
	if err != nil {
		return nil, err
	}
	desc, err := ve.ObtainSdkValue("Result.Description", *listenerResp)
	if err != nil {
		return nil, err
	}
	loadBalancerId, err := ve.ObtainSdkValue("Result.LoadBalancerId", *listenerResp)
	if err != nil {
		return nil, err
	}
	scheduler, err := ve.ObtainSdkValue("Result.Scheduler", *listenerResp)
	if err != nil {
		return nil, err
	}

	listenerAttrMap["EstablishedTimeout"] = timeout
	listenerAttrMap["Description"] = desc
	listenerAttrMap["LoadBalancerId"] = loadBalancerId
	listenerAttrMap["Scheduler"] = scheduler

	for k, v := range listenerAttrMap {
		data[k] = v
	}

	return data, err
}

func (s *VolcengineListenerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status interface{}
			)
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (*VolcengineListenerService) WithResourceResponseHandlers(listener map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return listener, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineListenerService) refreshAclStatus() ve.CallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) error {
		var aclIds []string
		for k, v := range *call.SdkParam {
			if strings.HasPrefix(k, "AclIds.") {
				aclIds = append(aclIds, v.(string))
			}
		}
		if len(aclIds) > 0 {
			if err := s.checkAcl(aclIds); err != nil {
				return err
			}
		}
		return nil
	}
}

func (s *VolcengineListenerService) checkAcl(aclIds []string) error {
	return resource.Retry(20*time.Minute, func() *resource.RetryError {
		action := "DescribeAcls"
		req := make(map[string]interface{})
		for index, id := range aclIds {
			req["AclIds."+strconv.Itoa(index+1)] = id
		}
		logger.Debug(logger.ReqFormat, "DescribeAcls", aclIds)
		// create 的时候上限为5个，无需翻页
		resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		logger.Debug(logger.RespFormat, "DescribeAcls", aclIds, *resp)

		statusOK := true
		aclIdMap := make(map[string]bool)
		results, err := ve.ObtainSdkValue("Result.Acls", *resp)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		acls, ok := results.([]interface{})
		if !ok {
			return resource.NonRetryableError(fmt.Errorf("checkAcl Result.Acls is not Slice"))
		}
		for _, element := range acls {
			aclMap, ok := element.(map[string]interface{})
			if !ok {
				return resource.NonRetryableError(fmt.Errorf("checkAcl Acl is not map"))
			}
			aclIdMap[aclMap["AclId"].(string)] = true
			if aclMap["Status"] == "Deleting" {
				return resource.NonRetryableError(fmt.Errorf("acl is in deleting status"))
			} else if aclMap["Status"] != "Active" { // Creating / Configuring
				statusOK = false
				break
			}
		}
		if !statusOK {
			return resource.RetryableError(fmt.Errorf("acl still in waiting status"))
		}

		for _, aclId := range aclIds {
			if _, exist := aclIdMap[aclId]; !exist {
				return resource.NonRetryableError(fmt.Errorf("Cannot find acl: %s ", aclId))
			}
		}
		return nil
	})
}

func (s *VolcengineListenerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateListener",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"acl_ids": {
					ConvertType: ve.ConvertWithN,
				},
				"health_check": {
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"un_healthy_threshold": {
							TargetField: "UnhealthyThreshold",
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				protocol := (*call.SdkParam)["Protocol"].(string)
				// 1. established_timeout
				if protocol == "HTTP" || protocol == "HTTPS" {
					// not allow establish_timeout
					if _, ok := (*call.SdkParam)["EstablishedTimeout"]; ok {
						return false, errors.New("established_timeout is not allowed for HTTP or HTTPS")
					}
				}

				// 2. certificate_id
				if protocol != "HTTPS" && (*call.SdkParam)["CertificateId"] != nil {
					return false, errors.New("certificate_id is only allowed for HTTPS")
				}

				// 3. connection_drain_timeout
				if d.Get("connection_drain_enabled") == "on" {
					timeout := d.Get("connection_drain_timeout").(int)
					if timeout == 0 {
						(*call.SdkParam)["ConnectionDrainTimeout"] = 0
					}
				}

				return true, nil
			},
			AfterLocked: s.refreshAclStatus(),
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建listener
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.ListenerId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active", "Disabled"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("load_balancer_id").(string),
				},
			},
			AfterRefresh: s.refreshAclStatus(),
			LockId: func(d *schema.ResourceData) string {
				return resourceData.Get("load_balancer_id").(string)
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineListenerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	clbId, err := s.queryLoadBalancerId(resourceData.Id())
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyListenerAttributes",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"acl_ids": {
					ConvertType: ve.ConvertWithN,
				},
				"health_check": {
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"un_healthy_threshold": {
							TargetField: "UnhealthyThreshold",
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				protocol := d.Get("protocol").(string)
				// 1. established_timeout
				if protocol == "HTTP" || protocol == "HTTPS" {
					// not allow establish_timeout
					if _, ok := (*call.SdkParam)["EstablishedTimeout"]; ok {
						return false, errors.New("established_timeout is not allowed for HTTP or HTTPS")
					}
				}

				// 2. certificate_id
				if protocol != "HTTPS" && (*call.SdkParam)["CertificateId"] != nil {
					return false, errors.New("certificate_id is only allowed for HTTPS")
				}

				(*call.SdkParam)["ListenerId"] = d.Id()
				aclStatus := d.Get("acl_status")
				if aclStatus, ok := aclStatus.(string); ok && aclStatus == "on" {
					(*call.SdkParam)["AclStatus"] = d.Get("acl_status").(string)
					(*call.SdkParam)["AclType"] = d.Get("acl_type").(string)
					if !d.HasChange("acl_ids") {
						if m, ok := d.Get("acl_ids").(*schema.Set); ok {
							aclIds := m.List()
							for i, aclId := range aclIds {
								k := fmt.Sprintf("AclIds.%d", i+1)
								(*call.SdkParam)[k] = aclId
							}
						}
					}
				}
				return true, nil
			},
			AfterLocked: s.refreshAclStatus(),
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改 listener 属性
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active", "Disabled"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			AfterRefresh: s.refreshAclStatus(),
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineListenerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	clbId, err := s.queryLoadBalancerId(resourceData.Id())
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteListener",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ListenerId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除 Listener
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading listener on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineListenerService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ListenerIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "ListenerName",
		IdField:      "ListenerId",
		CollectField: "listeners",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ListenerId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"HealthCheck.Enabled": {
				TargetField: "health_check_enabled",
			},
			"HealthCheck.Interval": {
				TargetField: "health_check_interval",
			},
			"HealthCheck.HealthyThreshold": {
				TargetField: "health_check_healthy_threshold",
			},
			"HealthCheck.UnHealthyThreshold": {
				TargetField: "health_check_un_healthy_threshold",
			},
			"HealthCheck.Timeout": {
				TargetField: "health_check_timeout",
			},
			"HealthCheck.Method": {
				TargetField: "health_check_method",
			},
			"HealthCheck.Uri": {
				TargetField: "health_check_uri",
			},
			"HealthCheck.Domain": {
				TargetField: "health_check_domain",
			},
			"HealthCheck.HttpCode": {
				TargetField: "health_check_http_code",
			},
			"HealthCheck.UdpRequest": {
				TargetField: "health_check_udp_request",
			},
			"HealthCheck.UdpExpect": {
				TargetField: "health_check_udp_expect",
			},
		},
	}
}

func (s *VolcengineListenerService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineListenerService) queryLoadBalancerId(listenerId string) (string, error) {
	if listenerId == "" {
		return "", fmt.Errorf("listener ID cannot be empty")
	}

	// 查询 LoadBalancerId
	action := "DescribeListenerAttributes"
	serverGroupResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{
		"ListenerId": listenerId,
	})
	if err != nil {
		return "", err
	}
	clbId, err := ve.ObtainSdkValue("Result.LoadBalancerId", *serverGroupResp)
	if err != nil {
		return "", err
	}
	return clbId.(string), nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

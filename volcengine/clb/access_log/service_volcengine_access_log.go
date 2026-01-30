package access_log

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineAccessLogService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAccessLogService(c *ve.SdkClient) *VolcengineAccessLogService {
	return &VolcengineAccessLogService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAccessLogService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAccessLogService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	data, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeLoadBalancers"

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
		results, err = ve.ObtainSdkValue("Result.LoadBalancers", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.LoadBalancers is not Slice")
		}
		return data, err
	})
	if err != nil {
		return data, err
	}
	return data, err
}

func (s *VolcengineAccessLogService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"LoadBalancerIds.1": id,
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
		return data, fmt.Errorf("clb %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineAccessLogService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d      map[string]interface{}
				status interface{}
			)
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineAccessLogService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableAccessLog",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"load_balancer_id": {
					TargetField: "LoadBalancerId",
				},
				"delivery_type": {
					TargetField: "DeliveryType",
				},
				"bucket_name": {
					TargetField: "BucketName",
				},
				"tls_project_id": {
					TargetField: "TlsProjectId",
				},
				"tls_topic_id": {
					TargetField: "TlsTopicId",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// Validate required fields based on delivery_type
				deliveryType, ok := d.GetOk("delivery_type")
				if !ok || deliveryType == "tos" {
					// For TOS delivery, bucket_name is required
					if _, ok := d.GetOk("bucket_name"); !ok {
						return false, fmt.Errorf("bucket_name is required when delivery_type is 'tos'")
					}
				} else if deliveryType == "tls" {
					// For TLS delivery, both tls_project_id and tls_topic_id are required
					if _, ok := d.GetOk("tls_project_id"); !ok {
						return false, fmt.Errorf("tls_project_id is required when delivery_type is 'tls'")
					}
					if _, ok := d.GetOk("tls_topic_id"); !ok {
						return false, fmt.Errorf("tls_topic_id is required when delivery_type is 'tls'")
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// Set resource ID as LoadBalancerId for access log
				loadBalancerId, _ := ve.ObtainSdkValue("LoadBalancerId", *call.SdkParam)
				d.SetId(loadBalancerId.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("load_balancer_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineAccessLogService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAccessLogService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// Access log doesn't support update operation
	return []ve.Callback{}
}

func (s *VolcengineAccessLogService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	clbId := resourceData.Get("load_balancer_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisableAccessLog",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["LoadBalancerId"] = clbId
				if deliveryType, ok := d.GetOk("delivery_type"); ok {
					(*call.SdkParam)["DeliveryType"] = deliveryType.(string)
				} else {
					(*call.SdkParam)["DeliveryType"] = "tos"
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineAccessLogService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineAccessLogService) ReadResourceId(id string) string {
	return id
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

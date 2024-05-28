package bucket_policy

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketPolicyService struct {
	Client *ve.SdkClient
}

func NewTosBucketPolicyService(c *ve.SdkClient) *VolcengineTosBucketPolicyService {
	return &VolcengineTosBucketPolicyService{
		Client: c,
	}
}

func (s *VolcengineTosBucketPolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketPolicyService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		action  string
		resp    *map[string]interface{}
		results interface{}
	)
	action = "GetBucketPolicy"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     condition[ve.BypassDomain].(string),
		UrlParam: map[string]string{
			"policy": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue(ve.BypassResponse, *resp)
	if err != nil {
		return data, err
	}

	if len(results.(map[string]interface{})) == 0 {
		return data, fmt.Errorf("bucket Policy %s not exist ", condition[ve.BypassDomain].(string))
	}

	data = append(data, map[string]interface{}{
		"Policy": results.(map[string]interface{}),
	})
	return data, err
}

func (s *VolcengineTosBucketPolicyService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	bucketName := resourceData.Get("bucket_name").(string)
	if instanceId == "" {
		instanceId = s.ReadResourceId(resourceData.Id())
	} else {
		instanceId = s.ReadResourceId(instanceId)
	}

	var (
		ok      bool
		results []interface{}
	)

	logger.Debug(logger.ReqFormat, "GetBucketPolicy", bucketName+":"+instanceId)
	condition := map[string]interface{}{
		ve.BypassDomain: bucketName,
	}
	results, err = s.ReadResources(condition)

	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, fmt.Errorf("Value is not map ")
		}
	}

	if len(data) == 0 {
		return data, fmt.Errorf("bucket Policy %s not exist ", instanceId)
	}

	return data, nil
}

func (s *VolcengineTosBucketPolicyService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineTosBucketPolicyService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"Policy": {
				Convert: func(i interface{}) interface{} {
					b, _ := json.Marshal(i)
					return string(b)
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketPolicyService) putBucketPolicy(data *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketPolicy",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					ForceGet:    isUpdate,
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
				},
				"policy": {
					ForceGet:    isUpdate,
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				j := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})["Policy"]
				data := map[string]interface{}{}
				err := json.Unmarshal([]byte(j.(string)), &data)
				if err != nil {
					return false, err
				}
				delete((*call.SdkParam)[ve.BypassParam].(map[string]interface{}), "Policy")
				for k, v := range data {
					(*call.SdkParam)[ve.BypassParam].(map[string]interface{})[k] = v
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod:  ve.PUT,
					ContentType: ve.ApplicationJSON,
					UrlParam: map[string]string{
						"policy": "",
					},
					Domain: (*call.SdkParam)[ve.BypassDomain].(string),
				}, &param)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string) + ":POLICY")
				return nil
			},
		},
	}
}

func (s *VolcengineTosBucketPolicyService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{s.putBucketPolicy(data, resource, false)}
}

func (s *VolcengineTosBucketPolicyService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{s.putBucketPolicy(data, resource, true)}
}

func (s *VolcengineTosBucketPolicyService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "DeleteBucketPolicy",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					ForceGet:    true,
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod: ve.DELETE,
					Domain:     (*call.SdkParam)[ve.BypassDomain].(string),
					Path:       []string{"?policy="},
					UrlParam: map[string]string{
						"policy": "",
					},
				}, nil)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading tos bucket policy on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketPolicyService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketPolicyService) ReadResourceId(id string) string {
	return id
}

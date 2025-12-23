package tos_bucket_customdomain

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

type VolcengineTosBucketCustomDomainService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketCustomDomainService(c *ve.SdkClient) *VolcengineTosBucketCustomDomainService {
	return &VolcengineTosBucketCustomDomainService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketCustomDomainService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketCustomDomainService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketCustomDomainService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	bucketName, customDomain, err := parseId(id)
	if err != nil {
		return data, err
	}

	action := "GetBucketCustomDomain"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     bucketName,
		UrlParam: map[string]string{
			"customdomain": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketCustomDomain Resp is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_customdomain %s not exist ", id)
	}

	// Find the specific custom domain
	if v, exist := data["CustomDomainRules"]; exist {
		if domains, ok := v.([]interface{}); ok {
			found := false
			for _, domain := range domains {
				if domainMap, ok := domain.(map[string]interface{}); ok {
					if domainName, exist := domainMap["Domain"]; exist && domainName.(string) == customDomain {
						data = domainMap
						data["BucketName"] = bucketName
						data["CustomDomain"] = customDomain
						found = true
						break
					}
				}
			}
			if !found {
				return data, fmt.Errorf("tos_bucket_customdomain %s not exist ", id)
			}
		}
	}

	return data, err
}

func (s *VolcengineTosBucketCustomDomainService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineTosBucketCustomDomainService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineTosBucketCustomDomainService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateCustomDomain(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketCustomDomainService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateCustomDomain(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketCustomDomainService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketCustomDomain",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BucketName"] = d.Id()
				bucketName, customDomain, err := parseId(d.Id())
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["CustomDomain"] = customDomain
				(*call.SdkParam)["BucketName"] = bucketName
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Domain:      (*call.SdkParam)["BucketName"].(string),
					UrlParam: map[string]string{
						"customdomain": (*call.SdkParam)["CustomDomain"].(string),
					},
				}, nil)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket customdomain on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketCustomDomainService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketCustomDomainService) createOrUpdateCustomDomain(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketCustomDomain",
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
				"custom_domain_rule": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "CustomDomainRule",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"domain": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Domain",
							ForceGet:    isUpdate,
						},
						"protocol": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Protocol",
							ForceGet:    isUpdate,
						},
						"cert_id": {
							ConvertType: ve.ConvertDefault,
							TargetField: "CertId",
							ForceGet:    isUpdate,
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
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
				param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"customdomain": "",
					},
				}, &param)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string) + ":" + (*call.SdkParam)[ve.BypassParam].(map[string]interface{})["CustomDomainRule"].(map[string]interface{})["Domain"].(string))
				return nil
			},
		},
	}

	return callback
}

func (s *VolcengineTosBucketCustomDomainService) ReadResourceId(id string) string {
	return id
}

func parseId(id string) (string, string, error) {
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid ID format, expected format: bucket_name:custom_domain")
	}
	return parts[0], parts[1], nil
}

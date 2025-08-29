package tos_bucket_encryption

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketEncryptionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketEncryptionService(c *ve.SdkClient) *VolcengineTosBucketEncryptionService {
	return &VolcengineTosBucketEncryptionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketEncryptionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketEncryptionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketEncryptionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketEncryption"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"encryption": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketEncryption Resp is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_encryption %s not exist ", id)
	}

	data["BucketName"] = id
	if v, exist := data["Rule"]; exist {
		if rule, ok := v.(map[string]interface{}); ok {
			if v1, exist1 := rule["ApplyServerSideEncryptionByDefault"]; exist1 {
				if encryption, ok1 := v1.(map[string]interface{}); ok1 {
					rule["ApplyServerSideEncryptionByDefault"] = []interface{}{encryption}
				}
			}
			data["Rule"] = []interface{}{rule}
		}
	}

	return data, err
}

func (s *VolcengineTosBucketEncryptionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineTosBucketEncryptionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"SSEAlgorithm": {
				TargetField: "sse_algorithm",
			},
			"KMSDataEncryption": {
				TargetField: "kms_data_encryption",
			},
			"KMSMasterKeyID": {
				TargetField: "kms_master_key_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketEncryptionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateEncryption(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketEncryptionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateEncryption(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketEncryptionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketEncryption",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BucketName"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Domain:      (*call.SdkParam)["BucketName"].(string),
					UrlParam: map[string]string{
						"encryption": "",
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket encryption on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketEncryptionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketEncryptionService) createOrUpdateEncryption(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketEncryption",
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
				"rule": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "Rule",
					NextLevelConvert: map[string]ve.RequestConvert{
						"apply_server_side_encryption_by_default": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "ApplyServerSideEncryptionByDefault",
							NextLevelConvert: map[string]ve.RequestConvert{
								"sse_algorithm": {
									ConvertType: ve.ConvertDefault,
									TargetField: "SSEAlgorithm",
								},
								"kms_data_encryption": {
									ConvertType: ve.ConvertDefault,
									TargetField: "KMSDataEncryption",
								},
								"kms_master_key_id": {
									ConvertType: ve.ConvertDefault,
									TargetField: "KMSMasterKeyID",
								},
							},
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

				if rule, ok := sourceParam["Rule"].(map[string]interface{}); ok {
					if applyServerSideEncryptionByDefault, ok := rule["ApplyServerSideEncryptionByDefault"].(map[string]interface{}); ok {
						if sseAlgorithm, ok := applyServerSideEncryptionByDefault["SSEAlgorithm"]; ok {
							if sseAlgorithm != "kms" {
								delete(applyServerSideEncryptionByDefault, "KMSDataEncryption")
								delete(applyServerSideEncryptionByDefault, "KMSMasterKeyID")
							}
						}
					}
				}

				(*call.SdkParam)[ve.BypassParam] = sourceParam

				bytes, err := json.Marshal((*call.SdkParam)[ve.BypassParam].(map[string]interface{}))
				if err != nil {
					return false, err
				}
				hash := md5.New()
				io.WriteString(hash, string(bytes))
				contentMd5 := base64.StdEncoding.EncodeToString(hash.Sum(nil))

				(*call.SdkParam)[ve.BypassHeader].(map[string]string)["Content-MD5"] = contentMd5

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)

				param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod:  ve.PUT,
					ContentType: ve.ApplicationJSON,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					Header:      (*call.SdkParam)[ve.BypassHeader].(map[string]string),
					UrlParam: map[string]string{
						"encryption": "",
					},
				}, &param)

				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string))
				return nil
			},
		},
	}

	return callback
}

func (s *VolcengineTosBucketEncryptionService) ReadResourceId(id string) string {
	return id
}

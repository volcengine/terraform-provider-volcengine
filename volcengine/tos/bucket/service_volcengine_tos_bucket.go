package bucket

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketService struct {
	Client *ve.SdkClient
}

func NewTosBucketService(c *ve.SdkClient) *VolcengineTosBucketService {
	return &VolcengineTosBucketService{
		Client: c,
	}
}

func (s *VolcengineTosBucketService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		action  string
		resp    *map[string]interface{}
		results interface{}
	)
	action = "ListBuckets"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
	}, nil)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue(ve.BypassResponse+".Buckets", *resp)
	if err != nil {
		return data, err
	}
	data = results.([]interface{})
	return data, err
}

func (s *VolcengineTosBucketService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		action  string
		resp    *map[string]interface{}
		ok      bool
		header  http.Header
		acl     map[string]interface{}
		version map[string]interface{}
		tags    map[string]interface{}
		buckets []interface{}
	)

	if instanceId == "" {
		instanceId = s.ReadResourceId(resourceData.Id())
	} else {
		instanceId = s.ReadResourceId(instanceId)
	}

	action = "HeadBucket"
	logger.Debug(logger.ReqFormat, action, instanceId)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.HEAD,
		Domain:     instanceId,
	}, nil)
	logger.Debug(logger.ReqFormat, action, *resp)
	logger.Debug(logger.ReqFormat, action, err)
	if err != nil {
		return data, err
	}

	buckets, err = s.ReadResources(nil)
	if err != nil {
		return data, err
	}
	var (
		local interface{}
		name  interface{}
	)
	for _, bucket := range buckets {
		local, err = ve.ObtainSdkValue("Location", bucket)
		if err != nil {
			return data, err
		}
		name, err = ve.ObtainSdkValue("Name", bucket)
		if err != nil {
			return data, err
		}
		if local.(string) == s.Client.Region && name.(string) == instanceId {
			data = bucket.(map[string]interface{})
		}
	}
	if data == nil {
		data = make(map[string]interface{})
	}

	if header, ok = (*resp)[ve.BypassHeader].(http.Header); ok {
		if header.Get("X-Tos-Storage-Class") != "" {
			data["StorageClass"] = header.Get("X-Tos-Storage-Class")
		}
		if header.Get("X-Tos-Az-Redundancy") != "" {
			data["AzRedundancy"] = header.Get("X-Tos-Az-Redundancy")
		}
	}

	action = "GetBucketAcl"
	req := map[string]interface{}{
		"acl": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     instanceId,
	}, &req)
	if err != nil {
		return data, err
	}
	if acl, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); ok {
		data["PublicAcl"] = acl
		data["AccountAcl"] = acl
		data["BucketAclDelivered"] = acl["BucketAclDelivered"]
	}

	action = "GetBucketVersioning"
	req = map[string]interface{}{
		"versioning": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     instanceId,
	}, &req)
	if err != nil {
		return data, err
	}
	if version, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); ok {
		data["EnableVersion"] = version
	}

	action = "GetBucketTagging"
	req = map[string]interface{}{
		"tagging": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     instanceId,
		//Path:       []string{"?tagging="},
	}, &req)
	if err != nil && !ve.ResourceNotFoundError(err) {
		return data, err
	}
	if tags, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); ok {
		if tagSet, exist := tags["TagSet"]; exist {
			if tagMap, ok := tagSet.(map[string]interface{}); ok {
				data["Tags"] = tagMap["Tags"]
			}
		}
	}

	if len(data) == 0 {
		return data, fmt.Errorf("bucket %s not exist ", instanceId)
	}
	return data, nil
}

func (s *VolcengineTosBucketService) RefreshResourceState(data *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return data, "Success", err
		},
	}
}

//func (s *VolcengineTosBucketService) getIdPermission(p string, grants []interface{}) []interface{} {
//	var result []interface{}
//	for _, grant := range grants {
//		permission, _ := ve.ObtainSdkValue("Permission", grant)
//		id, _ := ve.ObtainSdkValue("Grantee.ID", grant)
//		t, _ := ve.ObtainSdkValue("Grantee.Type", grant)
//		if id != nil && t.(string) == "CanonicalUser" && p == permission.(string) {
//			result = append(result, "Id="+id.(string))
//		}
//	}
//	return result
//}

func (s *VolcengineTosBucketService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"EnableVersion": {
				Convert: func(i interface{}) interface{} {
					status, _ := ve.ObtainSdkValue("Status", i)
					return status.(string) == "Enabled"
				},
			},
			"AccountAcl": {
				Convert: ve.ConvertTosAccountAcl(),
			},
			"PublicAcl": {
				Convert: ve.ConvertTosPublicAcl(),
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	//create bucket
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "CreateBucket",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
				},
				"public_acl": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-acl",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
				"storage_class": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-storage-class",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
				"az_redundancy": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-az-redundancy",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
				"project_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-project-name",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建Bucket
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod: ve.PUT,
					Domain:     (*call.SdkParam)[ve.BypassDomain].(string),
					Header:     (*call.SdkParam)[ve.BypassHeader].(map[string]string),
				}, nil)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string))
				return nil
			},
		},
	}
	callbacks = append(callbacks, callback)

	//version
	callbackVersion := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketVersioning",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
				},
				"enable_version": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Status",
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						if i.(bool) {
							return "Enabled"
						} else {
							return ""
						}
					},
					ForceGet: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				//if disable version,skip this call
				if (*call.SdkParam)[ve.BypassParam].(map[string]interface{})["Status"] == "" {
					return false, nil
				}
				return true, nil
			},
			ExecuteCall: s.executePutBucketVersioning(),
		},
	}
	callbacks = append(callbacks, callbackVersion)

	//acl
	callbackAcl := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketAcl",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
				},
				"account_acl": {
					ConvertType: ve.ConvertListN,
					TargetField: "Grants",
					NextLevelConvert: map[string]ve.RequestConvert{
						"account_id": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Grantee.ID",
						},
						"acl_type": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Grantee.Type",
						},
						"permission": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Permission",
						},
					},
				},
				"bucket_acl_delivered": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketAclDelivered",
				},
			},
			BeforeCall:  s.beforePutBucketAcl(),
			ExecuteCall: s.executePutBucketAcl(),
			//Refresh: &ve.StateRefresh{
			//	Target:  []string{"Success"},
			//	Timeout: resourceData.Timeout(schema.TimeoutCreate),
			//},
		},
	}
	callbacks = append(callbacks, callbackAcl)

	//tags
	if _, ok := resourceData.GetOk("tags"); ok {
		callbackTags := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceBypass,
				Action:          "PutBucketTagging",
				ConvertMode:     ve.RequestConvertInConvert,
				ContentType:     ve.ContentTypeJson,
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
				BeforeCall:  s.beforePutBucketTagging(),
				ExecuteCall: s.executePutBucketTagging(),
			},
		}
		callbacks = append(callbacks, callbackTags)
	}

	return callbacks
}

func (s *VolcengineTosBucketService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if data.HasChange("enable_version") {
		//version
		callbackVersion := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceBypass,
				Action:          "PutBucketVersioning",
				ConvertMode:     ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"bucket_name": {
						ConvertType: ve.ConvertDefault,
						TargetField: "BucketName",
						SpecialParam: &ve.SpecialParam{
							Type: ve.DomainParam,
						},
						ForceGet: true,
					},
					"enable_version": {
						ConvertType: ve.ConvertDefault,
						TargetField: "Status",
						Convert: func(data *schema.ResourceData, i interface{}) interface{} {
							if i.(bool) {
								return "Enabled"
							} else {
								return "Suspended"
							}
						},
						ForceGet: true,
					},
				},
				ExecuteCall: s.executePutBucketVersioning(),
			},
		}
		callbacks = append(callbacks, callbackVersion)
	}
	var grant = []string{
		"public_acl",
		"account_acl",
		"bucket_acl_delivered",
	}
	for _, v := range grant {
		if data.HasChange(v) {
			callbackAcl := ve.Callback{
				Call: ve.SdkCall{
					ServiceCategory: ve.ServiceBypass,
					Action:          "PutBucketAcl",
					ConvertMode:     ve.RequestConvertInConvert,
					Convert: map[string]ve.RequestConvert{
						"bucket_name": {
							ConvertType: ve.ConvertDefault,
							TargetField: "BucketName",
							SpecialParam: &ve.SpecialParam{
								Type: ve.DomainParam,
							},
							ForceGet: true,
						},
						"account_acl": {
							ConvertType: ve.ConvertListN,
							TargetField: "Grants",
							NextLevelConvert: map[string]ve.RequestConvert{
								"account_id": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Grantee.ID",
									ForceGet:    true,
								},
								"acl_type": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Grantee.Type",
									ForceGet:    true,
								},
								"permission": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Permission",
									ForceGet:    true,
								},
							},
							ForceGet: true,
						},
						"bucket_acl_delivered": {
							ConvertType: ve.ConvertDefault,
							TargetField: "BucketAclDelivered",
							ForceGet:    true,
						},
					},
					BeforeCall:  s.beforePutBucketAcl(),
					ExecuteCall: s.executePutBucketAcl(),
					Refresh: &ve.StateRefresh{
						Target:  []string{"Success"},
						Timeout: data.Timeout(schema.TimeoutCreate),
					},
				},
			}
			callbacks = append(callbacks, callbackAcl)
			break
		}
	}

	if data.HasChange("tags") {
		callbacks = s.setResourceTags(data, callbacks)
	}

	return callbacks
}

func (s *VolcengineTosBucketService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	if _, ok := resourceData.GetOk("tags"); ok {
		addCallback := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceBypass,
				Action:          "PutBucketTagging",
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
				BeforeCall:  s.beforePutBucketTagging(),
				ExecuteCall: s.executePutBucketTagging(),
			},
		}
		callbacks = append(callbacks, addCallback)
	} else {
		removeCallback := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceBypass,
				Action:          "DeleteBucketTagging",
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
						Path:       []string{"?tagging="},
						UrlParam: map[string]string{
							"tagging": "",
						},
					}, nil)
				},
			},
		}
		callbacks = append(callbacks, removeCallback)
	}

	return callbacks
}

func (s *VolcengineTosBucketService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucket",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"BucketName": s.ReadResourceId(resourceData.Id()),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除Bucket
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod: ve.DELETE,
					Domain:     (*call.SdkParam)["BucketName"].(string),
				}, nil)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading tos on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {

	name, ok := data.GetOk("bucket_name")
	return ve.DataSourceInfo{
		ServiceCategory: ve.ServiceBypass,
		RequestConverts: map[string]ve.RequestConvert{
			"bucket_name": {
				Ignore: true,
			},
		},
		NameField:        "Name",
		IdField:          "BucketId",
		CollectField:     "buckets",
		ResponseConverts: map[string]ve.ResponseConvert{},
		ExtraData: func(sourceData []interface{}) (extraData []interface{}, err error) {
			for _, v := range sourceData {
				if v.(map[string]interface{})["Location"].(string) != s.Client.Region {
					continue
				}
				if ok {
					if name.(string) == v.(map[string]interface{})["Name"].(string) {
						v.(map[string]interface{})["BucketId"] = v.(map[string]interface{})["Name"].(string)
						extraData = append(extraData, v)
						break
					} else {
						continue
					}
				} else {
					v.(map[string]interface{})["BucketId"] = v.(map[string]interface{})["Name"].(string)
					extraData = append(extraData, v)
				}

			}
			return extraData, err
		},
	}
}

func (s *VolcengineTosBucketService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineTosBucketService) beforePutBucketAcl() ve.BeforeCallFunc {

	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
		data, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			HttpMethod: ve.GET,
			Domain:     (*call.SdkParam)[ve.BypassDomain].(string),
			UrlParam: map[string]string{
				"acl": "",
			},
		}, nil)
		return ve.BeforeTosPutAcl(d, call, data, err)
	}
}

func (s *VolcengineTosBucketService) executePutBucketAcl() ve.ExecuteCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//PutAcl
		param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
		return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			HttpMethod:  ve.PUT,
			ContentType: ve.ApplicationJSON,
			Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
			Header:      (*call.SdkParam)[ve.BypassHeader].(map[string]string),
			UrlParam: map[string]string{
				"acl": "",
			},
		}, &param)
	}
}

func (s *VolcengineTosBucketService) executePutBucketVersioning() ve.ExecuteCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//PutVersion
		condition := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
		return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.PUT,
			Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
			UrlParam: map[string]string{
				"versioning": "",
			},
		}, &condition)
	}
}

func (s *VolcengineTosBucketService) beforePutBucketTagging() ve.BeforeCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
		var tagsArr []interface{}
		tags := d.Get("tags")
		tagSet, ok := tags.(*schema.Set)
		if !ok {
			return false, fmt.Errorf("tags is not set")
		}
		for _, v := range tagSet.List() {
			tagMap, ok := v.(map[string]interface{})
			if !ok {
				return false, fmt.Errorf("tags value is not set")
			}
			tagsArr = append(tagsArr, map[string]interface{}{
				"Key":   tagMap["key"],
				"Value": tagMap["value"],
			})
		}
		tagsParam := make(map[string]interface{})
		tagsParam["Tags"] = tagsArr

		(*call.SdkParam)[ve.BypassParam].(map[string]interface{})["TagSet"] = tagsParam

		bytes, err := json.Marshal((*call.SdkParam)[ve.BypassParam].(map[string]interface{}))
		if err != nil {
			return false, err
		}
		hash := md5.New()
		io.WriteString(hash, string(bytes))
		contentMd5 := base64.StdEncoding.EncodeToString(hash.Sum(nil))

		(*call.SdkParam)[ve.BypassHeader].(map[string]string)["Content-MD5"] = contentMd5
		return true, nil
	}
}

func (s *VolcengineTosBucketService) executePutBucketTagging() ve.ExecuteCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//PutBucketTagging
		condition := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
		return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.PUT,
			Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
			Header:      (*call.SdkParam)[ve.BypassHeader].(map[string]string),
			//Path:        []string{"?tagging="},
			UrlParam: map[string]string{
				"tagging": "",
			},
		}, &condition)
	}
}

func (s *VolcengineTosBucketService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "tos",
		ResourceType:         "bucket",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

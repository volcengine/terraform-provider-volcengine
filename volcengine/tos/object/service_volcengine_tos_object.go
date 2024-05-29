package object

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosObjectService struct {
	Client *ve.SdkClient
}

func NewTosObjectService(c *ve.SdkClient) *VolcengineTosObjectService {
	return &VolcengineTosObjectService{
		Client: c,
	}
}

func (s *VolcengineTosObjectService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosObjectService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		action  string
		resp    *map[string]interface{}
		results interface{}
	)
	action = "ListObjects"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     condition[ve.BypassDomain].(string),
	}, nil)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue(ve.BypassResponse+".Contents", *resp)
	if err != nil {
		return data, err
	}
	data = results.([]interface{})
	return data, err
}

func (s *VolcengineTosObjectService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	bucketName := resourceData.Get("bucket_name").(string)
	var (
		action        string
		resp          *map[string]interface{}
		respBody      *map[string]interface{}
		ok            bool
		header        http.Header
		acl           map[string]interface{}
		bucketVersion map[string]interface{}
		tags          map[string]interface{}
	)

	if instanceId == "" {
		instanceId = s.ReadResourceId(resourceData.Id())
	} else {
		instanceId = s.ReadResourceId(instanceId)
	}

	action = "HeadObject"
	logger.Debug(logger.ReqFormat, action, bucketName+":"+instanceId)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.HEAD,
		Domain:     bucketName,
		Path:       []string{instanceId},
	}, nil)
	if err != nil {
		return data, err
	}
	data = make(map[string]interface{})

	if header, ok = (*resp)[ve.BypassHeader].(http.Header); ok {
		if header.Get("X-Tos-Storage-Class") != "" {
			data["StorageClass"] = header.Get("x-tos-storage-class")
		}
		if header.Get("Content-Type") != "" {
			data["ContentType"] = header.Get("Content-Type")
			if strings.Contains(strings.ToLower(header.Get("Content-Type")), "application/json") ||
				strings.Contains(strings.ToLower(header.Get("Content-Type")), "application/xml") ||
				strings.Contains(strings.ToLower(header.Get("Content-Type")), "text/plain") {
				action = "GetObject"
				logger.Debug(logger.ReqFormat, action, bucketName+":"+instanceId)
				respBody, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod: ve.GET,
					Domain:     bucketName,
					Path:       []string{instanceId},
				}, nil)
				if err != nil {
					return data, err
				}
				data["Content"] = (*respBody)[ve.BypassResponseData]
			}
		}
		if header.Get("X-Tos-Server-Side-Encryption") != "" {
			data["Encryption"] = header.Get("X-Tos-Server-Side-Encryption")
		}
		if header.Get("x-tos-meta-content-md5") != "" {
			data["ContentMd5"] = strings.Replace(header.Get("x-tos-meta-content-md5"), "\"", "", -1)
		}

		if header.Get("X-Tos-Version-Id") != "" {
			action = "ListObjects"
			logger.Debug(logger.ReqFormat, action, bucketName+":"+instanceId)

			var (
				nextVersionIdMarker string
				versionIds          []string
			)

			for {
				urlParam := map[string]string{
					"prefix":   instanceId,
					"max-keys": "100",
					"versions": "",
				}
				if nextVersionIdMarker != "" {
					urlParam["key-marker"] = instanceId
					urlParam["version-id-marker"] = nextVersionIdMarker
				}

				resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod: ve.GET,
					Domain:     bucketName,
					UrlParam:   urlParam,
				}, nil)

				if err != nil {
					return data, err
				}
				versions, _ := ve.ObtainSdkValue(ve.BypassResponse+".Versions", *resp)
				next, _ := ve.ObtainSdkValue(ve.BypassResponse+".NextVersionIdMarker", *resp)

				if versions == nil || len(versions.([]interface{})) == 0 {
					break
				}

				if next == nil || next.(string) == "" {
					nextVersionIdMarker = ""
				} else {
					nextVersionIdMarker = next.(string)
				}

				for _, version := range versions.([]interface{}) {
					versionId, _ := ve.ObtainSdkValue("VersionId", version)
					versionIds = append(versionIds, versionId.(string))
				}

				if nextVersionIdMarker == "" {
					break
				}
			}
			logger.Debug(logger.ReqFormat, action, versionIds)
			data["VersionIds"] = versionIds
		}
	}

	action = "GetObjectAcl"
	req := map[string]interface{}{
		"acl": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     bucketName,
		Path:       []string{instanceId},
	}, &req)
	if err != nil {
		return data, err
	}
	if acl, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); ok {
		data["PublicAcl"] = acl
		data["AccountAcl"] = acl
	}

	action = "GetBucketVersioning"
	req = map[string]interface{}{
		"versioning": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     bucketName,
	}, &req)
	if err != nil {
		return data, err
	}
	if bucketVersion, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); ok {
		data["EnableVersion"] = bucketVersion
	}

	action = "GetObjectTagging"
	req = map[string]interface{}{
		"tagging": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     bucketName,
		//Path:       []string{instanceId, "?tagging="},
		Path: []string{instanceId},
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
		return data, fmt.Errorf("object %s not exist ", instanceId)
	}
	return data, nil
}

func (s *VolcengineTosObjectService) RefreshResourceState(data *schema.ResourceData, target []string, timeout time.Duration, instanceId string) *resource.StateChangeConf {
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

func (VolcengineTosObjectService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
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

func (s *VolcengineTosObjectService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	//create object
	callback := s.createOrReplaceObject(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	//acl
	callbackAcl := s.createOrUpdateObjectAcl(resourceData, resource, false)
	callbacks = append(callbacks, callbackAcl)

	//tags
	if _, ok := resourceData.GetOk("tags"); ok {
		callbackTags := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceBypass,
				Action:          "PutObjectTagging",
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
					"object_name": {
						ConvertType: ve.ConvertDefault,
						TargetField: "ObjectName",
						ForceGet:    true,
						SpecialParam: &ve.SpecialParam{
							Type:  ve.PathParam,
							Index: 0,
						},
					},
				},
				BeforeCall:  s.beforePutObjectTagging(),
				ExecuteCall: s.executePutObjectTagging(),
			},
		}
		callbacks = append(callbacks, callbackTags)
	}

	return callbacks
}

func (s *VolcengineTosObjectService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if data.HasChange("file_path") || data.HasChanges("content_md5") || data.HasChanges("content") {
		callbacks = append(callbacks, s.createOrReplaceObject(data, resource, true))
		callbacks = append(callbacks, s.createOrUpdateObjectAcl(data, resource, true))
		callbacks = s.setResourceTags(data, callbacks)
	} else {
		var grant = []string{
			"public_acl",
			"account_acl",
		}
		for _, v := range grant {
			if data.HasChange(v) {
				callbackAcl := s.createOrUpdateObjectAcl(data, resource, true)
				callbacks = append(callbacks, callbackAcl)
				break
			}
		}

		if data.HasChange("tags") {
			callbacks = s.setResourceTags(data, callbacks)
		}
	}

	return callbacks
}

func (s *VolcengineTosObjectService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	if _, ok := resourceData.GetOk("tags"); ok {
		addCallback := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceBypass,
				Action:          "PutObjectTagging",
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
					"object_name": {
						ConvertType: ve.ConvertDefault,
						TargetField: "ObjectName",
						ForceGet:    true,
						SpecialParam: &ve.SpecialParam{
							Type:  ve.PathParam,
							Index: 0,
						},
					},
				},
				BeforeCall:  s.beforePutObjectTagging(),
				ExecuteCall: s.executePutObjectTagging(),
			},
		}
		callbacks = append(callbacks, addCallback)
	} else {
		removeCallback := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceBypass,
				Action:          "DeleteObjectTagging",
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
					"object_name": {
						ConvertType: ve.ConvertDefault,
						TargetField: "ObjectName",
						ForceGet:    true,
						SpecialParam: &ve.SpecialParam{
							Type:  ve.PathParam,
							Index: 0,
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)[ve.BypassPath] = append((*call.SdkParam)[ve.BypassPath].([]string), "?tagging=")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
						HttpMethod: ve.DELETE,
						Domain:     (*call.SdkParam)[ve.BypassDomain].(string),
						Path:       (*call.SdkParam)[ve.BypassPath].([]string),
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

func (s *VolcengineTosObjectService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteObject",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"BucketName": resourceData.Get("bucket_name"),
				"ObjectName": s.ReadResourceId(resourceData.Id()),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {

				if d.Get("version_ids") != nil && len(d.Get("version_ids").(*schema.Set).List()) > 0 {
					for _, vv := range d.Get("version_ids").(*schema.Set).List() {
						condition := make(map[string]interface{})
						condition["versionId"] = vv
						//remove Object-with-version
						logger.Debug(logger.RespFormat, call.Action, condition)
						_, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
							HttpMethod: ve.DELETE,
							Domain:     (*call.SdkParam)["BucketName"].(string),
							Path:       []string{(*call.SdkParam)["ObjectName"].(string), fmt.Sprintf("?versionId=%s", vv.(string))},
						}, &condition)
						if err != nil {
							return nil, err
						}
					}
				} else {
					//remove Object-no-version
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
						HttpMethod: ve.DELETE,
						Domain:     (*call.SdkParam)["BucketName"].(string),
						Path:       []string{(*call.SdkParam)["ObjectName"].(string)},
					}, nil)
				}

				return nil, nil
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading tos object on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosObjectService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	name, ok := data.GetOk("object_name")
	bucketName, _ := data.GetOk("bucket_name")
	return ve.DataSourceInfo{
		ServiceCategory: ve.ServiceBypass,
		RequestConverts: map[string]ve.RequestConvert{
			"bucket_name": {
				ConvertType: ve.ConvertDefault,
				SpecialParam: &ve.SpecialParam{
					Type: ve.DomainParam,
				},
			},
			"object_name": {
				Ignore: true,
			},
		},
		NameField:    "Key",
		IdField:      "ObjectId",
		CollectField: "objects",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Key": {
				TargetField: "name",
			},
		},
		ExtraData: func(sourceData []interface{}) (extraData []interface{}, err error) {
			for _, v := range sourceData {
				if ok {
					if name.(string) == v.(map[string]interface{})["Key"].(string) {
						v.(map[string]interface{})["ObjectId"] = bucketName.(string) + ":" + v.(map[string]interface{})["Key"].(string)
						extraData = append(extraData, v)
						break
					} else {
						continue
					}
				} else {
					v.(map[string]interface{})["ObjectId"] = bucketName.(string) + ":" + v.(map[string]interface{})["Key"].(string)
					extraData = append(extraData, v)
				}

			}
			return extraData, err
		},

		EachResource: func(sourceData []interface{}, d *schema.ResourceData) ([]interface{}, error) {
			var newSourceData []interface{}
			for _, v := range sourceData {
				var (
					key     interface{}
					newData map[string]interface{}
					err     error
				)
				key, err = ve.ObtainSdkValue("Key", v)
				if err != nil {
					return nil, err
				}

				if str, ok1 := key.(string); ok1 {
					newData, err = s.ReadResource(d, str)
					if err != nil {
						return nil, err
					}
				}

				if v1, ok1 := v.(map[string]interface{}); ok1 {
					for k, value := range newData {
						if _, ok2 := v1[k]; !ok2 {
							v1[k] = value
						}
					}
					newSourceData = append(newSourceData, v1)
				}
			}
			return newSourceData, nil
		},
	}
}

func (s *VolcengineTosObjectService) ReadResourceId(id string) string {
	return id[strings.Index(id, ":")+1:]
}

func (s *VolcengineTosObjectService) beforePutObjectAcl() ve.BeforeCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		data, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			HttpMethod: ve.GET,
			Domain:     (*call.SdkParam)[ve.BypassDomain].(string),
			Path:       (*call.SdkParam)[ve.BypassPath].([]string),
			UrlParam: map[string]string{
				"acl": "",
			},
		}, nil)
		return ve.BeforeTosPutAcl(d, call, data, err)
	}
}

func (s *VolcengineTosObjectService) executePutObjectAcl() ve.ExecuteCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//PutAcl
		param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
		return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			HttpMethod:  ve.PUT,
			ContentType: ve.ApplicationJSON,
			Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
			Path:        (*call.SdkParam)[ve.BypassPath].([]string),
			UrlParam: map[string]string{
				"acl": "",
			},
		}, &param)
	}
}

func (s *VolcengineTosObjectService) createOrUpdateObjectAcl(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutObjectAcl",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
					ForceGet: isUpdate,
				},
				"object_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "ObjectName",
					SpecialParam: &ve.SpecialParam{
						Type:  ve.PathParam,
						Index: 0,
					},
					ForceGet: isUpdate,
				},
				"account_acl": {
					ConvertType: ve.ConvertListN,
					TargetField: "Grants",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"account_id": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Grantee.ID",
							ForceGet:    isUpdate,
						},
						"acl_type": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Grantee.Type",
							ForceGet:    isUpdate,
						},
						"permission": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Permission",
							ForceGet:    isUpdate,
						},
					},
				},
			},
			BeforeCall:  s.beforePutObjectAcl(),
			ExecuteCall: s.executePutObjectAcl(),
			//Refresh: &ve.StateRefresh{
			//	Target:  []string{"Success"},
			//	Timeout: resourceData.Timeout(schema.TimeoutCreate),
			//},
		},
	}
	//如果出现acl缓存的问题 这里再打开 暂时去掉 不再等待60s
	//if isUpdate && !resourceData.HasChange("file_path") && !resourceData.HasChanges("content_md5") {
	//	callback.Call.Refresh = &ve.StateRefresh{
	//		Target:  []string{"Success"},
	//		Timeout: resourceData.Timeout(schema.TimeoutCreate),
	//	}
	//}
	return callback
}

func (s *VolcengineTosObjectService) createOrReplaceObject(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	name := fmt.Sprintf("./%d-temp", time.Now().Unix())
	return ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutObject",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
					ForceGet: isUpdate,
				},
				"object_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "ObjectName",
					SpecialParam: &ve.SpecialParam{
						Type:  ve.PathParam,
						Index: 0,
					},
					ForceGet: isUpdate,
				},
				"public_acl": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-acl",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
					ForceGet: isUpdate,
				},
				"storage_class": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-storage-class",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
					ForceGet: isUpdate,
				},
				"content_type": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Content-Type",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
					ForceGet: isUpdate,
				},
				"content": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Content",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
				"content_md5": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Content-MD5",
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						b, _ := hex.DecodeString(i.(string))
						return base64.StdEncoding.EncodeToString(b)
					},
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
					ForceGet: isUpdate,
				},
				"file_path": {
					ConvertType: ve.ConvertDefault,
					TargetField: "file-path",
					SpecialParam: &ve.SpecialParam{
						Type: ve.FilePathParam,
					},
					ForceGet: isUpdate,
				},
				"encryption": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-server-side-encryption",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
					ForceGet: isUpdate,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if _, ok := (*call.SdkParam)[ve.BypassHeader].(map[string]string)["Content-MD5"]; ok {
					(*call.SdkParam)[ve.BypassHeader].(map[string]string)["x-tos-meta-content-md5"] = d.Get("content_md5").(string)
				}
				if _, ok := (*call.SdkParam)[ve.BypassHeader].(map[string]string)["Content"]; ok {
					content := []byte(d.Get("content").(string))
					err := os.WriteFile(name, content, 0644)
					if err != nil {
						return false, err
					}
					(*call.SdkParam)[ve.BypassFilePath] = name
					delete((*call.SdkParam)[ve.BypassHeader].(map[string]string), "Content")
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建Object
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					Header:      (*call.SdkParam)[ve.BypassHeader].(map[string]string),
					Path:        (*call.SdkParam)[ve.BypassPath].([]string),
					ContentPath: (*call.SdkParam)[ve.BypassFilePath].(string),
				}, nil)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string) + ":" + (*call.SdkParam)[ve.BypassPath].([]string)[0])
				_, err := os.Stat(name)
				if err == nil {
					return os.Remove(name)
				}
				if os.IsNotExist(err) {
					return nil
				}
				return err
			},
		},
	}
}

func (s *VolcengineTosObjectService) beforePutObjectTagging() ve.BeforeCallFunc {
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

		//(*call.SdkParam)[ve.BypassPath] = append((*call.SdkParam)[ve.BypassPath].([]string), "?tagging=")
		return true, nil
	}
}

func (s *VolcengineTosObjectService) executePutObjectTagging() ve.ExecuteCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//PutObjectTagging
		condition := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
		return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.PUT,
			Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
			Header:      (*call.SdkParam)[ve.BypassHeader].(map[string]string),
			Path:        (*call.SdkParam)[ve.BypassPath].([]string),
			UrlParam: map[string]string{
				"tagging": "",
			},
		}, &condition)
	}
}

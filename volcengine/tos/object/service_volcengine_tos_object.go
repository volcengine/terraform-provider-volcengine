package object

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
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
	tos := s.Client.TosClient
	var (
		action  string
		resp    *map[string]interface{}
		results interface{}
	)
	action = "ListObjects"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.GET,
		Domain:     condition[ve.TosDomain].(string),
	}, nil)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue(ve.TosResponse+".Contents", *resp)
	if err != nil {
		return data, err
	}
	data = results.([]interface{})
	return data, err
}

func (s *VolcengineTosObjectService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	tos := s.Client.TosClient
	bucketName := resourceData.Get("bucket_name").(string)
	var (
		action        string
		resp          *map[string]interface{}
		ok            bool
		header        http.Header
		acl           map[string]interface{}
		bucketVersion map[string]interface{}
	)

	if instanceId == "" {
		instanceId = s.ReadResourceId(resourceData.Id())
	} else {
		instanceId = s.ReadResourceId(instanceId)
	}

	action = "HeadObject"
	logger.Debug(logger.ReqFormat, action, bucketName+":"+instanceId)
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.HEAD,
		Domain:     bucketName,
		Path:       []string{instanceId},
	}, nil)
	if err != nil {
		return data, err
	}
	data = make(map[string]interface{})

	if header, ok = (*resp)[ve.TosHeader].(http.Header); ok {
		if header.Get("X-Tos-Storage-Class") != "" {
			data["StorageClass"] = header.Get("x-tos-storage-class")
		}
		if header.Get("Content-Type") != "" {
			data["ContentType"] = header.Get("Content-Type")
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

				resp, err = tos.DoTosCall(ve.TosInfo{
					HttpMethod: ve.GET,
					Domain:     bucketName,
					UrlParam:   urlParam,
				}, nil)

				if err != nil {
					return data, err
				}
				versions, _ := ve.ObtainSdkValue(ve.TosResponse+".Versions", *resp)
				next, _ := ve.ObtainSdkValue(ve.TosResponse+".NextVersionIdMarker", *resp)

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
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.GET,
		Domain:     bucketName,
		Path:       []string{instanceId},
	}, &req)
	if err != nil {
		return data, err
	}
	if acl, ok = (*resp)[ve.TosResponse].(map[string]interface{}); ok {
		data["PublicAcl"] = acl
		data["AccountAcl"] = acl
	}

	action = "GetBucketVersioning"
	req = map[string]interface{}{
		"versioning": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.GET,
		Domain:     bucketName,
	}, &req)
	if err != nil {
		return data, err
	}
	if bucketVersion, ok = (*resp)[ve.TosResponse].(map[string]interface{}); ok {
		data["EnableVersion"] = bucketVersion
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
	//create object
	callback := s.createOrReplaceObject(resourceData, resource, false)
	//acl
	callbackAcl := s.createOrUpdateObjectAcl(resourceData, resource, false)
	return []ve.Callback{callback, callbackAcl}
}

func (s *VolcengineTosObjectService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if data.HasChange("file_path") || data.HasChanges("content_md5") {
		callbacks = append(callbacks, s.createOrReplaceObject(data, resource, true))
		callbacks = append(callbacks, s.createOrUpdateObjectAcl(data, resource, true))
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
						_, err := s.Client.TosClient.DoTosCall(ve.TosInfo{
							HttpMethod: ve.DELETE,
							Domain:     (*call.SdkParam)["BucketName"].(string),
							Path:       []string{(*call.SdkParam)["ObjectName"].(string)},
						}, &condition)
						if err != nil {
							return nil, err
						}
					}
				} else {
					//remove Object-no-version
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.TosClient.DoTosCall(ve.TosInfo{
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
		ServiceCategory: ve.ServiceTos,
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
	}
}

func (s *VolcengineTosObjectService) ReadResourceId(id string) string {
	return id[strings.Index(id, ":")+1:]
}

func (s *VolcengineTosObjectService) beforePutObjectAcl() ve.BeforeCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		data, err := s.Client.TosClient.DoTosCall(ve.TosInfo{
			HttpMethod: ve.GET,
			Domain:     (*call.SdkParam)[ve.TosDomain].(string),
			Path:       (*call.SdkParam)[ve.TosPath].([]string),
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
		param := (*call.SdkParam)[ve.TosParam].(map[string]interface{})
		return s.Client.TosClient.DoTosCall(ve.TosInfo{
			HttpMethod:  ve.PUT,
			ContentType: ve.ApplicationJSON,
			Domain:      (*call.SdkParam)[ve.TosDomain].(string),
			Path:        (*call.SdkParam)[ve.TosPath].([]string),
			UrlParam: map[string]string{
				"acl": "",
			},
		}, &param)
	}
}

func (s *VolcengineTosObjectService) createOrUpdateObjectAcl(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceTos,
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
	return ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceTos,
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
				if _, ok := (*call.SdkParam)[ve.TosHeader].(map[string]string)["Content-MD5"]; ok {
					(*call.SdkParam)[ve.TosHeader].(map[string]string)["x-tos-meta-content-md5"] = d.Get("content_md5").(string)
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建Object
				return s.Client.TosClient.DoTosCall(ve.TosInfo{
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)[ve.TosDomain].(string),
					Header:      (*call.SdkParam)[ve.TosHeader].(map[string]string),
					Path:        (*call.SdkParam)[ve.TosPath].([]string),
					ContentPath: (*call.SdkParam)[ve.TosFilePath].(string),
				}, nil)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.TosDomain].(string) + ":" + (*call.SdkParam)[ve.TosPath].([]string)[0])
				return nil
			},
		},
	}
}

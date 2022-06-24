package object

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackTosObjectService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosObjectService(c *ve.SdkClient) *VestackTosObjectService {
	return &VestackTosObjectService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackTosObjectService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackTosObjectService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
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

func (s *VestackTosObjectService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	tos := s.Client.TosClient
	bucketName := resourceData.Get("bucket_name").(string)
	var (
		action string
		resp   *map[string]interface{}
		ok     bool
		header http.Header
		acl    map[string]interface{}
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

	if len(data) == 0 {
		return data, fmt.Errorf("object %s not exist ", instanceId)
	}
	return data, nil
}

func (s *VestackTosObjectService) RefreshResourceState(data *schema.ResourceData, target []string, timeout time.Duration, instanceId string) *resource.StateChangeConf {
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

func (VestackTosObjectService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
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

func (s *VestackTosObjectService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	//create object
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceTos,
			Action:          "CreateObject",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
				},
				"object_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "ObjectName",
					SpecialParam: &ve.SpecialParam{
						Type:  ve.PathParam,
						Index: 0,
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
				"content_type": {
					ConvertType: ve.ConvertDefault,
					TargetField: "content-type",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
				"file_path": {
					ConvertType: ve.ConvertDefault,
					TargetField: "file-path",
					SpecialParam: &ve.SpecialParam{
						Type: ve.FilePathParam,
					},
				},
				"encryption": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-server-side-encryption",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
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
	//acl
	callbackAcl := ve.Callback{
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
				},
				"object_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "ObjectName",
					SpecialParam: &ve.SpecialParam{
						Type:  ve.PathParam,
						Index: 0,
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
			},
			BeforeCall:  s.beforePutObjectAcl(),
			ExecuteCall: s.executePutObjectAcl(),
			//Refresh: &ve.StateRefresh{
			//	Target:  []string{"Success"},
			//	Timeout: resourceData.Timeout(schema.TimeoutCreate),
			//},
		},
	}
	return []ve.Callback{callback, callbackAcl}
}

func (s *VestackTosObjectService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	var grant = []string{
		"public_acl",
		"account_acl",
	}
	for _, v := range grant {
		if data.HasChange(v) {
			callbackAcl := ve.Callback{
				Call: ve.SdkCall{
					ServiceCategory: ve.ServiceTos,
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
						"object_name": {
							ConvertType: ve.ConvertDefault,
							TargetField: "ObjectName",
							SpecialParam: &ve.SpecialParam{
								Type:  ve.PathParam,
								Index: 0,
							},
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
					},
					BeforeCall:  s.beforePutObjectAcl(),
					ExecuteCall: s.executePutObjectAcl(),
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

	return callbacks
}

func (s *VestackTosObjectService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
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

func (s *VestackTosObjectService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
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

func (s *VestackTosObjectService) ReadResourceId(id string) string {
	return id[strings.Index(id, ":")+1:]
}

func (s *VestackTosObjectService) beforePutObjectAcl() ve.BeforeCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
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

func (s *VestackTosObjectService) executePutObjectAcl() ve.ExecuteCallFunc {
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

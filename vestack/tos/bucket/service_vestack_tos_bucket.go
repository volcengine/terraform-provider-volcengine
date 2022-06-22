package bucket

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

type VestackTosBucketService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketService(c *ve.SdkClient) *VestackTosBucketService {
	return &VestackTosBucketService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackTosBucketService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackTosBucketService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	tos := s.Client.TosClient
	var (
		action  string
		resp    *map[string]interface{}
		results interface{}
	)
	action = "ListBuckets"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.GET,
	}, nil)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue(ve.TosResponse+".Buckets", *resp)
	if err != nil {
		return data, err
	}
	data = results.([]interface{})
	return data, err
}

func (s *VestackTosBucketService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	tos := s.Client.TosClient
	var (
		action  string
		resp    *map[string]interface{}
		ok      bool
		header  *http.Header
		acl     map[string]interface{}
		version map[string]interface{}
	)

	if instanceId == "" {
		instanceId = s.ReadResourceId(resourceData.Id())
	} else {
		instanceId = s.ReadResourceId(instanceId)
	}

	action = "HeadBucket"
	logger.Debug(logger.ReqFormat, action, instanceId)
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.HEAD,
		Domain:     instanceId,
	}, nil)
	if err != nil {
		return data, err
	}
	data = make(map[string]interface{})

	if header, ok = (*resp)[ve.TosHeader].(*http.Header); ok {
		if header.Get("x-tos-storage-class") != "" {
			data["TosStorageClass"] = header.Get("x-tos-storage-class")
		}
	}

	action = "GetBucketAcl"
	req := map[string]interface{}{
		"acl": "",
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.GET,
		Domain:     instanceId,
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
		Domain:     instanceId,
	}, &req)
	if err != nil {
		return data, err
	}
	if version, ok = (*resp)[ve.TosResponse].(map[string]interface{}); ok {
		data["EnableVersion"] = version
	}

	if len(data) == 0 {
		return data, fmt.Errorf("bucket %s not exist ", instanceId)
	}
	return data, nil
}

func (s *VestackTosBucketService) RefreshResourceState(data *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VestackTosBucketService) getIdPermission(p string, grants []interface{}) []interface{} {
	var result []interface{}
	for _, grant := range grants {
		permission, _ := ve.ObtainSdkValue("Permission", grant)
		id, _ := ve.ObtainSdkValue("Grantee.ID", grant)
		t, _ := ve.ObtainSdkValue("Grantee.Type", grant)
		if id != nil && t.(string) == "CanonicalUser" && p == permission.(string) {
			result = append(result, "Id="+id.(string))
		}
	}
	return result
}

func (s *VestackTosBucketService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"EnableVersion": {
				Convert: func(i interface{}) interface{} {
					status, _ := ve.ObtainSdkValue("Status", i)
					if status.(string) != "Enabled" {
						return false
					}
					return true
				},
			},
			"AccountAcl": {
				Convert: func(i interface{}) interface{} {
					var accountAcl []interface{}
					owner, _ := ve.ObtainSdkValue("Owner.ID", i)
					grants, _ := ve.ObtainSdkValue("Grants", i)
					for _, grant := range grants.([]interface{}) {
						permission, _ := ve.ObtainSdkValue("Permission", grant)
						id, _ := ve.ObtainSdkValue("Grantee.ID", grant)
						if id == nil {
							continue
						}
						if id == owner && permission == "FULL_CONTROL" {
							continue
						}
						g := map[string]interface{}{
							"AccountId":  id,
							"AclType":    "CanonicalUser",
							"Permission": permission,
						}
						accountAcl = append(accountAcl, g)
					}
					return accountAcl
				},
			},
			"PublicAcl": {
				Convert: func(i interface{}) interface{} {
					owner, _ := ve.ObtainSdkValue("Owner.ID", i)
					grants, _ := ve.ObtainSdkValue("Grants", i)
					var (
						read  bool
						write bool
					)
					for _, grant := range grants.([]interface{}) {
						id, _ := ve.ObtainSdkValue("Grantee.ID", grant)
						canned, _ := ve.ObtainSdkValue("Grantee.Canned", grant)
						t, _ := ve.ObtainSdkValue("Grantee.Type", grant)
						permission, _ := ve.ObtainSdkValue("Permission", grant)
						if canned != nil && canned.(string) == "AllUsers" && t.(string) == "Group" {
							if permission.(string) == "READ" {
								read = true
								continue
							} else if permission.(string) == "WRITE" {
								write = true
								continue
							}
						}

						if canned != nil && canned.(string) == "AuthenticatedUsers" && t.(string) == "Group" {
							if permission.(string) == "READ" {
								return "authenticated-read"
							}
							break
						}

						if id != nil && id.(string) == owner.(string) && t.(string) == "CanonicalUser" {
							if permission.(string) == "FULL_CONTROL" {
								return "private"
							} else if permission.(string) == "READ" {
								return "bucket-owner-read"
							}
							break

						}

					}
					if read && !write {
						return "public-read"
					}
					if read && write {
						return "public-read-write"
					}
					return ""
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VestackTosBucketService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	//create bucket
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceTos,
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
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建Bucket
				return s.Client.TosClient.DoTosCall(ve.TosInfo{
					HttpMethod: ve.PUT,
					Domain:     (*call.SdkParam)[ve.TosDomain].(string),
					Header:     (*call.SdkParam)[ve.TosHeader].(map[string]string),
				}, nil)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(s.Client.Region + ":" + (*call.SdkParam)[ve.TosDomain].(string))
				return nil
			},
		},
	}
	//version
	callbackVersion := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceTos,
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
				if (*call.SdkParam)[ve.TosParam].(map[string]interface{})["Status"] == "" {
					return false, nil
				}
				return true, nil
			},
			ExecuteCall: s.executePutBucketVersioning(),
		},
	}
	//acl
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
			BeforeCall:  s.beforePutBucketAcl(),
			ExecuteCall: s.executePutBucketAcl(),
		},
	}
	return []ve.Callback{callback, callbackVersion, callbackAcl}
}

func (s *VestackTosBucketService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if data.HasChange("enable_version") {
		//version
		callbackVersion := ve.Callback{
			Call: ve.SdkCall{
				ServiceCategory: ve.ServiceTos,
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
					BeforeCall:  s.beforePutBucketAcl(),
					ExecuteCall: s.executePutBucketAcl(),
				},
			}
			callbacks = append(callbacks, callbackAcl)
			break
		}
	}

	return callbacks
}

func (s *VestackTosBucketService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
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
				return s.Client.TosClient.DoTosCall(ve.TosInfo{
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

func (s *VestackTosBucketService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {

	name, ok := data.GetOk("bucket_name")
	return ve.DataSourceInfo{
		ServiceCategory: ve.ServiceTos,
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
						v.(map[string]interface{})["BucketId"] = s.Client.Region + ":" + v.(map[string]interface{})["Name"].(string)
						extraData = append(extraData, v)
						break
					} else {
						continue
					}
				} else {
					v.(map[string]interface{})["BucketId"] = s.Client.Region + ":" + v.(map[string]interface{})["Name"].(string)
					extraData = append(extraData, v)
				}

			}
			return extraData, err
		},
	}
}

func (s *VestackTosBucketService) ReadResourceId(id string) string {
	if strings.HasPrefix(id, s.Client.Region+":") {
		return id[strings.Index(id, ":")+1:]
	}
	return id
}

func (s *VestackTosBucketService) mergePublicAcl(acl string, param *map[string]interface{}, ownerId string) {
	if _, ok := (*param)["Grants"]; !ok {
		(*param)["Grants"] = []interface{}{}
	}
	vs := (*param)["Grants"].([]interface{})

	defer func() {
		(*param)["Grants"] = vs
	}()

	switch acl {
	case "private":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Id":   ownerId,
				"Type": "CanonicalUser",
			},
			"Permission": "FULL_CONTROL",
		}
		vs = append(vs, m)
		return
	case "public-read":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Canned": "AllUsers",
				"Type":   "Group",
			},
			"Permission": "READ",
		}
		vs = append(vs, m)
		return
	case "public-read-write":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Canned": "AllUsers",
				"Type":   "Group",
			},
			"Permission": "WRITE",
		}
		vs = append(vs, m)
		return
	case "authenticated-read":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Canned": "AuthenticatedUsers",
				"Type":   "Group",
			},
			"Permission": "READ",
		}
		vs = append(vs, m)
		return
	case "bucket-owner-read":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Id":   ownerId,
				"Type": "CanonicalUser",
			},
			"Permission": "READ",
		}
		vs = append(vs, m)
		return
	}
}

func (s *VestackTosBucketService) beforePutBucketAcl() ve.BeforeCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
		//read acl before call and merge all acl
		data, err := s.Client.TosClient.DoTosCall(ve.TosInfo{
			HttpMethod: ve.GET,
			Domain:     (*call.SdkParam)[ve.TosDomain].(string),
			UrlParam: map[string]string{
				"acl": "",
			},
		}, nil)
		if err != nil {
			return false, err
		}

		sourceAclParam := ve.SortAndStartTransJson((*call.SdkParam)[ve.TosParam].(map[string]interface{}))
		ownerId, _ := ve.ObtainSdkValue("Owner.ID", (*data)[ve.TosResponse])

		grants, _ := ve.ObtainSdkValue("Grants", sourceAclParam)
		for _, grant := range grants.([]interface{}) {
			id, _ := ve.ObtainSdkValue("Grantee.ID", grant)
			p, _ := ve.ObtainSdkValue("Permission", grant)
			if id == ownerId && p == "FULL_CONTROL" {
				return false, fmt.Errorf("can not set FULL_CONTROL for owner")
			}
		}

		//merge owner
		owner, _ := ve.ObtainSdkValue("Owner", (*data)[ve.TosResponse])
		sourceAclParam["Owner"] = owner
		//merge public_acl
		s.mergePublicAcl(d.Get("public_acl").(string), &sourceAclParam, ownerId.(string))

		(*call.SdkParam)[ve.TosParam] = sourceAclParam
		return true, nil
	}
}

func (s *VestackTosBucketService) executePutBucketAcl() ve.ExecuteCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//PutAcl
		param := (*call.SdkParam)[ve.TosParam].(map[string]interface{})
		return s.Client.TosClient.DoTosCall(ve.TosInfo{
			HttpMethod:  ve.PUT,
			ContentType: ve.ApplicationJSON,
			Domain:      (*call.SdkParam)[ve.TosDomain].(string),
			Header:      (*call.SdkParam)[ve.TosHeader].(map[string]string),
			UrlParam: map[string]string{
				"acl": "",
			},
		}, &param)
	}
}

func (s *VestackTosBucketService) executePutBucketVersioning() ve.ExecuteCallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//PutVersion
		condition := (*call.SdkParam)[ve.TosParam].(map[string]interface{})
		return s.Client.TosClient.DoTosCall(ve.TosInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.PUT,
			Domain:      (*call.SdkParam)[ve.TosDomain].(string),
			UrlParam: map[string]string{
				"versioning": "",
			},
		}, &condition)
	}
}

package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

type VolcengineTlsProjectService struct {
	Client *ve.SdkClient
}

func NewTlsProjectService(c *ve.SdkClient) *VolcengineTlsProjectService {
	return &VolcengineTlsProjectService{
		Client: c,
	}
}

func (s *VolcengineTlsProjectService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTlsProjectService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	if tags, exist := condition["Tags"]; exist {
		tagsArr, ok := tags.([]interface{})
		if !ok {
			return data, fmt.Errorf(" Tags in condition is not slice ")
		}
		tagsBytes, err := json.Marshal(tagsArr)
		if err != nil {
			return data, fmt.Errorf(" json marshal tags error: %v", err)
		}
		condition["Tags"] = string(tagsBytes)
	}

	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeProjects"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.Projects", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("RESPONSE.Projects is not Slice")
		}

		return data, err
	})
}

func (s *VolcengineTlsProjectService) ReadResource(resourceData *schema.ResourceData, projectId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if projectId == "" {
		projectId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ProjectId": projectId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		var projectMap map[string]interface{}
		if projectMap, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if projectMap["ProjectId"].(string) == projectId {
			data = projectMap
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tls project %s is not exist ", projectId)
	}

	return data, err
}

func (s *VolcengineTlsProjectService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineTlsProjectService) WithResourceResponseHandlers(project map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return project, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineTlsProjectService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateProject",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"project_name": {
					TargetField: "ProjectName",
				},
				"description": {
					TargetField: "Description",
				},
				"iam_project_name": {
					TargetField: "IamProjectName",
				},
				"region": {
					TargetField: "Region",
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"key":   {TargetField: "Key"},
						"value": {TargetField: "Value"},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.ProjectId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTlsProjectService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyProject",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"project_name": {
					TargetField: "ProjectName",
				},
				"description": {
					TargetField: "Description",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["ProjectId"] = d.Id()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 更新Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcengineTlsProjectService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteProject",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ProjectId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				// 当前 project 包含 topic 时，不能删除，直接返回接口报错
				if tlsErr, ok := baseErr.(volcengineerr.Error); ok {
					if tlsErr.Code() == "ProjectDeletedHasTopic" {
						return tlsErr
					}
				}

				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading tls project on delete %q, %w", d.Id(), callErr))
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
	callbacks = append(callbacks, removeCallback)

	return callbacks
}

func (s *VolcengineTlsProjectService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"tags": {
				TargetField: "Tags",
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		NameField:    "ProjectName",
		IdField:      "ProjectId",
		CollectField: "tls_projects",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"ProjectId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineTlsProjectService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineTlsProjectService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveTagsFromResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceType"] = "project"
					(*call.SdkParam)["ResourcesList"] = []string{resourceData.Id()}
					(*call.SdkParam)["TagKeyList"] = make([]string, 0)
					for _, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeyList"] = append((*call.SdkParam)["TagKeyList"].([]string), tag.(map[string]interface{})["key"].(string))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddTagsToResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceType"] = "project"
					(*call.SdkParam)["ResourcesList"] = []string{resourceData.Id()}
					(*call.SdkParam)["Tags"] = make([]map[string]interface{}, 0)
					for _, tag := range addedTags.List() {
						(*call.SdkParam)["Tags"] = append((*call.SdkParam)["Tags"].([]map[string]interface{}), map[string]interface{}{
							"Key":   tag.(map[string]interface{})["key"],
							"Value": tag.(map[string]interface{})["value"],
						})
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func (s *VolcengineTlsProjectService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "tls",
		ResourceType:         "project",
		ProjectSchemaField:   "iam_project_name",
		ProjectResponseField: "IamProjectName",
	}
}

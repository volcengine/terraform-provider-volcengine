package apig_upstream_version

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineApigUpstreamVersionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewApigUpstreamVersionService(c *ve.SdkClient) *VolcengineApigUpstreamVersionService {
	return &VolcengineApigUpstreamVersionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineApigUpstreamVersionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineApigUpstreamVersionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListUpstreams"

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

		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		upstreams, ok := results.([]interface{})
		if !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		upstream := make(map[string]interface{})
		for _, v := range upstreams {
			if upstream, ok = v.(map[string]interface{}); !ok {
				return data, errors.New("Value is not map ")
			}
		}
		if len(upstream) == 0 {
			return data, err
		}

		if v, exist := upstream["VersionDetails"]; exist {
			if versions, ok := v.([]interface{}); ok {
				data = versions
			} else {
				return data, errors.New("VersionDetails is not Slice")
			}
		}

		return data, err
	})
}

func (s *VolcengineApigUpstreamVersionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid apig upstream version id: %s", id)
	}

	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []string{ids[0]},
		},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	version := make(map[string]interface{})
	for _, v := range results {
		if version, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if version["Name"] == ids[1] {
			data = make(map[string]interface{})
			data["UpstreamVersion"] = []interface{}{version}
			data["UpstreamId"] = ids[0]
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("apig_upstream_version %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineApigUpstreamVersionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineApigUpstreamVersionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineApigUpstreamVersionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateUpstreamVersion",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"upstream_version": {
					TargetField: "UpstreamVersion",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"labels": {
							TargetField: "Labels",
							ConvertType: ve.ConvertJsonObjectArray,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				upstreamId := d.Get("upstream_id").(string)
				versionName := d.Get("upstream_version.0.name").(string)
				d.SetId(fmt.Sprintf("%s:%s", upstreamId, versionName))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigUpstreamVersionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateUpstreamVersion",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"upstream_id": {
					TargetField: "UpstreamId",
					ForceGet:    true,
				},
				"upstream_version": {
					TargetField: "UpstreamVersion",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"name": {
							TargetField: "Name",
							ForceGet:    true,
						},
						"labels": {
							TargetField: "Labels",
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigUpstreamVersionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteUpstreamVersion",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid apig upstream version id: %s", d.Id())
				}
				(*call.SdkParam)["UpstreamId"] = ids[0]
				(*call.SdkParam)["UpstreamVersionName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading apig upstream version on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineApigUpstreamVersionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"upstream_id": {
				TargetField: "Filter.Ids",
				Convert: func(data *schema.ResourceData, i interface{}) interface{} {
					if id, ok := i.(string); ok {
						return []string{id}
					}
					return []string{}
				},
			},
		},
		NameField:    "Name",
		CollectField: "versions",
		ContentType:  ve.ContentTypeJson,
	}
}

func (s *VolcengineApigUpstreamVersionService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "apig",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

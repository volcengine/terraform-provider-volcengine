package tag

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsTagService struct {
	Client *ve.SdkClient
}

func NewTlsTagService(c *ve.SdkClient) *VolcengineTlsTagService {
	return &VolcengineTlsTagService{
		Client: c,
	}
}

func (v *VolcengineTlsTagService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsTagService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp     *map[string]interface{}
		results  interface{}
		tagsList []interface{}
		ok       bool
	)

	// Create a copy of the map to avoid modifying the original
	apiRequest := make(map[string]interface{})
	for k, v := range m {
		// Only copy valid parameters
		if k == "MaxResults" {
			// Skip MaxResults if it's 0, regardless of type
			var isZero bool
			switch val := v.(type) {
			case int:
				isZero = val == 0
			case int8:
				isZero = val == 0
			case int16:
				isZero = val == 0
			case int32:
				isZero = val == 0
			case int64:
				isZero = val == 0
			case uint:
				isZero = val == 0
			case uint8:
				isZero = val == 0
			case uint16:
				isZero = val == 0
			case uint32:
				isZero = val == 0
			case uint64:
				isZero = val == 0
			case float32:
				isZero = val == 0.0
			case float64:
				isZero = val == 0.0
			case string:
				isZero = val == "0" || val == "0.0"
			default:
				// For unknown types, skip MaxResults parameter
				isZero = true
			}
			if !isZero {
				apiRequest[k] = v
			}
		} else if k == "NextToken" {
			// Skip NextToken if it's empty
			if nextToken, ok := v.(string); ok {
				if nextToken != "" {
					apiRequest[k] = v
				}
			} else {
				// Only accept string type for NextToken
				continue
			}
		} else if k == "TagFilters" {
			// Convert tag_filters from {key, values} to {Key, Values} as required by API
			rawFilters := v.([]interface{})
			apiFilters := make([]interface{}, len(rawFilters))
			for i, rawFilter := range rawFilters {
				filterMap := rawFilter.(map[string]interface{})
				apiFilters[i] = map[string]interface{}{
					"Key":    filterMap["key"],
					"Values": filterMap["values"],
				}
			}
			apiRequest[k] = apiFilters
		} else {
			// Copy all other parameters
			apiRequest[k] = v
		}
	}

	action := "ListTagsForResources"
	logger.Debug(logger.ReqFormat, action, apiRequest)
	resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.ApplicationJSON,
		HttpMethod:  ve.POST,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, &apiRequest)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	results, err = ve.ObtainSdkValue("RESPONSE.ResourceTags", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}

	if tagsList, ok = results.([]interface{}); !ok {
		return data, errors.New("RESPONSE.ResourceTags is not Slice")
	}

	return tagsList, nil
}

func (v *VolcengineTlsTagService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}

	resourceType, exist := resourceData.GetOkExists("resource_type")
	if !exist {
		return data, fmt.Errorf("resource_type is required")
	}

	// Only pass necessary parameters for single resource query
	// Don't pass pagination parameters that might have invalid values
	req := map[string]interface{}{
		"ResourceType": resourceType,
		"ResourcesIds": []string{id},
		"MaxResults":   50,
	}

	if _, exists := req["NextToken"]; exists {
		delete(req, "NextToken")
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}

	var ResourcesList []string
	var TagKeyList []string
	var ResourceType string
	var tags []interface{}

	for _, v := range results {
		if tagData, ok := v.(map[string]interface{}); ok {
			if resourceId, ok := tagData["ResourceId"].(interface{}); ok {
				if id, ok := resourceId.(string); ok {
					ResourcesList = append(ResourcesList, id)
				}
			}
			if tagKey, ok := tagData["TagKey"].(interface{}); ok {
				if key, ok := tagKey.(string); ok {
					TagKeyList = append(TagKeyList, key)
				}
			}
			if rt, ok := tagData["ResourceType"].(string); ok {
				ResourceType = rt
			}

			t := map[string]interface{}{}
			if k, ok := tagData["TagKey"].(string); ok {
				t["key"] = k
			}
			if val, ok := tagData["TagValue"].(string); ok {
				t["value"] = val
			}
			tags = append(tags, t)
		}
	}

	// ReadResource expects a map, so we wrap the list in a map under "ResourceTags"
	// and also provide top-level fields from the first tag to satisfy schema
	data = make(map[string]interface{})

	data["ResourceType"] = ResourceType
	data["ResourcesList"] = ResourcesList
	data["TagKeyList"] = TagKeyList
	data["tags"] = tags

	logger.Debug(logger.ReqFormat, "data", data)

	return data, err
}

func (v *VolcengineTlsTagService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsTagService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		// Convert TagKey/TagValue to key/value for Terraform schema
		responseConverts := map[string]ve.ResponseConvert{
			"TagKey": {
				TargetField: "key",
			},
			"TagValue": {
				TargetField: "value",
			},
			"ResourceId": {
				TargetField: "resource_id",
			},
			"ResourceType": {
				TargetField: "resource_type",
			},
		}
		return m, responseConverts, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsTagService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	tagCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// Convert single resource_id to ResourcesIds array as required by API
				resourceId := d.Get("resource_id").(string)
				(*call.SdkParam)["ResourceType"] = d.Get("resource_type")
				(*call.SdkParam)["ResourcesIds"] = []string{resourceId}

				// Convert tags from {key, value} to {Key, Value} as required by API
				rawTags := d.Get("tags").([]interface{})
				apiTags := make([]interface{}, len(rawTags))
				for i, rawTag := range rawTags {
					tagMap := rawTag.(map[string]interface{})
					apiTags[i] = map[string]interface{}{
						"Key":   tagMap["key"],
						"Value": tagMap["value"],
					}
				}
				(*call.SdkParam)["Tags"] = apiTags

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// TagResources doesn't return an ID, use resource ID as tag resource ID
				resourceId := d.Get("resource_id").(string)
				d.SetId(resourceId)
				return nil
			},
		},
	}
	callbacks = append(callbacks, tagCallback)
	return callbacks
}

func (v *VolcengineTlsTagService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// Get all existing tags and prepare to remove them
				existingTags, err := v.ReadResource(d, "")
				if err != nil {
					return false, err
				}

				// 初始化默认值，避免空值问题
				resourceType := ""
				resourcesList := []string{}
				tagKeyList := []string{}

				// 1. 提取 ResourceType（string 类型）
				if rt, ok := existingTags["ResourceType"].(string); ok {
					resourceType = rt
				}

				if rl, ok := existingTags["ResourcesList"].([]string); ok {
					logger.Debug(logger.ReqFormat, "ResourcesList断言为[]interface{}成功，长度：%d", len(rl))
					rlUnique := make(map[string]struct{})
					for _, item := range rl {
						// 空值不加入（可选，根据业务需求调整）
						if item != "" {
							rlUnique[item] = struct{}{}
						}
					}
					// 将去重后的key转回切片
					for id := range rlUnique {
						resourcesList = append(resourcesList, id)
					}
				}
				logger.Debug(logger.ReqFormat, "resourcesList", resourcesList)
				// 3. 提取 TagKeyList（[]interface{} 转 []string）
				if tkl, ok := existingTags["TagKeyList"].([]string); ok {
					tklUnique := make(map[string]struct{})
					for _, item := range tkl {
						if item != "" {
							tklUnique[item] = struct{}{}
						}
					}
					// 将去重后的key转回切片
					for key := range tklUnique {
						tagKeyList = append(tagKeyList, key)
					}
				}
				logger.Debug(logger.ReqFormat, "tagKeyList", tagKeyList)
				// 4. 只有当 TagKeyList 非空时，才设置 SdkParam 并返回 true
				if len(tagKeyList) > 0 {
					// 设置 RemoveTagsFromResource API 的参数
					(*call.SdkParam)["ResourceType"] = resourceType
					// 优先使用 existingTags 中的 ResourcesList，若为空则使用 d.Get 的值
					if len(resourcesList) == 0 {
						if id, ok := d.Get("resource_id").(string); ok {
							resourcesList = []string{id}
						}
					}
					(*call.SdkParam)["ResourcesIds"] = resourcesList
					(*call.SdkParam)["TagKeys"] = tagKeyList

					logger.Debug(logger.ReqFormat, call.Action, "Set SdkParam success", (*call.SdkParam))
					return true, nil
				}
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// Verify tags are removed
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					tagsMap, err := v.ReadResource(d, "")
					if err != nil {
						// If resource not found (all tags removed and ReadResource returns error), it's success
						if fmt.Sprintf("%v", err) == fmt.Sprintf("tls tag %s not exist ", d.Get("resource_id").(string)) {
							return nil
						}
						return resource.NonRetryableError(err)
					}

					// Check if any of the removed tags still exist
					if tags, ok := tagsMap["ResourceTags"].([]map[string]interface{}); ok {
						removedTagKeys := (*call.SdkParam)["TagKeys"].([]string)
						for _, tag := range tags {
							if key, ok := tag["Key"].(string); ok {
								for _, removedKey := range removedTagKeys {
									if key == removedKey {
										return resource.RetryableError(fmt.Errorf("tag key %s still exists", key))
									}
								}
							}
						}
					}

					return nil
				})
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	return callbacks
}

func (v *VolcengineTlsTagService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineTlsTagService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"resource_type": {
				TargetField: "ResourceType",
			},
			"resource_ids": {
				TargetField: "ResourcesIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"max_results": {
				TargetField: "MaxResults",
			},
			"next_token": {
				TargetField: "NextToken",
			},
			"tag_filters": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		CollectField: "tags",
		IdField:      "ResourceId",
		NameField:    "ResourceId",
	}
}

func (v *VolcengineTlsTagService) ReadResourceId(id string) string {
	return id
}

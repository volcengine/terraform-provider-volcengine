package tag

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
			if rawFilters, ok := v.([]interface{}); ok {
				apiFilters := make([]interface{}, 0, len(rawFilters))
				for _, rawFilter := range rawFilters {
					if filterMap, ok := rawFilter.(map[string]interface{}); ok {
						apiFilters = append(apiFilters, map[string]interface{}{
							"Key":    filterMap["key"],
							"Values": filterMap["values"],
						})
					}
				}
				apiRequest[k] = apiFilters
			}
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
	// Get tag keys defined in the HCL configuration to filter out tags added via other means
	hclTagKeys := make(map[string]bool)
	if v, ok := resourceData.Get("tags").([]interface{}); ok {
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				if key, ok := m["key"].(string); ok {
					hclTagKeys[key] = true
				}
			}
		}
	}

	for _, v := range results {
		if tagData, ok := v.(map[string]interface{}); ok {
			tagKey, _ := tagData["TagKey"].(string)

			// Only include tags that are defined in the HCL configuration
			// If no tags are defined in HCL (e.g. during import), include all tags
			if len(hclTagKeys) > 0 {
				if _, ok := hclTagKeys[tagKey]; !ok {
					continue
				}
			}

			if resourceId, ok := tagData["ResourceId"].(interface{}); ok {
				if id, ok := resourceId.(string); ok {
					ResourcesList = append(ResourcesList, id)
				}
			}
			TagKeyList = append(TagKeyList, tagKey)
			if rt, ok := tagData["ResourceType"].(string); ok {
				ResourceType = rt
			}

			t := map[string]interface{}{
				"key":   tagKey,
				"value": tagData["TagValue"],
			}
			tags = append(tags, t)
		}
	}

	// ReadResource expects a map
	data = make(map[string]interface{})

	// Always preserve the resource ID and type
	data["resource_id"] = id
	data["resource_type"] = resourceType

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
				resourceId, ok := d.Get("resource_id").(string)
				if !ok {
					return false, fmt.Errorf("resource_id is not string")
				}
				(*call.SdkParam)["ResourceType"] = d.Get("resource_type")
				(*call.SdkParam)["ResourcesIds"] = []string{resourceId}

				// Convert tags from {key, value} to {Key, Value} as required by API
				vTags := d.Get("tags")
				rawTags, ok := vTags.([]interface{})
				if !ok {
					return false, fmt.Errorf("tags is not []interface{}")
				}
				apiTags := make([]interface{}, len(rawTags))
				for i, rawTag := range rawTags {
					tagMap, ok := rawTag.(map[string]interface{})
					if !ok {
						return false, fmt.Errorf("tag item is not map[string]interface{}")
					}
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
				resourceId, ok := d.Get("resource_id").(string)
				if !ok {
					return fmt.Errorf("resource_id is not string")
				}
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
				// Only remove tags defined in the resource configuration
				vTags := d.Get("tags")
				rawTags, ok := vTags.([]interface{})
				if !ok {
					return false, fmt.Errorf("tags is not []interface{}")
				}

				if len(rawTags) == 0 {
					return false, nil // No tags to remove
				}

				tagKeyList := make([]string, 0, len(rawTags))
				for _, rawTag := range rawTags {
					tagMap, ok := rawTag.(map[string]interface{})
					if !ok {
						return false, fmt.Errorf("tag item is not map[string]interface{}")
					}
					if key, ok := tagMap["key"].(string); ok && key != "" {
						tagKeyList = append(tagKeyList, key)
					}
				}

				if len(tagKeyList) > 0 {
					resourceId, ok := d.Get("resource_id").(string)
					if !ok {
						return false, fmt.Errorf("resource_id is not string")
					}
					resourceType := d.Get("resource_type")

					(*call.SdkParam)["ResourceType"] = resourceType
					(*call.SdkParam)["ResourcesIds"] = []string{resourceId}
					(*call.SdkParam)["TagKeys"] = tagKeyList

					logger.Debug(logger.ReqFormat, call.Action, "Set SdkParam success", (*call.SdkParam))
					return true, nil
				}
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
						vResId := d.Get("resource_id")
						resId, ok := vResId.(string)
						if ok && strings.Contains(fmt.Sprintf("%v", err), fmt.Sprintf("tls tag %s not exist", resId)) {
							return nil
						}
						return resource.NonRetryableError(err)
					}

					// Check if any of the removed tags still exist
					if tags, ok := tagsMap["tags"].([]interface{}); ok {
						if removedTagKeysRaw, ok := (*call.SdkParam)["TagKeys"]; ok {
							if removedTagKeys, ok := removedTagKeysRaw.([]string); ok {
								for _, tag := range tags {
									if tagMap, ok := tag.(map[string]interface{}); ok {
										if key, ok := tagMap["key"].(string); ok {
											for _, removedKey := range removedTagKeys {
												if key == removedKey {
													return resource.RetryableError(fmt.Errorf("tag key %s still exists", key))
												}
											}
										}
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
			},
			"max_results": {
				TargetField: "MaxResults",
			},
			"next_token": {
				TargetField: "NextToken",
			},
			"tag_filters": {
				TargetField: "TagFilters",
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
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

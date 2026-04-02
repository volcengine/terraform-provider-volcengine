package nlb_tag

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

type VolcengineNlbTagService struct {
	Client *ve.SdkClient
}

func NewNlbTagService(c *ve.SdkClient) *VolcengineNlbTagService {
	return &VolcengineNlbTagService{
		Client: c,
	}
}

func (s *VolcengineNlbTagService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbTagService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 100, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
		action := "ListTagsForNLBResources"

		// Flatten any slices in the condition map to .N format
		flattenedCondition := make(map[string]interface{})
		for k, v := range condition {
			if slice, ok := v.([]interface{}); ok {
				if len(slice) == 0 {
					continue
				}
				for i, item := range slice {
					flattenedCondition[fmt.Sprintf("%s.%d", k, i+1)] = item
				}
			} else if slice, ok := v.([]string); ok {
				if len(slice) == 0 {
					continue
				}
				for i, item := range slice {
					flattenedCondition[fmt.Sprintf("%s.%d", k, i+1)] = item
				}
			} else {
				flattenedCondition[k] = v
			}
		}

		// ListTagsForNLBResources specifically needs ResourceIds.N
		if v, ok := condition["ResourceIds"]; ok {
			if slice, ok := v.([]interface{}); ok {
				for i, item := range slice {
					flattenedCondition[fmt.Sprintf("ResourceIds.%d", i+1)] = item
				}
				delete(flattenedCondition, "ResourceIds")
			} else if slice, ok := v.([]string); ok {
				for i, item := range slice {
					flattenedCondition[fmt.Sprintf("ResourceIds.%d", i+1)] = item
				}
				delete(flattenedCondition, "ResourceIds")
			}
		}

		bytes, _ := json.Marshal(flattenedCondition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &flattenedCondition)
		if err != nil {
			return nil, "", err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err := ve.ObtainSdkValue("Result.ResourceTags", *resp)
		if err != nil {
			return nil, "", err
		}
		if results == nil {
			results = []interface{}{}
		}

		var nextTokenStr string
		nextToken, _ := ve.ObtainSdkValue("Result.NextToken", *resp)
		if nextToken != nil {
			if s, ok := nextToken.(string); ok {
				nextTokenStr = s
			}
		}

		if data, ok := results.([]interface{}); ok {
			return data, nextTokenStr, nil
		}
		return nil, "", errors.New("Result.ResourceTags is not Slice")
	})
}

func (s *VolcengineNlbTagService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	// Standalone tag resource is a bit special. We check if all specified tags exist on the resource.
	resourceId, ok := resourceData.Get("resource_id").(string)
	if !ok {
		return nil, errors.New("resource_id is not string")
	}
	resourceType, ok := resourceData.Get("resource_type").(string)
	if !ok {
		return nil, errors.New("resource_type is not string")
	}

	req := map[string]interface{}{
		"ResourceType":  resourceType,
		"ResourceIds.1": resourceId,
	}
	results, err := s.ReadResources(req)
	if err != nil {
		return nil, err
	}

	tags := make([]interface{}, 0)
	managedKeys := make(map[string]bool)
	if v := resourceData.Get("tags"); v != nil {
		if set, ok := v.(*schema.Set); ok {
			for _, t := range set.List() {
				if m, ok := t.(map[string]interface{}); ok {
					if key, ok := m["key"].(string); ok {
						managedKeys[key] = true
					}
				}
			}
		}
	}

	for _, v := range results {
		tagData, ok := v.(map[string]interface{})
		if !ok {
			return nil, errors.New("result item is not map")
		}
		resId, ok := tagData["ResourceId"].(string)
		if !ok {
			return nil, errors.New("result item ResourceId is not string")
		}
		if resId == resourceId {
			tagKey, ok := tagData["TagKey"].(string)
			if !ok {
				return nil, errors.New("result item TagKey is not string")
			}
			if strings.HasPrefix(tagKey, "sys:") {
				continue
			}
			if !managedKeys[tagKey] {
				continue
			}
			tags = append(tags, map[string]interface{}{
				"key":   tagKey,
				"value": tagData["TagValue"],
			})
		}
	}

	data = make(map[string]interface{})
	data["resource_id"] = resourceId
	data["resource_type"] = resourceType
	data["tags"] = tags
	return data, nil
}

func (s *VolcengineNlbTagService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineNlbTagService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "TagNLBResources",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				resourceId, ok := d.Get("resource_id").(string)
				if !ok {
					return nil, errors.New("resource_id is not string")
				}
				resourceType, ok := d.Get("resource_type").(string)
				if !ok {
					return nil, errors.New("resource_type is not string")
				}
				v := d.Get("tags")
				tagsSet, ok := v.(*schema.Set)
				if !ok {
					return nil, errors.New("tags is not *schema.Set")
				}
				tags := tagsSet.List()

				param := map[string]interface{}{
					"ResourceType":  resourceType,
					"ResourceIds.1": resourceId,
				}

				for index, t := range tags {
					tagIndex := index + 1
					tag, ok := t.(map[string]interface{})
					if !ok {
						return nil, errors.New("tag item is not map")
					}
					param[fmt.Sprintf("Tags.%d.Key", tagIndex)] = tag["key"]
					param[fmt.Sprintf("Tags.%d.Value", tagIndex)] = tag["value"]
				}

				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				resourceId, ok := d.Get("resource_id").(string)
				if !ok {
					return errors.New("resource_id is not string")
				}
				resourceType, ok := d.Get("resource_type").(string)
				if !ok {
					return errors.New("resource_type is not string")
				}
				d.SetId(fmt.Sprintf("%s:%s", resourceType, resourceId))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNlbTagService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	if removedTags != nil && len(removedTags.List()) > 0 {
		removeCallback := ve.Callback{
			Call: ve.SdkCall{
				Action: "UntagNLBResources",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					resourceId, ok := d.Get("resource_id").(string)
					if !ok {
						return nil, errors.New("resource_id is not string")
					}
					resourceType, ok := d.Get("resource_type").(string)
					if !ok {
						return nil, errors.New("resource_type is not string")
					}
					param := map[string]interface{}{
						"ResourceType":  resourceType,
						"ResourceIds.1": resourceId,
					}
					for index, t := range removedTags.List() {
						tag, ok := t.(map[string]interface{})
						if !ok {
							return nil, errors.New("removed tag item is not map")
						}
						key, ok := tag["key"].(string)
						if !ok {
							return nil, errors.New("tag key is not string")
						}
						param[fmt.Sprintf("TagKeys.%d", index+1)] = key
					}
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		}
		callbacks = append(callbacks, removeCallback)
	}

	if addedTags != nil && len(addedTags.List()) > 0 {
		addCallback := ve.Callback{
			Call: ve.SdkCall{
				Action: "TagNLBResources",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					resourceId, ok := d.Get("resource_id").(string)
					if !ok {
						return nil, errors.New("resource_id is not string")
					}
					resourceType, ok := d.Get("resource_type").(string)
					if !ok {
						return nil, errors.New("resource_type is not string")
					}
					param := map[string]interface{}{
						"ResourceType":  resourceType,
						"ResourceIds.1": resourceId,
					}
					for index, t := range addedTags.List() {
						tag, ok := t.(map[string]interface{})
						if !ok {
							return nil, errors.New("added tag item is not map")
						}
						param[fmt.Sprintf("Tags.%d.Key", index+1)] = tag["key"]
						param[fmt.Sprintf("Tags.%d.Value", index+1)] = tag["value"]
					}
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		}
		callbacks = append(callbacks, addCallback)
	}

	return callbacks
}

func (s *VolcengineNlbTagService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "UntagNLBResources",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				resourceId, ok := d.Get("resource_id").(string)
				if !ok {
					return nil, errors.New("resource_id is not string")
				}
				resourceType, ok := d.Get("resource_type").(string)
				if !ok {
					return nil, errors.New("resource_type is not string")
				}
				v := d.Get("tags")
				tagsSet, ok := v.(*schema.Set)
				if !ok {
					return nil, errors.New("tags is not *schema.Set")
				}
				tags := tagsSet.List()

				param := map[string]interface{}{
					"ResourceType":  resourceType,
					"ResourceIds.1": resourceId,
				}

				for index, t := range tags {
					tagIndex := index + 1
					tag, ok := t.(map[string]interface{})
					if !ok {
						return nil, errors.New("tag item is not map")
					}
					key, ok := tag["key"].(string)
					if !ok {
						return nil, errors.New("tag key is not string")
					}
					param[fmt.Sprintf("TagKeys.%d", tagIndex)] = key
				}

				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineNlbTagService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"resource_id": {
				TargetField: "resource_id",
			},
			"resource_type": {
				TargetField: "resource_type",
			},
			"tags": {
				TargetField: "tags",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNlbTagService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"resource_type": {
				TargetField: "ResourceType",
			},
			"resource_ids": {
				TargetField: "ResourceIds",
				ConvertType: ve.ConvertWithN,
			},
			"tag_type": {
				TargetField: "TagType",
			},
			"tag_filters": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"key": {
						TargetField: "Key",
					},
					"values": {
						TargetField: "Values",
						ConvertType: ve.ConvertWithN,
					},
				},
			},
		},
		CollectField: "tags",
		IdField:      "ResourceId",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ResourceType": {
				TargetField: "resource_type",
			},
			"ResourceId": {
				TargetField: "resource_id",
			},
			"TagKey": {
				TargetField: "key",
			},
			"TagValue": {
				TargetField: "value",
			},
		},
	}
}

func (s *VolcengineNlbTagService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

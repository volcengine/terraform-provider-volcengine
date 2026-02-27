package common

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/volcengine/terraform-provider-volcengine/logger"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TagsSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Tags.",
		Set:         TagsHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The Key of Tags.",
				},
				"value": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The Value of Tags.",
				},
			},
		},
	}
}

func TagsSchemaComputed() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Computed:    true,
		Description: "Tags.",
		Set:         TagsHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The Key of Tags.",
				},
				"value": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The Value of Tags.",
				},
			},
		},
	}
}

var TagsHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["key"], m["value"]))
	return hashcode.String(buf.String())
}

var VkeTagsResponseHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["key"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["value"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["type"].(string))))
	return hashcode.String(buf.String())
}

type GetUniversalInfo func(actionName string) UniversalInfo

func SetResourceTags(serviceClient *SdkClient, addAction, RemoveAction, resourceType string,
	resourceData *schema.ResourceData, getUniversalInfo GetUniversalInfo) []Callback {
	var callbacks []Callback
	addedTags, removedTags, _, _ := GetSetDifference("tags", resourceData, TagsHash, false)

	removeCallback := Callback{
		Call: SdkCall{
			Action:      RemoveAction,
			ConvertMode: RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds.1"] = resourceData.Id()
					(*call.SdkParam)["ResourceType"] = resourceType
					for index, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeys."+strconv.Itoa(index+1)] = tag.(map[string]interface{})["key"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return serviceClient.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := Callback{
		Call: SdkCall{
			Action:      addAction,
			ConvertMode: RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds.1"] = resourceData.Id()
					(*call.SdkParam)["ResourceType"] = resourceType
					for index, tag := range addedTags.List() {
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Key"] = tag.(map[string]interface{})["key"].(string)
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Value"] = tag.(map[string]interface{})["value"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return serviceClient.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func FilterSystemTags(tags []interface{}) []interface{} {
	var res []interface{}
	if len(tags) == 0 {
		return res
	}
	for _, tag := range tags {
		t := tag.(map[string]interface{})
		var tagKey string
		var tagValue interface{}
		if v, ok := t["Key"]; ok {
			tagKey = v.(string)
			tagValue = t["Value"]
		}
		if !tagIgnored(tagKey, tagValue) {
			res = append(res, tag)
		}
	}
	return res
}

func tagIgnored(tagKey string, tagValue interface{}) bool {
	filter := []string{"^volc:", "^sys:"}
	for _, v := range filter {
		ok, _ := regexp.MatchString(v, tagKey)
		if ok {
			return true
		}
	}
	return false
}

func TransTagsToResponse(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	// Filter system tags first using the original common method which expects capitalized Key/Value
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	filtered := FilterSystemTags(list)

	var tags []interface{}
	for _, v := range filtered {
		if m, ok := v.(map[string]interface{}); ok {
			tag := make(map[string]interface{})
			if key, ok := m["Key"].(string); ok {
				tag["key"] = key
			}
			if value, ok := m["Value"].(string); ok {
				tag["value"] = value
			}
			tags = append(tags, tag)
		}
	}
	return tags
}

func TransTagFiltersToRequest(d *schema.ResourceData, i interface{}) interface{} {
	if i == nil {
		return nil
	}
	var list []interface{}
	switch v := i.(type) {
	case *schema.Set:
		if v.Len() == 0 {
			return nil
		}
		list = v.List()
	case []interface{}:
		list = v
	default:
		return nil
	}
	var filters []map[string]interface{}
	for _, v := range list {
		if m, ok := v.(map[string]interface{}); ok {
			filter := make(map[string]interface{})
			if key, ok := m["key"].(string); ok {
				filter["Key"] = key
			}
			if value, ok := m["value"].(string); ok {
				filter["Value"] = value
			}
			filters = append(filters, filter)
		}
	}
	return filters
}

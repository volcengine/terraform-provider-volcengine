package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
)

func TagsSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Tags.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func TagsSchemaComputed() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Computed:    true,
		Description: "Tags.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func TagsMapToList(in interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	if tags, ok := in.(map[string]interface{}); ok {
		for k, v := range tags {
			m := map[string]interface{}{
				"key":   k,
				"value": v.(string),
			}
			result = append(result, m)
		}
	}
	return result
}

func TagsListToMap(in interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	if tags, ok := in.([]interface{}); ok {
		for _, tag := range tags {
			result[tag.(map[string]interface{})["Key"].(string)] = tag.(map[string]interface{})["Value"].(string)
		}
	}
	return result
}

func GetTagsDifference(key string, d *schema.ResourceData) (addedTags map[string]interface{}, removedTags map[string]interface{}) {
	if d.HasChange(key) {
		oldRaw, newRaw := d.GetChange(key)
		if oldRaw == nil {
			oldRaw = make(map[string]interface{})
		}
		if newRaw == nil {
			newRaw = make(map[string]interface{})
		}
		oldTags := oldRaw.(map[string]interface{})
		newTags := newRaw.(map[string]interface{})

		addedTags = getDifference(newTags, oldTags)
		removedTags = getDifference(oldTags, newTags)
		return addedTags, removedTags
	}
	return addedTags, removedTags
}

func getDifference(m, other map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		if otherValue, ok := other[k]; !ok || !reflect.DeepEqual(v, otherValue) {
			result[k] = v
		}
	}
	return result
}

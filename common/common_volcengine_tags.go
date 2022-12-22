package common

import (
	"bytes"
	"fmt"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"strconv"
	"strings"

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

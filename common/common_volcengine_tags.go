package common

import (
	"bytes"
	"fmt"
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

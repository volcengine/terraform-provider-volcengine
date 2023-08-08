package default_node_pool

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func defaultNodePoolNodeHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["instance_id"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(strconv.FormatBool(m["additional_container_storage_enabled"].(bool)))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["image_id"].(string))))
	if m["additional_container_storage_enabled"].(bool) {
		buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["container_storage_path"].(string))))
	}
	return hashcode.String(buf.String())
}

func defaultNodePoolDiffSuppress() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		key := strings.ReplaceAll(k, "container_storage_path", "") + "additional_container_storage_enabled"
		return d.Get(key) == false
		//if d.Get(key) == false {
		//	return true
		//}
		//return false
	}
}

func defaultNodePoolKeepNameDiffSuppress() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		key := strings.ReplaceAll(k, "keep_instance_name", "") + "id"
		if d.Get(key) == nil || len(d.Get(key).(string)) == 0 {
			return false
		}
		return true
	}
}

var kubernetesConfigLabelHash = func(v interface{}) int {
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

var defaultNodePoolImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	_ = data.Set("is_import", true)
	return []*schema.ResourceData{data}, nil
}

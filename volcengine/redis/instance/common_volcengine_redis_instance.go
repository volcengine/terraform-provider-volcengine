package instance

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var paramHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["name"], m["value"]))
	return hashcode.String(buf.String())
}

var tagsHash = func(v interface{}) int {
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

func redisInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// 不启用分片集群时，忽略 shard_number 的修改
	if k == "shard_number" {
		if d.Get("sharded_cluster").(int) == 1 {
			return false
		}
		return true
	}

	// 计费方式为 PostPaid 时，忽略相关参数的修改
	if (k == "purchase_months" || k == "auto_renew") && d.Get("charge_type").(string) == "PostPaid" {
		return true
	}

	// 只在修改时需要这些参数
	if d.Id() == "" {
		if k == "vpc_auth_mode" || k == "backup_hour" || k == "backup_active" || strings.Contains(k, "param_values") || strings.Contains(k, "backup_period") {
			return true
		}
	}

	// 只在修改实例规格时需要这些参数，其它情况均忽略修改
	if k == "create_backup" || k == "apply_immediately" {
		if d.HasChanges("node_number", "shard_number", "shard_capacity") {
			oldNum, _ := d.GetChange("shard_number")
			if oldNum.(int) == 1 {
				return true
			}
		}
		return true
	}

	return false
}

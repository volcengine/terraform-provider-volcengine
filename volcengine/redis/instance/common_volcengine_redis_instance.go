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

var configNodesHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["az"], m["az"]))
	return hashcode.String(buf.String())
}

func redisInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// 不启用分片集群时，忽略 shard_number 的修改
	if k == "shard_number" {
		return d.Get("sharded_cluster").(int) != 1
	}

	// 计费方式为 PostPaid 时，忽略相关参数的修改
	if (k == "purchase_months" || k == "auto_renew") && d.Get("charge_type").(string) == "PostPaid" {
		return true
	}

	// 单节点实例，忽略 backup plan 相关参数
	if k == "backup_hour" || k == "backup_active" || strings.Contains(k, "backup_period") {
		return d.Get("node_number").(int) == 1
	}

	// 只在修改实例规格时需要这些参数，其它情况均忽略修改
	if k == "create_backup" || k == "apply_immediately" {
		if d.HasChanges("configure_nodes", "multi_az", "node_number") {
			return false
		}
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

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func compareMaps(oldArr, newArr []interface{}) (added, removed []map[string]interface{}) {
	oldCount := make(map[string]int)
	newCount := make(map[string]int)

	// 统计 oldArr 中每个 "az" 值的出现次数
	for _, i := range oldArr {
		item := i.(map[string]interface{})
		if azValue, ok := item["az"].(string); ok {
			oldCount[azValue]++
		}
	}

	// 统计 newArr 中每个 "az" 值的出现次数
	for _, i := range newArr {
		item := i.(map[string]interface{})
		if azValue, ok := item["az"].(string); ok {
			newCount[azValue]++
		}
	}

	// 查找新增的元素
	for azValue, newCountValue := range newCount {
		if oldCountValue, exists := oldCount[azValue]; !exists || newCountValue > oldCountValue {
			// 如果新的计数超过旧的计数，表示新增
			for i := 0; i < newCountValue-oldCountValue; i++ {
				added = append(added, map[string]interface{}{"az": azValue})
			}
		}
	}

	// 查找移除的元素
	for azValue, oldCountValue := range oldCount {
		if newCountValue, exists := newCount[azValue]; !exists || oldCountValue > newCountValue {
			// 如果旧的计数超过新的计数，表示移除
			for i := 0; i < oldCountValue-newCountValue; i++ {
				removed = append(removed, map[string]interface{}{"az": azValue})
			}
		}
	}

	return added, removed
}

package instance

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var nodeSpecsAssignsDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	return !diffNodeSpecsAssigns(d)
}

var forceRestartAfterScaleDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	// 创建时不存在这个参数，修改 node_specs_assigns 时存在这个参数
	return !(d.Id() != "" && diffNodeSpecsAssigns(d))
}

func diffNodeSpecsAssigns(d *schema.ResourceData) bool {
	oldVal, newVal := d.GetChange("instance_configuration.0.node_specs_assigns")
	oldNodeConfigs := transListToMap(oldVal.([]interface{}))
	newNodeConfigs := transListToMap(newVal.([]interface{}))

	if !reflect.DeepEqual(oldNodeConfigs["Master"], newNodeConfigs["Master"]) {
		return true
	}
	if !reflect.DeepEqual(oldNodeConfigs["Hot"], newNodeConfigs["Hot"]) {
		return true
	}
	if !reflect.DeepEqual(oldNodeConfigs["Kibana"], newNodeConfigs["Kibana"]) {
		return true
	}

	return false
}

func transListToMap(nodeList []interface{}) map[string]map[string]interface{} {
	nodeMap := map[string]map[string]interface{}{}
	for _, value := range nodeList {
		nodeType := value.(map[string]interface{})["type"]
		nodeMap[nodeType.(string)] = map[string]interface{}{
			"StorageSpecName":  value.(map[string]interface{})["storage_spec_name"],
			"StorageSize":      value.(map[string]interface{})["storage_size"],
			"ResourceSpecName": value.(map[string]interface{})["resource_spec_name"],
			"Number":           value.(map[string]interface{})["number"],
		}
	}
	return nodeMap
}

package node_pool

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var prePaidDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	chargeType := d.Get("node_config").([]interface{})[0].(map[string]interface{})["instance_charge_type"].(string)
	return chargeType != "PrePaid"
}

var prePaidAndAutoNewDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	nodeConfig := d.Get("node_config").([]interface{})[0].(map[string]interface{})
	chargeType := nodeConfig["instance_charge_type"].(string)
	autoRenew := nodeConfig["auto_renew"].(bool)
	return chargeType != "PrePaid" || !autoRenew
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

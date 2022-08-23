package node_pool

import (
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

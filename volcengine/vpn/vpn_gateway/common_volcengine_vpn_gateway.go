package vpn_gateway

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var billingTypeRequestConvert = func(data *schema.ResourceData, old interface{}) interface{} {
	var ty int
	switch old.(string) {
	case "PrePaid":
		ty = 1
	default:
		ty = 0
	}
	return ty
}

var billingTypeResponseConvert = func(i interface{}) interface{} {
	var ty string
	switch i.(float64) {
	case 1:
		ty = "PrePaid"
	default:
		ty = fmt.Sprintf("%v", i)
	}
	return ty
}

var renewTypeResponseConvert = func(v interface{}) interface{} {
	if v == nil {
		return ""
	}
	ty := ""
	switch v.(float64) {
	case 1:
		ty = "ManualRenew"
	case 2:
		ty = "AutoRenew"
	case 3:
		ty = "NoneRenew"
	}
	return ty
}

var periodDiffSuppress = func(k, old, new string, d *schema.ResourceData) bool {
	if len(d.Id()) != 0 {
		return d.Get("renew_type").(string) != "ManualRenew"
	}
	return false
}

package cen_bandwidth_package

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var billingTypeRequestConvert = func(data *schema.ResourceData, old interface{}) interface{} {
	ty := 0
	switch old.(string) {
	case "PrePaid":
		ty = 1
	case "PayBy95Peak":
		ty = 4
	}
	return ty
}

var billingTypeResponseConvert = func(i interface{}) interface{} {
	var ty string
	switch i.(float64) {
	case 1:
		ty = "PrePaid"
	case 4:
		ty = "PayBy95Peak"
	default:
		ty = fmt.Sprintf("%v", i)
	}
	return ty
}

var periodDiffSuppress = func(k, old, new string, d *schema.ResourceData) bool {
	// 非包年包月实例时，忽略相关参数
	if d.Get("billing_type").(string) != "PrePaid" {
		return true
	}

	return len(d.Id()) != 0
}

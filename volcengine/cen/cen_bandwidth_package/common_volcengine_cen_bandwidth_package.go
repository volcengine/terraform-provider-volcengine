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

var periodDiffSuppress = func(k, old, new string, d *schema.ResourceData) bool {
	return len(d.Id()) != 0
}

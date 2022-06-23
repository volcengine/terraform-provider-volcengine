package vpn_gateway

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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

var renewTypeRequestConvert = func(data *schema.ResourceData, old interface{}) interface{} {
	ty := 0
	switch old.(string) {
	case "ManualRenew":
		ty = 1
	case "AutoRenew":
		ty = 2
	case "NoneRenew":
		ty = 3
	}
	return ty
}

var renewTypeResponseConvert = func(v interface{}) string {
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

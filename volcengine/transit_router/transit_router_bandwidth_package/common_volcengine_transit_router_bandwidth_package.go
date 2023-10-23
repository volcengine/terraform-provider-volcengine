package transit_router_bandwidth_package

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func transitRouterBandwidthPackageDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if d.Id() == "" && (k == "renew_type" || k == "renew_period" || k == "remain_renew_times") {
		return true
	}

	// 修改时，当续费方式不是手动续费的时候 period的变化会被忽略
	if d.Id() != "" && k == "period" && d.Get("renew_type").(string) != "Manual" {
		return true
	}

	if (k == "renew_period" || k == "remain_renew_times") && d.Get("renew_type").(string) != "Auto" {
		return true
	}

	return false
}

var billingTypeRequestConvert = func(data *schema.ResourceData, old interface{}) interface{} {
	ty := 0
	switch old.(string) {
	case "PrePaid":
		ty = 1
	case "PostPaidByBandwidth":
		ty = 2
	case "PostPaidByTraffic":
		ty = 3
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
	case 2:
		ty = "PostPaidByBandwidth"
	case 3:
		ty = "PostPaidByTraffic"
	case 4:
		ty = "PayBy95Peak"
	default:
		ty = fmt.Sprintf("%v", i)
	}
	return ty
}

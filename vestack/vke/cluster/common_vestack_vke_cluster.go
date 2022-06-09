package cluster

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var billingTypeRequestConvert = func(data *schema.ResourceData, old interface{}) interface{} {
	ty := 0
	switch old.(string) {
	case "PrePaid":
		ty = 1
	case "PostPaidByBandwidth":
		ty = 2
	case "PostPaidByTraffic":
		ty = 3
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
	default:
		ty = fmt.Sprintf("%v", i)
	}
	return ty
}

func ApiServerPublicAccessConfigFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	apiServerPublicAccessEnabled := d.Get("cluster_config").([]interface{})[0].(map[string]interface{})["api_server_public_access_enabled"].(bool)
	return apiServerPublicAccessEnabled == false
}

func FlannelFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	podNetworkMode := d.Get("pods_config").([]interface{})[0].(map[string]interface{})["pod_network_mode"].(string)
	return podNetworkMode != "Flannel"
}

func VpcCniConfigFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	podNetworkMode := d.Get("pods_config").([]interface{})[0].(map[string]interface{})["pod_network_mode"].(string)
	return podNetworkMode != "VpcCniShared"
}

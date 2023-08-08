package scaling_configuration

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

//var substituteDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
//	return d.Get("active").(bool)
//}

var eipDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	return d.Get("eip_bandwidth") == nil || d.Get("eip_bandwidth").(int) == 0
}

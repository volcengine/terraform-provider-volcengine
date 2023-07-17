package listener

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func HealthCheckHTTPOnlyFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	healthCheckEnabled := d.Get("health_check").([]interface{})[0].(map[string]interface{})["enabled"].(string)
	protocol := d.Get("protocol").(string)
	return healthCheckEnabled == "off" || protocol == "TCP" || protocol == "UDP"
}

func HealthCheckUDPOnlyFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	healthCheckEnabled := d.Get("health_check").([]interface{})[0].(map[string]interface{})["enabled"].(string)
	protocol := d.Get("protocol").(string)
	return !(healthCheckEnabled == "on" && protocol == "UDP")
}

func HealthCheckFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	healthCheckEnabled := d.Get("health_check").([]interface{})[0].(map[string]interface{})["enabled"].(string)
	return healthCheckEnabled == "off"
}

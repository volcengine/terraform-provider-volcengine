package rds_instance

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func RdsInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if d.Id() != "" {
		if k == "auto_renew" {
			return true
		}
		if k == "prepaid_period" {
			return true
		}
		if k == "used_time" {
			return true
		}
		if k == "project_name" {
			// project_name is not returned by the API
			return true
		}
	}
	return false
}

package rds_instance_v2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func RdsInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if d.Id() != "" {
		if k == "db_param_group_id" {
			// db_param_group_id is not returned by the API
			return true
		}
		if k == "db_time_zone" {
			// db_time_zone is not returned by the API
			return true
		}
	}
	return false
}

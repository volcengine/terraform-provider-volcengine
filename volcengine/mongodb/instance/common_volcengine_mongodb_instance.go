package instance

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func MongoDBInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_type").(string) == "PostPaid" && (k == "period" || k == "period_unit" || k == "auto_renew") {
		return true
	}

	return false
}

package rds_mysql_instance

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var parameterHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["parameter_name"], m["parameter_value"]))
	return hashcode.String(buf.String())
}

func RdsMysqlInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_info.0.charge_type").(string) == "PostPaid" && (k == "charge_info.0.period" || k == "charge_info.0.period_unit" || k == "charge_info.0.auto_renew") {
		return true
	}

	return false
}

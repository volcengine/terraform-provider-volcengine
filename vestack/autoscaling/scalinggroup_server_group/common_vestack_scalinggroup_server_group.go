package scalinggroup_server_group

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
)

func serverGroupAttributeHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v:", m["load_balancer_id"]))
	buf.WriteString(fmt.Sprintf("%v:", m["port"]))
	buf.WriteString(fmt.Sprintf("%v:", m["server_group_id"]))
	buf.WriteString(fmt.Sprintf("%v:", m["weight"]))
	return hashcode.String(buf.String())
}

package bucket

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
)

func tosAccountAclHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%s-", m["account_id"]))
	buf.WriteString(fmt.Sprintf("%s", m["permission"]))
	return hashcode.String(buf.String())
}

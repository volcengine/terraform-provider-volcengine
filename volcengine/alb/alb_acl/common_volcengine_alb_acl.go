package alb_acl

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"strings"
)

func aclEntryHashBase(m map[string]interface{}) (buf bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["entry"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["description"].(string))))
	return buf
}

func AclEntryHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := aclEntryHashBase(m)
	return hashcode.String(buf.String())
}

package common

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
)

func clbAclEntryHashBase(m map[string]interface{}) (buf bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["entry"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["description"].(string))))
	return buf
}

func ClbAclEntryHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := clbAclEntryHashBase(m)
	return hashcode.String(buf.String())
}

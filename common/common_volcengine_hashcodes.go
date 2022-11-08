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

func vkeTagsResponseHashBase(m map[string]interface{}) (buf bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["key"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["value"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["type"].(string))))
	return buf
}

var VkeTagsResponseHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := vkeTagsResponseHashBase(m)
	return hashcode.String(buf.String())
}

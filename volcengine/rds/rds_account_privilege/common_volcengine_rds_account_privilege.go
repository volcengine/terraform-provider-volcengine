package rds_account_privilege

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var rdsAccountPrivilegeImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("instance_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("account_name", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

func RdsAccountPrivilegeStrDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// 1. get account_privilege
	keys := strings.Split(k, ".")
	accountPrivilegeKey := fmt.Sprintf("%s.%s.%s", keys[0], keys[1], "account_privilege")

	// 2. if custom, compute if account_privilege_str is changed
	deepEqual := false
	if d.Get(accountPrivilegeKey).(string) == "Custom" {
		oldPrivileges := mappingAccountPrivilegeStr(old)
		newPrivileges := mappingAccountPrivilegeStr(new)
		deepEqual = reflect.DeepEqual(oldPrivileges, newPrivileges)
	}
	return d.Get(accountPrivilegeKey) != "Custom" || deepEqual
}

var rdsAccountPrivilegeStrMapping = map[string]string{
	"create tmp table": "create temporary tables",
	"CREATE TMP TABLE": "CREATE TEMPORARY TABLES",
}

func mappingAccountPrivilegeStr(accountPrivilegeStr string) []string {
	privileges := strings.Split(accountPrivilegeStr, ",")
	mappingPrivileges := make([]string, 0)
	for _, privilege := range privileges {
		if mappedPrivilege, ok := rdsAccountPrivilegeStrMapping[privilege]; ok {
			mappingPrivileges = append(mappingPrivileges, mappedPrivilege)
		} else {
			mappingPrivileges = append(mappingPrivileges, privilege)
		}
	}
	return mappingPrivileges
}

// mappingAndSortAccountPrivilegeStr RDS account privilege string mapping
func mappingAndSortAccountPrivilegeStr(accountPrivilegeStr string) string {
	mappingPrivileges := mappingAccountPrivilegeStr(accountPrivilegeStr)
	sort.Strings(mappingPrivileges)
	return strings.Join(mappingPrivileges, ",")
}

func rdsAccountPrivilegeHashBase(m map[string]interface{}) (buf bytes.Buffer) {
	dbName := strings.ToLower(m["db_name"].(string))
	accountPrivilege := strings.ToLower(m["account_privilege"].(string))
	buf.WriteString(fmt.Sprintf("%s-", dbName))
	buf.WriteString(fmt.Sprintf("%s-", accountPrivilege))
	if accountPrivilege == "custom" {
		buf.WriteString(fmt.Sprintf("%s-", mappingAndSortAccountPrivilegeStr(strings.ToLower(m["account_privilege_str"].(string)))))
	}
	return buf
}

func RdsAccountPrivilegeHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := rdsAccountPrivilegeHashBase(m)
	return hashcode.String(buf.String())
}

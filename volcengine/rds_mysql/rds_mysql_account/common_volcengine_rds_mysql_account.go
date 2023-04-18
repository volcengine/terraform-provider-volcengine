package rds_mysql_account

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

var rdsMysqlAccountImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
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

var rdsMysqlAccountPrivilegeStrMapping = map[string]string{
	"create tmp table": "create temporary tables",
	"CREATE TMP TABLE": "CREATE TEMPORARY TABLES",
}

func mappingAccountPrivilegeStr(accountPrivilegeStr string) []string {
	privileges := strings.Split(accountPrivilegeStr, ",")
	mappingPrivileges := make([]string, 0)
	for _, privilege := range privileges {
		if mappedPrivilege, ok := rdsMysqlAccountPrivilegeStrMapping[privilege]; ok {
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

func rdsMysqlAccountPrivilegeHashBase(m map[string]interface{}) (buf bytes.Buffer) {
	dbName := strings.ToLower(m["db_name"].(string))
	accountPrivilege := strings.ToLower(m["account_privilege"].(string))
	buf.WriteString(fmt.Sprintf("%s-", dbName))
	buf.WriteString(fmt.Sprintf("%s-", accountPrivilege))
	if accountPrivilege == "custom" {
		buf.WriteString(fmt.Sprintf("%s-", mappingAndSortAccountPrivilegeStr(strings.ToLower(m["account_privilege_detail"].(string)))))
	}
	return buf
}

func RdsMysqlAccountPrivilegeHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := rdsMysqlAccountPrivilegeHashBase(m)
	logger.DebugInfo("RdsMysqlAccountPrivilegeHash %s", buf.String())
	return hashcode.String(buf.String())
}

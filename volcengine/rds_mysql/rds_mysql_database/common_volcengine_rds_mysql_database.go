package rds_mysql_database

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

var databaseImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("instance_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("db_name", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

var rdsMysqlDatabasePrivilegeStrMapping = map[string]string{
	"create tmp table": "create temporary tables",
	"CREATE TMP TABLE": "CREATE TEMPORARY TABLES",
}

func mappingDatabasePrivilegeStr(databasePrivilegeStr string) []string {
	privileges := strings.Split(databasePrivilegeStr, ",")
	mappingPrivileges := make([]string, 0)
	for _, privilege := range privileges {
		if mappedPrivilege, ok := rdsMysqlDatabasePrivilegeStrMapping[privilege]; ok {
			mappingPrivileges = append(mappingPrivileges, mappedPrivilege)
		} else {
			mappingPrivileges = append(mappingPrivileges, privilege)
		}
	}
	return mappingPrivileges
}

// mappingAndSortDatabasePrivilegeStr RDS database privilege string mapping
func mappingAndSortDatabasePrivilegeStr(databasePrivilegeStr string) string {
	mappingPrivileges := mappingDatabasePrivilegeStr(databasePrivilegeStr)
	sort.Strings(mappingPrivileges)
	return strings.Join(mappingPrivileges, ",")
}

func rdsMysqlDatabasePrivilegeHashBase(m map[string]interface{}) (buf bytes.Buffer) {
	accountName := strings.ToLower(m["account_name"].(string))
	databasePrivilege := strings.ToLower(m["account_privilege"].(string))
	buf.WriteString(fmt.Sprintf("%s-", accountName))
	buf.WriteString(fmt.Sprintf("%s-", databasePrivilege))
	if databasePrivilege == "custom" {
		buf.WriteString(fmt.Sprintf("%s-", mappingAndSortDatabasePrivilegeStr(strings.ToLower(m["account_privilege_detail"].(string)))))
	}
	return buf
}

func RdsMysqlDatabasePrivilegeHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := rdsMysqlDatabasePrivilegeHashBase(m)
	logger.DebugInfo("RdsMysqlDatabasePrivilegeHash %s", buf.String())
	return hashcode.String(buf.String())
}

package waf_acl_rule

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func wafAclRuleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'ID:AclType'")
	}
	id := items[0]
	aclType := items[1]
	aclRuleIdIInt, err := strconv.Atoi(id)
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf(" ID cannot convert to int ")
	}
	_ = d.Set("id", aclRuleIdIInt)
	_ = d.Set("acl_type", aclType)

	return []*schema.ResourceData{d}, nil
}

package network_acl_associate

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

var aclAssociateImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("network_acl_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("resource_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

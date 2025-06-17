package waf_custom_page

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func wafCustomPageImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'ID:Host'")
	}
	id := items[0]
	host := items[1]
	customPageInt, err := strconv.Atoi(id)
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf(" ID cannot convert to int ")
	}
	_ = d.Set("id", customPageInt)
	_ = d.Set("host", host)

	return []*schema.ResourceData{d}, nil
}

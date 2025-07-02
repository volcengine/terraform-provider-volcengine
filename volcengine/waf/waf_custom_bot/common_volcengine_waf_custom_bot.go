package waf_custom_bot

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

func wafCustomBotImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'ID:Host'")
	}
	id := items[0]
	host := items[1]
	_ = d.Set("id", id)
	_ = d.Set("host", host)

	return []*schema.ResourceData{d}, nil
}

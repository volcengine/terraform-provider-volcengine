package index

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var tlsIndexImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) == 2 {
		if err := data.Set("topic_id", items[1]); err != nil {
			return []*schema.ResourceData{data}, err
		}
	} else {
		if err := data.Set("topic_id", data.Id()); err != nil {
			return []*schema.ResourceData{data}, err
		}
		data.SetId(fmt.Sprintf("index:%s", data.Id()))
	}
	return []*schema.ResourceData{data}, nil
}

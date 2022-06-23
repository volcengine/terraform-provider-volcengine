package cen_attach_instance

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

var cenAttachInstanceImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 4 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("cen_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("instance_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("instance_type", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("instance_region_id", items[3]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

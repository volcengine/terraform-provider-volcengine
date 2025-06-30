package cen_service_route_entry

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var cenServiceRouteEntryImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), "#")
	if len(items) != 4 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with '#'")
	}
	if err := data.Set("cen_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("destination_cidr_block", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("service_region_id", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("service_vpc_id", items[3]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

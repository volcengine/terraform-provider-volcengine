package endpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func mongoDBEndpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'instanceId:endpointId'")
	}
	instanceId := items[0]
	endpointId := items[1]
	d.Set("instance_id", instanceId)
	d.Set("endpoint_id", endpointId)

	return []*schema.ResourceData{d}, nil
}

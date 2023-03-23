package ssl_state

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func mongoDBSSLStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'ssl:instanceId'")
	}
	d.Set("instance_id", items[1])

	return []*schema.ResourceData{d}, nil
}

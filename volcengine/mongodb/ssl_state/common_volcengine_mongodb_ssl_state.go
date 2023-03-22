package ssl_state

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

/*

Import
mongosdb ssl state can be imported using the ssl:instanceId, e.g.
set `ssl_action` to `Update` will update ssl always when terraform apply.
```
$ terraform import volcengine_mongosdb_ssl_state.default ssl:mongo-shard-d050db19xxx
```

*/

func mongoDBSSLStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'ssl:instanceId'")
	}
	d.Set("instance_id", items[1])

	return []*schema.ResourceData{d}, nil
}

package ssl_state

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongosdb ssl state can be imported using the ssl:instanceId, e.g.
set `ssl_action` to `Update` will update ssl always when terraform apply.
```
$ terraform import volcengine_mongosdb_ssl_state.default ssl:mongo-shard-d050db19xxx
```

*/

func ResourceVolcengineMongoDBSSLState() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBSSLStateCreate,
		Read:   resourceVolcengineMongoDBSSLStateRead,
		Update: resourceVolcengineMongoDBSSLStateUpdate,
		Delete: resourceVolcengineMongoDBSSLStateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of mongodb instance.",
			},
			"ssl_action": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "The action of ssl,valid value contains `Update`.",
			},
		},
	}
	return resource
}

func resourceVolcengineMongoDBSSLStateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("Error on open ssl %q,%s", d.Id(), err)
	}
	return resourceVolcengineMongoDBSSLStateRead(d, meta)
}

func resourceVolcengineMongoDBSSLStateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("error on updating ssl  %q, %s", d.Id(), err)
	}
	return resourceVolcengineMongoDBSSLStateRead(d, meta)
}

func resourceVolcengineMongoDBSSLStateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("error on close ssl %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongoDBSSLStateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("Error on reading ssl state %q,%s", d.Id(), err)
	}
	return err
}

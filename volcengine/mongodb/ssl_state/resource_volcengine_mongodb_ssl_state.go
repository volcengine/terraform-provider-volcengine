package ssl_state

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongodb ssl state can be imported using the ssl:instanceId, e.g.
```
$ terraform import volcengine_mongodb_ssl_state.default ssl:mongo-shard-d050db19xxx
```

*/

func ResourceVolcengineMongoDBSSLState() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBSSLStateCreate,
		Read:   resourceVolcengineMongoDBSSLStateRead,
		Update: resourceVolcengineMongoDBSSLStateUpdate,
		Delete: resourceVolcengineMongoDBSSLStateDelete,
		Importer: &schema.ResourceImporter{
			State: mongoDBSSLStateImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of mongodb instance.",
			},
			"ssl_action": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				ValidateFunc: validation.StringInSlice([]string{
					"Update",
				}, false),
				Description: "The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will will trigger an SSL update operation when executing `terraform apply`." +
					"When the current time is less than 30 days from the `ssl_expired_time`, executing `terraform apply` will automatically renew the SSL.",
			},
		},
	}
	dataSource := DataSourceVolcengineMongoDBSSLStates().Schema["ssl_state"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineMongoDBSSLStateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("Error on opening ssl %q, %s ", d.Id(), err)
	}
	return resourceVolcengineMongoDBSSLStateRead(d, meta)
}

func resourceVolcengineMongoDBSSLStateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("Error on updating ssl %q, %s ", d.Id(), err)
	}
	return resourceVolcengineMongoDBSSLStateRead(d, meta)
}

func resourceVolcengineMongoDBSSLStateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("Error on reading ssl state %q, %s ", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongoDBSSLStateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineMongoDBSSLState())
	if err != nil {
		return fmt.Errorf("Error on deleting ssl state %q, %s ", d.Id(), err)
	}
	return err
}

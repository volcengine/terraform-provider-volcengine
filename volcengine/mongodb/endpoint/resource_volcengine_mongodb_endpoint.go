package endpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongosdb endpoint instance can be imported using the endpoint:instanceId, e.g.
```
$ terraform import volcengine_mongosdb_instance.default endpoint:mongo-replica-e405f8e2****
```

*/

func mongoDBEndpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 || items[0] != "endpoint" {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'endpoint:instanceId'")
	}
	return []*schema.ResourceData{d}, nil
}

func ResourceVolcengineMongoDBEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBEndpointCreate,
		Read:   resourceVolcengineMongoDBEndpointRead,
		Update: resourceVolcengineMongoDBEndpointUpdate,
		Delete: resourceVolcengineMongoDBEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance id.",
			},
			"object_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The object ID corresponding to the endpoint.",
			},
			"network_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Private",
				Description:  "The network type of endpoint.",
				ValidateFunc: validation.StringInSlice([]string{"Private", "Public"}, false),
			},
			"mongos_node_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The ID of the Mongos node that needs to apply for the endpoint.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"eip_ids": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A list of EIP IDs that need to be bound when applying for endpoint.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	return resource
}

func resourceVolcengineMongoDBEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineMongoDBEndpoint())
	if err != nil {
		return fmt.Errorf("Error on creating endpoint %q,%s", d.Id(), err)
	}
	return resourceVolcengineMongoDBEndpointRead(d, meta)
}

func resourceVolcengineMongoDBEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	return fmt.Errorf("mongodb endpoint does not support update")
}

func resourceVolcengineMongoDBEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineMongoDBEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting endpoint %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongoDBEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineMongoDBEndpoint())
	if err != nil {
		return fmt.Errorf("Error on reading endpoint %q,%s", d.Id(), err)
	}
	return err
}

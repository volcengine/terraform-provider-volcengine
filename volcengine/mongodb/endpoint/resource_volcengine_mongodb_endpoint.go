package endpoint

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongodb endpoint can be imported using the instanceId:endpointId
`instanceId`: represents the instance that endpoint related to.
`endpointId`: the id of endpoint.
e.g.
```
$ terraform import volcengine_mongodb_endpoint.default mongo-replica-e405f8e2****:BRhFA0pDAk0XXkxCZQ
```

*/

func ResourceVolcengineMongoDBEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBEndpointCreate,
		Read:   resourceVolcengineMongoDBEndpointRead,
		Update: resourceVolcengineMongoDBEndpointUpdate,
		Delete: resourceVolcengineMongoDBEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: mongoDBEndpointImporter,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance where the endpoint resides.",
			},
			"object_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "A list of the Mongos node that needs to apply for the endpoint.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"eip_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "A list of EIP IDs that need to be bound when applying for endpoint.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"endpoint_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of endpoint.",
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

package vpc_gateway_endpoint

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcGatewayEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_vpc_gateway_endpoint.default gwep-273yuq6q7bgn47fap8squ****
```

*/

func ResourceVolcengineVpcGatewayEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpcGatewayEndpointCreate,
		Read:   resourceVolcengineVpcGatewayEndpointRead,
		Update: resourceVolcengineVpcGatewayEndpointUpdate,
		Delete: resourceVolcengineVpcGatewayEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the vpc.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the gateway endpoint service.",
			},
			"endpoint_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the gateway endpoint.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the gateway endpoint.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the gateway endpoint.",
			},
			"vpc_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The vpc policy of the gateway endpoint.",
			},
			"tags": ve.TagsSchema(),
			// computed fields
			"endpoint_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the gateway endpoint.",
			},
		},
	}
	return resource
}

func resourceVolcengineVpcGatewayEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcGatewayEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVpcGatewayEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating vpc_gateway_endpoint %q, %w", d.Id(), err)
	}
	return resourceVolcengineVpcGatewayEndpointRead(d, meta)
}

func resourceVolcengineVpcGatewayEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcGatewayEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVpcGatewayEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading vpc_gateway_endpoint %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcGatewayEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcGatewayEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVpcGatewayEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating vpc_gateway_endpoint %q, %w", d.Id(), err)
	}
	return resourceVolcengineVpcGatewayEndpointRead(d, meta)
}

func resourceVolcengineVpcGatewayEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcGatewayEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVpcGatewayEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting vpc_gateway_endpoint %q, %w", d.Id(), err)
	}
	return err
}

package vpc_endpoint

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint.default ep-3rel74u229dz45zsk2i6l****
```

*/

func ResourceVolcenginePrivatelinkVpcEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateLinkVpcEndpointCreate,
		Read:   resourceVolcenginePrivateLinkVpcEndpointRead,
		Update: resourceVolcenginePrivateLinkVpcEndpointUpdate,
		Delete: resourceVolcenginePrivateLinkVpcEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"security_group_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "the security group ids of vpc endpoint.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of vpc endpoint service.",
			},
			"service_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The name of vpc endpoint service.",
			},
			"endpoint_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of vpc endpoint.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of vpc endpoint.",
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vpc id of vpc endpoint.",
			},
			"endpoint_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of vpc endpoint.",
			},
			"endpoint_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The domain of vpc endpoint.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of vpc endpoint.",
			},
			"business_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether the vpc endpoint is locked.",
			},
			"connection_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The connection  status of vpc endpoint.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of vpc endpoint.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of vpc endpoint.",
			},
			"deleted_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The delete time of vpc endpoint.",
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateLinkVpcEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	vpcEndpointService := NewVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(vpcEndpointService, d, ResourceVolcenginePrivatelinkVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating vpc endpoint %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivateLinkVpcEndpointRead(d, meta)
}

func resourceVolcenginePrivateLinkVpcEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	vpcEndpointService := NewVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(vpcEndpointService, d, ResourceVolcenginePrivatelinkVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading vpc endpoint %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateLinkVpcEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	vpcEndpointService := NewVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(vpcEndpointService, d, ResourceVolcenginePrivatelinkVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating vpc endoint %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivateLinkVpcEndpointRead(d, meta)
}

func resourceVolcenginePrivateLinkVpcEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	vpcEndpointService := NewVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(vpcEndpointService, d, ResourceVolcenginePrivatelinkVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting vpc endpoint %q, %w", d.Id(), err)
	}
	return err
}

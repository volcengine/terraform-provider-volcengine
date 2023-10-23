package transit_router_direct_connect_gateway_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterDirectConnectGatewayAttachment can be imported using the transitRouterId:attachmentId, e.g.
```
$ terraform import volcengine_transit_router_direct_connect_gateway_attachment.default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
```

*/

func ResourceVolcengineTransitRouterDirectConnectGatewayAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterDirectConnectGatewayAttachmentCreate,
		Read:   resourceVolcengineTransitRouterDirectConnectGatewayAttachmentRead,
		Update: resourceVolcengineTransitRouterDirectConnectGatewayAttachmentUpdate,
		Delete: resourceVolcengineTransitRouterDirectConnectGatewayAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				var err error
				items := strings.Split(d.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form TransitRouterId:AttachmentId")
				}
				err = d.Set("transit_router_id", items[0])
				if err != nil {
					return []*schema.ResourceData{d}, err
				}
				err = d.Set("transit_router_attachment_id", items[1])
				if err != nil {
					return []*schema.ResourceData{d}, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the transit router.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the direct connect gateway.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description.",
			},
			"transit_router_attachment_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the transit router direct connect gateway attachment.",
			},
			"transit_router_attachment_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the transit router direct connect gateway attachment.",
			},
		},
	}
	return resource
}

func resourceVolcengineTransitRouterDirectConnectGatewayAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterDirectConnectGatewayAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterDirectConnectGatewayAttachment())
	if err != nil {
		return fmt.Errorf("error on creating transit_router_direct_connect_gateway_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterDirectConnectGatewayAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterDirectConnectGatewayAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterDirectConnectGatewayAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterDirectConnectGatewayAttachment())
	if err != nil {
		return fmt.Errorf("error on reading transit_router_direct_connect_gateway_attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterDirectConnectGatewayAttachmentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterDirectConnectGatewayAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterDirectConnectGatewayAttachment())
	if err != nil {
		return fmt.Errorf("error on updating transit_router_direct_connect_gateway_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterDirectConnectGatewayAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterDirectConnectGatewayAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterDirectConnectGatewayAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterDirectConnectGatewayAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting transit_router_direct_connect_gateway_attachment %q, %s", d.Id(), err)
	}
	return err
}

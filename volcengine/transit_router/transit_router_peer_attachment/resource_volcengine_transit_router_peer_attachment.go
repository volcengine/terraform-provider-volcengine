package transit_router_peer_attachment

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterPeerAttachment can be imported using the id, e.g.
```
$ terraform import volcengine_transit_router_peer_attachment.default tr-attach-12be67d0yh2io17q7y1au****
```

*/

func ResourceVolcengineTransitRouterPeerAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterPeerAttachmentCreate,
		Read:   resourceVolcengineTransitRouterPeerAttachmentRead,
		Update: resourceVolcengineTransitRouterPeerAttachmentUpdate,
		Delete: resourceVolcengineTransitRouterPeerAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Description: "The id of the local transit router.",
			},
			"peer_transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the peer transit router.",
			},
			"peer_transit_router_region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region id of the peer transit router.",
			},
			"transit_router_attachment_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the transit router peer attachment.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the transit router peer attachment.",
			},
			"transit_router_bandwidth_package_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The bandwidth package id of the transit router peer attachment. When specifying this field, the field `bandwidth` must also be specified.",
			},
			"bandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("transit_router_bandwidth_package_id").(string) == ""
				},
				Description: "The bandwidth of the transit router peer attachment. Unit: Mbps.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the transit router peer attachment.",
			},
			"transit_router_route_table_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The route table id of the transit router peer attachment.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the transit router peer attachment.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the transit router peer attachment.",
			},
		},
	}
	return resource
}

func resourceVolcengineTransitRouterPeerAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterPeerAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterPeerAttachment())
	if err != nil {
		return fmt.Errorf("error on creating transit_router_peer_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterPeerAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterPeerAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterPeerAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterPeerAttachment())
	if err != nil {
		return fmt.Errorf("error on reading transit_router_peer_attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterPeerAttachmentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterPeerAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterPeerAttachment())
	if err != nil {
		return fmt.Errorf("error on updating transit_router_peer_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterPeerAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterPeerAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterPeerAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterPeerAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting transit_router_peer_attachment %q, %s", d.Id(), err)
	}
	return err
}

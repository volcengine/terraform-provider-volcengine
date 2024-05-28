package transit_router_vpn_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterVpnAttachment can be imported using the transitRouterId:attachmentId, e.g.
```
$ terraform import volcengine_transit_router_vpn_attachment.default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
```

*/

func ResourceVolcengineTransitRouterVpnAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterVpnAttachmentCreate,
		Read:   resourceVolcengineTransitRouterVpnAttachmentRead,
		Update: resourceVolcengineTransitRouterVpnAttachmentUpdate,
		Delete: resourceVolcengineTransitRouterVpnAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				var err error
				items := strings.Split(d.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form VolumeId:instanceId")
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
			"vpn_connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the IPSec connection.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the availability zone.",
			},
			"transit_router_attachment_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the transit router vpn attachment.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the transit router vpn attachment.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"transit_router_attachment_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the transit router vpn attachment.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the transit router.",
			},
		},
	}
	return resource
}

func resourceVolcengineTransitRouterVpnAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRVpnAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterVpnAttachment())
	if err != nil {
		return fmt.Errorf("error on creating transit router vpn attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterVpnAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterVpnAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRVpnAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterVpnAttachment())
	if err != nil {
		return fmt.Errorf("error on reading transit router vpn attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterVpnAttachmentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRVpnAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterVpnAttachment())
	if err != nil {
		return fmt.Errorf("error on updating transit router vpn attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterVpnAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterVpnAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRVpnAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterVpnAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting transit router vpn attachment %q, %s", d.Id(), err)
	}
	return err
}

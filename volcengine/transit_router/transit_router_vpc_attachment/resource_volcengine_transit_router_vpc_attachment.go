package transit_router_vpc_attachment

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterVpcAttachment can be imported using the transitRouterId:attachmentId, e.g.
```
$ terraform import volcengine_transit_router_vpc_attachment.default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
```

*/

func ResourceVolcengineTransitRouterVpcAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterVpcAttachmentCreate,
		Read:   resourceVolcengineTransitRouterVpcAttachmentRead,
		Update: resourceVolcengineTransitRouterVpcAttachmentUpdate,
		Delete: resourceVolcengineTransitRouterVpcAttachmentDelete,
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
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of vpc.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the transit router vpc attachment.",
			},
			"transit_router_attachment_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the transit router vpc attachment.",
			},
			"tags": ve.TagsSchema(),
			"attach_points": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The attach points of transit router vpc attachment.",
				Set:         hash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of subnet.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of zone.",
						},
					},
				},
			},

			"transit_router_attachment_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the transit router attachment.",
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

func hashBase(m map[string]interface{}) (buf bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["subnet_id"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["zone_id"].(string))))
	return buf
}

func hash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := hashBase(m)
	return hashcode.String(buf.String())
}

func resourceVolcengineTransitRouterVpcAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterVpcAttachment())
	if err != nil {
		return fmt.Errorf("error on creating transit router vpc attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterVpcAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterVpcAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterVpcAttachment())
	if err != nil {
		return fmt.Errorf("error on reading transit router vpc attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterVpcAttachmentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterVpcAttachment())
	if err != nil {
		return fmt.Errorf("error on updating transit router vpc attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterVpcAttachmentRead(d, meta)
}

func resourceVolcengineTransitRouterVpcAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterVpcAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting transit router vpc attachment %q, %s", d.Id(), err)
	}
	return err
}

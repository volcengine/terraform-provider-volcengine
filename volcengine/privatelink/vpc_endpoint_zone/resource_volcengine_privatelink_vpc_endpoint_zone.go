package vpc_endpoint_zone

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcEndpointZone can be imported using the endpointId:subnetId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_zone.default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
```

*/

func ResourceVolcenginePrivatelinkVpcEndpointZone() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateLinkVpcEndpointZoneCreate,
		Read:   resourceVolcenginePrivateLinkVpcEndpointZoneRead,
		Delete: resourceVolcenginePrivateLinkVpcEndpointZoneDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("endpoint_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("subnet_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"endpoint_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The endpoint id of vpc endpoint zone.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id of vpc endpoint zone.",
			},
			"private_ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The private ip address of vpc endpoint zone.",
			},

			"zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Id of vpc endpoint zone.",
			},
			"zone_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The domain of vpc endpoint zone.",
			},
			"zone_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of vpc endpoint zone.",
			},
			"network_interface_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network interface id of vpc endpoint.",
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateLinkVpcEndpointZoneCreate(d *schema.ResourceData, meta interface{}) (err error) {
	vpcEndpointZoneService := NewVpcEndpointZoneService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(vpcEndpointZoneService, d, ResourceVolcenginePrivatelinkVpcEndpointZone())
	if err != nil {
		return fmt.Errorf("error on creating vpc endpoint zone %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivateLinkVpcEndpointZoneRead(d, meta)
}

func resourceVolcenginePrivateLinkVpcEndpointZoneRead(d *schema.ResourceData, meta interface{}) (err error) {
	vpcEndpointZoneService := NewVpcEndpointZoneService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(vpcEndpointZoneService, d, ResourceVolcenginePrivatelinkVpcEndpointZone())
	if err != nil {
		return fmt.Errorf("error on reading vpc endpoint zone %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateLinkVpcEndpointZoneDelete(d *schema.ResourceData, meta interface{}) (err error) {
	vpcEndpointZoneService := NewVpcEndpointZoneService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(vpcEndpointZoneService, d, ResourceVolcenginePrivatelinkVpcEndpointZone())
	if err != nil {
		return fmt.Errorf("error on deleting vpc endpoint zone %q, %w", d.Id(), err)
	}
	return err
}

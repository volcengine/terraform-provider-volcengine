package ipv6_address_bandwidth

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Ipv6AddressBandwidth can be imported using the id, e.g.
```
$ terraform import volcengine_vpc_ipv6_address_bandwidth.default eip-2fede9fsgnr4059gp674m6ney
```

*/

func ResourceVolcengineIpv6AddressBandwidth() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIpv6AddressBandwidthCreate,
		Read:   resourceVolcengineIpv6AddressBandwidthRead,
		Update: resourceVolcengineIpv6AddressBandwidthUpdate,
		Delete: resourceVolcengineIpv6AddressBandwidthDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ipv6_address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Ipv6 address.",
			},
			"billing_type": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "BillingType of the Ipv6 bandwidth. Valid values: 3(Pay by Traffic).",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Peek bandwidth of the Ipv6 address. Valid values: 1 to 200. Unit: Mbit/s.",
			},
		},
	}
	dataSource := DataSourceVolcengineIpv6AddressBandwidths().Schema["ipv6_address_bandwidths"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineIpv6AddressBandwidthCreate(d *schema.ResourceData, meta interface{}) (err error) {
	ipv6AddressBandwidthService := NewIpv6AddressBandwidthService(meta.(*ve.SdkClient))
	err = ipv6AddressBandwidthService.Dispatcher.Create(ipv6AddressBandwidthService, d, ResourceVolcengineIpv6AddressBandwidth())
	if err != nil {
		return fmt.Errorf("error on creating Ipv6AddressBandwidth %q, %w", d.Id(), err)
	}
	return resourceVolcengineIpv6AddressBandwidthRead(d, meta)
}

func resourceVolcengineIpv6AddressBandwidthRead(d *schema.ResourceData, meta interface{}) (err error) {
	ipv6AddressBandwidthService := NewIpv6AddressBandwidthService(meta.(*ve.SdkClient))
	err = ipv6AddressBandwidthService.Dispatcher.Read(ipv6AddressBandwidthService, d, ResourceVolcengineIpv6AddressBandwidth())
	if err != nil {
		return fmt.Errorf("error on reading Ipv6AddressBandwidth %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineIpv6AddressBandwidthUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	ipv6AddressBandwidthService := NewIpv6AddressBandwidthService(meta.(*ve.SdkClient))
	err = ipv6AddressBandwidthService.Dispatcher.Update(ipv6AddressBandwidthService, d, ResourceVolcengineIpv6AddressBandwidth())
	if err != nil {
		return fmt.Errorf("error on updating Ipv6AddressBandwidth %q, %w", d.Id(), err)
	}
	return resourceVolcengineIpv6AddressBandwidthRead(d, meta)
}

func resourceVolcengineIpv6AddressBandwidthDelete(d *schema.ResourceData, meta interface{}) (err error) {
	ipv6AddressBandwidthService := NewIpv6AddressBandwidthService(meta.(*ve.SdkClient))
	err = ipv6AddressBandwidthService.Dispatcher.Delete(ipv6AddressBandwidthService, d, ResourceVolcengineIpv6AddressBandwidth())
	if err != nil {
		return fmt.Errorf("error on deleting Ipv6AddressBandwidth %q, %w", d.Id(), err)
	}
	return err
}

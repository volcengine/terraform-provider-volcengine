package dnat_entry

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Dnat entry can be imported using the id, e.g.
```
$ terraform import volcengine_dnat_entry.default dnat-3fvhk47kf56****
```

*/

func ResourceVolcengineDnatEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineDnatEntryCreate,
		Update: resourceVolcengineDnatEntryUpdate,
		Read:   resourceVolcengineDnatEntryRead,
		Delete: resourceVolcengineDnatEntryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"nat_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the nat gateway to which the entry belongs.",
			},
			"external_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Provides the public IP address for public network access.",
			},
			"external_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Provides the public port for public network access.",
			},
			"internal_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Provides the internal IP address.",
			},
			"internal_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Provides the internal port.",
			},
			"protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
				Description:  "The network protocol.",
			},
			"dnat_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the DNAT rule.",
			},
			"dnat_entry_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the DNAT rule.",
			},
		},
	}
}

func resourceVolcengineDnatEntryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnatEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineDnatEntry())
	if err != nil {
		return fmt.Errorf("error on creating dnat entry: %q, %w", d.Id(), err)
	}
	return resourceVolcengineDnatEntryRead(d, meta)
}

func resourceVolcengineDnatEntryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnatEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineDnatEntry())
	if err != nil {
		return fmt.Errorf("error on reading dnat entry: %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineDnatEntryUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnatEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineDnatEntry())
	if err != nil {
		return fmt.Errorf("error on updating dnat entry: %q, %w", d.Id(), err)
	}
	return resourceVolcengineDnatEntryRead(d, meta)
}

func resourceVolcengineDnatEntryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnatEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineDnatEntry())
	if err != nil {
		return fmt.Errorf("error on deleting dnat entry: %q, %w", d.Id(), err)
	}
	return nil
}

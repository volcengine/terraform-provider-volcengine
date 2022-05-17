package snat_entry

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Snat entry can be imported using the id, e.g.
```
$ terraform import vestack_snat_entry.default snat-3fvhk47kf56****
```

*/

func ResourceVestackSnatEntry() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVestackSnatEntryDelete,
		Create: resourceVestackSnatEntryCreate,
		Read:   resourceVestackSnatEntryRead,
		Update: resourceVestackSnatEntryUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"nat_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the nat gateway to which the entry belongs.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the subnet that is required to access the internet.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the public ip address used by the SNAT entry.",
			},
			"snat_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the SNAT entry.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the SNAT entry.",
			},
		},
	}
}

func resourceVestackSnatEntryCreate(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := snatEntryService.Dispatcher.Create(snatEntryService, d, ResourceVestackSnatEntry()); err != nil {
		return fmt.Errorf("error on creating snat entry  %q, %w", d.Id(), err)
	}
	return resourceVestackSnatEntryRead(d, meta)
}

func resourceVestackSnatEntryRead(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := snatEntryService.Dispatcher.Read(snatEntryService, d, ResourceVestackSnatEntry()); err != nil {
		return fmt.Errorf("error on reading snat entry %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackSnatEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := snatEntryService.Dispatcher.Update(snatEntryService, d, ResourceVestackSnatEntry()); err != nil {
		return fmt.Errorf("error on updating snat entry %q, %w", d.Id(), err)
	}
	return resourceVestackSnatEntryRead(d, meta)
}

func resourceVestackSnatEntryDelete(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := snatEntryService.Dispatcher.Delete(snatEntryService, d, ResourceVestackSnatEntry()); err != nil {
		return fmt.Errorf("error on deleting snat entry %q, %w", d.Id(), err)
	}
	return nil
}

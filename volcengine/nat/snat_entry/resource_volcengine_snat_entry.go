package snat_entry

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Snat entry can be imported using the id, e.g.
```
$ terraform import volcengine_snat_entry.default snat-3fvhk47kf56****
```

*/

func ResourceVolcengineSnatEntry() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineSnatEntryDelete,
		Create: resourceVolcengineSnatEntryCreate,
		Read:   resourceVolcengineSnatEntryRead,
		Update: resourceVolcengineSnatEntryUpdate,
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
			"subnet_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"subnet_id", "source_cidr"},
				Description:  "The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.",
			},
			"eip_id": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if len(old) != len(new) {
						return false
					}
					oldArr := strings.Split(old, ",")
					newArr := strings.Split(new, ",")
					sort.Strings(oldArr)
					sort.Strings(newArr)
					return reflect.DeepEqual(oldArr, newArr)
				},
				ConflictsWith: []string{"nat_ip_id"},
				Description:   "The id of the public ip address used by the SNAT entry. This field is required when the nat gateway is a internet NAT gateway.",
			},
			"nat_ip_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"eip_id"},
				Description:   "The ID of the intranet NAT gateway's transit IP. This field is required when the nat gateway is a intranet NAT gateway.",
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
			"source_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"subnet_id", "source_cidr"},
				Description:  "The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.",
			},
		},
	}
}

func resourceVolcengineSnatEntryCreate(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(snatEntryService, d, ResourceVolcengineSnatEntry()); err != nil {
		return fmt.Errorf("error on creating snat entry  %q, %w", d.Id(), err)
	}
	return resourceVolcengineSnatEntryRead(d, meta)
}

func resourceVolcengineSnatEntryRead(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(snatEntryService, d, ResourceVolcengineSnatEntry()); err != nil {
		return fmt.Errorf("error on reading snat entry %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineSnatEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(snatEntryService, d, ResourceVolcengineSnatEntry()); err != nil {
		return fmt.Errorf("error on updating snat entry %q, %w", d.Id(), err)
	}
	return resourceVolcengineSnatEntryRead(d, meta)
}

func resourceVolcengineSnatEntryDelete(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(snatEntryService, d, ResourceVolcengineSnatEntry()); err != nil {
		return fmt.Errorf("error on deleting snat entry %q, %w", d.Id(), err)
	}
	return nil
}

package acl

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/acl_entry"
)

/*

Import
Acl can be imported using the id, e.g.
```
$ terraform import vestack_acl.default acl-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVestackAcl() *schema.Resource {
	entry := acl_entry.ResourceVestackAclEntry().Schema
	for k, v := range entry {
		if k == "acl_id" {
			delete(entry, k)
		} else {
			v.ForceNew = false
		}
	}

	return &schema.Resource{
		Create: resourceVestackAclCreate,
		Read:   resourceVestackAclRead,
		Update: resourceVestackAclUpdate,
		Delete: resourceVestackAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of Acl.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Acl.",
			},
			"acl_entries": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The acl entry set of the Acl.",
				Set:         ve.ClbAclEntryHash,
				Elem: &schema.Resource{
					Schema: entry,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of Acl.",
			},
		},
	}
}

func resourceVestackAclCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Create(aclService, d, ResourceVestackAcl())
	if err != nil {
		return fmt.Errorf("error on creating acl %q, %w", d.Id(), err)
	}
	return resourceVestackAclRead(d, meta)
}

func resourceVestackAclRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Read(aclService, d, ResourceVestackAcl())
	if err != nil {
		return fmt.Errorf("error on reading acl %q, %w", d.Id(), err)
	}
	return err
}

func resourceVestackAclUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Update(aclService, d, ResourceVestackAcl())
	if err != nil {
		return fmt.Errorf("error on updating acl %q, %w", d.Id(), err)
	}
	return resourceVestackAclRead(d, meta)
}

func resourceVestackAclDelete(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Delete(aclService, d, ResourceVestackAcl())
	if err != nil {
		return fmt.Errorf("error on deleting acl %q, %w", d.Id(), err)
	}
	return err
}

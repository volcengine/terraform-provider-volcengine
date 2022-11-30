package acl

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl_entry"
)

/*

Import
Acl can be imported using the id, e.g.
```
$ terraform import volcengine_acl.default acl-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineAcl() *schema.Resource {
	entry := acl_entry.ResourceVolcengineAclEntry().Schema
	for k, v := range entry {
		if k == "acl_id" {
			delete(entry, k)
		} else {
			v.ForceNew = false
		}
	}

	return &schema.Resource{
		Create: resourceVolcengineAclCreate,
		Read:   resourceVolcengineAclRead,
		Update: resourceVolcengineAclUpdate,
		Delete: resourceVolcengineAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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

func resourceVolcengineAclCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Create(aclService, d, ResourceVolcengineAcl())
	if err != nil {
		return fmt.Errorf("error on creating acl %q, %w", d.Id(), err)
	}
	return resourceVolcengineAclRead(d, meta)
}

func resourceVolcengineAclRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Read(aclService, d, ResourceVolcengineAcl())
	if err != nil {
		return fmt.Errorf("error on reading acl %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineAclUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Update(aclService, d, ResourceVolcengineAcl())
	if err != nil {
		return fmt.Errorf("error on updating acl %q, %w", d.Id(), err)
	}
	return resourceVolcengineAclRead(d, meta)
}

func resourceVolcengineAclDelete(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Delete(aclService, d, ResourceVolcengineAcl())
	if err != nil {
		return fmt.Errorf("error on deleting acl %q, %w", d.Id(), err)
	}
	return err
}

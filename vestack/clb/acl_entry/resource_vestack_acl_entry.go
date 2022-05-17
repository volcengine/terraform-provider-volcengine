package acl_entry

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
AclEntry can be imported using the id, e.g.
```
$ terraform import vestack_acl_entry.default ID is a string concatenated with colons(AclId:Entry)
```

*/

func ResourceVestackAclEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackAclEntryCreate,
		Read:   resourceVestackAclEntryRead,
		Delete: resourceVestackAclEntryDelete,
		Importer: &schema.ResourceImporter{
			State: aclEntryImporter,
		},
		Schema: map[string]*schema.Schema{
			"acl_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of Acl.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the AclEntry.",
			},
			"entry": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The content of the AclEntry.",
			},
		},
	}
}

func resourceVestackAclEntryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclEntryService := NewAclEntryService(meta.(*ve.SdkClient))
	err = aclEntryService.Dispatcher.Create(aclEntryService, d, ResourceVestackAclEntry())
	if err != nil {
		return fmt.Errorf("error on creating acl entry %q, %w", d.Id(), err)
	}
	return resourceVestackAclEntryRead(d, meta)
}

func resourceVestackAclEntryRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclEntryService := NewAclEntryService(meta.(*ve.SdkClient))
	err = aclEntryService.Dispatcher.Read(aclEntryService, d, ResourceVestackAclEntry())
	if err != nil {
		return fmt.Errorf("error on reading acl entry %q, %w", d.Id(), err)
	}
	return err
}

func resourceVestackAclEntryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	aclEntryService := NewAclEntryService(meta.(*ve.SdkClient))
	err = aclEntryService.Dispatcher.Delete(aclEntryService, d, ResourceVestackAclEntry())
	if err != nil {
		return fmt.Errorf("error on deleting acl entry %q, %w", d.Id(), err)
	}
	return err
}

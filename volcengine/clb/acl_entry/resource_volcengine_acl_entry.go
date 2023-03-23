package acl_entry

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AclEntry can be imported using the id, e.g.
```
$ terraform import volcengine_acl_entry.default ID is a string concatenated with colons(AclId:Entry)
```

*/

func ResourceVolcengineAclEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineAclEntryCreate,
		Read:   resourceVolcengineAclEntryRead,
		Delete: resourceVolcengineAclEntryDelete,
		Importer: &schema.ResourceImporter{
			State: aclEntryImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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

func resourceVolcengineAclEntryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclEntryService := NewAclEntryService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(aclEntryService, d, ResourceVolcengineAclEntry())
	if err != nil {
		return fmt.Errorf("error on creating acl entry %q, %w", d.Id(), err)
	}
	return resourceVolcengineAclEntryRead(d, meta)
}

func resourceVolcengineAclEntryRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclEntryService := NewAclEntryService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(aclEntryService, d, ResourceVolcengineAclEntry())
	if err != nil {
		return fmt.Errorf("error on reading acl entry %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineAclEntryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	aclEntryService := NewAclEntryService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(aclEntryService, d, ResourceVolcengineAclEntry())
	if err != nil {
		return fmt.Errorf("error on deleting acl entry %q, %w", d.Id(), err)
	}
	return err
}
package alb_acl

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Acl can be imported using the id, e.g.
```
$ terraform import volcengine_alb_acl.default acl-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineAlbAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineAclCreate,
		Read:   resourceVolcengineAclRead,
		Update: resourceVolcengineAclUpdate,
		Delete: resourceVolcengineAclDelete,
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
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The project name of the Acl.",
			},
			"acl_entries": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The acl entry set of the Acl.",
				Set:         AclEntryHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the AclEntry.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description of the AclEntry.",
						},
					},
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
	err = aclService.Dispatcher.Create(aclService, d, ResourceVolcengineAlbAcl())
	if err != nil {
		return fmt.Errorf("error on creating acl %q, %w", d.Id(), err)
	}
	return resourceVolcengineAclRead(d, meta)
}

func resourceVolcengineAclRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Read(aclService, d, ResourceVolcengineAlbAcl())
	if err != nil {
		return fmt.Errorf("error on reading acl %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineAclUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Update(aclService, d, ResourceVolcengineAlbAcl())
	if err != nil {
		return fmt.Errorf("error on updating acl %q, %w", d.Id(), err)
	}
	return resourceVolcengineAclRead(d, meta)
}

func resourceVolcengineAclDelete(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewAclService(meta.(*ve.SdkClient))
	err = aclService.Dispatcher.Delete(aclService, d, ResourceVolcengineAlbAcl())
	if err != nil {
		return fmt.Errorf("error on deleting acl %q, %w", d.Id(), err)
	}
	return err
}

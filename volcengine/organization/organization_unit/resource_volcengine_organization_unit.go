package organization_unit

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
OrganizationUnit can be imported using the id, e.g.
```
$ terraform import volcengine_organization_unit.default ID
```

*/

func ResourceVolcengineOrganizationUnit() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineOrganizationUnitCreate,
		Read:   resourceVolcengineOrganizationUnitRead,
		Update: resourceVolcengineOrganizationUnitUpdate,
		Delete: resourceVolcengineOrganizationUnitDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"parent_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Parent Organization Unit ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the organization unit.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the organization unit.",
			},
			"org_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the organization.",
			},
			"org_type": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The organization type.",
			},
			"depth": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The depth of the organization unit.",
			},
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner of the organization unit.",
			},
		},
	}
	return resource
}

func resourceVolcengineOrganizationUnitCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewOrganizationUnitService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineOrganizationUnit())
	if err != nil {
		return fmt.Errorf("error on creating organization_unit %q, %s", d.Id(), err)
	}
	return resourceVolcengineOrganizationUnitRead(d, meta)
}

func resourceVolcengineOrganizationUnitRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewOrganizationUnitService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineOrganizationUnit())
	if err != nil {
		return fmt.Errorf("error on reading organization_unit %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineOrganizationUnitUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewOrganizationUnitService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineOrganizationUnit())
	if err != nil {
		return fmt.Errorf("error on updating organization_unit %q, %s", d.Id(), err)
	}
	return resourceVolcengineOrganizationUnitRead(d, meta)
}

func resourceVolcengineOrganizationUnitDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewOrganizationUnitService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineOrganizationUnit())
	if err != nil {
		return fmt.Errorf("error on deleting organization_unit %q, %s", d.Id(), err)
	}
	return err
}

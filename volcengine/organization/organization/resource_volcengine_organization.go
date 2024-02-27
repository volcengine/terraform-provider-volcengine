package organization

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Organization can be imported using the id, e.g.
```
$ terraform import volcengine_organization.default resource_id
```

*/

func ResourceVolcengineOrganization() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineOrganizationCreate,
		Read:   resourceVolcengineOrganizationRead,
		Delete: resourceVolcengineOrganizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			// computed fields
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner id of the organization.",
			},
			"type": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The type of the organization.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the organization.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the organization.",
			},
			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status of the organization.",
			},
			"delete_uk": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The delete uk of the organization.",
			},
			"account_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The account id of the organization owner.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account name of the organization owner.",
			},
			"main_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The main name of the organization owner.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The created time of the organization.",
			},
			"updated_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The updated time of the organization.",
			},
			"deleted_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The deleted time of the organization.",
			},
		},
	}
	return resource
}

func resourceVolcengineOrganizationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewOrganizationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineOrganization())
	if err != nil {
		return fmt.Errorf("error on creating organization %q, %s", d.Id(), err)
	}
	return resourceVolcengineOrganizationRead(d, meta)
}

func resourceVolcengineOrganizationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewOrganizationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineOrganization())
	if err != nil {
		return fmt.Errorf("error on reading organization %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineOrganizationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewOrganizationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineOrganization())
	if err != nil {
		return fmt.Errorf("error on deleting organization %q, %s", d.Id(), err)
	}
	return err
}

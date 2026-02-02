package iam_service_linked_role

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
IamServiceLinkedRole can be imported using the id, e.g.
```
$ terraform import volcengine_iam_service_linked_role.default service_name:role_name
```

*/

func ResourceVolcengineIamServiceLinkedRole() *schema.Resource {
	tagsSchema := ve.TagsSchema()
	tagsSchema.ForceNew = true
	resource := &schema.Resource{
		Create: resourceVolcengineIamServiceLinkedRoleCreate,
		Read:   resourceVolcengineIamServiceLinkedRoleRead,
		Delete: resourceVolcengineIamServiceLinkedRoleDelete,
		Importer: &schema.ResourceImporter{
			State: iamServiceLinkedRoleImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the service.",
			},

			// computed fields
			"role_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the role.",
			},
			"role_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The id of the role.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the role.",
			},
			"tags": tagsSchema,
		},
	}
	return resource
}

func resourceVolcengineIamServiceLinkedRoleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamServiceLinkedRoleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamServiceLinkedRole())
	if err != nil {
		return fmt.Errorf("error on creating iam_service_linked_role %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamServiceLinkedRoleRead(d, meta)
}

func resourceVolcengineIamServiceLinkedRoleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamServiceLinkedRoleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamServiceLinkedRole())
	if err != nil {
		return fmt.Errorf("error on reading iam_service_linked_role %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamServiceLinkedRoleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamServiceLinkedRoleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineIamServiceLinkedRole())
	if err != nil {
		return fmt.Errorf("error on updating iam_service_linked_role %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamServiceLinkedRoleRead(d, meta)
}

func resourceVolcengineIamServiceLinkedRoleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamServiceLinkedRoleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamServiceLinkedRole())
	if err != nil {
		return fmt.Errorf("error on deleting iam_service_linked_role %q, %s", d.Id(), err)
	}
	return err
}

var iamServiceLinkedRoleImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id is invalid")
	}
	if err := data.Set("service_name", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("role_name", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

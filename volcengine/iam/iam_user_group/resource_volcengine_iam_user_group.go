package iam_user_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
IamUserGroup can be imported using the id, e.g.
```
$ terraform import volcengine_iam_user_group.default user_group_name
```

*/

func ResourceVolcengineIamUserGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIamUserGroupCreate,
		Read:   resourceVolcengineIamUserGroupRead,
		Update: resourceVolcengineIamUserGroupUpdate,
		Delete: resourceVolcengineIamUserGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"user_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the user group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the user group.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The display name of the user group.",
			},
		},
	}
	return resource
}

func resourceVolcengineIamUserGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamUserGroup())
	if err != nil {
		return fmt.Errorf("error on creating iam_user_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamUserGroupRead(d, meta)
}

func resourceVolcengineIamUserGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamUserGroup())
	if err != nil {
		return fmt.Errorf("error on reading iam_user_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamUserGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineIamUserGroup())
	if err != nil {
		return fmt.Errorf("error on updating iam_user_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamUserGroupRead(d, meta)
}

func resourceVolcengineIamUserGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamUserGroup())
	if err != nil {
		return fmt.Errorf("error on deleting iam_user_group %q, %s", d.Id(), err)
	}
	return err
}

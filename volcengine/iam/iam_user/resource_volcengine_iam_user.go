package iam_user

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam user can be imported using the UserName, e.g.
```
$ terraform import volcengine_iam_user.default user_name
```

*/

func ResourceVolcengineIamUser() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIamUserCreate,
		Read:   resourceVolcengineIamUserRead,
		Update: resourceVolcengineIamUserUpdate,
		Delete: resourceVolcengineIamUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the user.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The display name of the user.",
			},
			"mobile_phone": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The mobile phone of the user.",
				DiffSuppressFunc: phoneDiffSuppressFunc,
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The email of the user.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the user.",
			},
		},
	}
	ve.MergeDateSourceToResource(DataSourceVolcengineIamUsers().Schema["users"].Elem.(*schema.Resource).Schema, &resource.Schema)
	return resource
}

func resourceVolcengineIamUserCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamUser())
	if err != nil {
		return fmt.Errorf("error on creating iam user  %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamUserRead(d, meta)
}

func resourceVolcengineIamUserRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamUser())
	if err != nil {
		return fmt.Errorf("error on reading iam user %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamUserUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineIamUser())
	if err != nil {
		return fmt.Errorf("error on updating iam user %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamUserRead(d, meta)
}

func resourceVolcengineIamUserDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamUser())
	if err != nil {
		return fmt.Errorf("error on deleting iam user %q, %s", d.Id(), err)
	}
	return err
}

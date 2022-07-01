package iam_login_profile

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Login profile can be imported using the UserName, e.g.
```
$ terraform import volcengine_iam_login_profile.default user_name
```

*/

func ResourceVolcengineIamLoginProfile() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIamLoginProfileCreate,
		Read:   resourceVolcengineIamLoginProfileRead,
		Update: resourceVolcengineIamLoginProfileUpdate,
		Delete: resourceVolcengineIamLoginProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The user name.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The password.",
			},
			"login_allowed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The flag of login allowed.",
			},
			"password_reset_required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is required reset password when next time login in.",
			},
		},
	}
	return resource
}

func resourceVolcengineIamLoginProfileCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamLoginProfile())
	if err != nil {
		return fmt.Errorf("error on creating login profile %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamLoginProfileRead(d, meta)
}

func resourceVolcengineIamLoginProfileRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamLoginProfile())
	if err != nil {
		return fmt.Errorf("error on reading login profile %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamLoginProfileUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineIamLoginProfile())
	if err != nil {
		return fmt.Errorf("error on updating login profile %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamLoginProfileRead(d, meta)
}

func resourceVolcengineIamLoginProfileDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamLoginProfile())
	if err != nil {
		return fmt.Errorf("error on deleting login profile %q, %s", d.Id(), err)
	}
	return err
}

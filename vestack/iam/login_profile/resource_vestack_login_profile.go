package login_profile

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Login profile can be imported using the UserName, e.g.
```
$ terraform import vestack_login_profile.default user_name
```

*/

func ResourceVestackLoginProfile() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackLoginProfileCreate,
		Read:   resourceVestackLoginProfileRead,
		Update: resourceVestackLoginProfileUpdate,
		Delete: resourceVestackLoginProfileDelete,
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

func resourceVestackLoginProfileCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVestackLoginProfile())
	if err != nil {
		return fmt.Errorf("error on creating login profile %q, %s", d.Id(), err)
	}
	return resourceVestackLoginProfileRead(d, meta)
}

func resourceVestackLoginProfileRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVestackLoginProfile())
	if err != nil {
		return fmt.Errorf("error on reading login profile %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackLoginProfileUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVestackLoginProfile())
	if err != nil {
		return fmt.Errorf("error on updating login profile %q, %s", d.Id(), err)
	}
	return resourceVestackLoginProfileRead(d, meta)
}

func resourceVestackLoginProfileDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewLoginProfileService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVestackLoginProfile())
	if err != nil {
		return fmt.Errorf("error on deleting login profile %q, %s", d.Id(), err)
	}
	return err
}

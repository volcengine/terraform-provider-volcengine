package cloud_identity_user

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudIdentityUser can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_user.default resource_id
```

*/

func ResourceVolcengineCloudIdentityUser() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudIdentityUserCreate,
		Read:   resourceVolcengineCloudIdentityUserRead,
		Update: resourceVolcengineCloudIdentityUserUpdate,
		Delete: resourceVolcengineCloudIdentityUserDelete,
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
				Description: "The name of the cloud identity user.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The display name of the cloud identity user.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the cloud identity user.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The email of the cloud identity user.",
			},
			"phone": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The phone of the cloud identity user. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},

			// computed fields
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source of the cloud identity user.",
			},
			"identity_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identity type of the cloud identity user.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudIdentityUserCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudIdentityUser())
	if err != nil {
		return fmt.Errorf("error on creating cloud_identity_user %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityUserRead(d, meta)
}

func resourceVolcengineCloudIdentityUserRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudIdentityUser())
	if err != nil {
		return fmt.Errorf("error on reading cloud_identity_user %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudIdentityUserUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudIdentityUser())
	if err != nil {
		return fmt.Errorf("error on updating cloud_identity_user %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityUserRead(d, meta)
}

func resourceVolcengineCloudIdentityUserDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudIdentityUser())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_identity_user %q, %s", d.Id(), err)
	}
	return err
}

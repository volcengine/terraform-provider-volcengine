package cloud_identity_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudIdentityGroup can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_group.default resource_id
```

*/

func ResourceVolcengineCloudIdentityGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudIdentityGroupCreate,
		Read:   resourceVolcengineCloudIdentityGroupRead,
		Update: resourceVolcengineCloudIdentityGroupUpdate,
		Delete: resourceVolcengineCloudIdentityGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the cloud identity group.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The display name of the cloud identity group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the cloud identity group.",
			},
			"join_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The user join type of the cloud identity group.",
			},

			// computed fields
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source of the cloud identity group.",
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The member user info of the cloud identity group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity user.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cloud identity user.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the cloud identity user.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cloud identity user.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the cloud identity user.",
						},
						"email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email of the cloud identity user.",
						},
						"phone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The phone of the cloud identity user.",
						},
						"identity_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identity type of the cloud identity user.",
						},
						"join_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The join time of the cloud identity user.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineCloudIdentityGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudIdentityGroup())
	if err != nil {
		return fmt.Errorf("error on creating cloud_identity_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityGroupRead(d, meta)
}

func resourceVolcengineCloudIdentityGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudIdentityGroup())
	if err != nil {
		return fmt.Errorf("error on reading cloud_identity_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudIdentityGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudIdentityGroup())
	if err != nil {
		return fmt.Errorf("error on updating cloud_identity_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityGroupRead(d, meta)
}

func resourceVolcengineCloudIdentityGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudIdentityGroup())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_identity_group %q, %s", d.Id(), err)
	}
	return err
}

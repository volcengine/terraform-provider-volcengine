package cloud_identity_user_provisioning

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudIdentityUserProvisioning can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_user_provisioning.default resource_id
```

*/

func ResourceVolcengineCloudIdentityUserProvisioning() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudIdentityUserProvisioningCreate,
		Read:   resourceVolcengineCloudIdentityUserProvisioningRead,
		Update: resourceVolcengineCloudIdentityUserProvisioningUpdate,
		Delete: resourceVolcengineCloudIdentityUserProvisioningDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"principal_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"User", "Group"}, false),
				Description:  "The principal type of the cloud identity user provisioning. Valid values: `User`, `Group`.",
			},
			"principal_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The principal id of the cloud identity user provisioning. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.",
			},
			"target_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target account id of the cloud identity user provisioning.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the cloud identity user provisioning.",
			},
			"identity_source_strategy": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Create", "Ignore"}, false),
				Description:  "The identity source strategy of the cloud identity user provisioning. Valid values: `Create`, `Ignore`.",
			},
			"duplication_strategy": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"KeepBoth", "Takeover"}, false),
				Description:  "The duplication strategy of the cloud identity user provisioning. Valid values: `KeepBoth`, `Takeover`.",
			},
			"duplication_suffix": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The duplication suffix of the cloud identity user provisioning. When the `duplication_strategy` is `KeepBoth`, this field must be specified.",
			},
			"deletion_strategy": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Keep", "Delete"}, false),
				Description:  "The deletion strategy of the cloud identity user provisioning. Valid values: `Keep`, `Delete`.",
			},
			"policy_name": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("principal_type").(string) != "User"
				},
				Description: "A list of policy name. Valid values: `AdministratorAccess`. This field is valid when the `principal_type` is `User`.",
			},

			// computed fields
			"provision_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the cloud identity user provisioning.",
			},
			"principal_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The principal name of the cloud identity user provisioning. When the `principal_type` is `User`, this field is specified to `UserName`. When the `principal_type` is `Group`, this field is specified to `GroupName`.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudIdentityUserProvisioningCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudIdentityUserProvisioning())
	if err != nil {
		return fmt.Errorf("error on creating cloud_identity_user_provisioning %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityUserProvisioningRead(d, meta)
}

func resourceVolcengineCloudIdentityUserProvisioningRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudIdentityUserProvisioning())
	if err != nil {
		return fmt.Errorf("error on reading cloud_identity_user_provisioning %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudIdentityUserProvisioningUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudIdentityUserProvisioning())
	if err != nil {
		return fmt.Errorf("error on updating cloud_identity_user_provisioning %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityUserProvisioningRead(d, meta)
}

func resourceVolcengineCloudIdentityUserProvisioningDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudIdentityUserProvisioning())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_identity_user_provisioning %q, %s", d.Id(), err)
	}
	return err
}

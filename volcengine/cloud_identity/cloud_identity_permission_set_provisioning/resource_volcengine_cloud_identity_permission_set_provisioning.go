package cloud_identity_permission_set_provisioning

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudIdentityPermissionSetProvisioning can be imported using the permission_set_id:target_id, e.g.
```
$ terraform import volcengine_cloud_identity_permission_set_provisioning.default resource_id
```

*/

func ResourceVolcengineCloudIdentityPermissionSetProvisioning() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudIdentityPermissionSetProvisioningCreate,
		Read:   resourceVolcengineCloudIdentityPermissionSetProvisioningRead,
		Update: resourceVolcengineCloudIdentityPermissionSetProvisioningUpdate,
		Delete: resourceVolcengineCloudIdentityPermissionSetProvisioningDelete,
		Importer: &schema.ResourceImporter{
			State: permissionSetProvisioningImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"permission_set_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the cloud identity permission set.",
			},
			"target_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target account id of the cloud identity permission set provisioning.",
			},
			"provisioning_status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Provisioned"}, false),
				Description: "The target provisioning status of the cloud identity permission set. This field must be specified as `Provisioned` in order to provision the updated permission set. \n" +
					"When deleting this resource, resource `volcengine_cloud_identity_permission_set_assignment` must be deleted first.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudIdentityPermissionSetProvisioningCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudIdentityPermissionSetProvisioning())
	if err != nil {
		return fmt.Errorf("error on creating cloud_identity_permission_set_provisioning %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityPermissionSetProvisioningRead(d, meta)
}

func resourceVolcengineCloudIdentityPermissionSetProvisioningRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudIdentityPermissionSetProvisioning())
	if err != nil {
		return fmt.Errorf("error on reading cloud_identity_permission_set_provisioning %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudIdentityPermissionSetProvisioningUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudIdentityPermissionSetProvisioning())
	if err != nil {
		return fmt.Errorf("error on updating cloud_identity_permission_set_provisioning %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityPermissionSetProvisioningRead(d, meta)
}

func resourceVolcengineCloudIdentityPermissionSetProvisioningDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetProvisioningService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudIdentityPermissionSetProvisioning())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_identity_permission_set_provisioning %q, %s", d.Id(), err)
	}
	return err
}

var permissionSetProvisioningImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("permission_set_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("target_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("provisioning_status", "Provisioned"); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

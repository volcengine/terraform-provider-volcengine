package iam_role

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam role can be imported using the id, e.g.
```
$ terraform import volcengine_iam_role.default TerraformTestRole
```

*/

func ResourceVolcengineIamRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamRoleCreate,
		Read:   resourceVolcengineIamRoleRead,
		Update: resourceVolcengineIamRoleUpdate,
		Delete: resourceVolcengineIamRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"trust_policy_document": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The trust policy document of the Role.",
			},
			"role_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Role.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The display name of the Role.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Role.",
			},
			"max_session_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max session duration of the Role.",
			},
			"trn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource name of the Role.",
			},
		},
	}
}

func resourceVolcengineIamRoleCreate(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(iamRoleService, d, ResourceVolcengineIamRole()); err != nil {
		return fmt.Errorf("error on creating iam role %q, %w", d.Id(), err)
	}
	return resourceVolcengineIamRoleRead(d, meta)
}

func resourceVolcengineIamRoleRead(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(iamRoleService, d, ResourceVolcengineIamRole()); err != nil {
		return fmt.Errorf("error on reading iam role %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineIamRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(iamRoleService, d, ResourceVolcengineIamRole()); err != nil {
		return fmt.Errorf("error on updating iam role %q, %w", d.Id(), err)
	}
	return resourceVolcengineIamRoleRead(d, meta)
}

func resourceVolcengineIamRoleDelete(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(iamRoleService, d, ResourceVolcengineIamRole()); err != nil {
		return fmt.Errorf("error on deleting iam role %q, %w", d.Id(), err)
	}
	return nil
}
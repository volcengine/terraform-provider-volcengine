package iam_role

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func ResourceVestackIamRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackIamRoleCreate,
		Read:   resourceVestackIamRoleRead,
		Update: resourceVestackIamRoleUpdate,
		Delete: resourceVestackIamRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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

func resourceVestackIamRoleCreate(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := iamRoleService.Dispatcher.Create(iamRoleService, d, ResourceVestackIamRole()); err != nil {
		return fmt.Errorf("error on creating iam role %q, %w", d.Id(), err)
	}
	return resourceVestackIamRoleRead(d, meta)
}

func resourceVestackIamRoleRead(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := iamRoleService.Dispatcher.Read(iamRoleService, d, ResourceVestackIamRole()); err != nil {
		return fmt.Errorf("error on reading iam role %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackIamRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := iamRoleService.Dispatcher.Update(iamRoleService, d, ResourceVestackIamRole()); err != nil {
		return fmt.Errorf("error on updating iam role %q, %w", d.Id(), err)
	}
	return resourceVestackIamRoleRead(d, meta)
}

func resourceVestackIamRoleDelete(d *schema.ResourceData, meta interface{}) error {
	iamRoleService := NewIamRoleService(meta.(*ve.SdkClient))
	if err := iamRoleService.Dispatcher.Delete(iamRoleService, d, ResourceVestackIamRole()); err != nil {
		return fmt.Errorf("error on deleting iam role %q, %w", d.Id(), err)
	}
	return nil
}

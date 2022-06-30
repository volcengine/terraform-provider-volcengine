package iam_policy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func ResourceVestackIamPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackIamPolicyCreate,
		Read:   resourceVestackIamPolicyRead,
		Update: resourceVestackIamPolicyUpdate,
		Delete: resourceVestackIamPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{},
	}
}

func resourceVestackIamPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewIamPolicyService(meta.(*ve.SdkClient))
	if err := iamPolicyService.Dispatcher.Create(iamPolicyService, d, ResourceVestackIamPolicy()); err != nil {
		return fmt.Errorf("error on creating iam policy %q, %w", d.Id(), err)
	}
	return resourceVestackIamPolicyRead(d, meta)
}

func resourceVestackIamPolicyRead(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewIamPolicyService(meta.(*ve.SdkClient))
	if err := iamPolicyService.Dispatcher.Read(iamPolicyService, d, ResourceVestackIamPolicy()); err != nil {
		return fmt.Errorf("error on reading iam policy %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackIamPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewIamPolicyService(meta.(*ve.SdkClient))
	if err := iamPolicyService.Dispatcher.Update(iamPolicyService, d, ResourceVestackIamPolicy()); err != nil {
		return fmt.Errorf("error on updating iam policy %q, %w", d.Id(), err)
	}
	return resourceVestackIamPolicyRead(d, meta)
}

func resourceVestackIamPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewIamPolicyService(meta.(*ve.SdkClient))
	if err := iamPolicyService.Dispatcher.Delete(iamPolicyService, d, ResourceVestackIamPolicy()); err != nil {
		return fmt.Errorf("error on deleting iam policy %q, %w", d.Id(), err)
	}
	return nil
}

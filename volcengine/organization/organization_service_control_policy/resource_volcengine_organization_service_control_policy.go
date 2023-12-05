package organization_service_control_policy

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Service Control Policy can be imported using the id, e.g.
```
$ terraform import volcengine_organization_service_control_policy.default 1000001
```

*/

func ResourceVolcengineServiceControlPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineServiceControlPolicyCreate,
		Read:   resourceVolcengineServiceControlPolicyRead,
		Update: resourceVolcengineServiceControlPolicyUpdate,
		Delete: resourceVolcengineServiceControlPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Policy.",
			},
			"statement": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The statement of the Policy.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					oldMap := make(map[string]interface{})
					newMap := make(map[string]interface{})

					_ = json.Unmarshal([]byte(old), &oldMap)
					_ = json.Unmarshal([]byte(new), &newMap)

					oldStr, _ := json.MarshalIndent(oldMap, "", "\t")
					newStr, _ := json.MarshalIndent(newMap, "", "\t")
					return string(oldStr) == string(newStr)
				},
			},
			"policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Policy.",
			},
			"policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the Policy.",
			},
			"create_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the Policy.",
			},
			"update_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the Policy.",
			},
		},
	}
}

func resourceVolcengineServiceControlPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(iamPolicyService, d, ResourceVolcengineServiceControlPolicy()); err != nil {
		return fmt.Errorf("error on creating policy %q, %w", d.Id(), err)
	}
	return resourceVolcengineServiceControlPolicyRead(d, meta)
}

func resourceVolcengineServiceControlPolicyRead(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(iamPolicyService, d, ResourceVolcengineServiceControlPolicy()); err != nil {
		return fmt.Errorf("error on reading policy %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineServiceControlPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(iamPolicyService, d, ResourceVolcengineServiceControlPolicy()); err != nil {
		return fmt.Errorf("error on updating policy %q, %w", d.Id(), err)
	}
	return resourceVolcengineServiceControlPolicyRead(d, meta)
}

func resourceVolcengineServiceControlPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(iamPolicyService, d, ResourceVolcengineServiceControlPolicy()); err != nil {
		return fmt.Errorf("error on deleting policy %q, %w", d.Id(), err)
	}
	return nil
}

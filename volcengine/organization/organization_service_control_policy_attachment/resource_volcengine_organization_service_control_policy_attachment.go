package organization_service_control_policy_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Service Control policy attachment can be imported using the id, e.g.
```
$ terraform import volcengine_organization_service_control_policy_attachment.default PolicyID:TargetID
```

*/

func ResourceVolcengineServiceControlPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineServiceControlPolicyAttachmentCreate,
		Read:   resourceVolcengineServiceControlPolicyAttachmentRead,
		Delete: resourceVolcengineServiceControlPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id is invalid")
				}
				if err := data.Set("policy_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("target_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of policy.",
			},
			"target_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of target.",
			},
			"target_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of target. Support Account or OU.",
			},
		},
	}
}

func resourceVolcengineServiceControlPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewServiceControlPolicyAttachmentService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineServiceControlPolicyAttachment()); err != nil {
		return fmt.Errorf("error on creating service_control_policy_attachment %q, %w", d.Id(), err)
	}
	return resourceVolcengineServiceControlPolicyAttachmentRead(d, meta)
}

func resourceVolcengineServiceControlPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	service := NewServiceControlPolicyAttachmentService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineServiceControlPolicyAttachment()); err != nil {
		return fmt.Errorf("error on reading service_control_policy_attachment %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineServiceControlPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewServiceControlPolicyAttachmentService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineServiceControlPolicyAttachment()); err != nil {
		return fmt.Errorf("error on deleting service_control_policy_attachment %q, %w", d.Id(), err)
	}
	return nil
}

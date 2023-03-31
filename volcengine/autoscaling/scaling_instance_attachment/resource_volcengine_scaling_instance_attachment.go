package scaling_instance_attachment

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Scaling instance attachment can be imported using the scaling_group_id and instance_id, e.g.
```
$ terraform import volcengine_scaling_instance_attachment.default scg-mizl7m1kqccg5smt1bdpijuj:i-l8u2ai4j0fauo6mrpgk8
```

*/

func ResourceVolcengineScalingInstanceAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingInstanceAttachmentCreate,
		Read:   resourceVolcengineScalingInstanceAttachmentRead,
		Update: resourceVolcengineScalingInstanceAttachmentUpdate,
		Delete: resourceVolcengineScalingInstanceAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: importScalingInstanceAttachment,
		},
		Schema: map[string]*schema.Schema{
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling group.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the instance.",
			},
			"entrusted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to host the instance to a scaling group. Default value is false.",
			},
			"delete_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Remove",
					"Detach",
				}, false),
				Description: "The type of delete activity. Valid values: Remove, Detach. Default value is Remove.",
			},
			"detach_option": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"both",
					"none",
				}, false),
				Description: "Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.",
			},
		},
	}
	return resource
}

func resourceVolcengineScalingInstanceAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachmentService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttachment())
	if err != nil {
		return fmt.Errorf("error on creating ScalingInstanceAttach %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingInstanceAttachmentRead(d, meta)
}

func resourceVolcengineScalingInstanceAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachmentService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttachment())
	if err != nil {
		return fmt.Errorf("error on reading ScalingInstanceAttach %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScalingInstanceAttachmentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachmentService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttachment())
	if err != nil {
		return fmt.Errorf("error on updating ScalingInstanceAttach %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingInstanceAttachmentRead(d, meta)
}

func resourceVolcengineScalingInstanceAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachmentService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingInstanceAttach %q, %s", d.Id(), err)
	}
	return err
}

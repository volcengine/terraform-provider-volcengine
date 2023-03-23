package scaling_group_enabler

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Scaling Group enabler can be imported using the scaling_group_id, e.g.
```
$ terraform import volcengine_scaling_group_enabler.default enable:scg-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineScalingGroupEnabler() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingGroupEnablerCreate,
		Read:   resourceVolcengineScalingGroupEnablerRead,
		Delete: resourceVolcengineScalingGroupEnablerDelete,
		Importer: &schema.ResourceImporter{
			State: scalingGroupEnablerImporter,
		},
		Schema: map[string]*schema.Schema{
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling group.",
			},
		},
	}
	return resource
}

func resourceVolcengineScalingGroupEnablerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingGroupEnablerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineScalingGroupEnabler())
	if err != nil {
		return fmt.Errorf("error on creating ScalingGroupEnable: %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingGroupEnablerRead(d, meta)
}

func resourceVolcengineScalingGroupEnablerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingGroupEnablerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineScalingGroupEnabler())
	if err != nil {
		return fmt.Errorf("error on reading ScalingGroupEnable: %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScalingGroupEnablerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingGroupEnablerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineScalingGroupEnabler())
	if err != nil {
		return fmt.Errorf("erron on deleting ScalingGroupEnabler: %q, %s", d.Id(), err)
	}
	return err
}

func scalingGroupEnablerImporter(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("scaling_group_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
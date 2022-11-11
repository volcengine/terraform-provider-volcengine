package scaling_group_enable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Scaling Group enable can be imported using the scaling_group_id, e.g.
```
$ terraform import volcengine_scaling_group_enable.default enable:scg-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineScalingGroupEnable() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingGroupEnableCreate,
		Read:   resourceVolcengineScalingGroupEnableRead,
		Delete: resourceVolcengineScalingGroupEnableDelete,
		Importer: &schema.ResourceImporter{
			State: scalingGroupEnableImporter,
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

func resourceVolcengineScalingGroupEnableCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingGroupEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineScalingGroupEnable())
	if err != nil {
		return fmt.Errorf("error on creating ScalingGroupEnable: %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingGroupEnableRead(d, meta)
}

func resourceVolcengineScalingGroupEnableRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingGroupEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineScalingGroupEnable())
	if err != nil {
		return fmt.Errorf("error on reading ScalingGroupEnable: %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScalingGroupEnableDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingGroupEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineScalingGroupEnable())
	if err != nil {
		return fmt.Errorf("erron on deleting ScalingGroupEnable: %q, %s", d.Id(), err)
	}
	return err
}

func scalingGroupEnableImporter(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("scaling_group_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

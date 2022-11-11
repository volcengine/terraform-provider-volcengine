package scaling_configuration_enable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Scaling Configuration enable can be imported using the scaling_group_id and scaling_configuration_id e.g.
```
$ terraform import volcengine_scaling_configuration_enable.default scg-mizl7m1kqccg5smt1bdpijuj:scc-ybrurj4uw6gh9zecj327
```

*/

func ResourceVolcengineScalingConfigurationEnable() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingConfigurationEnableCreate,
		Read:   resourceVolcengineScalingConfigurationEnableRead,
		Delete: resourceVolcengineScalingConfigurationEnableDelete,
		Importer: &schema.ResourceImporter{
			State: importScalingConfigurationEnable,
		},
		Schema: map[string]*schema.Schema{
			"scaling_configuration_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling configuration.",
			},
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

func resourceVolcengineScalingConfigurationEnableCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingConfigurationEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineScalingConfigurationEnable())
	if err != nil {
		return fmt.Errorf("error on creating ScalingConfigurationEnable: %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingConfigurationEnableRead(d, meta)
}

func resourceVolcengineScalingConfigurationEnableRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingConfigurationEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineScalingConfigurationEnable())
	if err != nil {
		return fmt.Errorf("error on reading ScalingConfigurationEnable: %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineScalingConfigurationEnableDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingConfigurationEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineScalingConfigurationEnable())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingConfigurationEnable: %q, %s", d.Id(), err)
	}
	return nil
}

func importScalingConfigurationEnable(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form ScalingGroupId:ScalingConfigurationId")
	}
	err = data.Set("scaling_group_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("scaling_configuration_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

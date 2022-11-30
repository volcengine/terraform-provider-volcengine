package scaling_configuration_attachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Scaling Configuration attachment can be imported using the scaling_configuration_id e.g.
```
$ terraform import volcengine_scaling_configuration_attachment.default enable:scc-ybrurj4uw6gh9zecj327
```

*/

func ResourceVolcengineScalingConfigurationAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingConfigurationAttachmentCreate,
		Read:   resourceVolcengineScalingConfigurationAttachmentRead,
		Delete: resourceVolcengineScalingConfigurationAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: importScalingConfigurationAttachment,
		},
		Schema: map[string]*schema.Schema{
			"scaling_configuration_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling configuration.",
			},
		},
	}
	return resource
}

func resourceVolcengineScalingConfigurationAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingConfigurationAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineScalingConfigurationAttachment())
	if err != nil {
		return fmt.Errorf("error on creating ScalingConfigurationEnable: %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingConfigurationAttachmentRead(d, meta)
}

func resourceVolcengineScalingConfigurationAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingConfigurationAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineScalingConfigurationAttachment())
	if err != nil {
		return fmt.Errorf("error on reading ScalingConfigurationEnable: %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineScalingConfigurationAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScalingConfigurationAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineScalingConfigurationAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingConfigurationEnable: %q, %s", d.Id(), err)
	}
	return nil
}

func importScalingConfigurationAttachment(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form enable:ScalingConfigurationId")
	}
	err = data.Set("scaling_configuration_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

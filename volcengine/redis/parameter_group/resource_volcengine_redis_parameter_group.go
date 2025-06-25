package parameter_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ParameterGroup can be imported using the id, e.g.
```
$ terraform import volcengine_parameter_group.default resource_id
```

*/

func ResourceVolcengineParameterGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineParameterGroupCreate,
		Read:   resourceVolcengineParameterGroupRead,
		Update: resourceVolcengineParameterGroupUpdate,
		Delete: resourceVolcengineParameterGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Parameter template name. The name needs to meet the following requirements simultaneously:" +
					" It cannot start with a number or a hyphen (-)." +
					" Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included." +
					" The length should be 2 to 64 characters.",
			},
			"engine_version": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The Redis database version adapted to the parameter template. The value range is as follows;" +
					" 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The remarks information of the parameter template should not exceed 200 characters in length.",
			},
			"param_values": {
				Type:     schema.TypeList,
				Required: true,
				Description: "The list of parameter information that needs to be included in the new parameter template. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The parameter names that need to be included in the parameter template.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The parameter values set for the corresponding parameters.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineParameterGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewParameterGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineParameterGroup())
	if err != nil {
		return fmt.Errorf("error on creating parameter_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineParameterGroupRead(d, meta)
}

func resourceVolcengineParameterGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewParameterGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineParameterGroup())
	if err != nil {
		return fmt.Errorf("error on reading parameter_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineParameterGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewParameterGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineParameterGroup())
	if err != nil {
		return fmt.Errorf("error on updating parameter_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineParameterGroupRead(d, meta)
}

func resourceVolcengineParameterGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewParameterGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineParameterGroup())
	if err != nil {
		return fmt.Errorf("error on deleting parameter_group %q, %s", d.Id(), err)
	}
	return err
}

package veecp_deployment

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpDeployment can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_deployment.default resource_id
```

*/

func ResourceVolcengineVeecpDeployment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpDeploymentCreate,
		Read:   resourceVolcengineVeecpDeploymentRead,
		Update: resourceVolcengineVeecpDeploymentUpdate,
		Delete: resourceVolcengineVeecpDeploymentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
		    // TODO: Add all your arguments and attributes.
			"replace_with_arguments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// TODO: See setting, getting, flattening, expanding examples below for this complex argument.
			"complex_argument": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_field_one": {
							Type:         schema.TypeString,
							Required:     true,
						},
						"sub_field_two": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVeecpDeploymentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on creating veecp_deployment %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpDeploymentRead(d, meta)
}

func resourceVolcengineVeecpDeploymentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on reading veecp_deployment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpDeploymentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on updating veecp_deployment %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpDeploymentRead(d, meta)
}

func resourceVolcengineVeecpDeploymentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_deployment %q, %s", d.Id(), err)
	}
	return err
}

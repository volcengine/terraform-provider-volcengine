package ecs_deployment_set

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ECS deployment set can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_deployment_set.default i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineEcsDeploymentSet() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsDeploymentSetCreate,
		Read:   resourceVolcengineEcsDeploymentSetRead,
		Update: resourceVolcengineEcsDeploymentSetUpdate,
		Delete: resourceVolcengineEcsDeploymentSetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"deployment_set_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of ECS DeploymentSet.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of ECS DeploymentSet.",
			},
			"granularity": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"switch",
					"host",
					"rack",
				}, false),
				Default:     "host",
				Description: "The granularity of ECS DeploymentSet.Valid values: switch, host, rack,Default is host.",
			},
			"strategy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Availability",
				}, false),
				Default:     "Availability",
				Description: "The strategy of ECS DeploymentSet.Valid values: Availability.Default is Availability.",
			},
			"deployment_set_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of ECS DeploymentSet.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsDeploymentSetCreate(d *schema.ResourceData, meta interface{}) (err error) {
	deploymentSetService := NewEcsDeploymentSetService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(deploymentSetService, d, ResourceVolcengineEcsDeploymentSet())
	if err != nil {
		return fmt.Errorf("error on creating ecs deployment set  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsDeploymentSetRead(d, meta)
}

func resourceVolcengineEcsDeploymentSetRead(d *schema.ResourceData, meta interface{}) (err error) {
	deploymentSetService := NewEcsDeploymentSetService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(deploymentSetService, d, ResourceVolcengineEcsDeploymentSet())
	if err != nil {
		return fmt.Errorf("error on reading ecs deployment set %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsDeploymentSetUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	deploymentSetService := NewEcsDeploymentSetService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(deploymentSetService, d, ResourceVolcengineEcsDeploymentSet())
	if err != nil {
		return fmt.Errorf("error on updating ecs deployment set  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsDeploymentSetRead(d, meta)
}

func resourceVolcengineEcsDeploymentSetDelete(d *schema.ResourceData, meta interface{}) (err error) {
	deploymentSetService := NewEcsDeploymentSetService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(deploymentSetService, d, ResourceVolcengineEcsDeploymentSet())
	if err != nil {
		return fmt.Errorf("error on deleting ecs deployment set %q, %s", d.Id(), err)
	}
	return err
}

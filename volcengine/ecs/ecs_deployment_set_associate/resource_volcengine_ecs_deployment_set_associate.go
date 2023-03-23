package ecs_deployment_set_associate

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ECS deployment set associate can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_deployment_set_associate.default dps-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineEcsDeploymentSetAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsDeploymentSetAssociateCreate,
		Read:   resourceVolcengineEcsDeploymentSetAssociateRead,
		Delete: resourceVolcengineEcsDeploymentSetAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("deployment_set_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("instance_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"deployment_set_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of ECS DeploymentSet Associate.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of ECS Instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsDeploymentSetAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	deploymentSetService := NewEcsDeploymentSetAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(deploymentSetService, d, ResourceVolcengineEcsDeploymentSetAssociate())
	if err != nil {
		return fmt.Errorf("error on creating ecs deployment set Associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsDeploymentSetAssociateRead(d, meta)
}

func resourceVolcengineEcsDeploymentSetAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	deploymentSetService := NewEcsDeploymentSetAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(deploymentSetService, d, ResourceVolcengineEcsDeploymentSetAssociate())
	if err != nil {
		return fmt.Errorf("error on reading ecs deployment set Associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsDeploymentSetAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	deploymentSetService := NewEcsDeploymentSetAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(deploymentSetService, d, ResourceVolcengineEcsDeploymentSetAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting ecs deployment set Associate %q, %s", d.Id(), err)
	}
	return err
}

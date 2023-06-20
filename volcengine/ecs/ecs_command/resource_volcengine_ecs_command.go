package ecs_command

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EcsCommand can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_command.default cmd-ychkepkhtim0tr3bcsw1
```

*/

func ResourceVolcengineEcsCommand() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsCommandCreate,
		Read:   resourceVolcengineEcsCommandRead,
		Update: resourceVolcengineEcsCommandUpdate,
		Delete: resourceVolcengineEcsCommandDelete,
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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the ecs command.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the ecs command.",
			},
			"command_content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The base64 encoded content of the ecs command.",
			},
			"working_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The working directory of the ecs command.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The user name of the ecs command.",
			},
			"timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(10, 600),
				Description:  "The timeout of the ecs command. Valid value range: 10-600.",
			},

			"invocation_times": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The invocation times of the ecs command. Public commands do not display the invocation times.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the ecs command.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the ecs command.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsCommandCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsCommandService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineEcsCommand())
	if err != nil {
		return fmt.Errorf("error on creating ecs command %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsCommandRead(d, meta)
}

func resourceVolcengineEcsCommandRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsCommandService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineEcsCommand())
	if err != nil {
		return fmt.Errorf("error on reading ecs command %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsCommandUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsCommandService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineEcsCommand())
	if err != nil {
		return fmt.Errorf("error on updating ecs command %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsCommandRead(d, meta)
}

func resourceVolcengineEcsCommandDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsCommandService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineEcsCommand())
	if err != nil {
		return fmt.Errorf("error on deleting ecs command %q, %s", d.Id(), err)
	}
	return err
}

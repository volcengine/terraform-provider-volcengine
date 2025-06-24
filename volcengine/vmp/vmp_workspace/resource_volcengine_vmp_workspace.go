package vmp_workspace

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Workspace can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_workspace.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

*/

func ResourceVolcengineVmpWorkspace() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpWorkspaceCreate,
		Read:   resourceVolcengineVmpWorkspaceRead,
		Update: resourceVolcengineVmpWorkspaceUpdate,
		Delete: resourceVolcengineVmpWorkspaceDelete,
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
				Description: "The name of the workspace.",
			},
			"instance_type_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance type id of the workspace.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the workspace.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username of the workspace.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password of the workspace.",
			},
			"delete_protection_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether enable delete protection.",
			},
			"prometheus_write_intranet_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The prometheus write intranet endpoint.",
			},
			"prometheus_query_intranet_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The prometheus query intranet endpoint.",
			},
			"overdue_reclaim_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The overdue reclaim time.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of workspace.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of workspace.",
			},
		},
	}
	return resource
}

func resourceVolcengineVmpWorkspaceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpWorkspace())
	if err != nil {
		return fmt.Errorf("error on creating workspace %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpWorkspaceRead(d, meta)
}

func resourceVolcengineVmpWorkspaceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpWorkspace())
	if err != nil {
		return fmt.Errorf("error on reading workspace %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpWorkspaceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpWorkspace())
	if err != nil {
		return fmt.Errorf("error on updating workspace %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpWorkspaceRead(d, meta)
}

func resourceVolcengineVmpWorkspaceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpWorkspace())
	if err != nil {
		return fmt.Errorf("error on deleting workspace %q, %s", d.Id(), err)
	}
	return err
}

package vmp_integration_task

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VMP Integration Task can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_integration_task.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

*/

func ResourceVolcengineVmpIntegrationTask() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpIntegrationTaskCreate,
		Read:   resourceVolcengineVmpIntegrationTaskRead,
		Update: resourceVolcengineVmpIntegrationTaskUpdate,
		Delete: resourceVolcengineVmpIntegrationTaskDelete,
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
				Description: "The name of the integration task. Length: 1-40 characters. Supports Chinese, English, numbers, and underscores.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the integration task. For example, `CloudMonitor` indicates a cloud monitoring integration task.",
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The deployment environment. Valid values: `Vke` or `Managed`.",
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The workspace ID. Required when Environment is `Managed`.",
			},
			"params": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parameters of the integration task. Must be a JSON-escaped string.",
			},
			"vke_cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the VKE cluster. Required when Environment is `Vke`.",
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the integration task. Valid values: `Creating`, `Updating`, `Active`, `Error`, `Deleting`.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the integration task.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the integration task.",
			},
		},
	}
	return resource
}

func resourceVolcengineVmpIntegrationTaskCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpIntegrationTask())
	if err != nil {
		return fmt.Errorf("error on creating integration task %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpIntegrationTaskRead(d, meta)
}

func resourceVolcengineVmpIntegrationTaskRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpIntegrationTask())
	if err != nil {
		return fmt.Errorf("error on reading integration task %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpIntegrationTaskUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpIntegrationTask())
	if err != nil {
		return fmt.Errorf("error on updating integration task %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpIntegrationTaskRead(d, meta)
}

func resourceVolcengineVmpIntegrationTaskDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpIntegrationTask())
	if err != nil {
		return fmt.Errorf("error on deleting integration task %q, %s", d.Id(), err)
	}
	return err
}

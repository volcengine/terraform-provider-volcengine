package workspace

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
$ terraform import volcengine_bioos_workspace.default *****
```

*/

func ResourceVolcengineBioosWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineBioosWorkspaceCreate,
		Read:   resourceVolcengineBioosWorkspaceRead,
		Delete: resourceVolcengineBioosWorkspaceDelete,
		Update: resourceVolcengineBioosWorkspaceUpdate,
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
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of the workspace.",
			},
			"s3_bucket": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "s3 bucket address.",
			},
			"cover_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cover path (relative path in tos bucket).",
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the workspace.",
			},
			"updated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the update complete.",
			},
		},
	}
}

func resourceVolcengineBioosWorkspaceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosWorkspaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineBioosWorkspace())
	if err != nil {
		return fmt.Errorf("error on creating volcengine bioos Workspace: %q, %w", d.Id(), err)
	}
	return resourceVolcengineBioosWorkspaceRead(d, meta)
}

func resourceVolcengineBioosWorkspaceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosWorkspaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineBioosWorkspace())
	if err != nil {
		return fmt.Errorf("error on reading volcengine bioos Workspace: %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineBioosWorkspaceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosWorkspaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineBioosWorkspace())
	if err != nil {
		return fmt.Errorf("error on updating volcengine bioos Workspace: %q, %w", d.Id(), err)
	}
	return resourceVolcengineBioosWorkspaceRead(d, meta)
}

func resourceVolcengineBioosWorkspaceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosWorkspaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineBioosWorkspace())
	if err != nil {
		return fmt.Errorf("error on deleting volcengine bioos Workspace: %q, %w", d.Id(), err)
	}
	return nil
}

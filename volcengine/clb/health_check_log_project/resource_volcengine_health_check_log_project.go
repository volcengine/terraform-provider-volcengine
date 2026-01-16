package health_check_log_project

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
HealthCheckLogProject can be imported using the id, e.g.
```
$ terraform import volcengine_health_check_log_project.default log_project_id(e.g. b8e16846-fb31-4a2c-a8c1-171434d41d15)
```

*/

func ResourceVolcengineHealthCheckLogProject() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineHealthCheckLogProjectCreate,
		Read:   resourceVolcengineHealthCheckLogProjectRead,
		Delete: resourceVolcengineHealthCheckLogProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"log_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the health check log project.",
			},
		},
	}
	return resource
}

func resourceVolcengineHealthCheckLogProjectCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewHealthCheckLogProjectService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineHealthCheckLogProject())
	if err != nil {
		return fmt.Errorf("error on creating health_check_log_project %q, %s", d.Id(), err)
	}
	return resourceVolcengineHealthCheckLogProjectRead(d, meta)
}

func resourceVolcengineHealthCheckLogProjectRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewHealthCheckLogProjectService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineHealthCheckLogProject())
	if err != nil {
		return fmt.Errorf("error on reading health_check_log_project %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineHealthCheckLogProjectDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewHealthCheckLogProjectService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineHealthCheckLogProject())
	if err != nil {
		return fmt.Errorf("error on deleting health_check_log_project %q, %s", d.Id(), err)
	}
	return err
}

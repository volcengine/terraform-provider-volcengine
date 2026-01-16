package alb_health_log

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The AlbHealthLog is not support import.


*/

func ResourceVolcengineAlbHealthLog() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbHealthLogCreate,
		Read:   resourceVolcengineAlbHealthLogRead,
		Delete: resourceVolcengineAlbHealthLogDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the LoadBalancer.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the Topic.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The project ID of the Topic.",
			},
		},
	}
	return resource
}

func resourceVolcengineAlbHealthLogCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbHealthLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbHealthLog())
	if err != nil {
		return fmt.Errorf("error on creating alb_health_log %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbHealthLogRead(d, meta)
}

func resourceVolcengineAlbHealthLogRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbHealthLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbHealthLog())
	if err != nil {
		return fmt.Errorf("error on reading alb_health_log %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbHealthLogDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbHealthLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbHealthLog())
	if err != nil {
		return fmt.Errorf("error on deleting alb_health_log %q, %s", d.Id(), err)
	}
	return err
}

package alb_access_log

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The AlbAccessLog is not support import.

*/

func ResourceVolcengineAlbAccessLog() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbAccessLogCreate,
		Read:   resourceVolcengineAlbAccessLogRead,
		Delete: resourceVolcengineAlbAccessLogDelete,
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
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the TOS bucket for storing access logs.",
			},
		},
	}
	return resource
}

func resourceVolcengineAlbAccessLogCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbAccessLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbAccessLog())
	if err != nil {
		return fmt.Errorf("error on creating alb_access_log %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbAccessLogRead(d, meta)
}

func resourceVolcengineAlbAccessLogRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbAccessLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbAccessLog())
	if err != nil {
		return fmt.Errorf("error on reading alb_access_log %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbAccessLogDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbAccessLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbAccessLog())
	if err != nil {
		return fmt.Errorf("error on deleting alb_access_log %q, %s", d.Id(), err)
	}
	return err
}

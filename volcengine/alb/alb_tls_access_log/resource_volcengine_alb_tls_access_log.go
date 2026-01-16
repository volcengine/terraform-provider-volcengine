package alb_tls_access_log

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The AlbTlsAccessLog is not support import.

*/

func ResourceVolcengineAlbTlsAccessLog() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbTlsAccessLogCreate,
		Read:   resourceVolcengineAlbTlsAccessLogRead,
		Delete: resourceVolcengineAlbTlsAccessLogDelete,
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

func resourceVolcengineAlbTlsAccessLogCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbTlsAccessLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbTlsAccessLog())
	if err != nil {
		return fmt.Errorf("error on creating alb_tls_access_log %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbTlsAccessLogRead(d, meta)
}

func resourceVolcengineAlbTlsAccessLogRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbTlsAccessLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbTlsAccessLog())
	if err != nil {
		return fmt.Errorf("error on reading alb_tls_access_log %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbTlsAccessLogDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbTlsAccessLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbTlsAccessLog())
	if err != nil {
		return fmt.Errorf("error on deleting alb_tls_access_log %q, %s", d.Id(), err)
	}
	return err
}

package alb_listener

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AlbListener can be imported using the id, e.g.
```
$ terraform import volcengine_alb_listener.default lsn-273yv0mhs5xj47fap8sehiiso
```

*/

func ResourceVolcengineAlbListener() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbListenerCreate,
		Read:   resourceVolcengineAlbListenerRead,
		Update: resourceVolcengineAlbListenerUpdate,
		Delete: resourceVolcengineAlbListenerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Id of the load balancer.",
			},
			"listener_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the Listener.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The protocol of the Listener. Optional choice contains `HTTP`, `HTTPS`.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The port receiving request of the Listener, the value range in 1~65535.",
			},
			"enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "on",
				Description: "The enable status of the Listener. Optional choice contains `on`, `off`. Default is `on`.",
			},
			"certificate_source": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "alb",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) == "HTTP"
				},
				Description: "The source of the certificate. Valid values: `alb`, `cert_center`. Default is `alb`.",
			},
			"cert_center_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) == "HTTP"
				},
				Description: "The certificate id associated with the listener. Source is `cert_center`.",
			},
			"certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) == "HTTP"
				},
				Description: "The certificate id associated with the listener. Source is `alb`.",
			},
			"ca_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) == "HTTP"
				},
				Description: "The CA certificate id associated with the listener.",
			},
			"server_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server group id associated with the listener.",
			},
			"enable_http2": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) == "HTTP"
				},
				Description: "The HTTP2 feature switch,valid value is on or off. Default is `off`.",
			},
			"enable_quic": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) == "HTTP"
				},
				Description: "The QUIC feature switch,valid value is on or off. Default is `off`.",
			},
			"acl_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "off",
				Description: "The enable status of Acl. Optional choice contains `on`, `off`. Default is `off`.",
			},
			"acl_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "The type of the Acl. Optional choice contains `white`, `black`. " +
					"When the AclStatus parameter is configured as on, AclType and AclIds.N are required.",
				ValidateFunc: validation.StringInSlice([]string{"white", "black"}, false),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("acl_status").(string) == "off"
				},
			},
			"acl_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: "The id list of the Acl. " +
					"When the AclStatus parameter is configured as on, AclType and AclIds.N are required.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("acl_status").(string) == "off"
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Listener.",
			},
			"customized_cfg_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Personalized configuration ID, with a value of \" \" when not bound.",
			},
			"listener_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the Listener.",
			},
		},
	}
	return resource
}

func resourceVolcengineAlbListenerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbListener())
	if err != nil {
		return fmt.Errorf("error on creating alb_listener %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbListenerRead(d, meta)
}

func resourceVolcengineAlbListenerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbListener())
	if err != nil {
		return fmt.Errorf("error on reading alb_listener %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbListenerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlbListener())
	if err != nil {
		return fmt.Errorf("error on updating alb_listener %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbListenerRead(d, meta)
}

func resourceVolcengineAlbListenerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbListener())
	if err != nil {
		return fmt.Errorf("error on deleting alb_listener %q, %s", d.Id(), err)
	}
	return err
}

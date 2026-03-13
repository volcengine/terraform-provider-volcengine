package nlb_listener

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*
Import
NlbListener can be imported using the listener ID, e.g.
```
$ terraform import volcengine_nlb_listener.foo lsn-2d6g5cxxx
```
*/
func ResourceVolcengineNlbListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNlbListenerCreate,
		Read:   resourceVolcengineNlbListenerRead,
		Update: resourceVolcengineNlbListenerUpdate,
		Delete: resourceVolcengineNlbListenerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the NLB.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The protocol of the listener. Valid values: `TCP`, `UDP`.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The port of the listener. Range: 0-65535.",
			},
			"server_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the server group.",
			},
			"listener_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the listener.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the listener.",
			},
			"connection_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     900,
				Description: "The connection timeout of the listener.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable the listener.",
			},
			"start_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The start port of the listener. Range: 0-65535.",
			},
			"end_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The end port of the listener. Range: 0-65535.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the listener.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the listener.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the listener.",
			},
			"tags": ve.TagsSchema(),
		},
	}
}

func resourceVolcengineNlbListenerCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbListenerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNlbListener())
	if err != nil {
		return fmt.Errorf("error on creating nlb listener %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbListenerRead(d, meta)
}

func resourceVolcengineNlbListenerRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbListenerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNlbListener())
	if err != nil {
		return fmt.Errorf("error on reading nlb listener %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNlbListenerUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbListenerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNlbListener())
	if err != nil {
		return fmt.Errorf("error on updating nlb listener %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbListenerRead(d, meta)
}

func resourceVolcengineNlbListenerDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbListenerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNlbListener())
	if err != nil {
		return fmt.Errorf("error on deleting nlb listener %q, %w", d.Id(), err)
	}
	return err
}

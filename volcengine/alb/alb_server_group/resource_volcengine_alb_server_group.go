package alb_server_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AlbServerGroup can be imported using the id, e.g.
```
$ terraform import volcengine_alb_server_group.default resource_id
```

*/

func ResourceVolcengineAlbServerGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbServerGroupCreate,
		Read:   resourceVolcengineAlbServerGroupRead,
		Update: resourceVolcengineAlbServerGroupUpdate,
		Delete: resourceVolcengineAlbServerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The vpc id of the Alb server group.",
			},
			"server_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the Alb server group.",
			},
			"server_group_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "instance",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"instance", "ip"}, false),
				Description:  "The type of the Alb server group. Valid values: `instance`, `ip`. Default is `instance`.",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "HTTP",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"HTTP", "HTTPS", "gRPC"}, false),
				Description:  "The backend protocol of the Alb server group. Valid values: `HTTP`, `HTTPS`, `gRPC`. Default is `HTTP`.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the Alb server group.",
			},
			"scheduler": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "wrr",
				ValidateFunc: validation.StringInSlice([]string{"wrr", "wlc", "sh"}, false),
				Description:  "The scheduling algorithm of the Alb server group. Valid values: `wrr`, `wlc`, `sh`.",
			},
			"cross_zone_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
				Description:  "Whether to enable cross-zone load balancing for the server group. Valid values: `on`, `off`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the Alb server group.",
			},
			"ip_address_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"IPv4"}, false),
				Description:  "The ip address type of the server group.",
			},
			"tags": ve.TagsSchema(),
			"health_check": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The health check config of the Alb server group. The enable status of health check function defaults to `on`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "on",
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Description:  "The enable status of health check function. Valid values: `on`, `off`. Default is `on`.",
						},
						"interval": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          2,
							Description:      "The interval executing health check. Unit: second. Valid value range in 1~300. Default is 2.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"timeout": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          2,
							Description:      "The response timeout of health check. Unit: second. Valid value range in 1~60. Default is 2.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"healthy_threshold": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          3,
							Description:      "The healthy threshold of health check. Valid value range in 2~10. Default is 3.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"unhealthy_threshold": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          3,
							Description:      "The unhealthy threshold of health check. Valid value range in 2~10. Default is 3.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"method": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "HEAD",
							ValidateFunc:     validation.StringInSlice([]string{"GET", "HEAD"}, false),
							Description:      "The method of health check. Valid values: `GET` or `HEAD`. Default is `HEAD`.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Description:      "The domain of health check.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"uri": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Description:      "The uri of health check.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"http_code": {
							Type:     schema.TypeString,
							Optional: true,
							Default:          "http_2xx,http_3xx",
							Description:      "The normal http status code of health check, the value can be `http_2xx`, `http_3xx`, `http_4xx` or `http_5xx`. Default is `http_2xx,http_3xx`.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"protocol": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "HTTP",
							ValidateFunc:     validation.StringInSlice([]string{"HTTP", "TCP"}, false),
							Description:      "The protocol of health check. Valid values: `HTTP`, `TCP`. Default is `HTTP`.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"port": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          0,
							Description:      "The port of health check. When the value is 0, it means use the backend server port for health check. Valid value range in 0~65535.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
						"http_version": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "HTTP1.0",
							ValidateFunc:     validation.StringInSlice([]string{"HTTP1.0", "HTTP1.1"}, false),
							Description:      "The http version of health check. Valid values: `HTTP1.0`, `HTTP1.1`. Default is `HTTP1.0`.",
							DiffSuppressFunc: albHealthCheckDiffSuppress,
						},
					},
				},
			},
			"sticky_session_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The sticky session config of the Alb server group. The enable status of sticky session function defaults to `off`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sticky_session_enabled": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Description:  "The enable status of sticky session. Valid values: `on`, `off`. Default is `off`.",
						},
						"sticky_session_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "insert",
							ValidateFunc:     validation.StringInSlice([]string{"insert", "server"}, false),
							Description:      "The cookie handle type of the sticky session. Valid values: `insert`, `server`. Default is `insert`. This field is required when the value of the `sticky_session_enabled` is `on`.",
							DiffSuppressFunc: albStickySessionTypeDiffSuppress,
						},
						"cookie": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Description:      "The cookie name of the sticky session. This field is required when the value of the `sticky_session_type` is `server`.",
							DiffSuppressFunc: albStickySessionCookieDiffSuppress,
						},
						"cookie_timeout": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          1000,
							Description:      "The cookie timeout of the sticky session. Unit: second. Valid value range in 1~86400. Default is 1000. This field is required when the value of the `sticky_session_type` is `insert`.",
							DiffSuppressFunc: albStickySessionCookieTimeoutDiffSuppress,
						},
					},
				},
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the Alb server group.",
			},
			"server_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server count of the Alb server group.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the Alb server group.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the Alb server group.",
			},
			"listeners": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The listener information of the Alb server group.",
			},
		},
	}
	return resource
}

func resourceVolcengineAlbServerGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbServerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on creating alb_server_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbServerGroupRead(d, meta)
}

func resourceVolcengineAlbServerGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbServerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on reading alb_server_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbServerGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbServerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on updating alb_server_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbServerGroupRead(d, meta)
}

func resourceVolcengineAlbServerGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbServerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on deleting alb_server_group %q, %s", d.Id(), err)
	}
	return err
}

func albHealthCheckDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	healthCheckEnabled := d.Get("health_check.0.enabled").(string)
	return healthCheckEnabled == "off"
}

func albStickySessionTypeDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	stickySessionEnabled := d.Get("sticky_session_config.0.sticky_session_enabled").(string)
	return stickySessionEnabled == "off"
}

func albStickySessionCookieDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	stickySessionTypeDiff := albStickySessionTypeDiffSuppress(k, old, new, d)
	stickySessionType := d.Get("sticky_session_config.0.sticky_session_type").(string)
	return stickySessionTypeDiff || stickySessionType == "insert"
}

func albStickySessionCookieTimeoutDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	stickySessionTypeDiff := albStickySessionTypeDiffSuppress(k, old, new, d)
	stickySessionType := d.Get("sticky_session_config.0.sticky_session_type").(string)
	return stickySessionTypeDiff || stickySessionType == "server"
}

package listener

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Listener can be imported using the id, e.g.
```
$ terraform import volcengine_listener.default lsn-273yv0mhs5xj47fap8sehiiso
```

*/

func ResourceVolcengineListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineListenerCreate,
		Read:   resourceVolcengineListenerRead,
		Update: resourceVolcengineListenerUpdate,
		Delete: resourceVolcengineListenerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the Listener.",
			},
			"load_balancer_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region of the request.",
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
				Description: "The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The port receiving request of the Listener, the value range in 1~65535.",
			},
			"scheduler": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.",
			},
			"enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The enable status of the Listener. Optional choice contains `on`, `off`.",
			},
			"established_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The connection timeout of the Listener.",
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The certificate id associated with the listener.",
			},
			"server_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server group id associated with the listener.",
			},
			"acl_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The enable status of Acl. Optional choice contains `on`, `off`.",
			},
			"acl_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The type of the Acl. Optional choice contains `white`, `black`.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("acl_status").(string) == "off"
				},
			},
			"acl_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The id list of the Acl.",
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
			"bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     -1,
				Description: "The bandwidth of the Listener. Unit: Mbps. Default is -1, indicating that the Listener does not specify a speed limit.",
			},
			"proxy_protocol_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) != "TCP" && d.Get("protocol").(string) != "UDP"
				},
				Description: "Whether to enable proxy protocol. Valid values: `off`, `standard`. Default is `off`.\n" +
					"This filed is valid only when the value of field `protocol` is `TCP` or `UDP`.",
			},
			"persistence_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					persistenceType := d.Get("persistence_type").(string)
					if persistenceType == "off" {
						return false
					}
					protocol := d.Get("protocol").(string)
					if persistenceType == "source_ip" && (protocol == "TCP" || protocol == "UDP") {
						return false
					}
					scheduler := d.Get("scheduler").(string)
					if (persistenceType == "insert" || persistenceType == "server") &&
						(protocol == "HTTP" || protocol == "HTTPS") && scheduler == "wrr" {
						return false
					}
					return true
				},
				Description: "The persistence type of the Listener. Valid values: `off`, `source_ip`, `insert`, `server`. Default is `off`.\n" +
					"`source_ip`: Represents the source IP address, only effective for TCP/UDP protocols. " +
					"`insert`: means implanting a cookie, only effective for HTTP/HTTPS protocol and when the scheduler is `wrr`. " +
					"`server`: Indicates rewriting cookies, only effective for HTTP/HTTPS protocols and when the scheduler is `wrr`.",
			},
			"persistence_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("persistence_type").(string) != "source_ip" && d.Get("persistence_type").(string) != "insert"
				},
				Description: "The persistence timeout of the Listener. Unit: second. Default is `1000`. When PersistenceType is configured as source_ip, " +
					"the value range is 1-3600. When PersistenceType is configured as insert, the value range is 1-86400. " +
					"This filed is valid only when the value of field `persistence_type` is `source_ip` or `insert`.",
			},
			"cookie": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("persistence_type").(string) != "server"
				},
				Description: "The name of the cookie for session persistence configured on the backend server. " +
					"When PersistenceType is configured as `server`, " +
					"this parameter is required. When PersistenceType is configured as any other value, " +
					"this parameter is not effective.",
			},
			"connection_drain_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "off",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("protocol").(string) != "TCP" && d.Get("protocol").(string) != "UDP"
				},
				Description: "Whether to enable connection drain of the Listener. Valid values: `off`, `on`. Default is `off`.\n" +
					"This filed is valid only when the value of field `protocol` is `TCP` or `UDP`.",
			},
			"connection_drain_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return (d.Get("protocol").(string) != "TCP" && d.Get("protocol").(string) != "UDP") || d.Get("connection_drain_enabled").(string) != "on"
				},
				Description: "The connection drain timeout of the Listener. Valid value range is `1-900`.\n" +
					"This filed is required when the value of field `connection_drain_enabled` is `on`.",
			},
			"health_check": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The config of health check.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The enable status of health check function. Optional choice contains `on`, `off`.",
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
						},
						"interval": {
							Type:             schema.TypeInt,
							Optional:         true,
							Description:      "The interval executing health check, default 2, range in 1~300.",
							DiffSuppressFunc: HealthCheckFieldDiffSuppress,
						},
						"timeout": {
							Type:             schema.TypeInt,
							Optional:         true,
							Description:      "The response timeout of health check, default 2, range in 1~60..",
							DiffSuppressFunc: HealthCheckFieldDiffSuppress,
						},
						"healthy_threshold": {
							Type:             schema.TypeInt,
							Optional:         true,
							Description:      "The healthy threshold of health check, default 3, range in 2~10.",
							DiffSuppressFunc: HealthCheckFieldDiffSuppress,
						},
						"un_healthy_threshold": {
							Type:             schema.TypeInt,
							Optional:         true,
							Description:      "The unhealthy threshold of health check, default 3, range in 2~10.",
							DiffSuppressFunc: HealthCheckFieldDiffSuppress,
						},
						"method": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The method of health check, the value can be `GET` or `HEAD`.",
							DiffSuppressFunc: HealthCheckHTTPOnlyFieldDiffSuppress,
						},
						"domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The domain of health check.",
							DiffSuppressFunc: HealthCheckHTTPOnlyFieldDiffSuppress,
						},
						"uri": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The uri of health check.",
							DiffSuppressFunc: HealthCheckHTTPOnlyFieldDiffSuppress,
						},
						"http_code": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The normal http status code of health check, the value can be `http_2xx` or `http_3xx` or `http_4xx` or `http_5xx`.",
							DiffSuppressFunc: HealthCheckHTTPOnlyFieldDiffSuppress,
						},
						"udp_request": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The UDP request of health check. This field must be specified simultaneously with field `udp_expect`.",
							DiffSuppressFunc: HealthCheckUDPOnlyFieldDiffSuppress,
						},
						"udp_expect": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The UDP expect of health check. This field must be specified simultaneously with field `udp_request`.",
							DiffSuppressFunc: HealthCheckUDPOnlyFieldDiffSuppress,
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineListenerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	listenerService := NewListenerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(listenerService, d, ResourceVolcengineListener())
	if err != nil {
		return fmt.Errorf("error on creating listener  %q, %w", d.Id(), err)
	}
	return resourceVolcengineListenerRead(d, meta)
}

func resourceVolcengineListenerRead(d *schema.ResourceData, meta interface{}) (err error) {
	listenerService := NewListenerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(listenerService, d, ResourceVolcengineListener())
	if err != nil {
		return fmt.Errorf("error on reading listener %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineListenerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	listenerService := NewListenerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(listenerService, d, ResourceVolcengineListener())
	if err != nil {
		return fmt.Errorf("error on updating listener  %q, %w", d.Id(), err)
	}
	return resourceVolcengineListenerRead(d, meta)
}

func resourceVolcengineListenerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	listenerService := NewListenerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(listenerService, d, ResourceVolcengineListener())
	if err != nil {
		return fmt.Errorf("error on deleting listener %q, %w", d.Id(), err)
	}
	return err
}

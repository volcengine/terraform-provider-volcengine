package nlb_server_group

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*
Import
NlbServerGroup can be imported using the NLB server group ID, e.g.
```
$ terraform import volcengine_nlb_server_group.foo rsp-2d6g5cxxx
```
*/
func ResourceVolcengineNlbServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNlbServerGroupCreate,
		Read:   resourceVolcengineNlbServerGroupRead,
		Update: resourceVolcengineNlbServerGroupUpdate,
		Delete: resourceVolcengineNlbServerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"server_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the server group.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the VPC to which the server group belongs.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "instance",
				Description: "The type of the server group. Valid values: `instance` (default), `ip`.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "TCP",
				Description: "The protocol of the server group. Valid values: `TCP` (default), `UDP`, `TCP_SSL`.",
			},
			"scheduler": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "wrr",
				Description: "The scheduling algorithm. Valid values: `wrr` (default), `wlc`, `sh`.",
			},
			"ip_address_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "ipv4",
				Description: "The ip address version of the server group. Valid values: `ipv4` (default), `ipv6`.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the server group.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the server group.",
			},
			"any_port_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable full port forwarding. Default is false.",
			},
			"connection_drain_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable connection graceful interruption. Default is false.",
			},
			"connection_drain_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Connection graceful interruption timeout. Unit: second. Value range: 0 ~ 900. Default is 0.",
			},
			"preserve_client_ip_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable source address retention. Default is true.",
			},
			"session_persistence_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable session persistence. Default is false.",
			},
			"session_persistence_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1000,
				Description: "Session persistence timeout. Unit: second. Value range: 1 ~ 3600. Default is 1000.",
			},
			"proxy_protocol_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "off",
				Description: "Whether to enable Proxy Protocol. Valid values: `off` (default), `standard`.",
			},
			"bypass_security_group_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable the function of passing through the backend security group. Default is false.",
			},
			"timestamp_remove_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable the function of removing the TCP/HTTP/HTTPS packet timestamp. Default is false.",
			},
			"server_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server count of the server group.",
			},
			"servers": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The backend servers of the server group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The instance ID of the backend server.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of the backend server. Valid values: `ecs`, `eni`, `ip`.",
						},
						"ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The IP address of the backend server.",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The port of the backend server.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The weight of the backend server. Value range: 0 ~ 100. Default is 100.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description of the backend server.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The zone ID of the backend server.",
						},
						"server_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the backend server.",
						},
					},
				},
			},
			"health_check": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The health check config of the server group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Whether to enable health check. Valid values: `true` (default), `false`.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "TCP",
							Description: "The type of the health check. Valid values: `TCP` (default), `HTTP`, `UDP`.",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The port of health check. Value range: 1 ~ 65535. Default is 0, which means using the port of the backend server.",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The method of health check. Valid values: `GET` (default), `HEAD`. Only available when `HealthCheck.Type` is `HTTP`.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if hc, ok := d.Get("health_check").([]interface{}); ok && len(hc) > 0 {
									if m, ok := hc[0].(map[string]interface{}); ok {
										if t, ok := m["type"].(string); ok && t != "HTTP" {
											return true
										}
									}
								}
								return false
							},
						},
						"uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The uri of health check. Must start with `/`. Only available when `HealthCheck.Type` is `HTTP`.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if hc, ok := d.Get("health_check").([]interface{}); ok && len(hc) > 0 {
									if m, ok := hc[0].(map[string]interface{}); ok {
										if t, ok := m["type"].(string); ok && t != "HTTP" {
											return true
										}
									}
								}
								return false
							},
						},
						"domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The domain of health check. Only available when `HealthCheck.Type` is `HTTP`.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if hc, ok := d.Get("health_check").([]interface{}); ok && len(hc) > 0 {
									if m, ok := hc[0].(map[string]interface{}); ok {
										if t, ok := m["type"].(string); ok && t != "HTTP" {
											return true
										}
									}
								}
								return false
							},
						},
						"http_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The http code of health check. Valid values: `http_2xx`, `http_3xx`, `http_4xx`, `http_5xx`. Default is `http_2xx,http_3xx`. Only available when `HealthCheck.Type` is `HTTP`.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if hc, ok := d.Get("health_check").([]interface{}); ok && len(hc) > 0 {
									if m, ok := hc[0].(map[string]interface{}); ok {
										if t, ok := m["type"].(string); ok && t != "HTTP" {
											return true
										}
									}
								}
								return false
							},
						},
						"timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     2,
							Description: "The response timeout of health check. Unit: second. Value range: 1 ~ 60. Default is 2.",
						},
						"interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     2,
							Description: "The interval of health check. Unit: second. Value range: 1 ~ 300. Default is 2.",
						},
						"healthy_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3,
							Description: "The healthy threshold of health check. Value range: 2 ~ 10. Default is 3.",
						},
						"unhealthy_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3,
							Description: "The unhealthy threshold of health check. Value range: 2 ~ 10. Default is 3.",
						},
						"udp_connect_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The udp connect timeout of health check. Unit: second. Value range: 1 ~ 60. Default is 3. Only available when `HealthCheck.Type` is `UDP`.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if hc, ok := d.Get("health_check").([]interface{}); ok && len(hc) > 0 {
									if m, ok := hc[0].(map[string]interface{}); ok {
										if t, ok := m["type"].(string); ok && t != "UDP" {
											return true
										}
									}
								}
								return false
							},
						},
						"udp_expect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The udp expect of health check. Only available when `HealthCheck.Type` is `UDP`.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if hc, ok := d.Get("health_check").([]interface{}); ok && len(hc) > 0 {
									if m, ok := hc[0].(map[string]interface{}); ok {
										if t, ok := m["type"].(string); ok && t != "UDP" {
											return true
										}
									}
								}
								return false
							},
						},
						"udp_request": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The udp request of health check. Only available when `HealthCheck.Type` is `UDP`.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if hc, ok := d.Get("health_check").([]interface{}); ok && len(hc) > 0 {
									if m, ok := hc[0].(map[string]interface{}); ok {
										if t, ok := m["type"].(string); ok && t != "UDP" {
											return true
										}
									}
								}
								return false
							},
						},
					},
				},
			},
			"tags": ve.TagsSchema(),
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the server group.",
			},
		},
	}
}

func resourceVolcengineNlbServerGroupCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbServerGroupService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on creating nlb server group %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbServerGroupRead(d, meta)
}

func resourceVolcengineNlbServerGroupRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbServerGroupService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on reading nlb server group %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNlbServerGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbServerGroupService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on updating nlb server group %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbServerGroupRead(d, meta)
}

func resourceVolcengineNlbServerGroupDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbServerGroupService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNlbServerGroup())
	if err != nil {
		return fmt.Errorf("error on deleting nlb server group %q, %w", d.Id(), err)
	}
	return err
}

package nlb_server_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNlbServerGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNlbServerGroupsRead,
		Schema: map[string]*schema.Schema{
			"server_group_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of server group IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the VPC.",
			},
			"server_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the server group.",
			},
			"instance_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of backend server IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"server_ips": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of backend server IPs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the server group. Valid values: `instance`, `ip`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the server group.",
			},
			"tags": ve.TagsSchema(),
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"server_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of server groups query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the server group.",
						},
						"server_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the server group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the server group.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of the server group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the server group.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the server group.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the server group.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPC to which the server group belongs.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the server group. Valid values: `instance`, `ip`.\n`instance`: Server type. Supports adding ECS instances and secondary ENIs bound to ECS instances.\n`ip`: IP address type. Supports adding any network-reachable servers in VPC or IDC.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol of the server group. Valid values: `TCP`, `UDP`, `TCP_SSL`.",
						},
						"scheduler": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scheduling algorithm. Valid values: `wrr`, `wlc`, `sh`.",
						},
						"ip_address_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address version of the server group. Valid values: `ipv4`, `ipv6`.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the server group.",
						},
						"any_port_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable full port forwarding. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"connection_drain_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable connection graceful interruption. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"connection_drain_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The timeout period of connection graceful interruption.",
						},
						"preserve_client_ip_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable source address retention. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"session_persistence_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable session persistence. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"session_persistence_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The timeout period of session persistence.",
						},
						"proxy_protocol_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to enable Proxy-Protocol. Valid values: `off`, `standard`.\n`off`: Disable.\n`standard`: Enable. NLB will forward the client source IP address to the server via Proxy-Protocol, and Proxy-Protocol needs to be configured on the server.",
						},
						"bypass_security_group_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the function of passing through the backend security group. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"timestamp_remove_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the function of removing the TCP/HTTP/HTTPS packet timestamp. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"server_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of backend servers in the server group.",
						},
						"related_load_balancer_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ID of the NLB instance associated with the server group.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"health_check": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The health check config of the server group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable health check. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of health check. Valid values: `TCP`, `HTTP`, `UDP`.\n`TCP`: Send SYN handshake packets to detect the port status of the backend server.\n`HTTP`: Send HEAD or GET requests to simulate browsing access behavior to detect whether the backend application is normal.\n`UDP`: Send ICMP or UDP detection packets to detect whether the backend server is normal.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port of health check. 0 indicates the port of the backend server.",
									},
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The method of health check. Valid values: `GET`, `HEAD`.",
									},
									"uri": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The uri of health check.",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of health check.",
									},
									"http_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The normal HTTP status code for health check.",
									},
									"timeout": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The timeout period of health check response.",
									},
									"interval": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The interval of health check.",
									},
									"healthy_threshold": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The healthy threshold of health check.",
									},
									"unhealthy_threshold": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The unhealthy threshold of health check.",
									},
									"udp_connect_timeout": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The udp connect timeout of the health check.",
									},
									"udp_expect": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The expected response string for UDP health check.",
									},
									"udp_request": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The request string for UDP health check.",
									},
								},
							},
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
			"total_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceVolcengineNlbServerGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbServerGroupService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNlbServerGroups())
}

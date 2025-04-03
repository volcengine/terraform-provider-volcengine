package rds_mysql_endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlEndpointsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the mysql instance.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"endpoints": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the mysql endpoint.",
						},
						"endpoint_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the mysql endpoint.",
						},
						"endpoint_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the mysql endpoint.",
						},
						"endpoint_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The endpoint type of the mysql endpoint.",
						},
						"read_write_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The read write mode.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the mysql endpoint.",
						},
						"auto_add_new_nodes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.",
						},
						"enable_read_write_splitting": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.",
						},
						"enable_read_only": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.",
						},
						"read_only_node_weight": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of nodes configured by the connection terminal and the corresponding read-only weights.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the node.",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the node.",
									},
									"weight": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The weight of the node.",
									},
								},
							},
						},
						"addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Address list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network address type, temporarily Private, Public, PublicService.",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Connect domain name.",
									},
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP Address.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Port.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subnet ID, valid only for private addresses.",
									},
									"eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the EIP, only valid for Public addresses.",
									},
									"dns_visibility": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "DNS Visibility.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsMysqlEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlEndpointService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlEndpoints())
}

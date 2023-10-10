package network_interface

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNetworkInterfaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNetworkInterfacesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list of ENI ids.",
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "A type of ENI, Optional choice contains `primary`, `secondary`.",
				ValidateFunc: validation.StringInSlice([]string{"primary", "secondary"}, false),
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Creating", "Available", "Attaching", "InUse", "Detaching", "Deleting"}, false),
				Description:  "A status of ENI, Optional choice contains `Creating`, `Available`, `Attaching`, `InUse`, `Detaching`, `Deleting`.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of the virtual private cloud (VPC) to which the ENI belongs.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of the subnet to which the ENI is connected.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of the instance to which the ENI is bound.",
			},
			"primary_ip_addresses": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsIPAddress,
				},
				Optional:    true,
				Description: "A list of primary IP address of ENI.",
			},
			"security_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of the security group to which the secondary ENI belongs.",
			},
			"network_interface_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "A list of network interface ids.",
			},
			"network_interface_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A name of ENI.",
			},
			"private_ip_addresses": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "A list of private IP addresses.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone ID.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the ENI.",
			},
			"tags": ve.TagsSchema(),

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of ENI query.",
			},
			"network_interfaces": {
				Description: "The collection of ENI.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ENI.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the ENI.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the ENI.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the ENI creator.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the ENI.",
						},
						"network_interface_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ENI.",
						},
						"network_interface_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ENI.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the virtual private cloud (VPC) to which the ENI belongs.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id of the ENI.",
						},
						"vpc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the virtual private cloud (VPC) to which the ENI belongs.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the subnet to which the ENI is connected.",
						},
						"mac_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The mac address of the ENI.",
						},
						"device_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the device to which the ENI is bound.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the ENI.",
						},
						"primary_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The primary IP address of the ENI.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the ENI.",
						},
						"security_group_ids": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of the security group id to which the secondary ENI belongs.",
							Computed:    true,
						},
						"port_security_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The enable of port security.",
						},
						"associated_elastic_ip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The allocation id of the EIP to which the ENI associates.",
						},
						"associated_elastic_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the EIP to which the ENI associates.",
						},
						"service_managed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the network card has been authorized to be used by other account services.",
						},
						"private_ip_sets": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The IP address of secondary private network interface.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"private_ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The secondary private network IP address of the network interface card.",
									},
									"associated_elastic_ip": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The public IP that secondary private network IP associated with.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allocation_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The public IP ID.",
												},
												"eip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The public IP address.",
												},
											},
										},
									},
									"primary": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the network interface is primary IP address.",
									},
								},
							},
						},
						"ipv6_sets": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The IPv6 address list of the ENI.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the ENI.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNetworkInterfacesRead(d *schema.ResourceData, meta interface{}) error {
	networkInterfaceService := NewNetworkInterfaceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(networkInterfaceService, d, DataSourceVolcengineNetworkInterfaces())
}

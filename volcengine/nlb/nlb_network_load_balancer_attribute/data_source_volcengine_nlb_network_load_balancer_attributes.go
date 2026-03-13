package network_load_balancer_attribute

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNlbNetworkLoadBalancerAttributes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNlbNetworkLoadBalancerAttributesRead,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the NLB.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"network_load_balancer_attributes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of NLB attributes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the NLB.",
						},
						"load_balancer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the NLB.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the NLB.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the NLB.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the NLB.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the NLB.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network type of the NLB.",
						},
						"ip_address_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ip address version of the NLB.",
						},
						"cross_zone_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The cross zone enabled of the NLB.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the NLB.",
						},
						"modification_protection_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification protection status of the NLB.",
						},
						"modification_protection_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification protection reason of the NLB.",
						},
						"security_group_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The security group ids of the NLB.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4_bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ipv4 bandwidth package id of the NLB.",
						},
						"ipv6_bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ipv6 bandwidth package id of the NLB.",
						},
						"tags": ve.TagsSchema(),
						"zone_mappings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The zone mappings of the NLB.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The zone id of the NLB.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The subnet id of the NLB.",
									},
									"ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ipv4 address of the NLB.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ipv6 address of the NLB.",
									},
									"eni_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The eni id of the NLB.",
									},
									"ipv4_eip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ipv4 eip address of the NLB.",
									},
									"ipv4_eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ipv4 eip id of the NLB.",
									},
									"ipv6_eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ipv6 eip id of the NLB.",
									},
									"ipv4_hc_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ipv4 hc status of the NLB.",
									},
									"ipv6_hc_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ipv6 hc status of the NLB.",
									},
									"ipv4_local_addresses": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The ipv4 local addresses of the NLB.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ipv6_local_addresses": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The ipv6 local addresses of the NLB.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the NLB.",
						},
						"dns_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dns name of the NLB.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the NLB.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the NLB.",
						},
						"billing_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing status of the NLB.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the NLB.",
						},
						"reclaimed_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reclaimed time of the NLB.",
						},
						"expected_overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected overdue time of the NLB.",
						},
						"managed_security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The managed security group id of the NLB.",
						},
						"ipv4_network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ipv4 network type of the NLB.",
						},
						"ipv6_network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ipv6 network type of the NLB.",
						},
						"access_log_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The access log config of the NLB.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The enabled of the access log config.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The project name of the access log config.",
									},
									"topic_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic name of the access log config.",
									},
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic id of the access log config.",
									},
								},
							},
						},
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

func dataSourceVolcengineNlbNetworkLoadBalancerAttributesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbNetworkLoadBalancerAttributeService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNlbNetworkLoadBalancerAttributes())
}

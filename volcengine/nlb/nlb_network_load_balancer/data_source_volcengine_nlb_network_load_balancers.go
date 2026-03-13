package network_load_balancer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNlbNetworkLoadBalancers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNlbNetworkLoadBalancersRead,
		Schema: map[string]*schema.Schema{
			"load_balancer_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of NLB IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the VPC.",
			},
			"load_balancer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the NLB instance.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the NLB instance. Valid values: `Inactive`, `Active`, `Creating`, `Provisioning`, `Configuring`, `Deleting`, `CreateFailed`.",
			},
			"ip_address_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IP address version of the NLB instance. Valid values: `ipv4`, `dualstack`.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the zone.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the NLB instance.",
			},
			"tags": ve.TagsSchema(),
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"network_load_balancers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of NLB query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the NLB instance.",
						},
						"load_balancer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the NLB instance.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the NLB instance.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of the NLB instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPC where the NLB instance is deployed.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the NLB instance.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network type of the NLB instance. Valid values: `internet`, `intranet`.\n`internet`: The NLB instance is an internet-facing instance.\n`intranet`: The NLB instance is an internal-facing instance.",
						},
						"ip_address_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address version of the NLB instance. Valid values: `ipv4`, `dualstack`.\n`ipv4`: The NLB instance supports IPv4.\n`dualstack`: The NLB instance supports both IPv4 and IPv6.",
						},
						"cross_zone_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable cross-zone load balancing for the NLB instance.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the NLB instance.",
						},
						"modification_protection_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification protection status of the NLB instance. Valid values: `NonProtection`, `ConsoleProtection`.\n`NonProtection`: No protection.\n`ConsoleProtection`: Console protection.",
						},
						"modification_protection_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason for modification protection.",
						},
						"security_group_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The security group IDs of the NLB instance.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4_bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the IPv4 bandwidth package.",
						},
						"ipv6_bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the IPv6 bandwidth package.",
						},
						"tags": ve.TagsSchema(),
						"zone_mappings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The zone mappings of the NLB instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the zone.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the subnet.",
									},
									"ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IPv4 address of the NLB instance.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IPv6 address of the NLB instance.",
									},
									"eni_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the ENI.",
									},
									"ipv4_eip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public IPv4 address.",
									},
									"ipv4_eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the EIP.",
									},
									"ipv6_eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the IPv6 EIP.",
									},
									"ipv4_hc_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IPv4 health check status. Valid values: `Healthy`, `Unhealthy`.",
									},
									"ipv6_hc_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IPv6 health check status. Valid values: `Healthy`, `Unhealthy`.",
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the NLB instance. Valid values: `Inactive`, `Active`, `Creating`, `Provisioning`, `Configuring`, `Deleting`, `CreateFailed`.\n`Inactive`: The NLB instance is stopped.\n`Active`: The NLB instance is running.\n`Creating`: The NLB instance is being created.\n`Provisioning`: The NLB instance is being created. This status is only available when creating an NLB instance by calling the API.\n`Configuring`: The NLB instance is being configured.\n`Deleting`: The NLB instance is being deleted.\n`CreateFailed`: The NLB instance failed to create.",
						},
						"dns_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The DNS name of the NLB instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the NLB instance.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the NLB instance.",
						},
						"billing_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing status of the NLB instance. Valid values: `Normal`, `FinancialLocked`.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the NLB instance.",
						},
						"reclaimed_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reclaimed time of the NLB instance.",
						},
						"expected_overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected overdue time of the NLB instance.",
						},
						"managed_security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The managed security group ID of the NLB instance.",
						},
						"ipv4_network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv4 network type of the NLB instance. Valid values: `internet`, `intranet`.",
						},
						"ipv6_network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv6 network type of the NLB instance. Valid values: `internet`, `intranet`.",
						},
						"access_log_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The access log configuration of the NLB instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable access logging.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The project name where the access log topic resides.",
									},
									"topic_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic name of the access log.",
									},
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic ID of the access log.",
									},
								},
							},
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of NLB query.",
			},
		},
	}
}

func dataSourceVolcengineNlbNetworkLoadBalancersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbNetworkLoadBalancerService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNlbNetworkLoadBalancers())
}

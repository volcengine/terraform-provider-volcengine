package clb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineClbs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineClbsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Clb IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Clb.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of Clb.",
			},
			// 开白参数，API文档已显示支持，API explore暂未支持
			// "exclusive_cluster_id": {
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "The ID of the exclusive cluster.",
			// },
			"instance_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The IDs of the backend server of the CLB.",
			},
			"instance_ips": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The IP address of the backend server of the CLB.",
			},
			"master_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The master zone ID of the CLB.",
			},
			"address_ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "DualStack"}, false),
				Description:  "The address IP version of the CLB.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the CLB.",
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"public", "private"}, false),
				Description:  "The network type of the CLB.",
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
				Description: "The total count of Clb query.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the VPC.",
			},
			"eni_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The private ip address of the Clb.",
			},
			"eip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The public ip address of the Clb.",
			},
			"load_balancer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Clb.",
			},
			"clbs": {
				Description: "The collection of Clb query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the Clb.",
						},
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Clb.",
						},
						"load_balancer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Clb.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Clb.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Clb.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Clb.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the Clb.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Clb.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc ID of the Clb.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet ID of the Clb.",
						},
						"modification_protection_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification protection status of the Clb.",
						},
						"modification_protection_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification protection reason of the Clb.",
						},
						"service_managed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the CLB instance is a managed resource.",
						},
						"address_ip_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The address ip version of the Clb.",
						},
						"eni_ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The eni ipv6 address of the Clb.",
						},
						"eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eip ID of the Clb.",
						},
						"eip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eip address of the Clb.",
						},
						"ipv6_eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Ipv6 Eip ID of the Clb.",
						},
						"eni_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eni ID of the Clb.",
						},
						"eni_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eni address of the Clb.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the Clb.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason why Clb is locked.",
						},
						"load_balancer_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The specifications of the Clb.",
						},
						"load_balancer_billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the Clb.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the Clb.",
						},
						"tags": ve.TagsSchemaComputed(),
						"master_zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The master zone ID of the CLB.",
						},
						"slave_zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The slave zone ID of the CLB.",
						},
						// 计费信息
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the CLB instance.",
						},
						"renew_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The renew type of the CLB. When the value of the load_balancer_billing_type is `PrePaid`, the query returns this field.",
						},
						"renew_period_times": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The renew period times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.",
						},
						"remain_renew_times": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The remain renew times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.",
						},
						"instance_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The billing status of the CLB.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time of the CLB.",
						},
						"reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reclaim time of the CLB.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the Clb.",
						},
						"overdue_reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The over reclaim time of the CLB.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected recycle time of the Clb.",
						},
						// 查询CLB实例详情的API返回
						"log_topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The log topic ID of the Clb.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the CLB instance is enabled.",
						},
						"eip_billing_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The eip billing config of the Clb.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ISP of the EIP assigned to CLB, the value can be `BGP`.",
									},
									"eip_billing_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The billing type of the EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.",
									},
									"bandwidth": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The peek bandwidth of the EIP assigned to CLB. Units: Mbps.",
									},
									"eip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public IP address of the CLB instance.",
									},
									"bandwidth_package_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bandwidth package id of the EIP assigned to CLB.",
									},
									"security_protection_types": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The security protection types of the EIP assigned to CLB.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"ipv6_address_bandwidth": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ipv6 address bandwidth information of the Clb.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The network type of the CLB Ipv6 address.",
									},
									"isp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ISP of the Ipv6 EIP assigned to CLB, the value can be `BGP`.",
									},
									"billing_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The billing type of the Ipv6 EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.",
									},
									"bandwidth": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The peek bandwidth of the Ipv6 EIP assigned to CLB. Units: Mbps.",
									},
									"bandwidth_package_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bandwidth package id of the Ipv6 EIP assigned to CLB.",
									},
								},
							},
						},
						// 查询CLB实例详情的API返回
						"listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information of the listeners in the CLB instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listener_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the Listener.",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the Listener.",
									},
								},
							},
						},
						"server_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information of the server groups in the CLB instance.",
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
								},
							},
						},
						"access_log": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The access log configuration of the CLB instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable the function of delivering access logs (Layer 7) to Object Storage TOS.",
									},
									"bucket_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the bucket to which the access logs are delivered.",
									},
									"tls_enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable the function of delivering access logs (layer 7) to the log service TLS.",
									},
									"tls_project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The project ID of the log service TLS.",
									},
									"tls_topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic ID of the log service TLS.",
									},
								},
							},
						},
						// 邀测中，查询CLB实例列表API返回
						"exclusive_cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the exclusive cluster to which the CLB instance belongs.",
						},
						"bypass_security_group_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the CLB instance has enabled the \"Allow Backend Security Groups\" function.",
						},
						"timestamp_remove_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to enable the function of clearing the timestamp of TCP/HTTP/HTTPS packets (i.e., time stamp).",
						},
						"eni_address_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ENI address num of the CLB.",
						},
						"eni_addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ENI addresses of the CLB.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eni_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The private IPv4 address of the CLB instance.",
									},
									"eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The eip ID of the public IP bound to the private IPv4 address.",
									},
									"eip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public IPv4 address bound to the private IPv4 address.",
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

func dataSourceVolcengineClbsRead(d *schema.ResourceData, meta interface{}) error {
	clbService := NewClbService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(clbService, d, DataSourceVolcengineClbs())
}

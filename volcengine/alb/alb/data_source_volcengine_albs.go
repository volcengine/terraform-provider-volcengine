package alb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Alb IDs.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id which Alb belongs to.",
			},
			"eni_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The private ip address of the Alb.",
			},
			"load_balancer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Alb.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project of the Alb.",
			},
			"tags": ve.TagsSchema(),

			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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

			"albs": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Alb.",
						},
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Alb.",
						},
						"load_balancer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Alb.",
						},
						"address_ip_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The address ip version of the Alb, valid value: `IPv4`, `DualStack`.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Alb.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Alb.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Alb.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the Alb.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Alb, valid value: `public`, `private`.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the Alb.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the Alb, valid value:`Normal`, `FinancialLocked`.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason why Alb is locked. This parameter has a query value only when the status of the Alb instance is `FinancialLocked`.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the Alb. This parameter has a query value only when the status of the Alb instance is `FinancialLocked`.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected deleted time of the Alb. This parameter has a query value only when the status of the Alb instance is `FinancialLocked`.",
						},
						"load_balancer_billing_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The billing type of the Alb.",
						},
						"dns_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The DNS name.",
						},
						"delete_protection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deletion protection function of the Alb instance is turned on or off.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the Alb.",
						},
						"tags": ve.TagsSchemaComputed(),
						"local_addresses": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The local addresses of the Alb.",
						},
						"listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The listener information of the Alb.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listener_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The listener id of the Alb.",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The listener name of the Alb.",
									},
								},
							},
						},
						"access_log": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The access log information of the Alb.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the access log function of the Alb is enabled.",
									},
									"bucket_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bucket name where the logs are stored.",
									},
								},
							},
						},
						"tls_access_log": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The tls access log information of the Alb.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the tls access log function is enabled.",
									},
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The TLS topic id bound to the access log.",
									},
									"project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The TLS project id bound to the access log.",
									},
								},
							},
						},
						"health_log": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The health log information of the Alb.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the health log function is enabled.",
									},
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The TLS topic id bound to the health check log.",
									},
									"project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The TLS project id bound to the health check log.",
									},
								},
							},
						},
						"zone_mappings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration information of the Alb instance in different Availability Zones.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The availability zone id of the Alb.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The subnet id of the Alb in this availability zone.",
									},
									"load_balancer_addresses": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The IP address information of the Alb in this availability zone.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"eni_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Eni address of the Alb in this availability zone.",
												},
												"eni_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Eni id of the Alb in this availability zone.",
												},
												"eip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Eip address of the Alb in this availability zone.",
												},
												"eip_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Eip id of alb instance in this availability zone.",
												},
												"eni_ipv6_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Eni Ipv6 address of the Alb in this availability zone.",
												},
												"ipv6_eip_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Ipv6 Eip id of alb instance in this availability zone.",
												},
												"eip": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The Eip information of the Alb in this availability zone.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"isp": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The ISP of the Eip assigned to Alb, the value can be `BGP`.",
															},
															"eip_billing_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The billing type of the Eip assigned to Alb. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.",
															},
															"bandwidth": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The peek bandwidth of the Eip assigned to Alb. Units: Mbps.",
															},
															"eip_address": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The Eip address of the Alb.",
															},
															"eip_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The Eip type of the Alb.",
															},
															"security_protection_types": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
																Description: "The security protection types of the Alb.",
															},
															"association_mode": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The association mode of the Alb. This parameter has a query value only when the type of the Eip is `anycast`.",
															},
															"pop_locations": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The pop locations of the Alb. This parameter has a query value only when the type of the Eip is `anycast`.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"pop_id": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The pop id of the Anycast Eip.",
																		},
																		"pop_name": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The pop name of the Anycast Eip.",
																		},
																	},
																},
															},
														},
													},
												},
												"ipv6_eip": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The Ipv6 Eip information of the Alb in this availability zone.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"isp": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The ISP of the Ipv6 Eip assigned to Alb, the value can be `BGP`.",
															},
															"billing_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The billing type of the Ipv6 Eip assigned to Alb. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.",
															},
															"bandwidth": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The peek bandwidth of the Ipv6 Eip assigned to Alb. Units: Mbps.",
															},
														},
													},
												},
											},
										},
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

func dataSourceVolcengineAlbsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbs())
}

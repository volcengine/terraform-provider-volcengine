package cluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackVkeVkeClusters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackVkeClustersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Cluster IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Cluster.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Cluster query.",
			},
			"page_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The page number of clusters query.",
			},
			"page_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The page size of clusters query.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the cluster.",
			},
			"delete_protection_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The delete protection of the cluster.",
			},
			"pods_config_pod_network_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network mode of the pod.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The statuses of the cluster",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The status of cluster.",
						},
						"conditions_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "State conditions in the current primary state of the cluster.",
						},
					},
				},
			},
			"create_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ClientToken when successfully created.",
			},
			"update_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ClientToken when the last update was successful.",
			},
			"clusters": {
				Description: "The collection of VkeCluster query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the Cluster.",
						},
						"create_client_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ClientToken when successfully created.",
						},
						"update_client_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ClientToken when the last update was successful.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Cluster.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time the cluster was last admitted and executed/completed.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cluster.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cluster.",
						},
						"delete_protection_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The delete protection of the cluster.",
						},
						"status": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The description of the cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of cluster.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "State conditions in the current primary state of the cluster.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "State conditions in the current primary state of the cluster.",
												},
											},
										},
									},
								},
							},
						},
						"cluster_config": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The config of the cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The VPC ID of the cluster control plane and the network of some nodes.",
									},
									"subnet_ids": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "The list of Subnet IDs.",
									},
									"security_group_ids": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "The list of Security Group IDs.",
									},
									"api_server_public_access_enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Cluster API Server public network access configuration.",
									},
									"api_server_public_access_config": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "Cluster API Server public network access configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"public_access_network_config": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Computed:    true,
													Description: "Public network access network configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"billing_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Billing type of public IP.",
															},
															"bandwidth": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Peak bandwidth of public IP.",
															},
															"isp": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The ISP of public IP.",
															},
														},
													},
												},
												"access_source_ipsv4": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Set:         schema.HashString,
													Description: "IPv4 public network access whitelist.",
												},
											},
										},
									},
									"resource_public_access_default_enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Node public network access configuration.",
									},
									"api_server_endpoints": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "Endpoint information accessed by the cluster API Server.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"private_ip": {
													Type:        schema.TypeList,
													Computed:    true,
													MaxItems:    1,
													Description: "Endpoint address of the cluster API Server private network.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv4": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Ipv4 address.",
															},
														},
													},
												},
												"public_ip": {
													Type:        schema.TypeList,
													Computed:    true,
													MaxItems:    1,
													Description: "Endpoint address of the cluster API Server public network.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv4": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Ipv4 address.",
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
						"pods_config": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The config of the pods.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pod_network_mode": {
										Type: schema.TypeString,
										Computed: true,
										Description: "Container Pod Network Type (CNI).",
									},
									"flannel_config": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "Flannel network configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pod_cidrs": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Set:         schema.HashString,
													Description: "Container Pod Network CIDR.",
												},
												"max_pods_per_node": {
													Type: schema.TypeInt,
													Computed: true,
													Description: "Maximum number of Pod instances on a single node.",
												},
											},
										},
									},
									"vpc_cni_config": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "VPC-CNI network configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vpc_id": {
													Type: schema.TypeString,
													Computed: true,
													Description: "Maximum number of Pod instances on a single node.",
												},
												"subnet_ids": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Set:         schema.HashString,
													Description: "List of subnets corresponding to the container Pod network.",
												},
											},
										},
									},
								},
							},
						},
						"services_config": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The config of the services.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_cidrsv4": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "The IPv4 private network address exposed by the service.",
									},
								},
							},
						},
						"node_statistics": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "Statistics on the number of nodes corresponding to each master state in the cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total_count": {
										Type: schema.TypeInt,
										Computed: true,
										Description: "Total number of nodes.",
									},
									"creating_count": {
										Type: schema.TypeInt,
										Computed: true,
										Description: "Phase=Creating total number of nodes.",
									},
									"running_count": {
										Type: schema.TypeInt,
										Computed: true,
										Description: "Phase=Running total number of nodes.",
									},
									"stopped_count": {
										Type: schema.TypeInt,
										Computed: true,
										Description: "Phase=Stopped total number of nodes.",
									},
									"updating_count": {
										Type: schema.TypeInt,
										Computed: true,
										Description: "Phase=Updating total number of nodes.",
									},
									"deleting_count": {
										Type: schema.TypeInt,
										Computed: true,
										Description: "Phase=Deleting total number of nodes.",
									},
									"failed_count": {
										Type: schema.TypeInt,
										Computed: true,
										Description: "Phase=Failed total number of nodes.",
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

func dataSourceVestackVkeClustersRead(d *schema.ResourceData, meta interface{}) error {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	return clusterService.Dispatcher.Data(clusterService, d, DataSourceVestackVkeVkeClusters())
}

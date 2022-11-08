package cluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVkeVkeClusters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVkeClustersRead,
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
				Description: "The delete protection of the cluster, the value is `true` or `false`.",
			},
			"pods_config_pod_network_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The status of cluster. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Stopped`, `Failed`.",
						},
						"conditions_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The state condition in the current main state of the cluster, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Degraded`, `SetByProvider`, `Balance`, `Security`, `CreateError`, `ResourceCleanupFailed`, `LimitedByQuota`, `StockOut`,`Unknown`.",
						},
					},
				},
			},
			"create_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
			},
			"update_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
			},
			"tags": ve.TagsSchema(),
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
							Description: "ClientToken on successful creation. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
						},
						"update_client_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ClientToken when the last update was successful. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster creation time. UTC+0 time in standard RFC3339 format.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last time a request was accepted by the cluster and executed or completed. UTC+0 time in standard RFC3339 format.",
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
							Description: "The delete protection of the cluster, the value is `true` or `false`.",
						},
						"kubernetes_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Kubernetes version information corresponding to the cluster, specific to the patch version.",
						},
						"status": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The status of the cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of cluster. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Stopped`, `Failed`.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The state condition in the current main state of the cluster, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Balance`, `CreateError`, `ResourceCleanupFailed`, `Unknown`.",
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
										Description: "The ID of the private network (VPC) where the network of the cluster control plane and some nodes is located.",
									},
									"subnet_ids": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "The subnet ID for the cluster control plane to communicate within the private network.",
									},
									"security_group_ids": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "The security group used by the cluster control plane and nodes.",
									},
									"api_server_public_access_enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Cluster API Server public network access configuration, the value is `true` or `false`.",
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
																Description: "Billing type of public IP, the value is `PostPaidByBandwidth` or `PostPaidByTraffic`.",
															},
															"bandwidth": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The peak bandwidth of the public IP, unit: Mbps.",
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
													Description: "IPv4 public network access whitelist. A null value means all network segments (0.0.0.0/0) are allowed to pass.",
												},
											},
										},
									},
									"resource_public_access_default_enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Node public network access configuration, the value is `true` or `false`.",
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
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Container Pod Network Type (CNI), the value is `Flannel` or `VpcCniShared`.",
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
													Description: "Pod CIDR for the Flannel container network.",
												},
												"max_pods_per_node": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The maximum number of single-node Pod instances for a Flannel container network.",
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
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The private network where the cluster control plane network resides.",
												},
												"subnet_ids": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Set:         schema.HashString,
													Description: "A list of Pod subnet IDs for the VPC-CNI container network.",
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
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Total number of nodes.",
									},
									"creating_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Creating total number of nodes.",
									},
									"running_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Running total number of nodes.",
									},
									"stopped_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Stopped total number of nodes.",
									},
									"updating_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Updating total number of nodes.",
									},
									"deleting_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Deleting total number of nodes.",
									},
									"failed_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Failed total number of nodes.",
									},
								},
							},
						},
						"kubeconfig_public": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kubeconfig data with public network access, returned in BASE64 encoding.",
						},
						"kubeconfig_private": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kubeconfig data with private network access, returned in BASE64 encoding.",
						},
						"eip_allocation_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Eip allocation Id.",
						},
						"tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Tags.",
							Set:         ve.VkeTagsResponseHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Key of Tags.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Value of Tags.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Type of Tags.",
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

func dataSourceVolcengineVkeClustersRead(d *schema.ResourceData, meta interface{}) error {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	return clusterService.Dispatcher.Data(clusterService, d, DataSourceVolcengineVkeVkeClusters())
}

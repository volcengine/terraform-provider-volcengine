package veecp_cluster

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVeecpClusters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVeecpClustersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Cluster ID. Supports exact matching." +
					" A maximum of 100 array elements can be filled in at a time." +
					" Note: When this parameter is an empty array, filtering is based on all clusters in the specified region under the account.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster name.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Cluster.",
			},
			"profiles": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Filter by cluster scenario: Cloud: non-edge cluster; Edge: edge cluster.",
			},
			"delete_protection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Cluster deletion protection. Values: " +
					"true: Enable deletion protection. false: Disable deletion protection",
			},
			"pods_config_pod_network_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.",
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Tags.",
				Set:         tagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Value of Tags.",
						},
					},
				},
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
			"clusters": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cluster.",
						},
						"create_client_token": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "ClientToken when creation is successful. " +
								"ClientToken is a string that guarantees request idempotency. " +
								"This string is passed in by the caller.",
						},
						"update_client_token": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "ClientToken when the last update was successful. " +
								"ClientToken is a string that guarantees request idempotency. " +
								"This string is passed in by the caller.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster creation time. UTC+0 time in standard RFC3339 format.",
						},
						"update_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The time when the cluster last accepted a request and executed or completed execution. " +
								"UTC+0 time in standard RFC3339 format.",
						},
						"kubernetes_version": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Kubernetes version information corresponding to the cluster," +
								" specific to the patch version.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster description information.",
						},
						"status": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cluster status. The value contains `Creating`, `Running`, `Updating`, `Deleting`, `Failed`.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "The state condition in the current main state of the cluster," +
														" that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Balance`, `CreateError`, `ResourceCleanupFailed`, `Unknown`.",
												},
											},
										},
									},
								},
							},
							Description: "Cluster status. For detailed instructions, " +
								"please refer to ClusterStatusResponse.",
						},
						"delete_protection_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The delete protection of the cluster, the value is `true` or `false`.",
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
												"ip_family": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "[SkipDoc]The IpFamily configuration,the value is `Ipv4` or `DualStack`.",
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
										Deprecated:  "This field has been deprecated and is not recommended for use.",
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
									"stopping_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Stopping total number of nodes.",
									},
									"starting_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Phase=Starting total number of nodes.",
									},
								},
							},
						},
						"logging_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cluster log configuration information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The TLS log item ID of the collection target.",
									},
									"log_setups": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Cluster logging options.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"log_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The currently enabled log type.",
												},
												"log_ttl": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The storage time of logs in Log Service. After the specified log storage time is exceeded, the expired logs in this log topic will be automatically cleared. The unit is days, and the default is 30 days. The value range is 1 to 3650, specifying 3650 days means permanent storage.",
												},
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to enable the log option, true means enable, false means not enable, the default is false. When Enabled is changed from false to true, a new Topic will be created.",
												},
											},
										},
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Tags of the Cluster.",
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

func dataSourceVolcengineVeecpClustersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVeecpClusterService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVeecpClusters())
}

var tagsHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["key"], m["value"]))
	return hashcode.String(buf.String())
}

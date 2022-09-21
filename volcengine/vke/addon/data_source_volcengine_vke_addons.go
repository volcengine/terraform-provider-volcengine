package addon

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVkeVkeAddons() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVkeAddonsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the addon, fuzzy matching.",
			},
			"pod_network_modes": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Description: "The container network model, the value is `Flannel` or `VpcCniShared`. " +
					"Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. " +
					"VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.",
			},
			"deploy_modes": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The deploy model, the value is `Managed` or `Unmanaged`.",
			},
			"deploy_node_types": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deploy_mode is `Unmanaged`.",
			},
			"necessaries": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The necessaries of addons, the value is `Required` or `Recommended` or `OnDemand`.",
			},
			"categories": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The categories of addons, the value is `Storage` or `Network` or `Monitor` or `Scheduler` or `Dns` or `Security` or `Gpu` or `Image`.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of addons query.",
			},
			"addons": {
				Description: "The collection of addons query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of addon.",
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
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVkeAddonsRead(d *schema.ResourceData, meta interface{}) error {
	clusterService := NewVkeAddonService(meta.(*ve.SdkClient))
	return clusterService.Dispatcher.Data(clusterService, d, DataSourceVolcengineVkeVkeAddons())
}

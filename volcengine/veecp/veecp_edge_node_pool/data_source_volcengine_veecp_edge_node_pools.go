package veecp_edge_node_pool

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVeecpNodePools() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVeecpNodePoolsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Status of NodePool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Phase of Status. The value can be `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Scaling`.",
						},
						"conditions_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the status condition of the node pool in the active state. The value can be `Progressing` or `Ok` or `VersionPartlyUpgraded` or `StockOut` or `LimitedByQuota` or `Balance` or `Degraded` or `ClusterVersionUpgrading` or `Cluster` or `ResourceCleanupFailed` or `Unknown` or `ClusterNotRunning` or `SetByProvider`.",
						},
					},
				},
			},
			"cluster_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ClusterIds of NodePool IDs.",
			},
			"create_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClientToken when successfully created.",
			},
			"update_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClientToken when last update was successful.",
			},
			"auto_scaling_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is enabled of AutoScaling.",
			},
			"node_pool_types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The NodePoolTypes of NodePool.",
			},
			"add_by_script": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Managed by script.",
			},
			"add_by_list": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Managed by list.",
			},
			"add_by_auto": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Managed by auto.",
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
			"node_pools": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of NodePool.",
						},
						"create_client_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ClientToken when successfully created.",
						},
						"update_client_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ClientToken when last update was successful.",
						},
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ClusterId of NodePool.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Name of NodePool.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CreateTime of NodePool.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UpdateTime time of NodePool.",
						},
						"phase": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Phase of Status.",
						},
						"condition_types": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The Condition of Status.",
						},
						"label_content": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The LabelContent of KubernetesConfig.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Key of KubernetesConfig.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Value of KubernetesConfig.",
									},
								},
							},
						},
						"taint_content": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Key of Taint.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Value of Taint.",
									},
									"effect": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Effect of Taint.",
									},
								},
							},
							Description: "The TaintContent of NodeConfig.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node pool type, machine-set: central node pool. edge-machine-set: edge node pool. edge-machine-pool: edge elastic node pool.",
						},
						"profile": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Edge: Edge node pool. If the return value is empty, it is the central node pool.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The static node pool specifies the node pool to associate with the VPC.",
						},
						"node_add_methods": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The method of adding nodes to the node pool.",
						},
						"node_statistics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The TotalCount of Node.",
									},
									"creating_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The CreatingCount of Node.",
									},
									"running_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The RunningCount of Node.",
									},
									"updating_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The UpdatingCount of Node.",
									},
									"deleting_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The DeletingCount of Node.",
									},
									"failed_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The FailedCount of Node.",
									},
									"stopped_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Deprecated:  "This field has been deprecated and is not recommended for use.",
										Description: "The StoppedCount of Node.",
									},
									"stopping_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Deprecated:  "This field has been deprecated and is not recommended for use.",
										Description: "The StoppingCount of Node.",
									},
									"starting_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Deprecated:  "This field has been deprecated and is not recommended for use.",
										Description: "The StartingCount of Node.",
									},
								},
							},
							Description: "The NodeStatistics of NodeConfig.",
						},
						"billing_configs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The billing configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pre_paid_period": {
										Type:     schema.TypeInt,
										Computed: true,
										Description: "The pre-paid period of the node pool, in months. " +
											"The value range is 1-9. " +
											"This parameter takes effect only when the billing_type is PrePaid.",
									},
									"pre_paid_period_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Prepaid period number.",
									},
									"auto_renew": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to automatically renew the node pool.",
									},
								},
							},
						},
						"elastic_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Elastic scaling configuration of node pool.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud_server_identity": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud server identity.",
									},
									"auto_scale_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The auto scaling configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to enable auto scaling.",
												},
												"min_replicas": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The minimum number of nodes.",
												},
												"max_replicas": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The maximum number of nodes.",
												},
												"desired_replicas": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The DesiredReplicas of AutoScaling.",
												},
												"priority": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The Priority of AutoScaling.",
												},
											},
										},
									},
									"instance_area": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The information of instance area.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"area_name": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "Region name. " +
														"You can obtain the regions and operators supported by instance specifications through the ListAvailableResourceInfo interface.",
												},
												"isp": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "Operator. " +
														"You can obtain the regions and operators supported by the instance specification through the ListAvailableResourceInfo interface.",
												},
												"cluster_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Cluster name.",
												},
												"default_isp": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "Default operator. When using three-line nodes, " +
														"this parameter can be configured. After configuration, " +
														"this operator will be used as the default export.",
												},
												"external_network_mode": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "Public network configuration of three-line nodes. " +
														"If it is a single-line node, this parameter will be ignored." +
														" Value range: single_interface_multi_ip: Single network card with multiple IPs. " +
														"single_interface_cmcc_ip: Single network card with China Mobile IP." +
														" Relevant permissions need to be opened by submitting a work order. " +
														"single_interface_cucc_ip: Single network card with China Unicom IP. " +
														"Relevant permissions need to be opened by submitting a work order. " +
														"single_interface_ctcc_ip: Single network card with China Telecom IP. " +
														"Relevant permissions need to be opened by submitting a work order. " +
														"multi_interface_multi_ip: Multiple network cards with multiple IPs. " +
														"Relevant permissions need to be opened by submitting a work order." +
														" no_interface: No public network network card. " +
														"Relevant permissions need to be opened by submitting a work order. " +
														"If this parameter is not configured: " +
														"When there is a public network network card, single_interface_multi_ip is used by default. " +
														"When there is no public network network card, no_interface is used by default.",
												},
												"vpc_identity": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "VPC ID.",
												},
												"subnet_identity": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Subnet ID.",
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

func dataSourceVolcengineVeecpNodePoolsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVeecpNodePools())
}

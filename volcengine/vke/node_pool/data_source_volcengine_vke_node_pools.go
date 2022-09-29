package node_pool

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNodePools() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNodePoolsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The IDs of NodePool.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of NodePool.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Returns the total amount of the data list.",
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
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of NodePool.",
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
			"node_pools": {
				Description: "The collection of NodePools query.",
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
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is Enabled of AutoScaling.",
						},
						"desired_replicas": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The DesiredReplicas of AutoScaling.",
						},
						"min_replicas": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The MinReplicas of AutoScaling.",
						},
						"max_replicas": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The MaxReplicas of AutoScaling.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Priority of AutoScaling.",
						},
						"subnet_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The SubnetId of NodeConfig.",
						},
						"image_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ImageId of NodeConfig.",
						},
						"security_group_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The SecurityGroupIds of NodeConfig.",
						},
						"security_strategy_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The SecurityStrategyEnabled of NodeConfig.",
						},
						"security_strategies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The SecurityStrategies of NodeConfig.",
						},
						"login_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The login type of NodeConfig.",
						},
						"login_key_pair_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The login SshKeyPairName of NodeConfig.",
						},
						"initialize_script": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The InitializeScript of NodeConfig.",
						},
						"system_volume": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Type of SystemVolume.",
									},
									"size": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Size of SystemVolume.",
									},
								},
							},
							Description: "The SystemVolume of NodeConfig.",
						},
						"node_statistics": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
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
										Description: "The StoppedCount of Node.",
									},
									"stopping_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The StoppingCount of Node.",
									},
									"starting_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The StartingCount of Node.",
									},
								},
							},
							Description: "The NodeStatistics of NodeConfig.",
						},
						"cordon": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The Cordon of KubernetesConfig.",
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
						"data_volumes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Type of DataVolume.",
									},
									"size": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Size of DataVolume.",
									},
									"mount_point": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The target mount directory of the disk.",
									},
								},
							},
							Description: "The DataVolume of NodeConfig.",
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
						"additional_container_storage_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is AdditionalContainerStorageEnabled of NodeConfig.",
						},
						"instance_type_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The InstanceTypeIds of NodeConfig.",
						},
						"instance_charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The InstanceChargeType of NodeConfig.",
						},
						"period": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The period of the PrePaid instance of NodeConfig.",
						},
						"auto_renew": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is auto renew of the PrePaid instance of NodeConfig.",
						},
						"auto_renew_period": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The AutoRenewPeriod of the PrePaid instance of NodeConfig.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNodePoolsRead(d *schema.ResourceData, meta interface{}) error {
	nodePoolsService := NewNodePoolService(meta.(*ve.SdkClient))
	return nodePoolsService.Dispatcher.Data(nodePoolsService, d, DataSourceVolcengineNodePools())
}

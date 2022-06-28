package node_pool

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackNodePools() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackNodePoolsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of NodePool IDs.",
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
				Description: "The total count of NodePools query.",
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
							Description: "The Phase of Status.",
						},
						"conditions_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Type of Status.",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of NodePool.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClusterId of NodePool.",
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
				Description: "The create client token of NodePool.",
			},
			"update_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The update client token of NodePool.",
			},
			"auto_scaling_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The Switch of AutoScaling.",
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
							Description: "The ID of NodePool.",
						},
						"create_client_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CreateClientToken of NodePool.",
						},
						"update_client_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UpdateClientToken of NodePool.",
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
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of NodePool.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CreateTime time of NodePool.",
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
							Description: "The switch of AutoScaling.",
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
										Description: "The type of SystemVolume.",
									},
									"size": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The size of SystemVolume.",
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
										Description: "The totalCount of Node.",
									},
									"creating_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The creatingCount of Node.",
									},
									"running_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The runningCount of Node.",
									},
									"updating_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The updatingCount of Node.",
									},
									"deleting_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The deletingCount of Node.",
									},
									"failed_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The failedCount of Node.",
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
							Description: "The Labels of KubernetesConfig.",
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
										Description: "The type of DataVolume.",
									},
									"size": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The size of DataVolume.",
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
										Description: "The key of Taint.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of Taint.",
									},
									"effect": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The effect of Taint.",
									},
								},
							},
							Description: "The taintContent of NodeConfig.",
						},
						"additional_container_storage_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The additionalContainerStorageEnabled of NodeConfig.",
						},
						"instance_type_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The InstanceTypeIds of NodeConfig.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVestackNodePoolsRead(d *schema.ResourceData, meta interface{}) error {
	nodePoolsService := NewNodePoolService(meta.(*ve.SdkClient))
	return nodePoolsService.Dispatcher.Data(nodePoolsService, d, DataSourceVestackNodePools())
}

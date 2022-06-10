package node_pool

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
NodePool can be imported using the id, e.g.
```
$ terraform import vestack_node_pools.default pcabe57vqtofgrbln3dp0
```

*/

func ResourceVestackNodePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackNodePoolCreate,
		Read:   resourceVestackNodePoolRead,
		Update: resourceVestackNodePoolUpdate,
		Delete: resourceVestackNodePoolDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
				Description: "The Status of filter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Phase of NodePool.",
						},
						"conditions_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Type of NodePool Condition.",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of NodePool.",
			},
			"create_client_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create client token of NodePool.",
			},
			"update_client_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update client token of NodePool.",
			},
			"auto_scaling_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The Switch of AutoScaling.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The clusterId  of NodePool.",
			},
			"client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClientToken of NodePool.",
			},
			"auto_scaling": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The Enabled of AutoScaling.",
						},
						"max_replicas": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The MaxReplicas of AutoScaling.",
						},
						"min_replicas": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The MinReplicas of AutoScaling.",
						},
						"desired_replicas": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The DesiredReplicas of AutoScaling.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The Priority of AutoScaling.",
						},
					},
				},
				Description: "The AutoScaling of NodePool.",
			},
			"node_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_type_ids": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The InstanceTypeIds of NodeConfig.",
						},
						"subnet_ids": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The SubnetIds of NodeConfig.",
						},
						"security": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"security_group_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The SecurityGroupIds of Security.",
									},
									"security_strategies": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The SecurityStrategies of Security.",
									},
									"login": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Password of Security.",
												},
												"ssh_key_pair_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The SshKeyPairName of Security.",
												},
											},
										},
										Description: "The Login of Security.",
									},
								},
							},
							Description: "The Security of NodeConfig.",
						},
						"system_volume": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The type of SystemVolume.",
									},
									"size": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The Size of SystemVolume.",
									},
								},
							},
							Description: "The SystemVolume of NodeConfig.",
						},
						"data_volumes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The type of DataVolumes.",
									},
									"size": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The Size of DataVolumes.",
									},
								},
							},
							Description: "The DataVolumes of NodeConfig.",
						},
						"initialize_script": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The InitializeScript of NodeConfig.",
						},
						"additional_container_storage_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The AdditionalContainerStorageEnabled of NodeConfig.",
						},
					},
				},
			},
			"kubernetes_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labels": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Key of Labels.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Value of Labels.",
									},
								},
							},
							Description: "The Labels of KubernetesConfig.",
						},
						"taints": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Key of Taints.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Value of Taints.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Effect of Taints.",
									},
								},
							},
							Description: "The Taints of KubernetesConfig.",
						},
						"cordon": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The Cordon of KubernetesConfig.",
						},
					},
				},
				Description: "The KubernetesConfig of NodeConfig.",
			},
		},
	}
}

func resourceVestackNodePoolCreate(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Create(nodePoolService, d, ResourceVestackNodePool())
	if err != nil {
		return fmt.Errorf("error on creating nodePoolService  %q, %w", d.Id(), err)
	}
	return resourceVestackNodePoolRead(d, meta)
}

func resourceVestackNodePoolRead(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Read(nodePoolService, d, ResourceVestackNodePool())
	if err != nil {
		return fmt.Errorf("error on reading nodePoolService %q, %w", d.Id(), err)
	}
	return err
}

func resourceVestackNodePoolUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Update(nodePoolService, d, ResourceVestackNodePool())
	if err != nil {
		return fmt.Errorf("error on updating nodePoolService  %q, %w", d.Id(), err)
	}
	return resourceVestackNodePoolRead(d, meta)
}

func resourceVestackNodePoolDelete(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewNodePoolService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Delete(nodePoolService, d, ResourceVestackNodePool())
	if err != nil {
		return fmt.Errorf("error on deleting nodePoolService %q, %w", d.Id(), err)
	}
	return err
}
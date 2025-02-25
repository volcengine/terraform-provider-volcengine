package veecp_node_pool

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpNodePool can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node_pool.default resource_id
```

*/

func ResourceVolcengineVeecpNodePool() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpNodePoolCreate,
		Read:   resourceVolcengineVeecpNodePoolRead,
		Update: resourceVolcengineVeecpNodePoolUpdate,
		Delete: resourceVolcengineVeecpNodePoolDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of NodePool.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ClusterId of NodePool.",
			},
			"client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClientToken of NodePool.",
			},
			"tags": ve.TagsSchema(),

			// TODO: 再考虑下这里怎么实现，三种节点池类型 ：https://www.volcengine.com/docs/6499/1122268
			"node_pool_type": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "If you want to create a control node pool, please ignore this field. If you want to create an edge node pool, please fill in this field correctly. " +
					"Node pool type, with the default being a static node pool. " +
					"edge - machine - set: Static node pool. " +
					"edge - machine - pool: Elastic node poolNode pool type, which is static node pool by default. " +
					"edge-machine-set: static node pool\nedge-machine-pool: elastic node pool.",
			},
			"kubernetes_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labels": {
							Type:     schema.TypeSet,
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
							Set:         kubernetesConfigLabelHash,
							Description: "The Labels of KubernetesConfig.",
						},
						"taints": {
							Type:     schema.TypeList,
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
										Default:     "NoSchedule",
										Description: "The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.",
									},
								},
							},
							Description: "The Taints of KubernetesConfig.",
						},
						//"cordon": {
						//	Type:        schema.TypeBool,
						//	Required:    true,
						//	Description: "The Cordon of KubernetesConfig.",
						//},
						//"name_prefix": {
						//	Type:        schema.TypeString,
						//	Optional:    true,
						//	Description: "The NamePrefix of node metadata.",
						//},
					},
				},
				Description: "The KubernetesConfig of NodeConfig.",
			},
			"node_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				//Required: true,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("node_pool_type").(string) != ""
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_type_ids": {
							Type:     schema.TypeList,
							Required: true,
							//ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The InstanceTypeIds of NodeConfig. The value can get from volcengine_veecp_support_resource_types datasource.",
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
										Computed: true,
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
										Description: "The SecurityStrategies of Security, the value can be empty or `Hids`.",
									},
									"login": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"password": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Password of Security, this field must be encoded with base64.",
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
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The Type of SystemVolume, the value can be `PTSSD` or `ESSD_PL0` or `ESSD_FlexPL`.",
									},
									"size": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "The Size of SystemVolume, the value range in 20~2048.",
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
										Default:     "ESSD_PL0",
										Description: "The Type of DataVolumes, the value can be `PTSSD` or `ESSD_PL0` or `ESSD_FlexPL`. Default value is `ESSD_PL0`.",
									},
									"size": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     20,
										Description: "The Size of DataVolumes, the value range in 20~32768. Default value is `20`.",
									},
									//"mount_point": {
									//	Type:        schema.TypeString,
									//	Optional:    true,
									//	Description: "The target mount directory of the disk. Must start with `/`.",
									//},
								},
							},
							Description: "The DataVolumes of NodeConfig.",
						},
						"initialize_script": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The initializeScript of NodeConfig.",
						},
						"additional_container_storage_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The AdditionalContainerStorageEnabled of NodeConfig.",
						},
						"image_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The ImageId of NodeConfig.",
						},
						"instance_charge_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "PostPaid",
							ForceNew:    true,
							Description: "The InstanceChargeType of PrePaid instance of NodeConfig. Valid values: PostPaid, PrePaid. Default value: PostPaid.",
						},
						"period": {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: prePaidDiffSuppressFunc,
							Description:      "The Period of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36. Unit: month. when InstanceChargeType is PrePaid, default value is 12.",
						},
						"auto_renew": {
							Type:             schema.TypeBool,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: prePaidDiffSuppressFunc,
							Description:      "Is AutoRenew of PrePaid instance of NodeConfig. Valid values: true, false. when InstanceChargeType is PrePaid, default value is true.",
						},
						"auto_renew_period": {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: prePaidAndAutoNewDiffSuppressFunc,
							Description:      "The AutoRenewPeriod of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 6, 12. Unit: month. when InstanceChargeType is PrePaid and AutoRenew enable, default value is 1.",
						},
						"name_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The NamePrefix of NodeConfig.",
						},
						"ecs_tags": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Tags for Ecs.",
							Set:         ve.TagsHash,
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
						"hpc_cluster_ids": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The IDs of HpcCluster, only one ID is supported currently.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
					},
				},
				Description: "The Config of NodePool. If you want to create a control node pool, this field is mandatory.",
			},
			"elastic_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Elastic scaling configuration. This field takes effect only when the node_pool_type is edge-machine-pool.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("node_pool_type").(string) != ""
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_server_identity": {
							Type:     schema.TypeString,
							Required: true,
							Description: "The ID of the edge service corresponding to the elastic node. " +
								"On the edge computing node's edge service page, obtain the edge service ID.",
						},
						"auto_scaling": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										Description: "Whether to enable the auto scaling function of the node pool. " +
											"When a node needs to be manually added to the node pool, the value of this field must be `false`.",
									},
									"max_replicas": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  10,
										Description: "The MaxReplicas of AutoScaling, default 10, range in 1~2000. " +
											"This field is valid when the value of `enabled` is `true`.",
									},
									"min_replicas": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										Description: "The MinReplicas of AutoScaling, default 0. " +
											"This field is valid when the value of `enabled` is `true`.",
									},
									"desired_replicas": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "The DesiredReplicas of AutoScaling, default 0, range in min_replicas to max_replicas.",
									},
									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										Description: "The Priority of AutoScaling, default 10, rang in 0~100. " +
											"This field is valid when the value of `enabled` is `true` and the value of `subnet_policy` is `Priority`.",
									},
								},
							},
							Description: "The node pool elastic scaling configuration information.",
						},
						"instance_area": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										Description: "Region name. " +
											"You can obtain the regions and operators supported by instance specifications through the ListAvailableResourceInfo interface.",
									},
									"isp": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										Description: "Operator. " +
											"You can obtain the regions and operators supported by instance specifications through the ListAvailableResourceInfo interface.",
									},
									"cluster_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The cluster name.",
									},
									"default_isp": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										Description: "Default operator." +
											" When using three-line nodes, this parameter can be configured. " +
											"After configuration, this operator will be used as the default export.",
									},
									"external_network_mode": {
										Type:     schema.TypeString,
										Optional: true,
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
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVeecpNodePoolCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on creating veecp_node_pool %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpNodePoolRead(d, meta)
}

func resourceVolcengineVeecpNodePoolRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on reading veecp_node_pool %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpNodePoolUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on updating veecp_node_pool %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpNodePoolRead(d, meta)
}

func resourceVolcengineVeecpNodePoolDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_node_pool %q, %s", d.Id(), err)
	}
	return err
}

var kubernetesConfigLabelHash = func(v interface{}) int {
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

var prePaidDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	chargeType := d.Get("node_config").([]interface{})[0].(map[string]interface{})["instance_charge_type"].(string)
	return chargeType != "PrePaid"
}

var prePaidAndAutoNewDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	nodeConfig := d.Get("node_config").([]interface{})[0].(map[string]interface{})
	chargeType := nodeConfig["instance_charge_type"].(string)
	autoRenew := nodeConfig["auto_renew"].(bool)
	return chargeType != "PrePaid" || !autoRenew
}

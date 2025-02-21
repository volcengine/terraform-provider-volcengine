package veecp_deployment

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpDeployment can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_deployment.default resource_id
```

*/

func ResourceVolcengineVeecpDeployment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpDeploymentCreate,
		Read:   resourceVolcengineVeecpDeploymentRead,
		Update: resourceVolcengineVeecpDeploymentUpdate,
		Delete: resourceVolcengineVeecpDeploymentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster ID.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specify the namespace for stateless workload deployment.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Set the name of the stateless workload.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set the description of the stateless workload.",
			},
			"replicas": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the number of replicas.",
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Set labels for stateless workloads.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the label.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the label.",
						},
					},
				},
			},
			"termination_grace_period_seconds": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "Graceful termination period in seconds. " +
					"The unit is second. The default value is 30.",
			},
			"strategy": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Update strategy. Optional values are as follows: " +
					"RollingUpdate (default value): Rolling update. Gradually replace old instances with new instances of the version. " +
					"Recreate: First delete the old version instances of the workload, and then install the specified new version instances." +
					" The business will be interrupted during the upgrade process.",
			},
			"max_unavailable": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("strategy").(string) != "RollingUpdate"
				},
				Description: "MaxUnavailable refers to the lower limit of the number of available instances (Pods) in a stateless workload during the shrinking process. " +
					"The default value is 25%. This parameter needs to be set only when the Strategy parameter is set to RollingUpdate.",
			},
			"max_surge": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("strategy").(string) != "RollingUpdate"
				},
				Description: "MaxSurge refers to the upper limit of the number of available instances (Pods) in a stateless load during the expansion process. " +
					"The default value is 25%. This parameter needs to be set only when the Strategy parameter is set to RollingUpdate.",
			},
			"pod_labels": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Set labels for Pods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the label.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the label.",
						},
					},
				},
			},
			"dns_policy": {
				Type:     schema.TypeString,
				Required: true,
				Description: "DNS policy. Optional values are as follows: " +
					"ClusterFirst (default value): In this mode, the information of Kube-dns or CoreDNS is regarded as preset parameters and written into the DNS configuration in this Pod. " +
					"Default: In this mode, the DNS configuration inside the Pod inherits the DNS configuration on the host machine." +
					" That is, the DNS configuration of this Pod is exactly the same as that of the host machine.",
			},
			"run_as_non_root": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Security policy, whether to run as a non-root user.",
			},
			"image_pull_secrets": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Image repository key name.",
			},
			"node_affinity": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Node affinity scheduling.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"required_terms": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Forced scheduling.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"match_expressions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Select according to expressions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the label.",
												},
												"operator": {
													Type:     schema.TypeString,
													Required: true,
													Description: "Matching relationship. " +
														"Optional values are as follows: In, NotIn, Exists, DoesNotExist",
												},
												"values": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the label.",
												},
											},
										},
									},
								},
							},
						},
						"preferred_terms": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Preferred scheduling.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weight": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Weight. The value range is 1 to 100.",
									},
									"match_expressions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Select according to expressions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the label.",
												},
												"operator": {
													Type:     schema.TypeString,
													Required: true,
													Description: "Matching relationship. " +
														"Optional values are as follows: In, NotIn, Exists, DoesNotExist",
												},
												"values": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the label.",
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
			"tolerations": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Node taint scheduling.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the label.",
						},
						"operator": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Matching relationship. Optional values are as follows: equal, exists.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tag value.",
						},
						"effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Stain effect. Optional values are as follows: NoSchedule, PreferNoSchedule, NoExecute.",
						},
						"toleration_seconds": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Toleration seconds, in seconds.",
						},
					},
				},
			},
			"pod_affinity": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Instance affinity scheduling.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"required_terms": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Forced scheduling.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"namespace": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The namespace.",
									},
									"match_expressions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Select according to expressions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the label.",
												},
												"operator": {
													Type:     schema.TypeString,
													Required: true,
													Description: "Matching relationship. " +
														"Optional values are as follows: In, NotIn, Exists, DoesNotExist",
												},
												"values": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the label.",
												},
											},
										},
									},
									"match_labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Select according to labels.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of the label.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of the label.",
												},
											},
										},
									},
									"topology_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Topology domain.",
									},
								},
							},
						},
						"preferred_terms": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Preferred scheduling.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weight": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Weight. The value range is 1 to 100.",
									},
									"topology_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Topology domain.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The namespace.",
									},
									"match_labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Select according to labels.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of the label.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of the label.",
												},
											},
										},
									},
									"match_expressions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Select according to expressions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the label.",
												},
												"operator": {
													Type:     schema.TypeString,
													Required: true,
													Description: "Matching relationship. " +
														"Optional values are as follows: In, NotIn, Exists, DoesNotExist",
												},
												"values": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the label.",
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
			"pod_anti_affinity": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Instance anti-affinity scheduling.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"required_terms": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Forced scheduling.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"namespace": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The namespace.",
									},
									"match_expressions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Select according to expressions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the label.",
												},
												"operator": {
													Type:     schema.TypeString,
													Required: true,
													Description: "Matching relationship. " +
														"Optional values are as follows: In, NotIn, Exists, DoesNotExist",
												},
												"values": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the label.",
												},
											},
										},
									},
									"match_labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Select according to labels.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of the label.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of the label.",
												},
											},
										},
									},
									"topology_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Topology domain.",
									},
								},
							},
						},
						"preferred_terms": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Preferred scheduling.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weight": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Weight. The value range is 1 to 100.",
									},
									"topology_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Topology domain.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The namespace.",
									},
									"match_labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Select according to labels.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of the label.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of the label.",
												},
											},
										},
									},
									"match_expressions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Select according to expressions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the label.",
												},
												"operator": {
													Type:     schema.TypeString,
													Required: true,
													Description: "Matching relationship. " +
														"Optional values are as follows: In, NotIn, Exists, DoesNotExist",
												},
												"values": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the label.",
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
			"volumes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Volume settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Volume Types. " +
								"The optional values are as follows: " +
								"- Configmap: Configuration item. If you mount all keys of the configuration item, " +
								"enter the configuration item name + volume name. " +
								"If you mount some keys of the configuration item, " +
								"enter the configuration item name + key of the configuration item + volume name (separate the key and path with a colon)." +
								" - Secret: Secret dictionary. If you mount all keys of the secret dictionary, " +
								"enter the secret dictionary name + volume name. If you mount some keys of the secret dictionary, " +
								"enter the secret dictionary name + key of the secret dictionary + volume name (separate the key and path with a colon). " +
								"- pvc: Persistent Volume Claim. Enter the name of the Persistent Volume Claim + volume name. " +
								"- emptydir: Temporary directory. Enter the volume name. " +
								"- hostpath: Host directory. Enter path + type + volume name.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Storage volume name.",
						},
						"source_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the storage volume claim/config item/secret dictionary.",
						},
						"config_keys": {
							Type:     schema.TypeSet,
							Optional: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The key of a configuration item or secret dictionary. If it is empty, all keys will be used.",
						},
						"path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The path of HostPath.",
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The type of HostPath. For the values of type," +
								" please refer to Kubernetes official documentation.",
						},
					},
				},
			},
			"annotations": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Annotations for stateless loads.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the annotation.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the annotation.",
						},
					},
				},
			},
			"pod_annotations": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Annotations for Pods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the annotation.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the annotation.",
						},
					},
				},
			},
			"dns_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "DNS configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nameservers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "DNS server list.",
						},
						"searches": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "DNS search domain list.",
						},
						"options": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "DNS configuration list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the configuration.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the configuration.",
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

func resourceVolcengineVeecpDeploymentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on creating veecp_deployment %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpDeploymentRead(d, meta)
}

func resourceVolcengineVeecpDeploymentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on reading veecp_deployment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpDeploymentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on updating veecp_deployment %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpDeploymentRead(d, meta)
}

func resourceVolcengineVeecpDeploymentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpDeploymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpDeployment())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_deployment %q, %s", d.Id(), err)
	}
	return err
}

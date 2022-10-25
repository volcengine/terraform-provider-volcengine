package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineMongoDBInstancesRead,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone ID to query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance ID to query.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance name to query.",
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The type of instance to query,the valid value contains `ReplicaSet` or `ShardedCluster`.",
				ValidateFunc: validation.StringInSlice([]string{"ReplicaSet", "ShardedCluster"}, false),
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance status to query.",
			},
			"db_engine": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"MongoDB"}, false),
				Description:  "The db engine to query,valid value contains `MongoDB`.",
			},
			"db_engine_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"MongoDB_4_0"}, false),
				Description:  "The version of db engine to query,valid value contains `MongoDB_4_0`.",
			},
			"create_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of creation to query.",
			},
			"create_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of creation to query.",
			},
			"update_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of update to query.",
			},
			"update_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of update to query.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id of instance to query.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of mongodb instances query.",
			},
			"instances": {
				Description: "The collection of mongodb instances query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_renew": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable automatic renewal.",
						},
						"charge_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge status.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of instance.",
						},
						"closed_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The planned close time.",
						},
						"config_servers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of config servers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config_server_node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The config server node ID.",
									},
									"node_role": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The config server node role.",
									},
									"node_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The config server node status.",
									},
									"total_memory_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The total memory in GB.",
									},
									"total_vcpu": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The total vCPU.",
									},
									"used_memory_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The used memory in GB.",
									},
									"used_vcpu": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The used vCPU.",
									},
								},
							},
						},
						"config_servers_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of config servers.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of instance.",
						},
						"db_engine": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The db engine.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of database engine.",
						},
						"db_engine_version_str": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version string of database engine.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time of instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance ID.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance name.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance status.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance type.",
						},
						"nodes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The node information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_delay_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The master-slave delay time.",
									},
									"node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node ID.",
									},
									"node_role": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The nod role.",
									},
									"node_spec": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node spec.",
									},
									"node_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node status.",
									},
									"total_memory_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The total memory in GB.",
									},
									"total_storage_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The total storage in GB.",
									},
									"total_vcpu": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The total vCPU.",
									},
									"used_memory_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The used memory in GB.",
									},
									"used_storage_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The used storage in GB.",
									},
									"used_vcpu": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The used vCPU.",
									},
								},
							},
						},
						"mongos": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of mongos.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mongos_node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The mongos node ID.",
									},
									"node_spec": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node spec.",
									},
									"node_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node status.",
									},
									"total_memory_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The total memory in GB.",
									},
									"total_vcpu": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The total vCPU.",
									},
									"used_memory_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The used memory in GB.",
									},
									"used_vcpu": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The used vCPU.",
									},
								},
							},
						},

						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name to which the instance belongs.",
						},
						"mongos_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of mongos.",
						},
						"reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The planned reclaim time of instance.",
						},
						"shards": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of shards.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nodes": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The node information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"node_delay_time": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The master-slave delay time.",
												},
												"node_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The node ID.",
												},
												"node_role": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The nod role.",
												},
												"node_spec": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The node spec.",
												},
												"node_status": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The node status.",
												},
												"total_memory_gb": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The total memory in GB.",
												},
												"total_storage_gb": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The total storage in GB.",
												},
												"total_vcpu": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The total vCPU.",
												},
												"used_memory_gb": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The used memory in GB.",
												},
												"used_storage_gb": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The used storage in GB.",
												},
												"used_vcpu": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The used vCPU.",
												},
											},
										},
									},
									"shard_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The shard ID.",
									},
								},
							},
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc ID.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of instance.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of instance.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							ForceNew:    true,
							Description: "The zone ID of instance.",
						},
						"ssl_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether ssl enabled.",
						},
						"ssl_is_valid": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether ssl is valid.",
						},
						"ssl_expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ssl expire time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineMongoDBInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineMongoDBInstances())
}

package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRedisDbInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRedisDbInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of redis instance to query. This field supports fuzzy queries.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of redis instance to query. This field supports fuzzy queries.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id of redis instance to query. This field supports fuzzy queries.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone id of redis instance to query. This field supports fuzzy queries.",
			},
			"sharded_cluster": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntInSlice([]int{0, 1}),
				Description:  "Whether enable sharded cluster for redis instance. Valid values: 0, 1.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of redis instance to query.",
			},
			"engine_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"4.0", "5.0", "6.0"}, false),
				Description:  "The engine version of redis instance to query. Valid values: `4.0`, `5.0`, `6.0`.",
			},
			"charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostPaid", "PrePaid"}, false),
				Description:  "The charge type of redis instance to query. Valid values: `PostPaid`, `PrePaid`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of redis instance to query.",
			},
			"tags": ve.TagsSchema(),

			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A name regex of redis.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of redis instances query.",
			},
			"instances": {
				Description: "The collection of redis instances query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// InstanceInfo
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the redis instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the redis instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the redis instance.",
						},
						"capacity": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The memory capacity information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The total memory capacity of the redis instance. Unit: MiB.",
									},
									"used": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The used memory capacity of the redis instance. Unit: MiB.",
									},
								},
							},
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of the redis instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the redis instance.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expire time of the redis instance, valid when charge type is `PrePaid`.",
						},
						"engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine version of the redis instance.",
						},
						"node_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of nodes in each shard.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id of the redis instance.",
						},
						"shard_capacity": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The memory capacity of each shard. Unit: GiB.",
						},
						"shard_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of shards in the redis instance.",
						},
						"sharded_cluster": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether enable sharded cluster for the redis instance.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the redis instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc ID of the redis instance.",
						},
						"zone_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of zone ID which the redis instance belongs.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the redis instance.",
						},
						"tags": ve.TagsSchemaComputed(),

						// InstanceDetail
						"deletion_protection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "whether enable deletion protection.",
						},
						"maintenance_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The maintainable time of the redis instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the redis instance.",
						},
						"visit_addrs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of connection information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The connection address type.",
									},
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The connection address.",
									},
									"eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The EIP ID bound to the instance's public network address.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The connection port.",
									},
								},
							},
						},
						"vpc_auth_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to enable password-free access when connecting to an instance through a private network.",
						},

						"params": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of params.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"current_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Current value of the configuration parameter.",
									},
									"default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value of the configuration parameter.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the configuration parameter.",
									},
									"editable_for_instance": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the current redis instance supports editing this parameter.",
									},
									"need_reboot": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether need to reboot the redis instance when modifying this parameter.",
									},
									"options": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of options. Valid when the configuration parameter type is `Radio`.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Optional item for `Radio` type parameters.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of this option item.",
												},
											},
										},
									},
									"param_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the configuration parameter.",
									},
									"range": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The valid value range of the numeric type configuration parameter.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the configuration parameter.",
									},
									"unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unit of the numeric type configuration parameter.",
									},
								},
							},
						},
						"backup_plan": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The list of backup plans.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"active": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether enable auto backup.",
									},
									"backup_hour": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The time period to start performing the backup. The value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on.",
									},
									"backup_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The backup type.",
									},
									"expect_next_backup_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The expected time for the next backup to be performed.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The instance ID.",
									},
									"last_update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The last time the backup policy was modified.",
									},
									"period": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The backup cycle. The value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"ttl": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of days to keep backups, the default is 7 days.",
									},
								},
							},
						},
						"node_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of redis instance node IDs.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"multi_az": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Set the availability zone deployment scheme for the instance. " +
								"The value range is as follows: \n" +
								"disabled: Single availability zone deployment scheme.\n " +
								"enabled: Multi-availability zone deployment scheme.\n " +
								"Description:\n When the newly created instance is a single-node instance" +
								" (that is, when the value of NodeNumber is 1), only the single availability zone deployment scheme is allowed. " +
								"At this time, the value of MultiAZ must be disabled.",
						},
						"configure_nodes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set the list of available zones to which the node belongs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"az": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Set the availability zone to which the node belongs. " +
											"The number of nodes of an instance (i.e., NodeNumber) and the availability zone deployment scheme (i.e., the value of the MultiAZ parameter) will affect the filling of the current parameter." +
											" Among them:\n When a new instance is a single-node instance (i.e., the value of NodeNumber is 1), " +
											"only a single availability zone deployment scheme is allowed (i.e., the value of MultiAZ must be disabled). " +
											"At this time, only one availability zone needs to be passed in AZ, " +
											"and all nodes in the instance will be deployed in this availability zone. " +
											"When creating a new instance as a primary-standby instance (that is, when the value of NodeNumber is greater than or equal to 2), " +
											"the number of availability zones passed in must be equal to the number of nodes in a single shard (that is, the value of the NodeNumber parameter), " +
											"and the value of AZ must comply with the multi-availability zone deployment scheme rules. " +
											"The specific rules are as follows: If the primary-standby instance selects the multi-availability zone deployment scheme (that is, the value of MultiAZ is enabled), " +
											"then at least two different availability zone IDs must be passed in in AZ, and the first availability zone is the availability zone where the primary node is located." +
											" If the primary and standby instances choose a single availability zone deployment scheme (that is, the value of MultiAZ is disabled), then the availability zones passed in for each node must be the same.",
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

func dataSourceVolcengineRedisDbInstancesRead(d *schema.ResourceData, meta interface{}) error {
	redisInstanceService := NewRedisDbInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(redisInstanceService, d, DataSourceVolcengineRedisDbInstances())
}

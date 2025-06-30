package instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
redis instance can be imported using the id, e.g.
```
$ terraform import volcengine_redis_instance.default redis-n769ewmjjqyqh5dv
```
Adding or removing nodes and migrating availability zones for multiple AZ instances are not supported to be orchestrated simultaneously, but it is possible for single AZ instances.
*/

func ResourceVolcengineRedisDbInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRedisDbInstanceCreate,
		Read:   resourceVolcengineRedisDbInstanceRead,
		Update: resourceVolcengineRedisDbInstanceUpdate,
		Delete: resourceVolcengineRedisDbInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"zone_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Deprecated:  "This field has been deprecated after version-0.0.152. Please use multi_az and configure_nodes to specify the availability zone.",
				Description: "The list of zone IDs of instance. When creating a single node instance, only one zone id can be specified.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The subnet id of the redis instance. The specified subnet id must belong to the zone ids.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the redis instance.",
			},
			"sharded_cluster": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Whether enable sharded cluster for the current redis instance. Valid values: 0, 1. 0 means disable, 1 means enable.",
			},
			"shard_number": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description:      "The number of shards in redis instance, the valid value range is `2-256`. This field is valid and required when the value of `ShardedCluster` is 1.",
			},
			"shard_capacity": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The memory capacity of each shard, unit is MiB. The valid value range is as fallows: When the value of `ShardedCluster` is 0: 256, 1024, 2048, 4096, 8192, 16384, 32768, 65536. When the value of `ShardedCluster` is 1: 1024, 2048, 4096, 8192, 16384. When the value of `node_number` is 1, the value of this field can not be 256.",
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Description: "The account password. When importing resources, this attribute will not be imported. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields. " +
					"If this parameter is left blank, it means that no password is set for the default account. " +
					"At this time, the system will automatically generate a password for the default account to ensure instance access security. " +
					"No account can obtain this random password. Therefore, " +
					"before connecting to the instance, you need to reset the password of the default account through the ModifyDBAccount interface." +
					"You can also set a new account and password through the CreateDBAccount interface according to business needs. " +
					"If you need to use password-free access function, you need to enable password-free access first through the ModifyDBInstanceVpcAuthMode interface.",
			},
			"node_number": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The number of nodes in each shard, the valid value range is `1-6`. When the value is 1, it means creating a single node instance, and this field can not be modified. When the value is greater than 1, it means creating a primary and secondary instance, and this field can be modified.",
			},
			"multi_az": {
				Type: schema.TypeString,
				// 新增required字段不兼容了
				// 改为optional，兼容改动
				Optional: true,
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
				Type:     schema.TypeList,
				Optional: true,
				// 新增required字段不兼容了
				// 改为optional，兼容改动
				Computed:    true,
				Description: "Set the list of available zones to which the node belongs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"az": {
							Type:     schema.TypeString,
							Required: true,
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
			"additional_bandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
				Description: "Modify the single-shard additional bandwidth of the target Redis instance. " +
					"Set the additional bandwidth of a single shard, that is, " +
					"the bandwidth that needs to be additionally increased on the basis of the default bandwidth. Unit: MB/s. " +
					"The value of additional bandwidth needs to meet the following conditions at the same time: " +
					"It must be greater than or equal to 0. When the value is 0, it means that no additional bandwidth is added, " +
					"and the bandwidth of a single shard is the default bandwidth. " +
					"The sum of additional bandwidth and default bandwidth cannot exceed the upper limit of bandwidth that can be modified for the current instance. " +
					"Different specification nodes have different upper limits of bandwidth that can be modified. " +
					"For more details, please refer to bandwidth modification range. " +
					"The upper limits of the total write bandwidth and the total read bandwidth of an instance are both 2048MB/s.",
			},
			"engine_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The engine version of redis instance. Valid value: `5.0`, `6.0`, `7.0`.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "PostPaid",
				Description: "The charge type of redis instance. Valid value: `PostPaid`, `PrePaid`.",
			},
			"purchase_months": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          1,
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description:      "The purchase months of redis instance, the unit is month. the valid value range is as fallows: `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36`. This field is valid and required when `ChargeType` is `Prepaid`. \nWhen importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"auto_renew": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description:      "Whether to enable automatic renewal. This field is valid only when `ChargeType` is `PrePaid`, the default value is false. \nWhen importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     6379,
				Description: "The port of custom define private network address. The valid value range is `1024-65535`. The default value is `6379`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name to which the redis instance belongs, if this parameter is empty, the new redis instance will be added to the `default` project.",
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
			"deletion_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "Whether enable deletion protection for redis instance. Valid values: `enabled`, `disabled`(default).",
			},
			"create_backup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to create a final backup when modify the instance configuration or destroy the redis instance.",
			},
			"apply_immediately": {
				Type:             schema.TypeBool,
				Optional:         true,
				Description:      "Whether to apply the instance configuration change operation immediately. The value of this field is false, means that the change operation will be applied within maintenance time.",
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
			},
			"vpc_auth_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable password-free access when connecting to an instance through a private network. Valid values: `open`, `close`.",
			},
			"param_values": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      paramHash,
				Description: "The configuration item information to be modified. This field can only be added or modified. Deleting this field is invalid.\n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields, or use the command `terraform apply` to perform a modification operation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of configuration parameter.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of configuration parameter.",
						},
					},
				},
			},
			"backup_period": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description: "The backup period. The valid value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on. \n" +
					"This field is valid and required when updating the backup plan of primary and secondary instance.",
			},
			"backup_hour": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description: "The time period to start performing the backup. The valid value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on. \n" +
					"This field is valid and required when updating the backup plan of primary and secondary instance.",
			},
			"backup_active": {
				Type:             schema.TypeBool,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description:      "Whether enable auto backup for redis instance. This field is valid and required when updating the backup plan of primary and secondary instance.",
			},
			"backup_point_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Set the backup name for the final backup of the instance to be deleted. " +
					"If the backup name is not set, the backup ID is used as the name by default. " +
					"Use lifecycle and ignore_changes in import.",
			},
			"time_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "The maintainable time period of the instance, in the format of HH:mm-HH:mm (UTC+8).",
			},
			"max_connections": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "Maximum number of connections per shard.",
			},
			"addr_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "The type of connection address that requires an address prefix. " +
					"Use lifecycle and ignore_changes in import.",
			},
			"new_address_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "The modified connection address prefix. " +
					"Use lifecycle and ignore_changes in import.",
			},
			"new_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "The modified connection address port number. " +
					"Use lifecycle and ignore_changes in import.",
			},
			"upgrade_region_domain": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "Whether to upgrade the domain suffix of the connection address. " +
					"Use lifecycle and ignore_changes in import.",
			},
		},
	}
	return resource
}

func resourceVolcengineRedisDbInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisInstanceService := NewRedisDbInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(redisInstanceService, d, ResourceVolcengineRedisDbInstance())
	if err != nil {
		return fmt.Errorf("Error on creating instance %q,%s", d.Id(), err)
	}
	return resourceVolcengineRedisDbInstanceRead(d, meta)
}

func resourceVolcengineRedisDbInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	redisInstanceService := NewRedisDbInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(redisInstanceService, d, ResourceVolcengineRedisDbInstance())
	if err != nil {
		return fmt.Errorf("error on updating instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineRedisDbInstanceRead(d, meta)
}

func resourceVolcengineRedisDbInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	redisInstanceService := NewRedisDbInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(redisInstanceService, d, ResourceVolcengineRedisDbInstance())
	if err != nil {
		return fmt.Errorf("error on deleting instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRedisDbInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisInstanceService := NewRedisDbInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(redisInstanceService, d, ResourceVolcengineRedisDbInstance())
	if err != nil {
		return fmt.Errorf("Error on reading instance %q,%s", d.Id(), err)
	}
	return err
}

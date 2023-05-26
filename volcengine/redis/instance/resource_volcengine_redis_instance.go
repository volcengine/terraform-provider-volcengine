package instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
redis instance can be imported using the id, e.g.
```
$ terraform import volcengine_redis_instance.default redis-n769ewmjjqyqh5dv
```

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
				Required:    true,
				ForceNew:    true,
				Description: "The list of zone IDs of instance. When creating a single node instance, only one zone id can be specified.",
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
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntInSlice([]int{0, 1}),
				Description:  "Whether enable sharded cluster for the current redis instance. Valid values: 0, 1. 0 means disable, 1 means enable.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"node_number": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 6),
				Description:  "The number of nodes in each shard, the valid value range is `1-6`. When the value is 1, it means creating a single node instance, and this field can not be modified. When the value is greater than 1, it means creating a primary and secondary instance, and this field can be modified.",
			},
			"shard_number": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.IntBetween(2, 256),
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description:      "The number of shards in redis instance, the valid value range is `2-256`. This field is valid and required when the value of `ShardedCluster` is 1.",
			},
			"shard_capacity": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The memory capacity of each shard, unit is MiB. The valid value range is as fallows: When the value of `ShardedCluster` is 0: 256, 1024, 2048, 4096, 8192, 16384, 32768, 65536. When the value of `ShardedCluster` is 1: 1024, 2048, 4096, 8192, 16384. When the value of `node_number` is 1, the value of this field can not be 256.",
			},
			"engine_version": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"4.0", "5.0", "6.0"}, false),
				Description:  "The engine version of redis instance. Valid value: `4.0`, `5.0`, `6.0`.",
			},
			"charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "PostPaid",
				ValidateFunc: validation.StringInSlice([]string{"PostPaid", "PrePaid"}, false),
				Description:  "The charge type of redis instance. Valid value: `PostPaid`, `PrePaid`.",
			},
			"purchase_months": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          1,
				ValidateFunc:     validation.IntInSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36}),
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
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				Default:      6379,
				ValidateFunc: validation.IntBetween(1024, 65535),
				Description:  "The port of custom define private network address. The valid value range is `1024-65535`. The default value is `6379`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The project name to which the redis instance belongs, if this parameter is empty, the new redis instance will not be added to any project.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "disabled",
				ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
				Description:  "Whether enable deletion protection for redis instance. Valid values: `enabled`, `disabled`(default).",
			},
			"create_backup": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				Description:      "Whether to create a final backup when modify the instance configuration or destroy the redis instance.",
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
			},
			"apply_immediately": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				Description:      "Whether to apply the instance configuration change operation immediately. The value of this field is false, means that the change operation will be applied within maintenance time.",
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
			},
			"vpc_auth_mode": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Description:      "Whether to enable password-free access when connecting to an instance through a private network. Valid values: `open`, `close`. Works only on modified scenes.",
				ValidateFunc:     validation.StringInSlice([]string{"open", "close"}, true),
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
			},
			"param_values": {
				Type:             schema.TypeSet,
				Optional:         true,
				Set:              paramHash,
				Description:      "The configuration item information to be modified. This field can only be added or modified. Deleting this field is invalid.",
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
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
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         schema.HashInt,
				Description: "The backup period. The valid value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on. \nThis field is valid and required when updating the backup plan of primary and secondary instance.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
			},
			"backup_hour": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description:      "The time period to start performing the backup. The valid value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on. \nThis field is valid and required when updating the backup plan of primary and secondary instance.",
			},
			"backup_active": {
				Type:             schema.TypeBool,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: redisInstanceImportDiffSuppress,
				Description:      "Whether enable auto backup for redis instance. This field is valid and required when updating the backup plan of primary and secondary instance.",
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

package instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongodb instance can be imported using the id, e.g.
```
$ terraform import volcengine_mongodb_instance.default mongo-replica-e405f8e2****
```

*/

func ResourceVolcengineMongoDBInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBInstanceCreate,
		Read:   resourceVolcengineMongoDBInstanceRead,
		Update: resourceVolcengineMongoDBInstanceUpdate,
		Delete: resourceVolcengineMongoDBInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Deprecated:  "This field has been deprecated after version-0.0.156. Please use `zone_ids` to deploy multiple availability zones.",
				Description: "The zone ID of instance.",
			},
			"zone_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of zone ids. If you need to deploy multiple availability zones for a newly created instance, you can specify three availability zone IDs at the same time. By default, the first available zone passed in is the primary available zone, and the two available zones passed in afterwards are the backup available zones.",
			},
			// 固定值，暂时不开放
			// "db_engine": {
			// 	Type:         schema.TypeString,
			// 	Optional:     true,
			// 	Computed:     true,
			// 	ValidateFunc: validation.StringInSlice([]string{"MongoDB"}, false),
			// 	Description:  "The db engine,valid value contains `MongoDB`.",
			// },
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The version of db engine, valid value contains `MongoDB_4_0`, `MongoDB_4_2`, `MongoDB_4_4`, `MongoDB_5_0`, `MongoDB_6_0`.",
			},
			"node_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The spec of node. When the instance_type is ReplicaSet, this parameter represents the computing node specification of the replica set instance. When the instance_type is ShardedCluster, this parameter represents the specification of the Shard node.",
			},
			// 固定值，暂时不开放
			// "node_number": {
			// 	Type:        schema.TypeInt,
			// 	Optional:    true,
			// 	Computed:    true,
			// 	Description: "If `InstanceType` is `ReplicaSet`,this parameter indicates the number of compute nodes of the replica set instance,if `InstanceType` is `ShardedCluster`,this parameter indicates the number of nodes in each shard.",
			// },
			"instance_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				//Default:      "ReplicaSet",
				Description:  "The type of instance, the valid value contains `ReplicaSet` or `ShardedCluster`. Default is `ReplicaSet`.",
				ValidateFunc: validation.StringInSlice([]string{"ReplicaSet", "ShardedCluster"}, false),
			},
			"mongos_node_spec": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 当实例类型为分片集群（即InstanceType取值为ShardedCluster）时，该参数必填。
					return d.Get("instance_type") == "ReplicaSet"
				},
				Description: "The mongos node spec of shard cluster, this parameter is required when the `InstanceType` is `ShardedCluster`.",
			},
			"mongos_node_number": {
				Type:     schema.TypeInt,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("instance_type") == "ReplicaSet"
				},
				Description: "The mongos node number of shard cluster, value range is `2~23`, this parameter is required when the `InstanceType` is `ShardedCluster`.",
			},
			"shard_number": {
				Type:     schema.TypeInt,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("instance_type") == "ReplicaSet"
				},
				Description: "The number of shards in shard cluster, value range is `2~32`, this parameter is required when the `InstanceType` is `ShardedCluster`.",
			},
			"storage_space_gb": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The total storage space of a replica set instance, or the storage space of a single shard in a sharded cluster. Unit: GiB.",
			},
			"config_server_node_spec": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("instance_type") == "ReplicaSet"
				},
				Description: "The config server node spec of shard cluster. Default is `mongo.config.1c2g`. This parameter is only effective when the `InstanceType` is `ShardedCluster`. \n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"config_server_storage_space_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("instance_type") == "ReplicaSet"
				},
				Description: "The config server storage space of shard cluster, Unit: GiB. Default is 20. This parameter is only effective when the `InstanceType` is `ShardedCluster`. \n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The vpc ID.",
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(i interface{}, s string) ([]string, []error) {
					return validation.StringIsNotEmpty(i, s)
				},
				Description: "The subnet id of instance.",
			},
			//"super_account_name": {
			//	Type:         schema.TypeString,
			//	ValidateFunc: validation.StringInSlice([]string{"root"}, false),
			//	Default:      "root",
			//	Optional:     true,
			//	Description:  "The name of database account,must be `root`.",
			//},
			"super_account_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password of database account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The instance name.",
			},
			"charge_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				//Default:      "PostPaid",
				ValidateFunc: validation.StringInSlice([]string{"Prepaid", "PostPaid"}, false),
				Description:  "The charge type of instance, valid value contains `Prepaid` or `PostPaid`. Default is `PostPaid`.",
			},
			"auto_renew": {
				Type:             schema.TypeBool,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: MongoDBInstanceImportDiffSuppress,
				Description:      "Whether to enable automatic renewal. This parameter is required when the `ChargeType` is `Prepaid`.",
			},
			"period_unit": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringInSlice([]string{"Year", "Month"}, false),
				DiffSuppressFunc: MongoDBInstanceImportDiffSuppress,
				Description:      "The period unit, valid value contains `Year` or `Month`. This parameter is required when the `ChargeType` is `Prepaid`.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: MongoDBInstanceImportDiffSuppress,
				Description:      "The instance purchase duration, the value range is `1~3` when `PeriodUtil` is `Year`, the value range is `1~9` when `PeriodUtil` is `Month`. This parameter is required when the `ChargeType` is `Prepaid`.",
			},
			"node_availability_zone": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "The readonly node of the instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The zone id of readonly nodes.",
						},
						"node_number": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
							Description: "The number of readonly nodes in current zone. Currently, only ReplicaSet instances and Shard in ShardedCluster instances support adding readonly nodes.\n" +
								"When the instance_type is ReplicaSet, this value represents the total number of readonly nodes in a single replica set instance. Each instance of the replica set supports adding up to 5 readonly nodes.\n" +
								"When the instance_type is ShardedCluster, this value represents the number of readonly nodes in each shard. Each shard can add up to 5 readonly nodes.",
						},
					},
				},
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name to which the instance belongs.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"private_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The private endpoint address of instance.",
			},
			"read_only_node_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of readonly node in instance.",
			},
			"shards": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The shards information of the ShardedCluster instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shard_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The shard id.",
						},
					},
				},
			},
			"config_servers_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The config servers id of the ShardedCluster instance.",
			},
			"mongos_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The mongos id of the ShardedCluster instance.",
			},
			"mongos": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The mongos information of the ShardedCluster instance.",
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
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineMongoDBInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("Error on creating instance %q,%s", d.Id(), err)
	}
	return resourceVolcengineMongoDBInstanceRead(d, meta)
}

func resourceVolcengineMongoDBInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("error on updating instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineMongoDBInstanceRead(d, meta)
}

func resourceVolcengineMongoDBInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("error on deleting instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongoDBInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("Error on reading instance %q,%s", d.Id(), err)
	}
	return err
}

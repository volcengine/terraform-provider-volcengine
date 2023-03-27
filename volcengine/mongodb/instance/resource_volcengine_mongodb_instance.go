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
				Description: "The zone ID of instance.",
			},
			// 固定值，暂时不开放
			// "db_engine": {
			// 	Type:         schema.TypeString,
			// 	Optional:     true,
			// 	Computed:     true,
			// 	ValidateFunc: validation.StringInSlice([]string{"MongoDB"}, false),
			// 	Description:  "The db engine,valid value contains `MongoDB`.",
			// },
			// "db_engine_version": {
			// 	Type:         schema.TypeString,
			// 	Optional:     true,
			// 	Computed:     true,
			// 	ValidateFunc: validation.StringInSlice([]string{"MongoDB_4_0"}, false),
			// 	Description:  "The version of db engine,valid value contains `MongoDB_4_0`.",
			// },
			"node_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The spec of node.",
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
				Description:  "The type of instance,the valid value contains `ReplicaSet` or `ShardedCluster`.",
				ValidateFunc: validation.StringInSlice([]string{"ReplicaSet", "ShardedCluster"}, false),
			},
			"mongos_node_spec": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The mongos node spec of shard cluster, this parameter is required when `InstanceType` is `ShardedCluster`.",
			},
			"mongos_node_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     2,
				Description: "The mongos node number of shard cluster,value range is `2~23`, this parameter is required when `InstanceType` is `ShardedCluster`.",
			},
			"shard_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The number of shards in shard cluster,value range is `2~32`, this parameter is required when `InstanceType` is `ShardedCluster`.",
			},
			"storage_space_gb": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The total storage space of a replica set instance, or the storage space of a single shard in a sharded cluster, in GiB.",
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
				Description: "The password of database account.",
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
				Description:  "The charge type of instance, valid value contains `Prepaid` or `PostPaid`.",
			},
			"auto_renew": {
				Type:             schema.TypeBool,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: MongoDBInstanceImportDiffSuppress,
				Description:      "Whether to enable automatic renewal.",
			},
			"period_unit": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringInSlice([]string{"Year", "Month"}, false),
				DiffSuppressFunc: MongoDBInstanceImportDiffSuppress,
				Description:      "The period unit,valid value contains `Year` or `Month`, this parameter is required when `ChargeType` is `Prepaid`.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: MongoDBInstanceImportDiffSuppress,
				Description:      "The instance purchase duration,the value range is `1~3` when `PeriodUtil` is `Year`, the value range is `1~9` when `PeriodUtil` is `Month`, this parameter is required when `ChargeType` is `Prepaid`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name to which the instance belongs.",
			},
			"tags": ve.TagsSchema(),
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

package instance

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongosdb instance can be imported using the id, e.g.
```
$ terraform import volcengine_mongosdb_instance.default mongo-replica-e405f8e2****
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
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The zone ID of instance.",
			},
			"db_engine": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"MongoDB"}, false),
				Description:  "The db engine.",
			},
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of db engine.",
			},
			"node_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The spec of node.",
			},
			"node_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of node.",
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ReplicaSet",
				Description:  "The type of instance.",
				ValidateFunc: validation.StringInSlice([]string{"ReplicaSet", "ShardedCluster"}, false),
			},
			"mongos_node_spec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The mongos node spec of shard cluster.",
			},
			"mongos_node_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The mongos node number of shard cluster.",
			},
			"shard_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of shards in shard cluster.",
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
				Description: "The vpc ID.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The subnet id of instance.",
			},
			"supper_account_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"root"}, false),
				Description:  "The name of database account.",
			},
			"supper_account_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The passwotd of database account.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance name.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "PostPaid",
				Description: "The charge type of instance.",
			},
			"auto_renew": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable automatic renewal.",
			},
			"period_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The period unit.",
				ValidateFunc: validation.StringInSlice([]string{"Year", "Month"}, false),
			},
			"period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The instance purchase duration.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name to which the instance belongs.",
			},
		},
	}
	return resource
}

func resourceVolcengineMongoDBInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("Error on creating instance %q,%s", d.Id(), err)
	}
	return resourceVolcengineMongoDBInstanceRead(d, meta)
}

func resourceVolcengineMongoDBInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("error on updating instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineMongoDBInstanceRead(d, meta)
}

func resourceVolcengineMongoDBInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("error on deleting instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongoDBInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineMongoDBInstance())
	if err != nil {
		return fmt.Errorf("Error on reading instance %q,%s", d.Id(), err)
	}
	return err
}

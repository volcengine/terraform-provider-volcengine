package rds_mysql_instance_readonly_node

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Rds Mysql Instance Readonly Node can be imported using the instance_id:node_id, e.g.
```
$ terraform import volcengine_rds_mysql_instance_readonly_node.default mysql-72da4258c2c7:mysql-72da4258c2c7-r7f93
```

*/

func ResourceVolcengineRdsMysqlInstanceReadonlyNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsMysqlInstanceReadonlyNodeCreate,
		Read:   resourceVolcengineRdsMysqlInstanceReadonlyNodeRead,
		Update: resourceVolcengineRdsMysqlInstanceReadonlyNodeUpdate,
		Delete: resourceVolcengineRdsMysqlInstanceReadonlyNodeDelete,
		Importer: &schema.ResourceImporter{
			State: rdsMysqlInstanceReadonlyNodeImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The RDS mysql instance id of the readonly node.",
			},
			"node_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The specification of readonly node.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The available zone of readonly node.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the readonly node.",
			},
			"delay_replication_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The delay time of the readonly node.",
			},
		},
	}
}

func resourceVolcengineRdsMysqlInstanceReadonlyNodeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceReadonlyNodeService := NewRdsMysqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceReadonlyNodeService.Dispatcher.Create(rdsMysqlInstanceReadonlyNodeService, d, ResourceVolcengineRdsMysqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on creating RDS mysql instance readonly node %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlInstanceReadonlyNodeRead(d, meta)
}

func resourceVolcengineRdsMysqlInstanceReadonlyNodeRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceReadonlyNodeService := NewRdsMysqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceReadonlyNodeService.Dispatcher.Read(rdsMysqlInstanceReadonlyNodeService, d, ResourceVolcengineRdsMysqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on reading RDS mysql instance readonly node %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlInstanceReadonlyNodeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceReadonlyNodeService := NewRdsMysqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceReadonlyNodeService.Dispatcher.Update(rdsMysqlInstanceReadonlyNodeService, d, ResourceVolcengineRdsMysqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on updating RDS mysql instance readonly node %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlInstanceReadonlyNodeRead(d, meta)
}

func resourceVolcengineRdsMysqlInstanceReadonlyNodeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceReadonlyNodeService := NewRdsMysqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceReadonlyNodeService.Dispatcher.Delete(rdsMysqlInstanceReadonlyNodeService, d, ResourceVolcengineRdsMysqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on deleting RDS mysql instance readonly node %q, %w", d.Id(), err)
	}
	return err
}

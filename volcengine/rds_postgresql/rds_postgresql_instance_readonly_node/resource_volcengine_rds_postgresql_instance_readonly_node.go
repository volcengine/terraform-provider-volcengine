package rds_postgresql_instance_readonly_node

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlInstanceReadonlyNode can be imported using the instance_id:node_id, e.g.
```
$ terraform import volcengine_rds_postgresql_instance_readonly_node.default postgres-21a3333b****:postgres-ca7b7019****
```

*/

func ResourceVolcengineRdsPostgresqlInstanceReadonlyNode() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlInstanceReadonlyNodeCreate,
		Read:   resourceVolcengineRdsPostgresqlInstanceReadonlyNodeRead,
		Update: resourceVolcengineRdsPostgresqlInstanceReadonlyNodeUpdate,
		Delete: resourceVolcengineRdsPostgresqlInstanceReadonlyNodeDelete,
		Importer: &schema.ResourceImporter{
			State: rdsPostgreSQLInstanceReadonlyNodeImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The RDS PostgreSQL instance id of the readonly node.",
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
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlInstanceReadonlyNodeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_instance_readonly_node %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlInstanceReadonlyNodeRead(d, meta)
}

func resourceVolcengineRdsPostgresqlInstanceReadonlyNodeRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_instance_readonly_node %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlInstanceReadonlyNodeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_instance_readonly_node %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlInstanceReadonlyNodeRead(d, meta)
}

func resourceVolcengineRdsPostgresqlInstanceReadonlyNodeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceReadonlyNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlInstanceReadonlyNode())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_instance_readonly_node %q, %s", d.Id(), err)
	}
	return err
}

var rdsPostgreSQLInstanceReadonlyNodeImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("instance_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("node_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

package rds_postgresql_schema

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlSchema can be imported using the instance id, database name and schema name, e.g.
```
$ terraform import volcengine_rds_postgresql_schema.default instance_id:db_name:schema_name
```

*/

func ResourceVolcengineRdsPostgresqlSchema() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlSchemaCreate,
		Read:   resourceVolcengineRdsPostgresqlSchemaRead,
		Update: resourceVolcengineRdsPostgresqlSchemaUpdate,
		Delete: resourceVolcengineRdsPostgresqlSchemaDelete,
		Importer: &schema.ResourceImporter{
			State: schemaImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the postgresql instance.",
			},
			"db_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the database.",
			},
			"schema_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the schema.",
			},
			"owner": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The owner of the schema." +
					"The instance read-only account, a high-privilege account with DDL permissions disabled, " +
					"or a normal account with DDL permissions disabled cannot be used as the owner of the schema.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlSchemaCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlSchemaService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlSchema())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_schema %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlSchemaRead(d, meta)
}

func resourceVolcengineRdsPostgresqlSchemaRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlSchemaService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlSchema())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_schema %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlSchemaUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlSchemaService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlSchema())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_schema %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlSchemaRead(d, meta)
}

func resourceVolcengineRdsPostgresqlSchemaDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlSchemaService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlSchema())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_schema %q, %s", d.Id(), err)
	}
	return err
}

func schemaImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'instanceId:db_name:schema_name'")
	}
	instanceId := items[0]
	dbName := items[1]
	schemaName := items[2]
	_ = d.Set("instance_id", instanceId)
	_ = d.Set("db_name", dbName)
	_ = d.Set("schema_name", schemaName)
	return []*schema.ResourceData{d}, nil
}

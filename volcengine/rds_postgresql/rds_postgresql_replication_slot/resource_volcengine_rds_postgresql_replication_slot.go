package rds_postgresql_replication_slot

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlReplicationSlot can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_replication_slot.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlReplicationSlot() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlReplicationSlotCreate,
		Read:   resourceVolcengineRdsPostgresqlReplicationSlotRead,
		Delete: resourceVolcengineRdsPostgresqlReplicationSlotDelete,
		Importer: &schema.ResourceImporter{State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			parts := strings.Split(d.Id(), ":")
			if len(parts) < 2 {
				return []*schema.ResourceData{d}, fmt.Errorf("import id must be 'instance_id:slot_name'")
			}
			_ = d.Set("instance_id", parts[0])
			_ = d.Set("slot_name", parts[1])
			return []*schema.ResourceData{d}, nil
		}},
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
				Description: "The id of the PostgreSQL instance.",
			},
			"slot_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the slot.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlReplicationSlotCreate(d *schema.ResourceData, meta interface{}) (err error) {
	instanceId := d.Get("instance_id").(string)
	slotName := d.Get("slot_name").(string)
	d.SetId(fmt.Sprintf("%s:%s", instanceId, slotName))

	service := NewRdsPostgresqlReplicationSlotService(meta.(*ve.SdkClient))
	_, err = service.ReadResource(d, d.Id())
	if err != nil {
		return fmt.Errorf("replication_slot %s not exist or not readable: %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRdsPostgresqlReplicationSlotRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlReplicationSlotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlReplicationSlot())
	if err != nil {
		if ve.ResourceNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error on reading rds_postgresql_replication_slot %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRdsPostgresqlReplicationSlotDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlReplicationSlotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlReplicationSlot())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_replication_slot %q, %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

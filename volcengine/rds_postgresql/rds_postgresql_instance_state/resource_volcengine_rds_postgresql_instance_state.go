package rds_postgresql_instance_state

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlInstanceState can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_instance_state.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlInstanceState() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlInstanceStateCreate,
		Read:   resourceVolcengineRdsPostgresqlInstanceStateRead,
		Update: resourceVolcengineRdsPostgresqlInstanceStateUpdate,
		Delete: resourceVolcengineRdsPostgresqlInstanceStateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Restart"}, false),
				Description:  "The action to perform on the instance. Valid value: Restart.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the RDS PostgreSQL instance to perform the action on.",
			},
			"apply_scope": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "AllNode",
				ValidateFunc: validation.StringInSlice([]string{"AllNode", "CustomNode"}, false),
				Description:  "The scope of the action. Valid values: AllNode, CustomNode. Default value: AllNode.",
			},
			"custom_node_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The ID of the read-only node(s) to restart. " +
					"Required if apply_scope is CustomNode.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlInstanceStateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceStateService(meta.(*ve.SdkClient))
	if err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlInstanceState()); err != nil {
		return fmt.Errorf("error on creating rds_postgresql_instance_state %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRdsPostgresqlInstanceStateRead(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}

func resourceVolcengineRdsPostgresqlInstanceStateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}

func resourceVolcengineRdsPostgresqlInstanceStateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}

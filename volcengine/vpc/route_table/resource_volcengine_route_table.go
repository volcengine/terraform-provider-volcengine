package route_table

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Route table can be imported using the id, e.g.
```
$ terraform import volcengine_route_table.default vtb-274e0syt9av407fap8tle16kb
```

*/

func ResourceVolcengineRouteTable() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineRouteTableDelete,
		Create: resourceVolcengineRouteTableCreate,
		Read:   resourceVolcengineRouteTableRead,
		Update: resourceVolcengineRouteTableUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the VPC.",
			},
			"route_table_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the route table.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the route table.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the route table.",
			},
		},
	}
}

func resourceVolcengineRouteTableCreate(d *schema.ResourceData, meta interface{}) error {
	routeTableService := NewRouteTableService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(routeTableService, d, ResourceVolcengineRouteTable()); err != nil {
		return fmt.Errorf("error on creating route table  %q, %w", d.Id(), err)
	}
	return resourceVolcengineRouteTableRead(d, meta)
}

func resourceVolcengineRouteTableRead(d *schema.ResourceData, meta interface{}) error {
	routeTableService := NewRouteTableService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(routeTableService, d, ResourceVolcengineRouteTable()); err != nil {
		return fmt.Errorf("error on reading route table %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRouteTableUpdate(d *schema.ResourceData, meta interface{}) error {
	routeTableService := NewRouteTableService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(routeTableService, d, ResourceVolcengineRouteTable()); err != nil {
		return fmt.Errorf("error on updating route table %q, %w", d.Id(), err)
	}
	return resourceVolcengineRouteTableRead(d, meta)
}

func resourceVolcengineRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
	routeTableService := NewRouteTableService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(routeTableService, d, ResourceVolcengineRouteTable()); err != nil {
		return fmt.Errorf("error on deleting route table %q, %w", d.Id(), err)
	}
	return nil
}

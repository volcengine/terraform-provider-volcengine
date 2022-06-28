package scalinggroup_server_group

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
ScalinggroupServerGroup can be imported using the scaling_group_id, e.g.
```
$ terraform import vestack_scalinggroup_server_group.default scg-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVestackScalinggroupServerGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackScalinggroupServerGroupCreate,
		Read:   resourceVestackScalinggroupServerGroupRead,
		Update: resourceVetackScalinggroupServerGroupUpdate,
		Delete: resourceVetackScalinggroupServerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the scaling group.",
			},
			"server_group_attributes": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The load balancer server group attributes of the scaling group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:         schema.TypeInt,
							Required:     true,
							Description:  "The port receiving request of the server group.",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the server group.",
						},
						"weight": {
							Type:         schema.TypeInt,
							Required:     true,
							Description:  "The weight of the instance.",
							ValidateFunc: validation.IntBetween(0, 100),
						},
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The load balancer id.",
						},
					},
				},
				Set: serverGroupAttributeHash,
			},
		},
	}
	return resource
}

func resourceVestackScalinggroupServerGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewScalinggroupServerGroupService(meta.(*ve.SdkClient))
	err = serverGroupService.Dispatcher.Create(serverGroupService, d, ResourceVestackScalinggroupServerGroup())
	if err != nil {
		return fmt.Errorf("error on creating ScalinggroupServerGroup %q, %s", d.Id(), err)
	}
	return resourceVestackScalinggroupServerGroupRead(d, meta)
}

func resourceVestackScalinggroupServerGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewScalinggroupServerGroupService(meta.(*ve.SdkClient))
	err = serverGroupService.Dispatcher.Read(serverGroupService, d, ResourceVestackScalinggroupServerGroup())
	if err != nil {
		return fmt.Errorf("error on reading ScalinggroupServerGroup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVetackScalinggroupServerGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewScalinggroupServerGroupService(meta.(*ve.SdkClient))
	err = serverGroupService.Dispatcher.Update(serverGroupService, d, ResourceVestackScalinggroupServerGroup())
	if err != nil {
		return fmt.Errorf("error on updating ScalinggroupServerGroup %q, %s", d.Id(), err)
	}
	return resourceVestackScalinggroupServerGroupRead(d, meta)
}

func resourceVetackScalinggroupServerGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewScalinggroupServerGroupService(meta.(*ve.SdkClient))
	err = serverGroupService.Dispatcher.Delete(serverGroupService, d, ResourceVestackScalinggroupServerGroup())
	if err != nil {
		return fmt.Errorf("error on deleting ScalinggroupServerGroup %q, %s", d.Id(), err)
	}
	return err
}

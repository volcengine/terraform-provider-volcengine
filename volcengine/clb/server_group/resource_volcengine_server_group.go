package server_group

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ServerGroup can be imported using the id, e.g.
```
$ terraform import volcengine_server_group.default rsp-273yv0kir1vk07fap8tt9jtwg
```

*/

func ResourceVolcengineServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineServerGroupCreate,
		Read:   resourceVolcengineServerGroupRead,
		Update: resourceVolcengineServerGroupUpdate,
		Delete: resourceVolcengineServerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"server_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID of the ServerGroup.",
			},
			"load_balancer_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the Clb.",
			},
			"server_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the ServerGroup.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of ServerGroup.",
			},
			"address_ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "ipv4",
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
				Description:  "The address ip version of the ServerGroup. Valid values: `ipv4`, `ipv6`. Default is `ipv4`.",
			},
		},
	}
}

func resourceVolcengineServerGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewServerGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(serverGroupService, d, ResourceVolcengineServerGroup())
	if err != nil {
		return fmt.Errorf("error on creating serverGroup  %q, %w", d.Id(), err)
	}
	return resourceVolcengineServerGroupRead(d, meta)
}

func resourceVolcengineServerGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewServerGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(serverGroupService, d, ResourceVolcengineServerGroup())
	if err != nil {
		return fmt.Errorf("error on reading serverGroup %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineServerGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewServerGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(serverGroupService, d, ResourceVolcengineServerGroup())
	if err != nil {
		return fmt.Errorf("error on updating serverGroup  %q, %w", d.Id(), err)
	}
	return resourceVolcengineServerGroupRead(d, meta)
}

func resourceVolcengineServerGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	serverGroupService := NewServerGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(serverGroupService, d, ResourceVolcengineServerGroup())
	if err != nil {
		return fmt.Errorf("error on deleting serverGroup %q, %w", d.Id(), err)
	}
	return err
}

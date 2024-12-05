package rocketmq_allow_list

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RocketmqAllowList can be imported using the id, e.g.
```
$ terraform import volcengine_rocketmq_allow_list.default resource_id
```

*/

func ResourceVolcengineRocketmqAllowList() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRocketmqAllowListCreate,
		Read:   resourceVolcengineRocketmqAllowListRead,
		Update: resourceVolcengineRocketmqAllowListUpdate,
		Delete: resourceVolcengineRocketmqAllowListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allow_list_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the allow list.",
			},
			"allow_list": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of ip addresses. Enter an IP address or a range of IP addresses in CIDR format.",
			},

			// computed fields
			"allow_list_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the rocketmq allow list.",
			},
			"allow_list_ip_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of ip address in the rocketmq allow list.",
			},
			"associated_instance_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of the rocketmq instances associated with the allow list.",
			},
			"associated_instances": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The associated instance information of the allow list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rocketmq instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rocketmq instance.",
						},
						"vpc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the rocketmq instance.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRocketmqAllowListCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRocketmqAllowList())
	if err != nil {
		return fmt.Errorf("error on creating rocketmq_allow_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqAllowListRead(d, meta)
}

func resourceVolcengineRocketmqAllowListRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRocketmqAllowList())
	if err != nil {
		return fmt.Errorf("error on reading rocketmq_allow_list %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRocketmqAllowListUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRocketmqAllowList())
	if err != nil {
		return fmt.Errorf("error on updating rocketmq_allow_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqAllowListRead(d, meta)
}

func resourceVolcengineRocketmqAllowListDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRocketmqAllowList())
	if err != nil {
		return fmt.Errorf("error on deleting rocketmq_allow_list %q, %s", d.Id(), err)
	}
	return err
}

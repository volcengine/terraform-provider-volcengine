package waf_ip_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafIpGroup can be imported using the id, e.g.
```
$ terraform import volcengine_waf_ip_group.default resource_id
```

*/

func ResourceVolcengineWafIpGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafIpGroupCreate,
		Read:   resourceVolcengineWafIpGroupRead,
		Update: resourceVolcengineWafIpGroupUpdate,
		Delete: resourceVolcengineWafIpGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of ip group.",
			},
			"add_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The way of addition.",
			},
			"ip_list": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The IP address to be added.",
			},
			"ip_group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the ip group.",
			},
			"ip_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of IP addresses within the address group.",
			},
			"related_rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of associated rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rule.",
						},
						"rule_tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the rule.",
						},
						"rule_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the rule.",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The information of the protected domain names associated with the rules.",
						},
					},
				},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ip group update time.",
			},
		},
	}
	return resource
}

func resourceVolcengineWafIpGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafIpGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafIpGroup())
	if err != nil {
		return fmt.Errorf("error on creating waf_ip_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafIpGroupRead(d, meta)
}

func resourceVolcengineWafIpGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafIpGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafIpGroup())
	if err != nil {
		return fmt.Errorf("error on reading waf_ip_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineWafIpGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafIpGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafIpGroup())
	if err != nil {
		return fmt.Errorf("error on updating waf_ip_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafIpGroupRead(d, meta)
}

func resourceVolcengineWafIpGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafIpGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafIpGroup())
	if err != nil {
		return fmt.Errorf("error on deleting waf_ip_group %q, %s", d.Id(), err)
	}
	return err
}

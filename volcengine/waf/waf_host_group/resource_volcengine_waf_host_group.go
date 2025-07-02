package waf_host_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafHostGroup can be imported using the id, e.g.
```
$ terraform import volcengine_waf_host_group.default resource_id
```

*/

func ResourceVolcengineWafHostGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafHostGroupCreate,
		Read:   resourceVolcengineWafHostGroupRead,
		Update: resourceVolcengineWafHostGroupUpdate,
		Delete: resourceVolcengineWafHostGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Domain name group description.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project of Domain name group.",
			},
			"host_list": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Domain names that need to be added to this domain name group.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the domain name group.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Domain name list modification action. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"host_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of domain names contained in the domain name group.",
			},
			"host_group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the domain name group.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Domain name group update time.",
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
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineWafHostGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafHostGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafHostGroup())
	if err != nil {
		return fmt.Errorf("error on creating waf_host_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafHostGroupRead(d, meta)
}

func resourceVolcengineWafHostGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafHostGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafHostGroup())
	if err != nil {
		return fmt.Errorf("error on reading waf_host_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineWafHostGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafHostGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafHostGroup())
	if err != nil {
		return fmt.Errorf("error on updating waf_host_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafHostGroupRead(d, meta)
}

func resourceVolcengineWafHostGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafHostGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafHostGroup())
	if err != nil {
		return fmt.Errorf("error on deleting waf_host_group %q, %s", d.Id(), err)
	}
	return err
}

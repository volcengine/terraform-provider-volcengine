package private_zone_resolver_rule

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateZoneResolverRule can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_resolver_rule.default resource_id
```

*/

func ResourceVolcenginePrivateZoneResolverRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateZoneResolverRuleCreate,
		Read:   resourceVolcenginePrivateZoneResolverRuleRead,
		Update: resourceVolcenginePrivateZoneResolverRuleUpdate,
		Delete: resourceVolcenginePrivateZoneResolverRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
		    // TODO: Add all your arguments and attributes.
			"replace_with_arguments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// TODO: See setting, getting, flattening, expanding examples below for this complex argument.
			"complex_argument": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_field_one": {
							Type:         schema.TypeString,
							Required:     true,
						},
						"sub_field_two": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateZoneResolverRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on creating private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneResolverRuleRead(d, meta)
}

func resourceVolcenginePrivateZoneResolverRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on reading private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateZoneResolverRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on updating private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneResolverRuleRead(d, meta)
}

func resourceVolcenginePrivateZoneResolverRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on deleting private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return err
}

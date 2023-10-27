package transit_router_grant_rule

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterGrantRule can be imported using the transit router id and accountId, e.g.
```
$ terraform import volcengine_transit_router_grant_rule.default trId:accountId
```

*/

func ResourceVolcengineTransitRouterGrantRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterGrantRuleCreate,
		Read:   resourceVolcengineTransitRouterGrantRuleRead,
		Update: resourceVolcengineTransitRouterGrantRuleUpdate,
		Delete: resourceVolcengineTransitRouterGrantRuleDelete,
		Importer: &schema.ResourceImporter{
			State: trGrantRuleImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the transit router.",
			},
			"grant_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Account ID awaiting authorization for intermediate router instance.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the rule.",
			},
		},
	}
	return resource
}

func resourceVolcengineTransitRouterGrantRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterGrantRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterGrantRule())
	if err != nil {
		return fmt.Errorf("error on creating transit_router_grant_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterGrantRuleRead(d, meta)
}

func resourceVolcengineTransitRouterGrantRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterGrantRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterGrantRule())
	if err != nil {
		return fmt.Errorf("error on reading transit_router_grant_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterGrantRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterGrantRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterGrantRule())
	if err != nil {
		return fmt.Errorf("error on updating transit_router_grant_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterGrantRuleRead(d, meta)
}

func resourceVolcengineTransitRouterGrantRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTransitRouterGrantRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterGrantRule())
	if err != nil {
		return fmt.Errorf("error on deleting transit_router_grant_rule %q, %s", d.Id(), err)
	}
	return err
}

var trGrantRuleImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("transit_router_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("grant_account_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

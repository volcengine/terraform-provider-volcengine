package rule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Rule can be imported using the id, e.g.
Notice: resourceId is ruleId, due to the lack of describeRuleAttributes in openapi, for import resources, please use ruleId:listenerId to import.
we will fix this problem later.
```
$ terraform import vestack_clb_rule.foo rule-273zb9hzi1gqo7fap8u1k3utb:lsn-273ywvnmiu70g7fap8u2xzg9d
```

*/

func resourceParseId(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of Id (%s), expected ruleId:listenerId", id)
	}
	return parts[0], parts[1], nil
}

func ResourceVestackRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackRuleCreate,
		Read:   resourceVestackRuleRead,
		Update: resourceVestackRuleUpdate,
		Delete: resourceVestackRuleDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				ruleId, listenerId, err := resourceParseId(d.Id())
				if err != nil {
					return nil, err
				}
				d.Set("listener_id", listenerId)
				d.SetId(ruleId)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of listener.",
			},
			"domain": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"domain", "url"},
				Description:  "The domain of Rule.",
			},
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"domain", "url"},
				// 若指定Domain，则Url不传入数值时，默认为“/”
				Default:     "/",
				Description: "The Url of Rule.",
			},
			"server_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Server Group Id.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Rule.",
			},
		},
	}
}

func resourceVestackRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ruleService.Dispatcher.Create(ruleService, d, ResourceVestackRule())
	if err != nil {
		return fmt.Errorf("error on creating rule %q, %w", d.Id(), err)
	}
	return resourceVestackRuleRead(d, meta)
}

func resourceVestackRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ruleService.Dispatcher.Read(ruleService, d, ResourceVestackRule())
	if err != nil {
		return fmt.Errorf("error on reading rule %q, %w", d.Id(), err)
	}
	return err
}

func resourceVestackRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ruleService.Dispatcher.Update(ruleService, d, ResourceVestackRule())
	if err != nil {
		return fmt.Errorf("error on updating rule %q, %w", d.Id(), err)
	}
	return resourceVestackRuleRead(d, meta)
}

func resourceVestackRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ruleService.Dispatcher.Delete(ruleService, d, ResourceVestackRule())
	if err != nil {
		return fmt.Errorf("error on deleting rule %q, %w", d.Id(), err)
	}
	return err
}

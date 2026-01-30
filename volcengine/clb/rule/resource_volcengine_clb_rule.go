package rule

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Rule can be imported using the id, e.g.
Notice: resourceId is ruleId, due to the lack of describeRuleAttributes in openapi, for import resources, please use ruleId:listenerId to import.
we will fix this problem later.
```
$ terraform import volcengine_clb_rule.foo rule-273zb9hzi1gqo7fap8u1k3utb:lsn-273ywvnmiu70g7fap8u2xzg9d
```

*/

func resourceParseId(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of Id (%s), expected ruleId:listenerId", id)
	}
	return parts[0], parts[1], nil
}

func ResourceVolcengineRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRuleCreate,
		Read:   resourceVolcengineRuleRead,
		Update: resourceVolcengineRuleUpdate,
		Delete: resourceVolcengineRuleDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				ruleId, listenerId, err := resourceParseId(d.Id())
				if err != nil {
					return nil, err
				}
				_ = d.Set("listener_id", listenerId)
				d.SetId(ruleId)
				return []*schema.ResourceData{d}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
			"action_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Forward", // 默认动作是 Forward：转发至
				ValidateFunc: validation.StringInSlice([]string{"Forward", "Redirect"}, false),
				Description:  "The action type of Rule, valid values: `Forward`, `Redirect`.",
			},
			"server_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Server Group Id. Required when action_type is Forward.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Rule.",
			},
			"tags": ve.TagsSchema(),
			"redirect_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"HTTP", "HTTPS"}, false),
							Default:      "HTTPS",
							Description:  "The redirect protocol. Valid values: `HTTP`, `HTTPS`.",
						},
						"host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The redirect host, i.e. the domain name redirected by the rule.",
						},
						"path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The redirect path.",
						},
						"port": {
							Type:     schema.TypeString,
							Optional: true,
							// ValidateFunc: validation.StringMatch(regexp.MustCompile("^[1-9]\\d{0,4}$"), "must be a valid port"),
							Description: "The redirect port, valid range: 1~65535.",
						},
						"status_code": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"301", "302", "307", "308"}, false),
							Default:      "301",
							Description:  "The redirect status code. Valid values: 301, 302, 307, 308.",
						},
					},
				},
				Description: "The redirect configuration. Required when action_type is `Redirect`.",
			},
		},
	}
}

func resourceVolcengineRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(ruleService, d, ResourceVolcengineRule())
	if err != nil {
		return fmt.Errorf("error on creating rule %q, %w", d.Id(), err)
	}
	return resourceVolcengineRuleRead(d, meta)
}

func resourceVolcengineRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(ruleService, d, ResourceVolcengineRule())
	if err != nil {
		return fmt.Errorf("error on reading rule %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(ruleService, d, ResourceVolcengineRule())
	if err != nil {
		return fmt.Errorf("error on updating rule %q, %w", d.Id(), err)
	}
	return resourceVolcengineRuleRead(d, meta)
}

func resourceVolcengineRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(ruleService, d, ResourceVolcengineRule())
	if err != nil {
		return fmt.Errorf("error on deleting rule %q, %w", d.Id(), err)
	}
	return err
}

package alb_rule

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AlbRule can be imported using the listener id and rule id, e.g.
```
$ terraform import volcengine_alb_rule.default lsn-273yv0mhs5xj47fap8sehiiso:rule-****
```

*/

func ResourceVolcengineAlbRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbRuleCreate,
		Read:   resourceVolcengineAlbRuleRead,
		Update: resourceVolcengineAlbRuleUpdate,
		Delete: resourceVolcengineAlbRuleDelete,
		Importer: &schema.ResourceImporter{
			State: importAlbRule,
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
				Computed:     true,
				AtLeastOneOf: []string{"domain", "url"},
				Description:  "The domain of Rule.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of rule.",
			},
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"domain", "url"},
				Computed:     true,
				Description:  "The Url of Rule.",
			},
			"rule_action": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The forwarding rule action, if this parameter is empty(`\"\"`), forward to server group, if value is `Redirect`, will redirect.",
			},
			"server_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("rule_action").(string) == "Redirect"
				},
				Description: "Server group ID, this parameter is required if `rule_action` is empty.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Rule.",
			},
			"traffic_limit_enabled": {
				Type:        schema.TypeString,
				Default:     "off",
				Optional:    true,
				Description: "Forwarding rule QPS rate limiting switch:\n on: enable.\n off: disable (default).",
			},
			"traffic_limit_qps": {
				Type:     schema.TypeInt,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("traffic_limit_enabled").(string) == "off"
				},
				Description: "When Rules.N.TrafficLimitEnabled is turned on, this field is required. " +
					"Requests per second. Valid values are between 100 and 100000.",
			},
			"rewrite_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("rule_action").(string) != "Redirect"
				},
				Default: "off",
				Description: "Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). " +
					"Only available for whitelist users, please submit an application to experience. " +
					"Supported values are as follows:\non: enable.\noff: disable.",
			},
			"redirect_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("rule_action").(string) != "Redirect"
				},
				Description: "The redirect related configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redirect_domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The redirect domain, only support exact domain name.",
						},
						"redirect_uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The redirect URI.",
						},
						"redirect_port": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The redirect port.",
						},
						"redirect_http_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "301",
							Description: "The redirect http code, support 301(default), 302, 307, 308.",
						},
						"redirect_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "HTTPS",
							Description: "The redirect protocol, support HTTP, HTTPS(default).",
						},
					},
				},
			},
			"rewrite_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("rewrite_enabled").(string) == "off"
				},
				Description: "The list of rewrite configurations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rewrite_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Rewrite path.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineAlbRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbRule())
	if err != nil {
		return fmt.Errorf("error on creating alb_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbRuleRead(d, meta)
}

func resourceVolcengineAlbRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbRule())
	if err != nil {
		return fmt.Errorf("error on reading alb_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlbRule())
	if err != nil {
		return fmt.Errorf("error on updating alb_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbRuleRead(d, meta)
}

func resourceVolcengineAlbRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbRule())
	if err != nil {
		return fmt.Errorf("error on deleting alb_rule %q, %s", d.Id(), err)
	}
	return err
}

func importAlbRule(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form listenerId:ruleId")
	}
	err = data.Set("listener_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("rule_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

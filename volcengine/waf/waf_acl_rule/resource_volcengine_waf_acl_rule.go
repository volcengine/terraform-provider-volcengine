package waf_acl_rule

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafAclRule can be imported using the id, e.g.
```
$ terraform import volcengine_waf_acl_rule.default resource_id:AclType
```

*/

func ResourceVolcengineWafAclRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafAclRuleCreate,
		Read:   resourceVolcengineWafAclRuleRead,
		//Update: resourceVolcengineWafAclRuleUpdate,
		Delete: resourceVolcengineWafAclRuleDelete,
		Importer: &schema.ResourceImporter{
			State: wafAclRuleImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Action to be taken on requests that match the rule.",
			},
			"host_group_id": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Description: "The ID of the domain group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Rule description.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Rule name.",
			},
			"ip_location_country": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Country or region code.",
			},
			"ip_location_subregion": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Domestic region code.",
			},
			"accurate_group": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Advanced conditions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accurate_rules": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: "Details of advanced conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_obj": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "The HTTP object to be added to the advanced conditions.",
									},
									"obj_type": {
										Type:        schema.TypeInt,
										Required:    true,
										ForceNew:    true,
										Description: "The matching field for HTTP objects.",
									},
									"opretar": {
										Type:        schema.TypeInt,
										Required:    true,
										ForceNew:    true,
										Description: "The logical operator for the condition.",
									},
									"property": {
										Type:        schema.TypeInt,
										Required:    true,
										ForceNew:    true,
										Description: "Operate the properties of the http object.",
									},
									"value_string": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "The value to be matched.",
									},
								},
							},
						},
						"logic": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "The logical relationship of advanced conditions.",
						},
					},
				},
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The path of Matching.",
			},
			"ip_add_type": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Type of IP address addition.",
			},
			"host_add_type": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Type of domain name addition.",
			},
			"enable": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Whether to enable the rule.",
			},
			"advanced": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Whether to set advanced conditions.",
			},
			"acl_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of access control rules.",
			},
			"host_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Required if HostAddType = 3. Single or multiple domain names are supported.",
			},
			"ip_group_id": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Description: "Required if IpAddType = 2.",
			},
			"ip_list": {
				Type:     schema.TypeSet,
				MinItems: 2,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Required if IpAddType = 3. Single or multiple IP addresses are supported.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the project to which your domain names belong.",
			},
			"host_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of domain name groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of host group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Name of host group.",
						},
					},
				},
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rule unique identifier.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update time of the rule.",
			},
			"ip_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of domain name groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the IP address group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Name of the IP address group.",
						},
					},
				},
			},
			"client_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP address.",
			},
		},
	}
	return resource
}

func resourceVolcengineWafAclRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafAclRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafAclRule())
	if err != nil {
		return fmt.Errorf("error on creating waf_acl_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafAclRuleRead(d, meta)
}

func resourceVolcengineWafAclRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafAclRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafAclRule())
	if err != nil {
		return fmt.Errorf("error on reading waf_acl_rule %q, %s", d.Id(), err)
	}
	return err
}

//func resourceVolcengineWafAclRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
//	service := NewWafAclRuleService(meta.(*ve.SdkClient))
//	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafAclRule())
//	if err != nil {
//		return fmt.Errorf("error on updating waf_acl_rule %q, %s", d.Id(), err)
//	}
//	return resourceVolcengineWafAclRuleRead(d, meta)
//}

func resourceVolcengineWafAclRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafAclRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafAclRule())
	if err != nil {
		return fmt.Errorf("error on deleting waf_acl_rule %q, %s", d.Id(), err)
	}
	return err
}

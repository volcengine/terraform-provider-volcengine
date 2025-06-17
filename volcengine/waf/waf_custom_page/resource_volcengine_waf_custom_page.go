package waf_custom_page

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafCustomPage can be imported using the id, e.g.
```
$ terraform import volcengine_waf_custom_page.default resource_id:Host
```

*/

func ResourceVolcengineWafCustomPage() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafCustomPageCreate,
		Read:   resourceVolcengineWafCustomPageRead,
		Update: resourceVolcengineWafCustomPageUpdate,
		Delete: resourceVolcengineWafCustomPageDelete,
		Importer: &schema.ResourceImporter{
			State: wafCustomPageImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain name to be protected.",
			},
			"policy": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Action to be taken on requests that match the rule.",
			},
			"client_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Fill in ALL, which means this rule will take effect on all IP addresses.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Rule name.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Rule description.",
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Match the path.",
			},
			"enable": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Whether to enable the rule.",
			},
			"code": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Custom HTTP code returned when the request is blocked. Required if PageMode=0 or 1.",
			},
			"page_mode": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The layout template of the response page.",
			},
			"content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The layout template of the response page. Required if PageMode=0 or 1.",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The layout content of the response page.",
			},
			"advanced": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to configure advanced conditions.",
			},
			"redirect_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The path where users should be redirected.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project to which your domain names belong.",
			},
			"accurate": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "Advanced conditions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accurate_rules": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Details of advanced conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_obj": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The HTTP object to be added to the advanced conditions.",
									},
									"obj_type": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The matching field for HTTP objects.",
									},
									"opretar": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The logical operator for the condition.",
									},
									"property": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Operate the properties of the http object.",
									},
									"value_string": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value to be matched.",
									},
								},
							},
						},
						"logic": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The logical relationship of advanced conditions.",
						},
					},
				},
			},
			"group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the advanced conditional rule group.",
			},
			"header": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Request header information.",
			},
			"isolation_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of Region.",
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identification of the rules.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rule update time.",
			},
		},
	}
	return resource
}

func resourceVolcengineWafCustomPageCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCustomPageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafCustomPage())
	if err != nil {
		return fmt.Errorf("error on creating waf_custom_page %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafCustomPageRead(d, meta)
}

func resourceVolcengineWafCustomPageRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCustomPageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafCustomPage())
	if err != nil {
		return fmt.Errorf("error on reading waf_custom_page %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineWafCustomPageUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCustomPageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafCustomPage())
	if err != nil {
		return fmt.Errorf("error on updating waf_custom_page %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafCustomPageRead(d, meta)
}

func resourceVolcengineWafCustomPageDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCustomPageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafCustomPage())
	if err != nil {
		return fmt.Errorf("error on deleting waf_custom_page %q, %s", d.Id(), err)
	}
	return err
}

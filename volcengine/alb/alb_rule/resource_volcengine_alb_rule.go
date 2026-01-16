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
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The priority of the Rule.Only the standard version is supported.",
			},
			"sticky_session_enabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable group session stickiness. Valid values are 'on' and 'off'.",
			},
			"sticky_session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The group session stickiness timeout, in seconds.",
			},
			"server_group_tuples": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Weight forwarded to the corresponding backend server group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The server group ID. The priority of this parameter is higher than that of `server_group_id`.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     100,
							Description: "The weight of the server group.",
						},
					},
				},
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
			"rule_conditions": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The rule conditions for standard edition forwarding rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of rule condition. Valid values: Host, Path, Header, Method, QueryString.",
						},
						"host_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Host configuration for Host type condition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"values": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The list of domain names.",
									},
								},
							},
						},
						"path_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Path configuration for Path type condition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"values": {
										Type:        schema.TypeList,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The list of absolute paths.",
									},
								},
							},
						},
						"header_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Header configuration for Header type condition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The header key.",
									},
									"values": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The list of header values.",
									},
								},
							},
						},
						"method_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Method configuration for Method type condition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"values": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The values of the method. Vaild values: HEAD,GET,POST,OPTIONS,PUT,PATCH,DELETE.",
									},
								},
							},
						},
						"query_string_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Query string configuration for QueryString type condition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of query string values.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The query string key.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The query string value.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"rule_actions": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The rule actions for standard edition forwarding rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The type of rule action. Valid values: ForwardGroup, Redirect, Rewrite, TrafficLimit.",
						},
						"traffic_limit_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Traffic limit configuration for TrafficLimit type action.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"qps": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "The QPS limit.",
									},
								},
							},
						},
						"forward_group_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Forward group configuration for ForwardGroup type action.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_group_sticky_session": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										MaxItems:    1,
										Description: "The config of group session stickiness.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"timeout": {
													Type:        schema.TypeInt,
													Optional:    true,
													Computed:    true,
													Description: "The sticky session timeout, in seconds.",
												},
												"enabled": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Whether to enable sticky session stickiness. Valid values are 'on' and 'off'.",
												},
											},
										},
									},
									"server_group_tuples": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "The server group tuples.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"server_group_id": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The server group ID.",
												},
												"weight": {
													Type:        schema.TypeInt,
													Optional:    true,
													Computed:    true,
													Description: "The weight of the server group.",
												},
											},
										},
									},
								},
							},
						},
						"redirect_config": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Computed:    true,
							Description: "Redirect configuration for Redirect type action.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The redirect HTTP code.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The redirect protocol.",
									},
									"host": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain name to which the request was redirected.",
									},
									"path": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path to which the request was redirected.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect port.",
									},
								},
							},
						},
						"rewrite_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							MaxItems:    1,
							Description: "Rewrite configuration for Rewrite type action.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The rewrite path.",
									},
								},
							},
						},
						"fixed_response_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							MaxItems:    1,
							Description: "Fixed response configuration for fixed response type rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"response_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The fixed response HTTP status code.",
									},
									"response_message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The fixed response message.",
									},
									"content_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content type of the fixed response.",
									},
									"response_body": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The response body of the fixed response.",
									},
								},
							},
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

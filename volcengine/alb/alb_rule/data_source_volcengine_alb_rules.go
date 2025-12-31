package alb_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbRulesRead,
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Id of listener.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Rule query.",
			},
			"rules": {
				Description: "The collection of Rule query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of Rule.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of Rule.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Domain of Rule.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Url of Rule.",
						},
						"rule_action": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The forwarding rule action, if this parameter is empty, " +
								"forward to server group, if value is `Redirect`, will redirect.",
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of Server Group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Description of Rule.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of the Rule. Only the standard version is supported.",
						},
						"traffic_limit_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forwarding rule QPS rate limiting switch:\n on: enable.\noff: disable (default).",
						},
						"traffic_limit_qps": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: "When Rules.N.TrafficLimitEnabled is turned on, this field is required. " +
								"Requests per second. Valid values are between 100 and 100000.",
						},
						"rewrite_enabled": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). " +
								"Only available for whitelist users, please submit an application to experience. " +
								"Supported values are as follows:\non: enable.\noff: disable.",
						},
						"rewrite_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of rewrite configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rewrite_path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rewrite path.",
									},
								},
							},
						},
						"forward_group_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of forward group configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_group_tuples": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of destination server groups to forward to.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"server_group_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The destination server group ID to forward to.",
												},
												"weight": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Server group weight.",
												},
											},
										},
									},
									"sticky_session_enabled": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Whether to enable inter-group session hold.",
									},
									"sticky_session_timeout": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The group session stickiness timeout, in seconds.",
									},
								},
							},
						},
						"redirect_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Redirect related configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redirect_domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect domain.",
									},
									"redirect_uri": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect URI.",
									},
									"redirect_port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect port.",
									},
									"redirect_http_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect HTTP code,support 301(default), 302, 307, 308.",
									},
									"redirect_protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect protocol,support HTTP,HTTPS(default).",
									},
								},
							},
						},
						"rule_conditions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The rule conditions for standard edition forwarding rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of rule condition. Valid values: Host, Path, Header.",
									},
									"host_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Host configuration for host type condition.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"values": {
													Type:        schema.TypeList,
													Computed:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "The list of domain names.",
												},
											},
										},
									},
									"path_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Path configuration for Path type condition.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"values": {
													Type:        schema.TypeList,
													Computed:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "The list of absolute paths.",
												},
											},
										},
									},
									"header_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Header configuration for Header type condition.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The header key.",
												},
												"values": {
													Type:        schema.TypeList,
													Computed:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "The list of header values.",
												},
											},
										},
									},
									"method_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Method configuration for Method type condition.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"values": {
													Type:        schema.TypeList,
													Computed:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "The list of HTTP methods.",
												},
											},
										},
									},
									"query_string_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Query string configuration.",
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
							Computed:    true,
							Description: "The rule actions for standard edition forwarding rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of rule action. Valid values: ForwardGroup, Redirect, Rewrite, TrafficLimit.",
									},
									"traffic_limit_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Traffic limit configuration for TrafficLimit type action.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"qps": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The QPS limit.",
												},
											},
										},
									},
									"forward_group_config": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Forward group configuration for ForwardGroup type action.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"server_group_sticky_session": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The config of group session stickiness.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"timeout": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The sticky session timeout, in seconds.",
															},
															"enabled": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Whether to enable sticky session stickiness. Valid values are 'on' and 'off'.",
															},
														},
													},
												},
												"server_group_tuples": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The server group tuples.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server_group_id": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The server group ID.",
															},
															"weight": {
																Type:        schema.TypeInt,
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
										Computed:    true,
										Description: "Redirect configuration for Redirect type action.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"redirect_http_code": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The redirect HTTP code.",
												},
												"redirect_protocol": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The redirect protocol.",
												},
												"redirect_domain": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The redirect domain.",
												},
												"redirect_uri": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The redirect URI.",
												},
												"redirect_port": {
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
				},
			},
		},
	}
}

func dataSourceVolcengineAlbRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbRules())
}

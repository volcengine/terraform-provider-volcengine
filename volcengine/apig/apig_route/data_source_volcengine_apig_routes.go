package apig_route

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineApigRoutes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineApigRoutesRead,
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of api gateway.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of api gateway service.",
			},
			"upstream_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of api gateway upstream.",
			},
			"upstream_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of api gateway upstream.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource type of route. Valid values: `Console`, `Ingress`.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of api gateway route. This field support fuzzy query.",
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The path of api gateway route.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"routes": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the api gateway route.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the api gateway route.",
						},
						"service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the api gateway service.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the api gateway service.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the api gateway route.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the api gateway route.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type of route. Valid values: `Console`, `Ingress`.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the api gateway route.",
						},
						"reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason of the api gateway route.",
						},
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the api gateway route is enabled.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of the api gateway route.",
						},
						"domains": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The domains of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of the api gateway route.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the domain.",
									},
								},
							},
						},
						"custom_domains": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The custom domains of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the custom domain.",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The custom domain of the api gateway route.",
									},
								},
							},
						},
						"match_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The match rule of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The path of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"match_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The match type of the api gateway route.",
												},
												"match_content": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The match content of the api gateway route.",
												},
											},
										},
									},
									"method": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The method of the api gateway route.",
									},
									"query_string": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The query string of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key of the query string.",
												},
												"value": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The path of the api gateway route.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"match_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The match type of the api gateway route.",
															},
															"match_content": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The match content of the api gateway route.",
															},
														},
													},
												},
											},
										},
									},
									"header": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The header of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key of the header.",
												},
												"value": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The path of the api gateway route.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"match_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The match type of the api gateway route.",
															},
															"match_content": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The match content of the api gateway route.",
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
						"upstream_list": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The upstream list of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"upstream_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the api gateway upstream.",
									},
									"version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of the api gateway upstream.",
									},
									"weight": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The weight of the api gateway upstream.",
									},
									"ai_provider_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The ai provider settings of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"model": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The model of the ai provider.",
												},
												"target_path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The target path of the ai provider.",
												},
											},
										},
									},
								},
							},
						},
						"advanced_setting": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The advanced setting of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeout_setting": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The timeout setting of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether the timeout setting is enabled.",
												},
												"timeout": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The timeout of the api gateway route.",
												},
											},
										},
									},
									"cors_policy_setting": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The cors policy setting of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether the cors policy setting is enabled.",
												},
											},
										},
									},
									"url_rewrite_setting": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The url rewrite setting of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether the url rewrite setting is enabled.",
												},
												"url_rewrite": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The url rewrite path of the api gateway route.",
												},
											},
										},
									},
									"retry_policy_setting": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The retry policy setting of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether the retry policy setting is enabled.",
												},
												"attempts": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The attempts of the api gateway route.",
												},
												"per_try_timeout": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The per try timeout of the api gateway route.",
												},
												"retry_on": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The retry on of the api gateway route.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"http_codes": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The http codes of the api gateway route.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"header_operations": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The header operations of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operation": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The operation of the header.",
												},
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key of the header.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value of the header.",
												},
												"direction_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The direction type of the header.",
												},
											},
										},
									},
									"mirror_policies": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The mirror policies of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"upstream": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The upstream of the mirror policy.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"upstream_id": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The id of the api gateway upstream.",
															},
															"version": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The version of the api gateway upstream.",
															},
															"type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the api gateway upstream.",
															},
														},
													},
												},
												"percent": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The percent of the mirror policy.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The percent value of the mirror policy.",
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
				},
			},
		},
	}
}

func dataSourceVolcengineApigRoutesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewApigRouteService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineApigRoutes())
}

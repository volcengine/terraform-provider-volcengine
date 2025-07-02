package apig_route

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ApigRoute can be imported using the id, e.g.
```
$ terraform import volcengine_apig_route.default resource_id
```

*/

func ResourceVolcengineApigRoute() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineApigRouteCreate,
		Read:   resourceVolcengineApigRouteRead,
		Update: resourceVolcengineApigRouteUpdate,
		Delete: resourceVolcengineApigRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the apig route.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The service id of the apig route.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The resource type of the apig route. Valid values: `Console`, `Ingress` Default is `Console`.",
			},
			"enable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether the apig route is enabled. Default is `false`.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The priority of the apig route. Valid values: 0~100.",
			},
			"upstream_list": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The upstream list of the api gateway route.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"upstream_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the api gateway upstream.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The version of the api gateway upstream.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The weight of the api gateway upstream. Valid values: 0~10000.",
						},
						"ai_provider_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The ai provider settings of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"model": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The model of the ai provider.",
									},
									"target_path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The target path of the ai provider.",
									},
								},
							},
						},
					},
				},
			},
			"match_rule": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The match rule of the api gateway route.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "The path of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"match_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The match type of the api gateway route. Valid values: `Prefix`, `Exact`, `Regex`.",
									},
									"match_content": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The match content of the api gateway route.",
									},
								},
							},
						},
						"method": {
							Type:     schema.TypeSet,
							Optional: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The method of the api gateway route. Valid values: `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `CONNECT`.",
						},
						"query_string": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The query string of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of the query string.",
									},
									"value": {
										Type:        schema.TypeList,
										Required:    true,
										MaxItems:    1,
										Description: "The path of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"match_type": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The match type of the api gateway route. Valid values: `Prefix`, `Exact`, `Regex`.",
												},
												"match_content": {
													Type:        schema.TypeString,
													Required:    true,
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
							Optional:    true,
							Description: "The header of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of the header.",
									},
									"value": {
										Type:        schema.TypeList,
										Required:    true,
										MaxItems:    1,
										Description: "The path of the api gateway route.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"match_type": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The match type of the api gateway route. Valid values: `Prefix`, `Exact`, `Regex`.",
												},
												"match_content": {
													Type:        schema.TypeString,
													Required:    true,
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
			"advanced_setting": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The advanced setting of the api gateway route.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout_setting": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The timeout setting of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether the timeout setting is enabled.",
									},
									"timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The timeout of the api gateway route. Unit: s.",
									},
								},
							},
						},
						"cors_policy_setting": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The cors policy setting of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether the cors policy setting is enabled.",
									},
								},
							},
						},
						"url_rewrite_setting": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The url rewrite setting of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether the url rewrite setting is enabled.",
									},
									"url_rewrite": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The url rewrite path of the api gateway route.",
									},
								},
							},
						},
						"retry_policy_setting": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The retry policy setting of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether the retry policy setting is enabled.",
									},
									"attempts": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The attempts of the api gateway route.",
									},
									"per_try_timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The per try timeout of the api gateway route.",
									},
									"retry_on": {
										Type:        schema.TypeSet,
										Optional:    true,
										Set:         schema.HashString,
										Description: "The retry on of the api gateway route. Valid values: `5xx`, `reset`, `connect-failure`, `refused-stream`, `cancelled`, `deadline-exceeded`, `internal`, `resource-exhausted`, `unavailable`.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"http_codes": {
										Type:        schema.TypeSet,
										Optional:    true,
										Set:         schema.HashString,
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
							Optional:    true,
							Description: "The header operations of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"operation": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operation of the header. Valid values: `set`, `add`, `remove`.",
									},
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of the header.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The value of the header.",
									},
									"direction_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The direction type of the header. Valid values: `request`, `response`.",
									},
								},
							},
						},
						"mirror_policies": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The mirror policies of the api gateway route.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"upstream": {
										Type:        schema.TypeList,
										Required:    true,
										MaxItems:    1,
										Description: "The upstream of the mirror policy.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"upstream_id": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The id of the api gateway upstream.",
												},
												"version": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The version of the api gateway upstream.",
												},
												"type": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The type of the api gateway upstream.",
												},
											},
										},
									},
									"percent": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "The percent of the mirror policy.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:        schema.TypeInt,
													Required:    true,
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

			// common fields
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
		},
	}
	return resource
}

func resourceVolcengineApigRouteCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigRouteService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineApigRoute())
	if err != nil {
		return fmt.Errorf("error on creating apig_route %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigRouteRead(d, meta)
}

func resourceVolcengineApigRouteRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigRouteService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineApigRoute())
	if err != nil {
		return fmt.Errorf("error on reading apig_route %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineApigRouteUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigRouteService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineApigRoute())
	if err != nil {
		return fmt.Errorf("error on updating apig_route %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigRouteRead(d, meta)
}

func resourceVolcengineApigRouteDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigRouteService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineApigRoute())
	if err != nil {
		return fmt.Errorf("error on deleting apig_route %q, %s", d.Id(), err)
	}
	return err
}

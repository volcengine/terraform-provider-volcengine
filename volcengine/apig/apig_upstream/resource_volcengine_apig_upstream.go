package apig_upstream

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ApigUpstream can be imported using the id, e.g.
```
$ terraform import volcengine_apig_upstream.default resource_id
```

*/

func ResourceVolcengineApigUpstream() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineApigUpstreamCreate,
		Read:   resourceVolcengineApigUpstreamRead,
		Update: resourceVolcengineApigUpstreamUpdate,
		Delete: resourceVolcengineApigUpstreamDelete,
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
				Description: "The name of the apig upstream.",
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The gateway id of the apig upstream.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comments of the apig upstream.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The source type of the apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The resource type of the apig upstream. Valid values: `Console`, `Ingress`.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The protocol of the apig upstream. Valid values: `HTTP`, `HTTP2`, `GRPC`.",
			},
			"load_balancer_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The load balancer settings of apig upstream.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lb_policy": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The load balancer policy of apig upstream. Valid values: `SimpleLB`, `ConsistentHashLB`.",
						},
						"simple_lb": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The simple load balancer of apig upstream. Valid values: `ROUND_ROBIN`, `LEAST_CONN`, `RANDOM`.",
						},
						"warmup_duration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The warmup duration of apig upstream lb. This field is valid when the simple_lb is `ROUND_ROBIN` or `LEAST_CONN`.",
						},
						"consistent_hash_lb": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The consistent hash lb of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hash_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The hash key of apig upstream consistent hash lb. Valid values: `HTTPCookie`, `HttpHeaderName`, `HttpQueryParameterName`, `UseSourceIp`.",
									},
									"use_source_ip": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "The use source ip of apig upstream consistent hash lb.",
									},
									"http_query_parameter_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The http query parameter name of apig upstream consistent hash lb.",
									},
									"http_header_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The http header name of apig upstream consistent hash lb.",
									},
									"http_cookie": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "The http cookie of apig upstream consistent hash lb.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of apig upstream consistent hash lb http cookie.",
												},
												"path": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The path of apig upstream consistent hash lb http cookie.",
												},
												"ttl": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The ttl of apig upstream consistent hash lb http cookie.",
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
			"tls_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The tls settings of apig upstream.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tls_mode": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The tls mode of apig upstream tls setting. Valid values: `DISABLE`, `SIMPLE`.",
						},
						"sni": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The sni of apig upstream tls setting.",
						},
					},
				},
			},
			"circuit_breaking_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The circuit breaking settings of apig upstream.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the circuit breaking is enabled.",
						},
						"consecutive_errors": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("circuit_breaking_settings.0.enable").(bool)
							},
							Description: "The consecutive errors of circuit breaking. Default is 5.",
						},
						"interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("circuit_breaking_settings.0.enable").(bool)
							},
							Description: "The interval of circuit breaking. Unit: ms. Default is 10s.",
						},
						"base_ejection_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("circuit_breaking_settings.0.enable").(bool)
							},
							Description: "The base ejection time of circuit breaking. Unit: ms. Default is 10s.",
						},
						"max_ejection_percent": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("circuit_breaking_settings.0.enable").(bool)
							},
							Description: "The max ejection percent of circuit breaking. Default is 20%.",
						},
						"min_health_percent": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("circuit_breaking_settings.0.enable").(bool)
							},
							Description: "The min health percent of circuit breaking. Default is 60%.",
						},
					},
				},
			},
			"upstream_spec": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The upstream spec of apig upstream.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ve_faas": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							//ExactlyOneOf: []string{"ve_faas", "k8s_service", "ecs_list", "fixed_ip_list", "domain", "nacos_service", "ve_mlp", "ai_provider"},
							Description: "The vefaas of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"function_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The function id of vefaas.",
									},
								},
							},
						},
						"k8s_service": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The k8s service of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"namespace": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The namespace of k8s service.",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of k8s service.",
									},
									"port": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The port of k8s service.",
									},
								},
							},
						},
						"ecs_list": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The ecs list of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ecs_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The instance id of ecs.",
									},
									"ip": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The ip of ecs.",
									},
									"port": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The port of ecs.",
									},
								},
							},
						},
						"fixed_ip_list": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The fixed ip list of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The ip of apig upstream.",
									},
									"port": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The port of apig upstream.",
									},
								},
							},
						},
						"domain": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The domain of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The protocol of apig upstream. Valid values: `HTTP`, `HTTPS`.",
									},
									"domain_list": {
										Type:        schema.TypeSet,
										Required:    true,
										MaxItems:    1,
										Description: "The domain list of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The domain of apig upstream.",
												},
												"port": {
													Type:        schema.TypeInt,
													Optional:    true,
													Computed:    true,
													Description: "The port of domain. Default is 80 for HTTP, 443 for HTTPS.",
												},
											},
										},
									},
								},
							},
						},
						"nacos_service": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The nacos service of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"upstream_source_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The upstream source id.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The namespace of nacos service.",
									},
									"group": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The group of nacos service.",
									},
									"service": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The service of nacos service.",
									},
									"namespace_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The namespace id of nacos service.",
									},
								},
							},
						},
						"ve_mlp": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The mlp of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The service id of mlp.",
									},
									"service_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The service name of mlp.",
									},
									"service_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The service url of mlp.",
									},
									"service_discover_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The service discover type of mlp.",
									},
									"upstream_source_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The upstream source id.",
									},
									"k8s_service": {
										Type:        schema.TypeList,
										Required:    true,
										MaxItems:    1,
										Description: "The k8s service of mlp.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The namespace of k8s service.",
												},
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of k8s service.",
												},
												"port": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The port of k8s service.",
												},
												"cluster_info": {
													Type:        schema.TypeList,
													Required:    true,
													MaxItems:    1,
													Description: "The cluster info of k8s service.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cluster_name": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "The cluster name of k8s service.",
															},
															"account_id": {
																Type:        schema.TypeInt,
																Required:    true,
																Description: "The account id of k8s service.",
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
						"ai_provider": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The ai provider of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of ai provider.",
									},
									"token": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The token of ai provider.",
									},
									"base_url": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The base url of ai provider.",
									},
									"custom_model_service": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "The custom model service of ai provider.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The namespace of custom model service.",
												},
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of custom model service.",
												},
												"port": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The port of custom model service.",
												},
											},
										},
									},
									"custom_header_params": {
										Type:        schema.TypeMap,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The custom header params of ai provider.",
									},
									"custom_body_params": {
										Type:        schema.TypeMap,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The custom body params of ai provider.",
									},
								},
							},
						},
					},
				},
			},

			// computed fields
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of apig upstream.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of apig upstream.",
			},
			"version_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The version details of apig upstream.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of apig upstream version.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of apig upstream version.",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The labels of apig upstream version.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of apig upstream version label.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of apig upstream version label.",
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

func resourceVolcengineApigUpstreamCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineApigUpstream())
	if err != nil {
		return fmt.Errorf("error on creating apig_upstream %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigUpstreamRead(d, meta)
}

func resourceVolcengineApigUpstreamRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineApigUpstream())
	if err != nil {
		return fmt.Errorf("error on reading apig_upstream %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineApigUpstreamUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineApigUpstream())
	if err != nil {
		return fmt.Errorf("error on updating apig_upstream %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigUpstreamRead(d, meta)
}

func resourceVolcengineApigUpstreamDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineApigUpstream())
	if err != nil {
		return fmt.Errorf("error on deleting apig_upstream %q, %s", d.Id(), err)
	}
	return err
}

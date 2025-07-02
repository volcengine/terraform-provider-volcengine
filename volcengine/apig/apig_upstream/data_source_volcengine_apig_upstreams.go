package apig_upstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineApigUpstreams() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineApigUpstreamsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of apig upstream IDs.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of apig upstream. This field support fuzzy query.",
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of api gateway.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource type of apig upstream. Valid values: `Console`, `Ingress`.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The source type of apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.",
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
			"upstreams": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of apig upstream.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of apig upstream.",
						},
						"gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of api gateway.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comments of apig upstream.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type of apig upstream.",
						},
						"source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source type of apig upstream.",
						},
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
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol of apig upstream.",
						},
						"load_balancer_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The load balancer settings of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lb_policy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The load balancer policy of apig upstream.",
									},
									"simple_lb": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The simple load balancer of apig upstream.",
									},
									"warmup_duration": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The warmup duration of apig upstream lb.",
									},
									"consistent_hash_lb": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The consistent hash lb of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hash_key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The hash key of apig upstream consistent hash lb.",
												},
												"use_source_ip": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "The use source ip of apig upstream consistent hash lb.",
												},
												"http_query_parameter_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The http query parameter name of apig upstream consistent hash lb.",
												},
												"http_header_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The http header name of apig upstream consistent hash lb.",
												},
												"http_cookie": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The http cookie of apig upstream consistent hash lb.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of apig upstream consistent hash lb http cookie.",
															},
															"path": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The path of apig upstream consistent hash lb http cookie.",
															},
															"ttl": {
																Type:        schema.TypeInt,
																Computed:    true,
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
							Computed:    true,
							Description: "The tls settings of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tls_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tls mode of apig upstream tls setting.",
									},
									"sni": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The sni of apig upstream tls setting.",
									},
								},
							},
						},
						"circuit_breaking_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The circuit breaking settings of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the circuit breaking is enabled.",
									},
									"consecutive_errors": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The consecutive errors of circuit breaking.",
									},
									"interval": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The interval of circuit breaking. Unit: ms.",
									},
									"base_ejection_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The base ejection time of circuit breaking. Unit: ms.",
									},
									"max_ejection_percent": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max ejection percent of circuit breaking.",
									},
									"min_health_percent": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The min health percent of circuit breaking.",
									},
								},
							},
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
						"backend_target_list": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The backend target list of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ip of apig upstream backend.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port of apig upstream backend.",
									},
									"health_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The health status of apig upstream backend.",
									},
								},
							},
						},
						"upstream_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The upstream spec of apig upstream.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ve_faas": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The vefaas of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"function_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The function id of vefaas.",
												},
											},
										},
									},
									"k8s_service": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The k8s service of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The namespace of k8s service.",
												},
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of k8s service.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The port of k8s service.",
												},
											},
										},
									},
									"ecs_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The ecs list of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ecs_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The instance id of ecs.",
												},
												"ip": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ip of ecs.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The port of ecs.",
												},
											},
										},
									},
									"fixed_ip_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The fixed ip list of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ip of apig upstream.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The port of apig upstream.",
												},
											},
										},
									},
									"domain": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The domain of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The protocol of apig upstream.",
												},
												"domain_list": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The domain list of apig upstream.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"domain": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The domain of apig upstream.",
															},
															"port": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The port of domain.",
															},
														},
													},
												},
											},
										},
									},
									"nacos_service": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The nacos service of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"upstream_source_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The upstream source id.",
												},
												"namespace": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The namespace of nacos service.",
												},
												"group": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The group of nacos service.",
												},
												"service": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The service of nacos service.",
												},
												"namespace_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The namespace id of nacos service.",
												},
											},
										},
									},
									"ve_mlp": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The mlp of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"service_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The service id of mlp.",
												},
												"service_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The service name of mlp.",
												},
												"service_url": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The service url of mlp.",
												},
												"service_discover_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The service discover type of mlp.",
												},
												"upstream_source_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The upstream source id.",
												},
												"k8s_service": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The k8s service of mlp.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"namespace": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The namespace of k8s service.",
															},
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of k8s service.",
															},
															"port": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The port of k8s service.",
															},
															"cluster_info": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The cluster info of k8s service.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cluster_name": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The cluster name of k8s service.",
																		},
																		"account_id": {
																			Type:        schema.TypeInt,
																			Computed:    true,
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
										Computed:    true,
										Description: "The ai provider of apig upstream.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of ai provider.",
												},
												"token": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The token of ai provider.",
												},
												"base_url": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The base url of ai provider.",
												},
												"custom_model_service": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The custom model service of ai provider.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"namespace": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The namespace of custom model service.",
															},
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of custom model service.",
															},
															"port": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The port of custom model service.",
															},
														},
													},
												},
												"custom_header_params": {
													Type:        schema.TypeMap,
													Computed:    true,
													Elem:        schema.TypeString,
													Description: "The custom header params of ai provider.",
												},
												"custom_body_params": {
													Type:        schema.TypeMap,
													Computed:    true,
													Elem:        schema.TypeString,
													Description: "The custom body params of ai provider.",
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

func dataSourceVolcengineApigUpstreamsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewApigUpstreamService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineApigUpstreams())
}

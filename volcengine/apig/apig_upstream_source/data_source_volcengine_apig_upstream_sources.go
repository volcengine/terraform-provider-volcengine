package apig_upstream_source

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineApigUpstreamSources() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineApigUpstreamSourcesRead,
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of api gateway.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of nacos source.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The source type of apig upstream source. Valid values: `K8S`, `Nacos`.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of apig upstream source. Valid values: `Syncing`, `SyncedSucceed`, `SyncedFailed`.",
			},
			"enable_ingress": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The enable ingress of apig upstream source.",
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
			"upstream_sources": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of apig upstream source.",
						},
						"gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of api gateway.",
						},
						"source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source type of apig upstream source.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comments of apig upstream source.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of apig upstream source.",
						},
						"status_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status message of apig upstream source.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of apig upstream source.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of apig upstream source.",
						},
						"source_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The source spec of apig upstream source.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"k8s_source": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The k8s source of apig upstream source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The cluster id of k8s source.",
												},
												"cluster_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The cluster type of k8s source.",
												},
											},
										},
									},
									"nacos_source": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The nacos source of apig upstream source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nacos_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The nacos id of nacos source.",
												},
												"nacos_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The nacos name of nacos source.",
												},
												"address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The address of nacos source.",
												},
												"http_port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The http port of nacos source.",
												},
												"grpc_port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The grpc port of nacos source.",
												},
												"context_path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The context path of nacos source.",
												},
												"auth_config": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The auth config of nacos source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"basic": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The basic auth config of nacos source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"username": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The username of basic auth config.",
																		},
																		"password": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The password of basic auth config.",
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
						"ingress_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ingress settings of apig upstream source.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_ingress": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable ingress.",
									},
									"enable_all_ingress_classes": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable all ingress classes.",
									},
									"enable_ingress_without_ingress_class": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable ingress without ingress class.",
									},
									"enable_all_namespaces": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable all namespaces.",
									},
									"update_status": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The update status of ingress settings.",
									},
									"ingress_classes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The ingress classes of ingress settings.",
									},
									"watch_namespaces": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The watch namespaces of ingress settings.",
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

func dataSourceVolcengineApigUpstreamSourcesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewApigUpstreamSourceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineApigUpstreamSources())
}

package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineESCloudInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineESCloudInstancesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of instance IDs.",
			},
			"statuses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list status of instance.",
			},
			"charge_types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The charge types of instance.",
			},
			"names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The names of instance.",
			},
			"versions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The versions of instance.",
			},
			"zone_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The available zone IDs of instance.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of instance query.",
			},
			"instances": {
				Description: "The collection of instance query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of instance.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of instance.",
						},
						"user_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user id of instance.",
						},
						"charge_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The charge status of instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of instance.",
						},
						"expire_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expire time of instance.",
						},
						"maintenance_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The maintenance time of instance.",
						},
						"maintenance_day": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The maintenance day of instance.",
						},
						"namespace": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The namespace of instance.",
						},
						"total_nodes": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total nodes of instance.",
						},
						"enable_es_public_network": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "whether enable es public network.",
						},
						"enable_es_private_network": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "whether enable es private network.",
						},
						"es_public_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The es public domain of instance.",
						},
						"es_private_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The es private domain of instance.",
						},
						"es_public_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The es public endpoint of instance.",
						},
						"es_private_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The es private endpoint of instance.",
						},
						"es_inner_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The es inner endpoint of instance.",
						},
						"enable_kibana_public_network": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "whether enable kibana public network.",
						},
						"enable_kibana_private_network": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "whether enable kibana private network.",
						},
						"kibana_private_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The kibana private domain of instance.",
						},
						"kibana_public_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The kibana public domain of instance.",
						},
						"instance_configuration": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The configuration of instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of instance.",
									},
									"instance_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of instance.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of project.",
									},
									"period": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The period of project.",
									},
									"region_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region info of instance.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The zoneId of instance.",
									},
									"zone_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The zone number of instance.",
									},
									"enable_https": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "whether enable https.",
									},
									"admin_user_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user name of instance.",
									},
									"charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge type of instance.",
									},
									"enable_pure_master": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether enable pure master.",
									},
									"master_node_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The node number of master.",
									},
									"hot_node_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The node number of host.",
									},
									"kibana_node_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The node number of kibana.",
									},

									"kibana_node_resource_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node resource spec of kibana.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of resource spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of resource spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of resource spec.",
												},
												"cpu": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The cpu info of resource spec.",
												},
												"memory": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The memory info of resource spec.",
												},
											},
										},
									},
									"kibana_node_storage_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node storage spec of kibana.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of storage spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of storage spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of storage spec.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of storage spec.",
												},
												"min_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The min size of storage spec.",
												},
												"max_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The max size of storage spec.",
												},
												"size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The size of storage spec.",
												},
											},
										},
									},
									"hot_node_resource_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node resource spec of host.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of resource spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of resource spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of resource spec.",
												},
												"cpu": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The cpu info of resource spec.",
												},
												"memory": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The memory info of resource spec.",
												},
											},
										},
									},
									"hot_node_storage_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node storage spec of host.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of storage spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of storage spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of storage spec.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of storage spec.",
												},
												"min_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The min size of storage spec.",
												},
												"max_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The max size of storage spec.",
												},
												"size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The size of storage spec.",
												},
											},
										},
									},
									"master_node_resource_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node resource spec of master.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of resource spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of resource spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of resource spec.",
												},
												"cpu": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The cpu info of resource spec.",
												},
												"memory": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The memory info of resource spec.",
												},
											},
										},
									},
									"master_node_storage_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node storage spec of master.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of storage spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of storage spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of storage spec.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of storage spec.",
												},
												"min_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The min size of storage spec.",
												},
												"max_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The max size of storage spec.",
												},
												"size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The size of storage spec.",
												},
											},
										},
									},

									"vpc": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The vpc info.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vpc_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The id of vpc.",
												},
												"vpc_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of vpc.",
												},
											},
										},
									},
									"subnet": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The subnet info.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"subnet_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The id of subnet.",
												},
												"subnet_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of subnet.",
												},
											},
										},
									},
								},
							},
						},
						"nodes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The nodes info of instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of node.",
									},
									"node_display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The show name of node.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of node.",
									},
									"start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The start time of node.",
									},
									"restart_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The restart times of node.",
									},
									"is_master": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is master node.",
									},
									"is_hot": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is hot node.",
									},
									"is_warm": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is warm node.",
									},
									"is_cold": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is cold node.",
									},
									"is_kibana": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is kibana node.",
									},
									"resource_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node resource spec of master.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of resource spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of resource spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of resource spec.",
												},
												"cpu": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The cpu info of resource spec.",
												},
												"memory": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The memory info of resource spec.",
												},
											},
										},
									},
									"storage_spec": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "The node storage spec of master.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of storage spec.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of storage spec.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The show name of storage spec.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of storage spec.",
												},
												"min_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The min size of storage spec.",
												},
												"max_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The max size of storage spec.",
												},
												"size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The size of storage spec.",
												},
											},
										},
									},
								},
							},
						},
						"plugins": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The plugin info of instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of plugin.",
									},
									"plugin_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of plugin.",
									},
									"version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of plugin.",
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

func dataSourceVolcengineESCloudInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewESCloudInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineESCloudInstances())
}

package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineInstancesRead,
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
			"names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of instance names.",
			},
			"cloud_server_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of cloud server ids.",
			},
			"statuses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
				Description: "The list of instance status. The value can be `opening` or `starting` or `running` or " +
					"`stopping` or `stop` or `rebooting` or `terminating`.",
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
						"instance_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of instance.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of instance.",
						},
						"cloud_server_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of cloud server.",
						},
						"cloud_server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of cloud server.",
						},
						"vpc_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of vpc.",
						},
						"subnet_cidr": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet cidr.",
						},
						"spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The spec of instance.",
						},
						"spec_display": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The spec display of instance.",
						},
						"cpu": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cpu of instance.",
						},
						"mem": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The memory of instance.",
						},
						"creator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creator of instance.",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start time of instance.",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end time of instance.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The create time of instance.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of instance.",
						},
						"delete_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The delete time of instance.",
						},
						"gpu": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of gpu.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gpus": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list gpu info.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"num": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of gpu.",
												},
												"gpu_spec": {
													Type:        schema.TypeList,
													Computed:    true,
													MaxItems:    1,
													Description: "The spec of gpu.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gpu_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of gpu.",
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
						"network": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of network.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vf_passthrough": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The passthrough info.",
									},
									"enable_ipv6": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether enable ipv6.",
									},
									"internal_interface": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "The internal interface of network.",
										Elem:        networkDataSource,
									},
									"external_interface": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "The external interface of network.",
										Elem:        networkDataSource,
									},
								},
							},
						},
						"image": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of image.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image_identity": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of image.",
									},
									"image_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of image.",
									},
									"system_arch": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The arch of system.",
									},
									"system_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of system.",
									},
									"system_bit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bit of system.",
									},
									"system_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of system.",
									},
									"property": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The property of system.",
									},
								},
							},
						},
						"storage": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of storage.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_disk": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "The disk info of system.",
										Elem:        diskSpec,
									},
									"data_disk": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "The disk info of data.",
										Elem:        diskSpec,
									},
									"data_disk_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The disk list info of data.",
										Elem:        diskSpec,
									},
								},
							},
						},
						"cluster": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The cluster info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of cluster.",
									},
									"country": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The country of cluster.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region of cluster.",
									},
									"province": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The province of cluster.",
									},
									"city": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The city of cluster.",
									},
									"isp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The isp of cluster.",
									},
									"level": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The level of cluster.",
									},
									"alias": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The alias of cluster.",
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

var diskSpec = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"storage_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The type of storage.",
		},
		"capacity": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The capacity of storage.",
		},
	},
}

var networkDataSource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"ip_addr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ip address.",
		},
		"mask": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The mask info.",
		},
		"ip6_addr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ipv6 address.",
		},
		"mask6": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ipv6 mask info.",
		},
		"mac_addr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The mac address.",
		},
		"bandwidth_peak": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The peak of bandwidth.",
		},
		"ips": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The ips of network.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mask": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The mask info.",
					},
					"addr": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The ip address.",
					},
					"isp": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The isp info.",
					},
					"ip_version": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The version of ip address.",
					},
				},
			},
		},
	},
}

func dataSourceVolcengineInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineInstances())
}

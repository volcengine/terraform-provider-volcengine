package cloud_server

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVeenedgeCloudServers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVeenedgeCloudServersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cloud server IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Cloud Server.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cloud servers query.",
			},
			"cloud_servers": {
				Description: "The collection of cloud servers query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of cloud server.",
						},
						"cloud_server_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of cloud server.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of cloud server.",
						},
						"server_area_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The server area count number.",
						},
						"instance_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of instances.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The create time info.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time info.",
						},
						"server_area_level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The area level of cloud server.",
						},
						"spec_sum": {
							Type:        schema.TypeMap,
							Computed:    true,
							Elem:        schema.TypeInt,
							Description: "The spec summary of cloud server.",
						},
						"spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The spec info of cloud server.",
						},
						"spec_display": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Chinese spec info of cloud server.",
						},
						"cpu": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cpu info of cloud server.",
						},
						"mem": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The memory info of cloud server.",
						},
						"instance_status": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The status of instances.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status info.",
									},
									"instance_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The count of instance.",
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
						"network": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of network.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bandwidth_peak": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The peak of bandwidth.",
									},
									"internal_bandwidth_peak": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The internal peak of bandwidth.",
									},
									"enable_ipv6": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether enable ipv6.",
									},
								},
							},
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
						"secret_config": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of secret.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"secret_type": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The type of secret.",
									},
									"secret_data": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The data of secret.",
									},
								},
							},
						},
						"custom_data": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of custom data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The data info.",
									},
								},
							},
						},
						"billing_config": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of billing.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"computing_billing_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The computing billing method.",
									},
									"bandwidth_billing_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bandwidth billing method.",
									},
								},
							},
						},
						"schedule_strategy_configs": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of schedule strategy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule_strategy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The schedule strategy.",
									},
									"price_strategy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The price strategy.",
									},
								},
							},
						},
						"server_areas": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The server areas info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The area info.",
									},
									"isp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The isp info.",
									},
									"instance_num": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of instance.",
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

func dataSourceVolcengineVeenedgeCloudServersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudServerService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVeenedgeCloudServers())
}
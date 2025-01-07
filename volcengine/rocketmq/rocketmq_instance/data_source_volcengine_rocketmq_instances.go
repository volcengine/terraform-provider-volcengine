package rocketmq_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRocketmqInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRocketmqInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of rocketmq instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of rocketmq instance. This field support fuzzy query.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of rocketmq instance.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone id of rocketmq instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id of rocketmq instance.",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of rocketmq instance. Valid values: `4.8`.",
			},
			"spec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The spec of rocketmq instance.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The charge type of rocketmq instance. Valid values: `PostPaid`, `PrePaid`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of rocketmq instance.",
			},
			"tags": ve.TagsSchema(),
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

			"rocketmq_instances": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rocketmq instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rocketmq instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rocketmq instance.",
						},
						"instance_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the rocketmq instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the rocketmq instance.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id of the rocketmq instance.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id of the rocketmq instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the rocketmq instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the rocketmq instance.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the rocketmq instance.",
						},
						"compute_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The compute spec of the rocketmq instance.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the rocketmq instance.",
						},
						"eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The eip id of the rocketmq instance.",
						},
						"ssl_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ssl mode of the rocketmq instance.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the rocketmq instance.",
						},
						"tags": ve.TagsSchemaComputed(),
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the rocketmq instance.",
						},
						"used_topic_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used topic number of the rocketmq instance.",
						},
						"used_storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used storage space of the rocketmq instance.",
						},
						"storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total storage space of the rocketmq instance.",
						},
						"file_reserved_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The reserved time of messages on the RocketMQ server of the message queue. Messages that exceed the reserved time will be cleared after expiration. The unit is in hours.",
						},
						"available_queue_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The available queue number of the rocketmq instance.",
						},
						"used_queue_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used queue number of the rocketmq instance.",
						},
						"used_group_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used group number of the rocketmq instance.",
						},
						"enable_ssl": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the ssl authentication is enabled for the rocketmq instance.",
						},
						"apply_private_dns_to_public": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the private dns to public function is enabled for the rocketmq instance.",
						},
						"charge_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The charge detail information of the rocketmq instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge type of the rocketmq instance.",
									},
									"charge_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge status of the rocketmq instance.",
									},
									"charge_start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge start time of the rocketmq instance.",
									},
									"charge_expire_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge expire time of the rocketmq instance.",
									},
									"overdue_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The overdue time of the rocketmq instance.",
									},
									"overdue_reclaim_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The overdue reclaim time of the rocketmq instance.",
									},
									"period_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The period unit of the rocketmq instance.",
									},
									"auto_renew": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable automatic renewal.",
									},
								},
							},
						},
						"connection_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The connection information of the rocketmq.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The endpoint type of the rocketmq.",
									},
									"network_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The network type of the rocketmq.",
									},
									"internal_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The internal endpoint of the rocketmq.",
									},
									"public_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public endpoint of the rocketmq.",
									},
									"endpoint_address_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The endpoint address ip of the rocketmq.",
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

func dataSourceVolcengineRocketmqInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRocketmqInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRocketmqInstances())
}

package rabbitmq_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRabbitmqInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRabbitmqInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of rabbitmq instance. This field supports fuzzy query.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of rabbitmq instance. This field supports fuzzy query.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of rabbitmq instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id of rabbitmq instance. This field supports fuzzy query.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone id of rabbitmq instance. This field supports fuzzy query.",
			},
			"charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostPaid", "PrePaid"}, false),
				Description:  "The charge type of rabbitmq instance.",
			},
			"spec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The calculation specification of rabbitmq instance.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of rabbitmq instance.",
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
			"rabbitmq_instances": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rabbitmq instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rabbitmq instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rabbitmq instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the rabbitmq instance.",
						},
						"instance_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the rabbitmq instance.",
						},
						"apply_private_dns_to_public": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether enable the public network parsing function of the rabbitmq instance.",
						},
						"arch_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the rabbitmq instance.",
						},
						"compute_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The compute specification of the rabbitmq instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the rabbitmq instance.",
						},
						"eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The eip id of the rabbitmq instance.",
						},
						"init_user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The WebUI admin user name of the rabbitmq instance.",
						},
						//"is_encrypted": {
						//	Type:        schema.TypeBool,
						//	Computed:    true,
						//	Description: "Whether enable the volume encryption function.",
						//},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the rabbitmq instance.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id of the rabbitmq instance.",
						},
						"region_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region description of the rabbitmq instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the rabbitmq instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the rabbitmq instance.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id of the rabbitmq instance.",
						},
						"zone_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone description of the rabbitmq instance.",
						},
						"storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total storage space of the rabbitmq instance. Unit: GiB.",
						},
						"used_storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used storage space of the rabbitmq instance. Unit: GiB.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the rabbitmq instance.",
						},
						"charge_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The charge detail information of the rabbitmq instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge type of the rabbitmq instance.",
									},
									"charge_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge status of the rabbitmq instance.",
									},
									"auto_renew": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to automatically renew in prepaid scenarios.",
									},
									"charge_start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge start time of the rabbitmq instance.",
									},
									"charge_end_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge end time of the rabbitmq instance.",
									},
									"charge_expire_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge expire time of the rabbitmq instance.",
									},
									"overdue_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The overdue time of the rabbitmq instance.",
									},
									"overdue_reclaim_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The overdue reclaim time of the rabbitmq instance.",
									},
								},
							},
						},
						"endpoints": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The endpoint info of the rabbitmq instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The endpoint type of the rabbitmq instance.",
									},
									"internal_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The internal endpoint of the rabbitmq instance.",
									},
									"network_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The network type of the rabbitmq instance.",
									},
									"public_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public endpoint of the rabbitmq instance.",
									},
								},
							},
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the rabbitmq instance.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRabbitmqInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRabbitmqInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRabbitmqInstances())
}

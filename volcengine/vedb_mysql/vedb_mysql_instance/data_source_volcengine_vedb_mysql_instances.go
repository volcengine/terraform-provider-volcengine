package vedb_mysql_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVedbMysqlInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVedbMysqlInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the veDB Mysql instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the veDB Mysql instance.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the veDB Mysql instance.",
			},
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the veDB Mysql instance.",
			},
			"create_time_start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of creating veDB Mysql instance.",
			},
			"create_time_end": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of creating veDB Mysql instance.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone of the veDB Mysql instance.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The charge type of the veDB Mysql instance.",
			},
			"tags": ve.TagsSchema(),
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the veDB Mysql instance.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of veDB mysql instance.",
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
			"instances": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the veDB Mysql instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the veDB Mysql instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the veDB Mysql instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the veDB Mysql instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the veDB Mysql instance.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine version of the veDB Mysql instance.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the veDB Mysql instance.",
						},
						"zone_ids": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available zone of the veDB Mysql instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc ID of the veDB Mysql instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet ID of the veDB Mysql instance.",
						},
						"time_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time zone.",
						},
						"lower_case_table_names": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the table name is case sensitive, the default value is 1.\nRanges:\n0: Table names are stored as fixed and table names are case-sensitive.\n1: Table names will be stored in lowercase and table names are not case sensitive.",
						},
						"tags": ve.TagsSchemaComputed(),
						"storage_charge_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Storage billing type. " +
								"Values:\nPostPaid: Pay-as-you-go (postpaid).\n" +
								"PrePaid: Monthly/yearly subscription (prepaid).",
						},
						"pre_paid_storage_in_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total storage capacity in GiB for prepaid services.",
						},
						"storage_used_gib": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Used storage size, unit: GiB.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id.",
						},
						"charge_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Calculate the billing type. Values:\n" +
								"PostPaid: Pay-as-you-go (postpaid).\n" +
								"PrePaid: Monthly/yearly subscription (prepaid).",
						},
						"charge_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Payment status:\nNormal: Normal.\nOverdue: In arrears.\nShutdown: Shut down.",
						},
						"overdue_reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expected release time when shut down due to arrears. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Overdue shutdown time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"auto_renew": {
							Type:     schema.TypeBool,
							Computed: true,
							Description: "Whether auto-renewal is enabled in the prepaid scenario. " +
								"Values:\ntrue: Auto-renewal is enabled.\nfalse: Auto-renewal is not enabled.",
						},
						"charge_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when billing starts. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"charge_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing expiration time in the prepaid scenario, in the format: yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"nodes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Detailed information of instance nodes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the node.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The zone id.",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node type. Values:\nPrimary: Primary node.\nReadOnly: Read-only node.",
									},
									"v_cpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "CPU size. For example, when the value is 1, it means the CPU size is 1U.",
									},
									"memory": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Memory size, in GiB.",
									},
									"node_spec": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node specification of an instance.",
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

func dataSourceVolcengineVedbMysqlInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVedbMysqlInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVedbMysqlInstances())
}

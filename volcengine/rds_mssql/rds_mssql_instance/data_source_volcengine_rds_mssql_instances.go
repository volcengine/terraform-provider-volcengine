package sqlserver_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMssqlInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineSqlserverInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Id of the instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the instance.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Status of the instance.",
			},
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Compatible version. Currently only supports the value SQLServer_2019_Std, which represents SQL Server 2019 Standard Edition.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance type. Currently only supports the value HA, which represents high availability type.",
			},
			"create_time_start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of creating the instance, using UTC time format.",
			},
			"create_time_end": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of creating the instance, using UTC time format.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the zone.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The charge type.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Tag query array object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The tag key of the instance tag.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The tag value of the instance tag.",
						},
					},
				},
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
							Description: "The id of the instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the instance.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The db engine version.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the instance.",
						},
						"node_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node spec.",
						},
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The port of the instance.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id.",
						},
						"server_collation": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server sorting rules.",
						},
						"storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The storage space.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The storage type.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id.",
						},
						"time_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time zone.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id.",
						},
						"charge_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The charge detail.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge type.",
									},
									"auto_renew": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable automatic renewal in the prepaid scenario. This parameter can be set when ChargeType is Prepaid.",
									},
									"period_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Purchase cycle in prepaid scenarios. This parameter can be set when ChargeType is Prepaid.\nMonth: represents monthly (default).\nYear: represents yearly.",
									},
									"period": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Purchase duration in a prepaid scenario.",
									},
									"charge_start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Charge start time.",
									},
									"charge_end_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Charge end time.",
									},
									"charge_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The charge status.",
									},
									"overdue_reclaim_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Expected release time when overdue fees are shut down.",
									},
									"overdue_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Time for Disconnection due to Unpaid Fees",
									},
								},
							},
						},
						"node_detail_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Node detail information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"create_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node creation time.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Instance ID.",
									},
									"memory": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The Memory.",
									},
									"node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Node ID.",
									},
									"node_spec": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node spec.",
									},
									"node_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node status.",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node type.",
									},
									"region_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region id.",
									},
									"update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The update time.",
									},
									"vcpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "CPU size. For example: 1 represents 1U.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The zone id.",
									},
									"node_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node ip.",
									},
								},
							},
						},
						"parameter_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of instance parameters.",
						},
						"parameters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of instance parameters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the parameter.",
									},
									"parameter_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of the parameter.",
									},
									"parameter_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the parameter.",
									},
									"parameter_description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the parameter.",
									},
									"parameter_default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The default value of the parameter.",
									},
									"checking_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The valid value range of the parameter.",
									},
									"force_modify": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether the parameter running value can be modified.",
									},
									"force_restart": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether the instance needs to be restarted to take effect after modifying the running value of the parameter.",
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

func dataSourceVolcengineSqlserverInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMssqlInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMssqlInstances())
}

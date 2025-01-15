package rds_mysql_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlInstancesRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of RDS instance.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of RDS instance query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the RDS instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the RDS instance.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the RDS instance.",
			},
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the RDS instance.",
			},
			"create_time_start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of creating RDS instance.",
			},
			"create_time_end": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of creating RDS instance.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone of the RDS instance.",
			},
			"charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The charge type of the RDS instance.",
				ValidateFunc: validation.StringInSlice([]string{"PostPaid", "PrePaid"}, false),
			},
			"tags": ve.TagsSchema(),

			"rds_mysql_instances": {
				Description: "The collection of RDS instance query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RDS instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the RDS instance.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the RDS instance.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available zone of the RDS instance.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine version of the RDS instance.",
						},
						"v_cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "CPU size.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory size.",
						},
						"node_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The specification of primary node.",
						},
						"node_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of nodes.",
						},
						"zone_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "List of availability zones where each node of the instance is located.",
						},
						"node_cpu_used_percentage": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Average CPU usage of the instance master node in nearly one minute.",
						},
						"node_memory_used_percentage": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Average memory usage of the instance master node in nearly one minute.",
						},
						"node_space_used_percentage": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Average disk usage of the instance master node in nearly one minute.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the RDS instance.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the RDS instance.",
						},
						"storage_use": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The instance has used storage space. Unit: GB.",
						},
						"backup_use": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The instance has used backup space. Unit: GB.",
						},
						"storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total instance storage space. Unit: GB.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance storage type.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc ID of the RDS instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet ID of the RDS instance.",
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
						"data_sync_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data synchronization mode.",
						},
						"allow_list_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of allow list.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the RDS instance.",
						},
						"tags": ve.TagsSchemaComputed(),
						"charge_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							MinItems:    1,
							Description: "Payment methods.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Payment type. Value:\nPostPaid - Pay-As-You-Go\nPrePaid - Yearly and monthly (default).",
									},
									"auto_renew": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to automatically renew in prepaid scenarios.\nAutorenew_Enable\nAutorenew_Disable (default).",
									},
									"period_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The purchase cycle in the prepaid scenario.\nMonth - monthly subscription (default)\nYear - Package year.",
									},
									"period": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Purchase duration in prepaid scenarios. Default: 1.",
									},
									"charge_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Pay status. Value:\nnormal - normal\noverdue - overdue\n.",
									},
									"charge_start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Billing start time (pay-as-you-go & monthly subscription).",
									},
									"charge_end_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Billing expiry time (yearly and monthly only).",
									},
									"overdue_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Shutdown time in arrears (pay-as-you-go & monthly subscription).",
									},
									"overdue_reclaim_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Estimated release time when arrears are closed (pay-as-you-go & monthly subscription).",
									},
									"temp_modify_start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Temporary upgrade start time.",
									},
									"temp_modify_end_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Restore time of temporary upgrade.",
									},
								},
							},
						},
						"maintenance_window": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Maintenance Window.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"maintenance_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The maintainable time of the RDS instance.",
									},
									"day_kind": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "DayKind of maintainable window. Value: Week. Month.",
									},
									"day_of_week": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Days of maintainable window of the week.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"day_of_month": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Days of maintainable window of the month.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"connection_pool_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Connection pool type.",
						},
						"binlog_dump": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Does it support the binlog capability? This parameter is returned only when the database proxy is enabled. Values:\ntrue: Yes.\nfalse: No.",
						},
						"global_read_only": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable global read-only.\ntrue: Yes.\nfalse: No.",
						},
						"db_proxy_status": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The running status of the proxy instance. " +
								"This parameter is returned only when the database proxy is enabled. " +
								"Values:\nCreating: The proxy is being started.\n" +
								"Running: The proxy is running.\nShutdown: The proxy is closed.\n" +
								"Deleting: The proxy is being closed.",
						},
						"check_modify_db_proxy_allowed": {
							Type:     schema.TypeList,
							Computed: true,
							Description: "Is execution of the ModifyDBProxy interface allowed:\n" +
								"Allowed: If it is closed, return whether the proxy can be enabled. " +
								"If it is enabled, return whether the proxy can be disabled. Values: " +
								"true (yes); false (no).\nReason: When Allowed is false, " +
								"return the specific reason.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the ModifyDBProxy interface can be executed.",
									},
									"reason": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The reason why the ModifyDBProxy interface cannot be executed.",
									},
								},
							},
						},
						"feature_states": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Feature status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"feature_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature name.",
									},
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether it is enabled. Values:\ntrue: Enabled.\nfalse: Disabled.",
									},
									"support": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether it support this function. Value:\ntrue: Supported.\nfalse: Not supported.",
									},
								},
							},
						},
						"endpoints": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The endpoint info of the RDS instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Instance connection terminal ID.",
									},
									"endpoint_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The instance connection terminal name.",
									},
									"endpoint_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Terminal type:\nCluster: The default terminal. (created by default)\nPrimary: Primary node terminal.\nCustom: Custom terminal.\nDirect: Direct connection to the terminal. (Only the operation and maintenance side)\nAllNode: All node terminals. (Only the operation and maintenance side).",
									},
									"read_write_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Read and write mode:\nReadWrite: read and write\nReadOnly: read only (default).",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Address description.",
									},
									"auto_add_new_nodes": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.",
									},
									"enable_read_write_splitting": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.",
									},
									"enable_read_only": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.",
									},
									"idle_connection_reclaim": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the idle connection reclaim function is enabled. true: Enabled. false: Disabled.",
									},
									"node_weight": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of nodes configured by the connection terminal and the corresponding read-only weights.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"node_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ID of the node.",
												},
												"node_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of the node.",
												},
												"weight": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The weight of the node.",
												},
											},
										},
									},
									"addresses": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Address list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"network_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Network address type, temporarily Private, Public, PublicService.",
												},
												"domain": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Connect domain name.",
												},
												"ip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP Address.",
												},
												"port": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Port.",
												},
												"subnet_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Subnet ID, valid only for private addresses.",
												},
												"eip_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ID of the EIP, only valid for Public addresses.",
												},
												"dns_visibility": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "DNS Visibility.",
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
							Description: "Instance node information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Instance ID.",
									},
									"node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node ID.",
									},
									"region_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Region ID, you can call the DescribeRegions query and use this parameter to specify the region where the instance is to be created.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Availability zone ID. Subsequent support for multi-availability zones can be separated and displayed by an English colon.",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node type. Value: Primary: Primary node.\nSecondary: Standby node.\nReadOnly: Read-only node.",
									},
									"node_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node state, value: aligned with instance state.",
									},
									"node_spec": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "General instance type, different from Custom instance type.",
									},
									"v_cpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "CPU size. For example: 1 means 1U.",
									},
									"memory": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Memory size in GB.",
									},
									"create_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node creation local time.",
									},
									"update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node updates local time.",
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

func dataSourceVolcengineRdsMysqlInstancesRead(d *schema.ResourceData, meta interface{}) error {
	rdsMysqlInstanceService := NewRdsMysqlInstanceService(meta.(*ve.SdkClient))
	return rdsMysqlInstanceService.Dispatcher.Data(rdsMysqlInstanceService, d, DataSourceVolcengineRdsMysqlInstances())
}

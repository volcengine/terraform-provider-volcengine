package rds_postgresql_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstancesRead,
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
				Description: "The id of the RDS PostgreSQL instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the RDS PostgreSQL instance.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the RDS PostgreSQL instance.",
			},
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the RDS PostgreSQL instance.",
			},
			"create_time_start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of creating RDS PostgreSQL instance.",
			},
			"create_time_end": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of creating RDS PostgreSQL instance.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone of the RDS PostgreSQL instance.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The charge type of the RDS instance.",
			},
			"tags": ve.TagsSchema(),
			"instances": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS PostgreSQL instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS PostgreSQL instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RDS PostgreSQL instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the RDS PostgreSQL instance.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the RDS PostgreSQL instance.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available zone of the RDS PostgreSQL instance.",
						},
						"zone_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "ID of the availability zone where each instance is located.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine version of the RDS PostgreSQL instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the RDS PostgreSQL instance.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the RDS PostgreSQL instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance type of the RDS PostgreSQL instance.",
						},
						"node_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Master node specifications.",
						},
						"node_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of nodes.",
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
							Description: "The vpc ID of the RDS PostgreSQL instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet ID of the RDS PostgreSQL instance.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the RDS PostgreSQL instance.",
						},
						"tags": ve.TagsSchemaComputed(),
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
						"backup_use": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The instance has used backup space. Unit: GB.",
						},
						"data_sync_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data synchronization mode.",
						},
						"charge_detail": {
							Type:        schema.TypeList,
							Computed:    true,
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
										Description: "Start time of temporary upgrade.",
									},
									"temp_modify_end_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Temporary upgrade of restoration time.",
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
									"read_only_node_weight": {
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
									"address": {
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
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstances())
}

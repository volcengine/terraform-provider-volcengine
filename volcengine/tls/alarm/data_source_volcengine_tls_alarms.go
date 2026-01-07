package alarm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsAlarms() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsAlarmsRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project id.",
			},
			"alarm_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The alarm id.",
			},
			"alarm_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The alarm name.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The topic id.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The topic name.",
			},
			"status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The status.",
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
			"alarms": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of alarms.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alarm_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the alarm.",
						},
						"severity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The severity of the alarm.",
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project id.",
						},
						"status": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the alert policy. The default value is true, that is, on.",
						},
						"trigger_period": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.",
						},
						"alarm_period": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.",
						},
						"alarm_notify_group": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of notification groups corresponding to the alarm.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alarm_notify_group_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the notification group.",
									},
									"alarm_notify_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the notify group.",
									},
									"notify_type": {
										Type:        schema.TypeList,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Computed:    true,
										Description: "The notify group type.",
									},
									"create_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The create time the notification.",
									},
									"modify_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The modification time the notification.",
									},
									"iam_project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The iam project name.",
									},
									"receivers": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of IAM users to receive alerts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"receiver_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The receiver type.",
												},
												"receiver_names": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of the receiver names.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"receiver_channels": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The list of the receiver channels.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"start_time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The start time.",
												},
												"end_time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The end time.",
												},
											},
										},
									},
								},
							},
						},
						"user_define_msg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Customize the alarm notification content.",
						},
						"query_request": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Search and analyze sentences, 1~3 can be configured.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the topic.",
									},
									"topic_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the topic.",
									},
									"query": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Query statement, the maximum supported length is 1024.",
									},
									"number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Alarm object sequence number; increments from 1.",
									},
									"start_time_offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The start time of the query range is relative to the current historical time, in minutes. The value is non-positive, the maximum value is 0, and the minimum value is -1440.",
									},
									"end_time_offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The end time of the query range is relative to the current historical time. The unit is minutes. The value is not positive and must be greater than StartTimeOffset. The maximum value is 0 and the minimum value is -1440.",
									},
									"time_span_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The time span type.",
									},
									"truncated_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The truncated time.",
									},
									"end_time_offset_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The end time offset unit.",
									},
									"start_time_offset_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The start time offset unit.",
									},
								},
							},
						},
						"condition": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Alarm trigger condition.",
						},
						"request_cycle": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The execution period of the alarm task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Execution cycle type.",
									},
									"time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The cycle of alarm task execution, or the time point of periodic execution. The unit is minutes, and the value range is 1~1440.",
									},
									"cron_tab": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The cron tab.",
									},
								},
							},
						},
						"alarm_period_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sms": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "SMS alarm cycle, the unit is minutes, and the value range is 10~1440.",
									},
									"phone": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Telephone alarm cycle, the unit is minutes, and the value range is 10~1440.",
									},
									"email": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Email alarm period, the unit is minutes, and the value range is 1~1440.",
									},
									"general_webhook": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Customize the webhook alarm period, the unit is minutes, and the value range is 1~1440.",
									},
								},
							},
						},
						"trigger_conditions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of trigger conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The condition.",
									},
									"count_condition": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The count condition.",
									},
									"severity": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The severity.",
									},
									"no_data": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The no data.",
									},
								},
							},
						},
						"join_configurations": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of join configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The condition.",
									},
									"set_operation_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The set operation type.",
									},
								},
							},
						},
						"send_resolved": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to send resolved.",
						},
						"alarm_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The alarm id.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modify time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsAlarmsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineTlsAlarmService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsAlarms())
}

package alarm_notify_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsAlarmNotifyGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsAlarmNotifyGroupsRead,
		Schema: map[string]*schema.Schema{
			"alarm_notify_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the alarm notify group.",
			},
			"alarm_notify_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the alarm notify group.",
			},
			"receiver_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the receiver.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the iam project.",
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
			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of the notify groups.",
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
									"end_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The end time.",
									},
									"start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The start time.",
									},
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
									"general_webhook_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The webhook url.",
									},
									"general_webhook_body": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The webhook body.",
									},
									"alarm_webhook_at_users": {
										Type:        schema.TypeList,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Computed:    true,
										Description: "The alarm webhook at users.",
									},
									"alarm_webhook_is_at_all": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The alarm webhook is at all.",
									},
									"general_webhook": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The general webhook.",
									},
									"general_webhook_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The general webhook method.",
									},
									"general_webhook_headers": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The general webhook headers.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key of the header.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value of the header.",
												},
											},
										},
									},
									"alarm_content_template_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The alarm content template id.",
									},
									"alarm_webhook_integration_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The alarm webhook integration id.",
									},
									"alarm_webhook_integration_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The alarm webhook integration name.",
									},
								},
							},
						},
						"notice_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of the notice rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"has_next": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to continue to the next level of condition judgment.",
									},
									"rule_node": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The rule node.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of the rule node.",
												},
												"value": {
													Type:        schema.TypeList,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Computed:    true,
													Description: "The value of the rule node.",
												},
												"children": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The children of the rule node.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of the rule node.",
															},
															"value": {
																Type:        schema.TypeList,
																Elem:        &schema.Schema{Type: schema.TypeString},
																Computed:    true,
																Description: "The value of the rule node.",
															},
														},
													},
												},
											},
										},
									},
									"has_end_node": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether there is an end node behind.",
									},
									"receiver_infos": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of IAM users to receive alerts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The end time.",
												},
												"start_time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The start time.",
												},
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
												"general_webhook_url": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The webhook url.",
												},
												"general_webhook_body": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The webhook body.",
												},
												"alarm_webhook_at_users": {
													Type:        schema.TypeList,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Computed:    true,
													Description: "The alarm webhook at users.",
												},
												"alarm_webhook_is_at_all": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "The alarm webhook is at all.",
												},
												"general_webhook": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The general webhook.",
												},
												"general_webhook_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The general webhook method.",
												},
												"general_webhook_headers": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The general webhook headers.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The key of the header.",
															},
															"value": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value of the header.",
															},
														},
													},
												},
												"alarm_content_template_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The alarm content template id.",
												},
												"alarm_webhook_integration_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The alarm webhook integration id.",
												},
												"alarm_webhook_integration_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The alarm webhook integration name.",
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

func dataSourceVolcengineTlsAlarmNotifyGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmNotifyGroupService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsAlarmNotifyGroups())
}

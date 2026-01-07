package alarm_notify_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls alarm notify group can be imported using the id, e.g.
```
$ terraform import volcengine_tls_alarm_notify_group.default fa************
```

*/

func ResourceVolcengineTlsAlarmNotifyGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceVolcengineTlsAlarmNotifyGroupRead,
		Create: resourceVolcengineTlsAlarmNotifyGroupCreate,
		Update: resourceVolcengineTlsAlarmNotifyGroupUpdate,
		Delete: resourceVolcengineTlsAlarmNotifyGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alarm_notify_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the notify group.",
			},
			"notify_type": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         schema.HashString,
				Description: "The notify type.\nTrigger: Alarm Trigger\nRecovery: Alarm Recovery.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "The name of the iam project.",
			},
			"receivers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of IAM users to receive alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"receiver_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The receiver type, Can be set as: `User`(The id of user).",
						},
						"receiver_names": {
							Type:        schema.TypeSet,
							Required:    true,
							Set:         schema.HashString,
							Description: "List of the receiver names.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"receiver_channels": {
							Type:        schema.TypeSet,
							Required:    true,
							Set:         schema.HashString,
							Description: "The list of the receiver channels. Currently supported channels: Email, Sms, Phone.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"start_time": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The start time.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The end time.",
						},
						"general_webhook_url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The webhook url.",
						},
						"general_webhook_body": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The webhook body.",
						},
						"alarm_webhook_at_users": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The alarm webhook at users.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"alarm_webhook_is_at_all": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The alarm webhook is at all.",
						},
						"general_webhook_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The general webhook method.",
						},
						"general_webhook_headers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The general webhook headers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The key of the header.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The value of the header.",
									},
								},
							},
						},
						"alarm_content_template_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The alarm content template id.",
						},
						"alarm_webhook_integration_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The alarm webhook integration id.",
						},
						"alarm_webhook_integration_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The alarm webhook integration name.",
						},
					},
				},
			},
			"notice_rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of the notice rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"has_next": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to continue to the next level of condition judgment.",
						},
						"rule_node": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The rule node.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The type of the rule node.",
									},
									"value": {
										Type:        schema.TypeList,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Optional:    true,
										Description: "The value of the rule node.",
									},
									"children": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The children of the rule node.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The type of the rule node.",
												},
												"value": {
													Type:        schema.TypeList,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Optional:    true,
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
							Optional:    true,
							Description: "Whether there is an end node behind.",
						},
						"receiver_infos": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of IAM users to receive alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The end time.",
									},
									"start_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The start time.",
									},
									"receiver_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The receiver type.",
									},
									"receiver_names": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of the receiver names.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"receiver_channels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of the receiver channels.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"general_webhook_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The webhook url.",
									},
									"general_webhook_body": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The webhook body.",
									},
									"alarm_webhook_at_users": {
										Type:        schema.TypeList,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Optional:    true,
										Description: "The alarm webhook at users.",
									},
									"alarm_webhook_is_at_all": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "The alarm webhook is at all.",
									},
									"general_webhook_method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The general webhook method.",
									},
									"general_webhook_headers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The general webhook headers.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The key of the header.",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the header.",
												},
											},
										},
									},
									"alarm_content_template_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The alarm content template id.",
									},
									"alarm_webhook_integration_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The alarm webhook integration id.",
									},
									"alarm_webhook_integration_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The alarm webhook integration name.",
									},
								},
							},
						},
					},
				},
			},
			"alarm_notify_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The alarm notification group id.",
			},
		},
	}
}

func resourceVolcengineTlsAlarmNotifyGroupCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmNotifyGroupService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsAlarmNotifyGroup()); err != nil {
		return fmt.Errorf("error on creating tls Alarm Notify Group %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsAlarmNotifyGroupRead(d, meta)
}

func resourceVolcengineTlsAlarmNotifyGroupRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmNotifyGroupService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsAlarmNotifyGroup()); err != nil {
		return fmt.Errorf("error on reading tls Alarm Notify Group %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineTlsAlarmNotifyGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmNotifyGroupService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(service, d, ResourceVolcengineTlsAlarmNotifyGroup()); err != nil {
		return fmt.Errorf("error on creating tls Alarm Notify Group %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsAlarmNotifyGroupRead(d, meta)
}

func resourceVolcengineTlsAlarmNotifyGroupDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmNotifyGroupService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsAlarmNotifyGroup()); err != nil {
		return fmt.Errorf("error on deleting tls Alarm Notify Group %q, %w", d.Id(), err)
	}
	return nil
}

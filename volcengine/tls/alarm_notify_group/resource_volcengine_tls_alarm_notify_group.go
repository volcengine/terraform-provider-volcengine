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
				Required:    true,
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
				Required:    true,
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

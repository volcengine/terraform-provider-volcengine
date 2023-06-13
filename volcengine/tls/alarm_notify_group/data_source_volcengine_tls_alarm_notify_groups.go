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
		},
	}
}

func dataSourceVolcengineTlsAlarmNotifyGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmNotifyGroupService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsAlarmNotifyGroups())
}

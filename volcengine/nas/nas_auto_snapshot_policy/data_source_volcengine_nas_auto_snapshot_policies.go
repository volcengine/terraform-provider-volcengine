package nas_auto_snapshot_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNasAutoSnapshotPolicys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNasAutoSnapshotPolicysRead,
		Schema: map[string]*schema.Schema{
			"auto_snapshot_policy_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of auto snapshot policy.",
			},
			"auto_snapshot_policy_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of auto snapshot policy.",
			},
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
			"auto_snapshot_polices": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of auto snapshot policy.",
						},
						"auto_snapshot_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of auto snapshot policy.",
						},
						"auto_snapshot_policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of auto snapshot policy.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of auto snapshot policy.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of auto snapshot policy.",
						},
						"time_points": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time points of auto snapshot policy. Unit: hour.",
						},
						"retention_days": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The retention days of auto snapshot policy. Unit: day.",
						},
						"repeat_weekdays": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat weekdays of auto snapshot policy. Unit: day.",
						},
						"file_system_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of file system which auto snapshot policy bind.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNasAutoSnapshotPolicysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNasAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineNasAutoSnapshotPolicys())
}

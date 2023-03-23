package scaling_lifecycle_hook

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineScalingLifecycleHooks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineScalingLifecycleHooksRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of lifecycle hook ids.",
			},
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Set:         schema.HashString,
				Description: "An id of scaling group id.",
			},
			"lifecycle_hook_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of lifecycle hook names.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of lifecycle hook.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of lifecycle hook query.",
			},
			"lifecycle_hooks": {
				Description: "The collection of lifecycle hook query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the lifecycle hook.",
						},
						"lifecycle_hook_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the lifecycle hook.",
						},
						"scaling_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling group.",
						},
						"lifecycle_hook_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the lifecycle hook.",
						},
						"lifecycle_hook_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The timeout of the lifecycle hook.",
						},
						"lifecycle_hook_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the lifecycle hook.",
						},
						"lifecycle_hook_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The policy of the lifecycle hook.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineScalingLifecycleHooksRead(d *schema.ResourceData, meta interface{}) error {
	lifecycleHookService := NewScalingLifecycleHookService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(lifecycleHookService, d, DataSourceVolcengineScalingLifecycleHooks())
}
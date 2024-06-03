package cloud_identity_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudIdentityGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudIdentityGroupsRead,
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"display_name"},
				Description:   "The name of cloud identity group.",
			},
			"display_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"group_name"},
				Description:   "The display name of cloud identity group.",
			},
			"join_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The join type of cloud identity group. Valid values: `Auto`, `Manual`.",
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
			"groups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity group.",
						},
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity group.",
						},
						"group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cloud identity group.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the cloud identity group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cloud identity group.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the cloud identity group.",
						},
						"join_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email of the cloud identity group.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the cloud identity group.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the cloud identity group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudIdentityGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudIdentityGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudIdentityGroups())
}

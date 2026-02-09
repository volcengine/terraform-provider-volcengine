package iam_user_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamUserGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamUserGroupsRead,
		Schema: map[string]*schema.Schema{
			"query": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Fuzzy query. Can query by user group name, display name or description.",
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
			"user_groups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the user group.",
						},
						"account_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the account.",
						},
						"user_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the user group.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the user group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the user group.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation date of the user group.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update date of the user group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamUserGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamUserGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineIamUserGroups())
}

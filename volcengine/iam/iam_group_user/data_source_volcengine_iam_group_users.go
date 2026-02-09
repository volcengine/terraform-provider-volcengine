package iam_group_user

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamGroupUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamGroupUsersRead,
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of user.",
			},
			"query": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Fuzzy search, supports searching for user group names, display names, and remarks.",
			},
			"user_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of user group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the user group.",
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
						"join_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The join date of the user group.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
		},
	}
}

func dataSourceVolcengineIamGroupUsersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamGroupUserService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineIamGroupUsers())
}

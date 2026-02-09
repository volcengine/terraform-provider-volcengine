package iam_user_group_attachment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamUserGroupUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamUserGroupUsersRead,
		Schema: map[string]*schema.Schema{
			"user_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of user group.",
			},
			"users": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of user.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the user.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the user.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the user.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the user.",
						},
						"join_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The join date of the user.",
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

func dataSourceVolcengineIamUserGroupUsersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamUserGroupAttachmentService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineIamUserGroupUsers())
}

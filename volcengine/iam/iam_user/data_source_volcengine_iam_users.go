package iam_user

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamUsersRead,
		Schema: map[string]*schema.Schema{
			"query": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Fuzzy query. Can query by user name, display name or description.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of user query.",
			},
			"users": {
				Description: "The collection of user.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the user.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the user.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create date of the user.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update date of the user.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Main account ID to which the sub-user belongs.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The trn of the user.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the user.",
						},
						"mobile_phone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The mobile phone of the user.",
						},
						"mobile_phone_is_verify": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number has been verified.",
						},
						"email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email of the user.",
						},
						"email_is_verify": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the email has been verified.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the user.",
						},
						"tags": ve.TagsSchema(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamUsersRead(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewIamUserService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(eipAddressService, d, DataSourceVolcengineIamUsers())
}

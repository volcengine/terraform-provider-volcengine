package iam_user

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamUsersRead,
		Schema: map[string]*schema.Schema{
			"user_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of user names.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of IAM.",
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
							Description: "The account id of the user.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The trn of the user.",
						},
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
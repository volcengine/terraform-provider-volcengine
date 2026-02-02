package iam_login_profile

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamLoginProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamLoginProfilesRead,
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user name.",
			},
			"login_profiles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of login profiles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user name.",
						},
						"login_allowed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The flag of login allowed.",
						},
						"password_reset_required": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is required reset password when next time login in.",
						},
						"password_expire_at": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The password expire at.",
						},
						"last_reset_password_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The last reset password time.",
						},
						"last_login_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last login date.",
						},
						"last_login_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last login ip.",
						},
						"login_locked": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The flag of login locked.",
						},
						"safe_auth_flag": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The flag of safe auth.",
						},
						"safe_auth_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of safe auth.",
						},
						"safe_auth_exempt_required": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The flag of safe auth exempt required.",
						},
						"safe_auth_exempt_unit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The unit of safe auth exempt.",
						},
						"safe_auth_exempt_duration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The duration of safe auth exempt.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create date.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update date.",
						},
						"user_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The user id.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamLoginProfilesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamLoginProfileService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamLoginProfiles())
}

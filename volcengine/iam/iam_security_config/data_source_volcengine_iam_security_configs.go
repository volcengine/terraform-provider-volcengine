package iam_security_config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamSecurityConfigs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamSecurityConfigsRead,
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user name.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"security_configs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of security configs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user name.",
						},
						"user_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The user id.",
						},
						"safe_auth_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of safe auth.",
						},
						"safe_auth_exempt_duration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The exempt duration of safe auth.",
						},
						"safe_auth_close": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of safe auth.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamSecurityConfigsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamSecurityConfigService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamSecurityConfigs())
}

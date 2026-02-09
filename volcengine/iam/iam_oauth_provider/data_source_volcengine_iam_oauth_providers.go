package iam_oauth_provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamOAuthProviders() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamOAuthProvidersRead,
		Schema: map[string]*schema.Schema{
			"oauth_provider_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the OAuth provider.",
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
			"providers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of OAuth providers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oauth_provider_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the OAuth provider.",
						},
						"provider_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the OAuth provider.",
						},
						"sso_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The SSO type of the OAuth provider.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of the OAuth provider.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the OAuth provider.",
						},
						"client_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The client id of the OAuth provider.",
						},
						"client_secret": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The client secret of the OAuth provider.",
						},
						"user_info_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user info url of the OAuth provider.",
						},
						"token_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The token url of the OAuth provider.",
						},
						"authorize_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The authorize url of the OAuth provider.",
						},
						"authorize_template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The authorize template of the OAuth provider.",
						},
						"scope": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scope of the OAuth provider.",
						},
						"identity_map_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The identity map type of the OAuth provider.",
						},
						"idp_identity_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The idp identity key of the OAuth provider.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The trn of the OAuth provider.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create date of the OAuth provider.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update date of the OAuth provider.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamOAuthProvidersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOAuthProviderService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamOAuthProviders())
}

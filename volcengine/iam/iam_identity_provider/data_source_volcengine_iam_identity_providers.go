package iam_identity_provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamIdentityProviders() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamIdentityProvidersRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of identity providers.",
			},
			"providers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of identity providers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The TRN of the identity provider.",
						},
						"provider_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the identity provider.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the identity provider.",
						},
						"idp_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The type of the identity provider.",
						},
						"sso_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The SSO type of the identity provider.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of the identity provider.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create date of the identity provider.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update date of the identity provider.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamIdentityProvidersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamIdentityProviderService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamIdentityProviders())
}

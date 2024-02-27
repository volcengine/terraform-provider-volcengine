package iam_saml_provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamSamlProviders() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamSamlProvidersRead,
		Schema: map[string]*schema.Schema{
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
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"saml_provider_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the SAML provider.",
						},
						"encoded_saml_metadata_document": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Metadata document, encoded in Base64.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the SAML provider.",
						},
						"sso_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "SSO types, 1. Role-based SSO, 2. User-based SSO.",
						},
						"status": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: "User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, " +
								"3. Disabled, is a required field when creating user SSO.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identity provider creation time, such as 20150123T123318Z.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identity provider update time, such as: 20150123T123318Z.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamSamlProvidersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamSamlProviderService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineIamSamlProviders())
}

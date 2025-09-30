package iam_oidc_provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamOidcProviders() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamOidcProvidersRead,
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
			"oidc_providers": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The trn of OIDC provider.",
						},
						"provider_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the OIDC provider.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the OIDC provider.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create date of the OIDC provider.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update date of the OIDC provider.",
						},
						"issuer_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the OIDC provider.",
						},
						"issuance_limit_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The issuance limit time of the OIDC provider.",
						},
						"client_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The client IDs of the OIDC provider.",
						},
						"thumbprints": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The thumbprints of the OIDC provider.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamOidcProvidersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOidcProviderService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineIamOidcProviders())
}

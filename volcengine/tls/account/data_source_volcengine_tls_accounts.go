package account

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsAccountsRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of tls account query.",
			},
			"tls_accounts": {
				Description: "The collection of tls account query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arch_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the log service architecture. Valid values: 2.0 (new architecture), 1.0 (old architecture).",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the log service. Valid values: Activated (already activated), NonActivated (not activated).",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsAccountsRead(d *schema.ResourceData, meta interface{}) error {
	tlsAccountService := NewTlsAccountService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(tlsAccountService, d, DataSourceVolcengineTlsAccounts())
}

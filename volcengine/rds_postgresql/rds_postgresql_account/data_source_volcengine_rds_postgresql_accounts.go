package rds_postgresql_account

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlAccountsRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of database account query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the RDS instance.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the database account. This field supports fuzzy query.",
			},
			"accounts": {
				Description: "The collection of RDS instance account query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the database account.",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the database account.",
						},
						"account_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the database account.",
						},
						"account_privileges": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The privileges of the database account.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlAccountsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlAccountService(meta.(*volc.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlAccounts())
}

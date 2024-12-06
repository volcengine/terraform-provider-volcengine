package vedb_mysql_account

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVedbMysqlAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVedbMysqlAccountsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the veDB Mysql instance.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the database account. This field supports fuzzy query.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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
			"accounts": {
				Description: "The collection of query.",
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
						"account_privileges": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The privilege detail list of RDS mysql instance account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"db_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of database.",
									},
									"account_privilege": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The privilege type of the account.",
									},
									"account_privilege_detail": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The privilege detail of the account.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVedbMysqlAccountsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVedbMysqlAccountService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVedbMysqlAccounts())
}

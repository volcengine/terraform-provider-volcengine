package rds_mysql_account_table_column_info

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlAccountTableColumnInfos() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlAccountTableColumnInfosRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the mysql instance.",
			},
			"db_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the database.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the account.",
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Specify the IP address for the account to access the database." +
					" The default value is %.",
			},
			"table_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the table.",
			},
			"column_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the column.",
			},
			"table_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "Specify the number of tables in the table column permission information to be returned." +
					" If it exceeds the setting, it will be truncated.",
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
			"table_infos": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_privileges": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The table permissions of the account.",
						},
						"table_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the table.",
						},
						"column_infos": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The column permission information of the account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"column_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the column.",
									},
									"account_privileges": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The column privileges of the account.",
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

func dataSourceVolcengineRdsMysqlAccountTableColumnInfosRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlAccountTableColumnInfoService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlAccountTableColumnInfos())
}

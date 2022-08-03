package rds_database_v2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsDatabases() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsDatabasesRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of RDS database.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of RDS database query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the RDS instance.",
			},
			"db_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the RDS database.",
			},
			"rds_databases": {
				Description: "The collection of RDS instance account query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the RDS database.",
						},
						"db_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RDS database.",
						},
						"db_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the RDS database.",
						},
						"character_set_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The character set of the RDS database.",
						},
						"databases_privileges_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of database account privileges.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the account to be authorized.",
									},
									"account_privilege": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Authorized database permission type, value: ReadWrite: read and write permission. ReadOnly: Read-only permission. DDLOnly: DDL permissions only. DMLOnly: DML permissions only. Custom: Custom permissions.",
									},
									"account_privilege_custom": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Database privilege string, required when AccountPrivilege is Custom, value: SELECTINSERTUPDATEDELETECREATEDROPREFERENCESINDEXALTERCREATE TEMPORARY TABLESLOCK TABLESEXECUTECREATE VIEWSHOW VIEWCREATE ROUTINEALTER ROUTINEEVENTTRIGGER\nillustrate:\nMultiple strings are separated by commas.",
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

func dataSourceVolcengineRdsDatabasesRead(d *schema.ResourceData, meta interface{}) error {
	databaseService := NewRdsDatabaseService(meta.(*volc.SdkClient))
	return databaseService.Dispatcher.Data(databaseService, d, DataSourceVolcengineRdsDatabases())
}

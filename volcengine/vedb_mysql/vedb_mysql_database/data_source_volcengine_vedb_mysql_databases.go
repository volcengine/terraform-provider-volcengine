package vedb_mysql_database

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVedbMysqlDatabases() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVedbMysqlDatabasesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance id.",
			},
			"db_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Database name.",
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
			"databases": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The name of the database. Naming rules:\n " +
								"Unique name. Start with a lowercase letter and end with a letter or number. " +
								"The length is within 2 to 64 characters.\n " +
								"Consist of lowercase letters, numbers, underscores (_), or hyphens (-).\n " +
								"The name cannot contain certain reserved words.",
						},
						"character_set_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Database character set: utf8mb4 (default), utf8, latin1, ascii.",
						},
						"databases_privileges": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Account name that requires authorization.",
									},
									"account_privilege": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Authorization database privilege types: " +
											"\nReadWrite: Read and write privilege.\n " +
											"ReadOnly: Read-only privilege.\n " +
											"DDLOnly: Only DDL privilege.\n " +
											"DMLOnly: Only DML privilege.\n " +
											"Custom: Custom privilege.",
									},
									"account_privilege_detail": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "The specific SQL operation permissions contained in the permission type are separated by English commas (,) between multiple strings.\n " +
											"When used as a request parameter in the CreateDatabase interface, " +
											"when the AccountPrivilege value is Custom, this parameter is required. " +
											"Value range (multiple selections allowed): SELECT, INSERT, UPDATE, DELETE," +
											" CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, " +
											"EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER. " +
											"When used as a return parameter in the DescribeDatabases interface, " +
											"regardless of the value of AccountPrivilege, the details of the SQL operation permissions contained in this permission type are returned. " +
											"For the specific SQL operation permissions contained in each permission type, " +
											"please refer to the account permission list.",
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

func dataSourceVolcengineVedbMysqlDatabasesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVedbMysqlDatabaseService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVedbMysqlDatabases())
}

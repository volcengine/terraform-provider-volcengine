package rds_database

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDatabases() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDatabasesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of database IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of database.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of database query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the RDS instance.",
			},
			"db_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the database.",
			},
			"databases": {
				Description: "The collection of RDS instance account query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the database.",
						},
						"db_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the database.",
						},
						"db_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the database.",
						},
						"character_set_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The character set of the database.",
						},
						"account_names": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account names of the database.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDatabasesRead(d *schema.ResourceData, meta interface{}) error {
	databaseService := NewDatabaseService(meta.(*volc.SdkClient))
	return databaseService.Dispatcher.Data(databaseService, d, DataSourceVolcengineDatabases())
}

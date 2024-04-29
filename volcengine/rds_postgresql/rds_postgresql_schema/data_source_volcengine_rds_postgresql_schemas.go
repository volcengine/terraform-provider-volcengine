package rds_postgresql_schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlSchemas() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlSchemasRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the instance.",
			},
			"db_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the database.",
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
			"schemas": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the database.",
						},
						"schema_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the schema.",
						},
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner of the schema.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlSchemasRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlSchemaService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlSchemas())
}

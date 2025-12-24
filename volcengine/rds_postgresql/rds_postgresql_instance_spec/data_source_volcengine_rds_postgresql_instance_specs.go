package rds_postgresql_instance_spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceSpecs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceSpecsRead,
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
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Primary availability zone ID.",
			},
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the RDS PostgreSQL instance.",
			},
			"spec_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance specification code.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Storage type, fixed to LocalSSD.",
			},
			"instance_specs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Available instance specs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of connections supported by the instance.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the RDS PostgreSQL instance.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The memory size of the instance. Unit: GB.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the region.",
						},
						"spec_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance specification code.",
						},
						"v_cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of vCPUs of the instance.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Supported availability zone ID.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Storage type, fixed to LocalSSD.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstanceSpecsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceSpecService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceSpecs())
}

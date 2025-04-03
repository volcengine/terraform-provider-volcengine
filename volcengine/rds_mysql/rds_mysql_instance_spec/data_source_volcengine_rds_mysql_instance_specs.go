package rds_mysql_instance_spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlInstanceSpecs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlInstanceSpecsRead,
		Schema: map[string]*schema.Schema{
			"db_engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Compatible version. Values:\nMySQL_5_7: MySQL 5.7 version. Default value.\nMySQL_8_0: MySQL 8.0 version.",
			},
			"spec_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance specification code.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Availability zone ID.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance type. The value is DoubleNode.",
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
			"instance_specs": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Default value of maximum number of connections.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compatible version. Values:\nMySQL_5_7: MySQL 5.7 version. Default value.\nMySQL_8_0: MySQL 8.0 version.",
						},
						"iops": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum IOPS per second.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance type. The value is DoubleNode.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory size, in GB.",
						},
						"qps": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Queries Per Second (QPS).",
						},
						"spec_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance specification code.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the region.",
						},
						"spec_family": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance specification type. Values:\nGeneral: Exclusive specification (formerly \"General Purpose\").\nShared: General specification (formerly \"Shared Type\").",
						},
						"spec_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the available zone where the specification is located includes the following statuses:\nNormal: On sale.\nSoldout: Sold out.",
						},
						"storage_max": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum storage space, in GB.",
						},
						"storage_min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minimum storage space, in GB.",
						},
						"storage_step": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Disk step size, in GB.",
						},
						"vcpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of vCPUs.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Availability zone ID.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsMysqlInstanceSpecsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlInstanceSpecService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlInstanceSpecs())
}

package rds_postgresql_database_endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlDatabaseEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlDatabaseEndpointsRead,
		Schema: map[string]*schema.Schema{
			// 可以保留一下，只用来根据 InstanceId 查询 Endpoints 的相关信息，忽略其他返回项
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the RDS PostgreSQL instance.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "The name of the endpoint to filter.",
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
			"endpoints": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS PostgreSQL database endpoint.",
						},
						"endpoint_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RDS PostgreSQL database endpoint.",
						},
						"endpoint_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The type of the RDS PostgreSQL database endpoint. " +
								"Valid values: `Custom`(custom endpoint), `Cluster`(default endpoint).",
						},
						"read_write_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ReadWrite or ReadOnly. Default value is ReadOnly.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Connect domain name.",
						},
						"dns_visibility": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable public network resolution.",
						},
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The endpoint port.",
						},
						"cross_region_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cross-region domain for private address.",
						},
						"read_only_node_distribution_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The distribution type of the read-only nodes.",
						},
						"read_only_node_max_delay_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ReadOnly node max delay seconds.",
						},
						"read_write_proxy_connection": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of proxy connections set for the terminal.",
						},
						"write_node_halt_writing": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the endpoint sends write requests to the write node.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlDatabaseEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlDatabaseEndpointService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlDatabaseEndpoints())
}

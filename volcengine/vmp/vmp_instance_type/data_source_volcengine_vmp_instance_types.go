package vmp_instance_type

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpInstanceTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpInstanceTypesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Instance Type IDs.",
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
			"instance_types": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of instance type.",
						},
						"retention_period": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Maximum data retention time.",
						},
						"dedicated": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the workspace is exclusive.",
						},
						"active_series": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of active sequences.",
						},
						"ingest_samples_per_second": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum write samples per second.",
						},
						"query_per_second": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum Query QPS.",
						},
						"availability_zone_replicas": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of zone.",
						},
						"replicas_per_zone": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Data replicas per az.",
						},
						"query_concurrency": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of concurrent queries.",
						},
						"scan_series_per_second": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of scan sequences per second.",
						},
						"scan_samples_per_second": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum scan samples per second.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVmpInstanceTypesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpInstanceTypes())
}

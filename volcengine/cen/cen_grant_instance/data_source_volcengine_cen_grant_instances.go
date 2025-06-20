package cen_grant_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCenGrantInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCenGrantInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the instance.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the instance. Valid values: `VPC`, `DCGW`.",
			},
			"instance_region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region ID of the instance.",
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
			"grant_rules": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cen_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen.",
						},
						"cen_owner_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner ID of the cen.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the instance.",
						},
						"instance_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region ID of the instance.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the grant rule.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCenGrantInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCenGrantInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCenGrantInstances())
}

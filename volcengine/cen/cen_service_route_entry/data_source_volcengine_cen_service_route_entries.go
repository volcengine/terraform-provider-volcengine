package cen_service_route_entry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCenServiceRouteEntries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCenServiceRouteEntriesRead,
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A cen ID.",
			},
			"destination_cidr_block": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsCIDR,
				Description:  "A destination cidr block.",
			},
			"service_region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A service region id.",
			},
			"service_vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A service VPC id.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cen service route entry.",
			},
			"service_route_entries": {
				Description: "The collection of cen service route entry query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cen_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cen ID of the cen service route entry.",
						},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination cidr block of the cen service route entry.",
						},
						"service_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service region id of the cen service route entry.",
						},
						"service_vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service VPC id of the cen service route entry.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cen service route entry.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the cen service route entry.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cen service route entry.",
						},
						"publish_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.",
						},
						"publish_to_instances": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The publish instances. A maximum of 100 can be uploaded in one request.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_region_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region where the cloud service access route needs to be published.",
									},
									"instance_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The network instance type that needs to be published for cloud service access routes. The values are as follows: `VPC`, `DCGW`.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud service access routes need to publish the network instance ID.",
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

func dataSourceVolcengineCenServiceRouteEntriesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCenServiceRouteEntryService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCenServiceRouteEntries())
}

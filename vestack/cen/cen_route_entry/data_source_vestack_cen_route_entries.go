package cen_route_entry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackCenRouteEntries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackCenRouteEntriesRead,
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A cen ID.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An instance ID.",
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"VPC", "DCGW"}, false),
				Description:  "An instance type.",
			},
			"instance_region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An instance region ID.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A destination cidr block.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cen route entry.",
			},
			"cen_route_entries": {
				Description: "The collection of cen route entry query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cen_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cen ID of the cen route entry.",
						},
						//"type": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "The type of the cen route entry.",
						//},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination cidr block of the cen route entry.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance type of the next hop of the cen route entry.",
						},
						"instance_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance region id of the next hop of the cen route entry.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance id of the next hop of the cen route entry.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cen route entry.",
						},
						"publish_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The publish status of the cen route entry.",
						},
						"as_path": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The AS path of the cen route entry.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVestackCenRouteEntriesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCenRouteEntryService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVestackCenRouteEntries())
}

package vpc_endpoint_zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivatelinkVpcEndpointZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivatelinkVpcEndpointZonesRead,
		Schema: map[string]*schema.Schema{
			"endpoint_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The endpoint id of query.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Returns the total amount of the data list.",
			},
			"vpc_endpoint_zones": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of vpc endpoint zone.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of vpc endpoint zone.",
						},
						"zone_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain of vpc endpoint zone.",
						},
						"zone_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of vpc endpoint zone.",
						},
						"service_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of vpc endpoint service.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of vpc endpoint.",
						},
						"network_interface_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network interface id of vpc endpoint.",
						},
						"network_interface_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network interface ip of vpc endpoint.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcenginePrivatelinkVpcEndpointZonesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVpcEndpointZoneService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcenginePrivatelinkVpcEndpointZones())
}

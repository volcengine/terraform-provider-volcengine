package vpc_endpoint_connection

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivatelinkVpcEndpointConnections() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivatelinkVpcEndpointConnectionsRead,
		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the vpc endpoint service.",
			},
			"endpoint_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the vpc endpoint.",
			},
			"endpoint_owner_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account id of the vpc endpoint.",
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
			"connections": {
				Description: "The list of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc endpoint service.",
						},
						"endpoint_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc endpoint.",
						},
						"endpoint_owner_account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the vpc endpoint.",
						},
						"endpoint_vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the vpc endpoint.",
						},
						"connection_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the connection.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the connection.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the connection.",
						},
						"zones": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The available zones.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the zone.",
									},
									"zone_domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of the zone.",
									},
									"zone_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the zone.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the subnet.",
									},
									"network_interface_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the network interface.",
									},
									"network_interface_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ip address of the network interface.",
									},
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the resource.",
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

func dataSourceVolcenginePrivatelinkVpcEndpointConnectionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVpcEndpointConnectionService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcenginePrivatelinkVpcEndpointConnections())
}

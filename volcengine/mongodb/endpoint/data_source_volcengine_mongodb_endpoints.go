package endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineMongoDBEndpointsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance ID to query.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of mongodb endpoint query.",
			},
			"endpoints": {
				Description: "The collection of mongodb endpoints query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of mongodb addresses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of mongodb connection.",
									},
									"address_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP of mongodb connection.",
									},
									"address_port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The port of mongodb connection.",
									},
									"address_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The connection type of mongodb.",
									},
									"eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The EIP ID bound to the instance's public network address.",
									},
									"node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node ID.",
									},
								},
							},
						},
						"endpoint_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of endpoint.",
						},
						"endpoint_str": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The endpoint information.",
						},
						"endpoint_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node type corresponding to the endpoint.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network type of endpoint.",
						},
						"object_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The object ID corresponding to the endpoint.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet ID.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPC ID.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineMongoDBEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewMongoDBEndpointService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineMongoDBEndpoints())
}
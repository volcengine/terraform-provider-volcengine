package private_zone_resolver_endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivateZoneResolverEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivateZoneResolverEndpointsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the private zone resolver endpoint.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc ID of the private zone resolver endpoint.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the private zone resolver endpoint.",
			},
			"direction": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The direction of the private zone resolver endpoint.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the endpoint.",
						},
						"endpoint_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The endpoint id.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the endpoint.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the endpoint.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the endpoint.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the endpoint.",
						},
						"direction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The direction of the endpoint.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the endpoint.",
						},
						"vpc_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc region of the endpoint.",
						},
						"security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The security group id of the endpoint.",
						},
						"ip_configs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of IP configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"az_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The availability zone id of the endpoint.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The subnet id of the endpoint.",
									},
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address of the endpoint.",
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

func dataSourceVolcenginePrivateZoneResolverEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewPrivateZoneResolverEndpointService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcenginePrivateZoneResolverEndpoints())
}

package vpc_gateway_endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVpcGatewayEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVpcGatewayEndpointsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of gateway endpoint IDs.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the vpc.",
			},
			"endpoint_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the gateway endpoint.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the gateway endpoint.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of gateway endpoint.",
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
			"tags": ve.TagsSchema(),
			"vpc_gateway_endpoints": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the gateway endpoint.",
						},
						"endpoint_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the gateway endpoint.",
						},
						"endpoint_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the gateway endpoint.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the gateway endpoint.",
						},
						"service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the gateway endpoint service.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the gateway endpoint service.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the gateway endpoint.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the gateway endpoint.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the gateway endpoint.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the gateway endpoint.",
						},
						"vpc_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc policy of the gateway endpoint.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVpcGatewayEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVpcGatewayEndpointService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVpcGatewayEndpoints())
}

package ipv6_gateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIpv6Gateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIpv6GatewaysRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ID list of the Ipv6Gateways.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Ipv6Gateway.",
			},
			"vpc_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ID list of the VPC which the Ipv6Gateway belongs to.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of the Ipv6Gateway.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Ipv6Gateway query.",
			},
			"ipv6_gateways": {
				Description: "The collection of Ipv6Gateway query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Ipv6Gateway.",
						},
						"ipv6_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Ipv6Gateway.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Name of the Ipv6Gateway.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Ipv6Gateway.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the VPC which the Ipv6Gateway belongs to.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Status of the Ipv6Gateway.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the Ipv6Gateway.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of the Ipv6Gateway.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIpv6GatewaysRead(d *schema.ResourceData, meta interface{}) error {
	ipv6GatewayService := NewIpv6GatewayService(meta.(*ve.SdkClient))
	return ipv6GatewayService.Dispatcher.Data(ipv6GatewayService, d, DataSourceVolcengineIpv6Gateways())
}

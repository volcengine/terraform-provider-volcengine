package vpn_gateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVpnGateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVpnGatewaysRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of VPN gateway ids.",
			},
			"vpn_gateway_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of VPN gateway names.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A VPC ID of the VPN gateway.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A subnet ID of the VPN gateway.",
			},
			"ip_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
				Description:  "A IP address of the VPN gateway.",
			},
			"tags": ve.TagsSchema(),
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of VPN gateway.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of VPN gateway query.",
			},
			"vpn_gateways": {
				Description: "The collection of VPN gateway query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPN gateway.",
						},
						"vpn_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPN gateway.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of the VPN gateway.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPC ID of the VPN gateway.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The subnet ID of the VPN gateway.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of VPN gateway.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of VPN gateway.",
						},
						"vpn_gateway_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the VPN gateway.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the VPN gateway.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the VPN gateway.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the VPN gateway.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth of the VPN gateway.",
						},
						"connection_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The connection count of the VPN gateway.",
						},
						"route_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The route count of the VPN gateway.",
						},
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The BillingType of the VPN gateway.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the VPN gateway.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lock reason of the VPN gateway.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time of the VPN gateway.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the VPN gateway.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVpnGatewaysRead(d *schema.ResourceData, meta interface{}) error {
	vpnGatewayService := NewVpnGatewayService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(vpnGatewayService, d, DataSourceVolcengineVpnGateways())
}

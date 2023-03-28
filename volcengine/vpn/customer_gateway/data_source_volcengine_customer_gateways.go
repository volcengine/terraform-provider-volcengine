package customer_gateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCustomerGateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCustomerGatewaysRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of customer gateway ids.",
			},
			"customer_gateway_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of customer gateway names.",
			},
			"ip_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
				Description:  "A IP address of the customer gateway.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of customer gateway.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of customer gateway query.",
			},
			"customer_gateways": {
				Description: "The collection of customer gateway query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of the customer gateway.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the customer gateway.",
						},
						"customer_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the customer gateway.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the customer gateway.",
						},
						"connection_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The connection count of the customer gateway.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of customer gateway.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of customer gateway.",
						},
						"customer_gateway_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the customer gateway.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the customer gateway.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the customer gateway.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCustomerGatewaysRead(d *schema.ResourceData, meta interface{}) error {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(customerGatewayService, d, DataSourceVolcengineCustomerGateways())
}

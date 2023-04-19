package ipv6_address

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIpv6Addresses() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIpv6AddressesRead,
		Schema: map[string]*schema.Schema{
			"associated_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the ECS instance that is assigned the IPv6 address.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Ipv6Address query.",
			},
			"ipv6_addresses": {
				Description: "The collection of Ipv6Address query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv6 address.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIpv6AddressesRead(d *schema.ResourceData, meta interface{}) error {
	ipv6AddressService := NewIpv6AddressService(meta.(*ve.SdkClient))
	return ipv6AddressService.Dispatcher.Data(ipv6AddressService, d, DataSourceVolcengineIpv6Addresses())
}

package ipv6_address_bandwidth

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIpv6AddressBandwidths() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIpv6AddressBandwidthsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Allocation IDs of the Ipv6 address width.",
			},
			"associated_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the associated instance.",
			},
			"associated_instance_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the associated instance.",
			},
			"isp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ISP of the ipv6 address.",
			},
			"ipv6_addresses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ipv6 addresses.",
			},
			"network_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network type of the ipv6 address.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of Vpc the ipv6 address in.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Ipv6AddressBandwidth query.",
			},
			"ipv6_address_bandwidths": {
				Description: "The collection of Ipv6AddressBandwidth query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Ipv6AddressBandwidth.",
						},
						"allocation_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Ipv6AddressBandwidth.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Peek bandwidth of the Ipv6 address.",
						},
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BillingType of the Ipv6 bandwidth.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The BusinessStatus of the Ipv6AddressBandwidth.",
						},
						"isp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ISP of the Ipv6AddressBandwidth.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the associated instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the associated instance.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv6 address.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The BusinessStatus of the Ipv6AddressBandwidth.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network type of the Ipv6AddressBandwidth.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Ipv6AddressBandwidth.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the Ipv6AddressBandwidth.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of the Ipv6AddressBandwidth.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Overdue time of the Ipv6AddressBandwidth.",
						},
						"delete_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete time of the Ipv6AddressBandwidth.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIpv6AddressBandwidthsRead(d *schema.ResourceData, meta interface{}) error {
	ipv6AddressBandwidthService := NewIpv6AddressBandwidthService(meta.(*ve.SdkClient))
	return ipv6AddressBandwidthService.Dispatcher.Data(ipv6AddressBandwidthService, d, DataSourceVolcengineIpv6AddressBandwidths())
}

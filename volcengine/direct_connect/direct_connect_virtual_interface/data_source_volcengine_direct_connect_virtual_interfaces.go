package direct_connect_virtual_interface

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDirectConnectVirtualInterfaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDirectConnectVirtualInterfacesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
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
			"direct_connect_connection_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The direct connect connection ID that associated with this virtual interface.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The direct connect gateway ID that associated with this virtual interface.",
			},
			"virtual_interface_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of virtual interface.",
			},
			"route_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The route type of virtual interface.",
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The VLAN ID of virtual interface.",
			},
			"local_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The local IP that associated with this virtual interface.",
			},
			"peer_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The peer IP that associated with this virtual interface.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The filter tag of direct connect virtual interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The tag key of cloud resource instance.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The tag value of cloud resource instance.",
						},
					},
				},
			},
			"virtual_interfaces": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID which this virtual interface belongs.",
						},
						"virtual_interface_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The virtual interface ID.",
						},
						"virtual_interface_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of virtual interface.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the virtual interface.",
						},
						"direct_connect_connection_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The direct connect connection ID which associated with this virtual interface.",
						},
						"direct_connect_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The direct connect gateway ID which associated with this virtual interface.",
						},
						"route_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The route type of this virtual interface.",
						},
						"vlan_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The VLAN ID of virtual interface.",
						},
						"local_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The local IP that associated with this virtual interface.",
						},
						"peer_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The peer IP that associated with this virtual interface.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of virtual interface.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of virtual interface.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of virtaul interface.",
						},
						"enable_bfd": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether enable BFD detect.",
						},
						"bfd_detect_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The BFD detect interval.",
						},
						"bfd_detect_multiplier": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The BFD detect times.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The band width limit of virtual interface,in Mbps.",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The tags that direct connect gateway added.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tag key.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tag value.",
									},
								},
							},
						},
						"enable_nqa": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether enable NQA detect.",
						},
						"nqa_detect_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The NQA detect interval.",
						},
						"nqa_detect_multiplier": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The NAQ detect times.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDirectConnectVirtualInterfacesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDirectConnectVirtualInterfaceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDirectConnectVirtualInterfaces())
}

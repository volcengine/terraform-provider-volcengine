package dnat_entry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDnatEntries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDnatEntriesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of DNAT entry ids.",
			},
			"dnat_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the DNAT entry.",
			},
			"nat_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the NAT gateway.",
			},
			"external_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Provides the public IP address for public network access.",
			},
			"external_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.",
			},
			"internal_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Provides the internal IP address.",
			},
			"internal_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port or port segment on which the cloud server instance provides services to the public network.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network protocol.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of snat entries query.",
			},
			"dnat_entries": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of DNAT entries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dnat_entry_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the DNAT entry.",
						},
						"dnat_entry_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the DNAT entry.",
						},
						"nat_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the NAT gateway.",
						},
						"external_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Provides the public IP address for public network access.",
						},
						"external_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.",
						},
						"internal_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Provides the internal IP address.",
						},
						"internal_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The port or port segment on which the cloud server instance provides services to the public network.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network protocol.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDnatEntriesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDnatEntryService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineDnatEntries())
}
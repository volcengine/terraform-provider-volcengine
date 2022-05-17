package snat_entry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackSnatEntries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackSnatEntriesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of SNAT entry ids.",
			},
			"snat_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A name of SNAT entry.",
			},
			"nat_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of the nat gateway to which the entry belongs.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of the subnet that is required to access the Internet.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of the public ip address used by the SNAT entry.",
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
			"snat_entries": {
				Description: "The collection of snat entries.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the SNAT entry.",
						},
						"snat_entry_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the SNAT entry.",
						},
						"snat_entry_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the SNAT entry.",
						},
						"nat_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the nat gateway to which the entry belongs.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the subnet that is required to access the internet.",
						},
						"eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the public ip address used by the SNAT entry.",
						},
						"eip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public ip address used by the SNAT entry.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the SNAT entry.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVestackSnatEntriesRead(d *schema.ResourceData, meta interface{}) error {
	snatEntryService := NewSnatEntryService(meta.(*ve.SdkClient))
	return snatEntryService.Dispatcher.Data(snatEntryService, d, DataSourceVestackSnatEntries())
}

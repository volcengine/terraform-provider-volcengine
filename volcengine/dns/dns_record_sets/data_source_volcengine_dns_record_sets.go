package dns_record_sets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDnsRecordSets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDnsRecordSetsRead,
		Schema: map[string]*schema.Schema{
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
			"record_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The record set ID.",
			},
			"zid": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The domain ID.",
			},
			"search_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"like", "exact"}, false),
				Description:  "The matching mode for Host.",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain prefix of the record set.",
			},
			"record_sets": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the DNS record set.",
						},
						"pqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain prefix contained in the DNS record set, in PQDN (Partially Qualified Domain Name) format.",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host record contained in the DNS record set.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of DNS records in the DNS record set.",
						},
						"line": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The line code corresponding to the DNS record set.",
						},
						"weight_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether load balancing is enabled for the DNS record set.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDnsRecordSetsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDnsRecordSetsService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDnsRecordSets())
}

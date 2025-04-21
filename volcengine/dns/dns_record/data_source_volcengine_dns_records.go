package dns_record

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDnsRecords() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDnsRecordsRead,
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
			"search_order": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Method to sort the returned list of DNS records.",
			},
			"zid": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the domain.",
			},
			"search_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The matching mode for the Host parameter.",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Domain prefix of the DNS record.",
			},
			"line": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Line of the DNS record.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the DNS record.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Value of the DNS record.",
			},
			"records": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the domain.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The most recent update time of the domain.",
						},
						"record_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the DNS record.",
						},
						"pqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hostname included in the DNS record, in PQDN (Partially Qualified Domain Name) format.",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host record included in the DNS record.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the DNS record.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Time to Live (TTL) of the DNS record. The unit is seconds.",
						},
						"line": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The line code corresponding to the DNS record.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The record value contained in the DNS record.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The weight of the DNS record.",
						},
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the DNS record is enabled.",
						},
						"record_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the record set to which the DNS record belongs.",
						},
						"tags": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The tag information of the DNS record.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remark of the DNS record.",
						},
						"operators": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The account ID that called this API.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDnsRecordsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDnsRecordService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDnsRecords())
}

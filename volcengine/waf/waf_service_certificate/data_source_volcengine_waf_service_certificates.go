package waf_service_certificate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafServiceCertificates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafServiceCertificatesRead,
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
			"data": {
				Description: "The Information of the certificate.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"applicable_domains": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Associate the domain name of this certificate.",
						},
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the certificate.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the certificate.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the certificate.",
						},
						"expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expiration time of the certificate.",
						},
						"insert_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the certificate was added.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineWafServiceCertificatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafServiceCertificateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafServiceCertificates())
}

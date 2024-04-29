package cdn_certificate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCdnCertificates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCdnCertificatesRead,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Specify the location for storing the certificate. " +
					"The parameter can take the following values: " +
					"`volc_cert_center`: indicates that the certificate will be stored in the certificate center." +
					"`cdn_cert_hosting`: indicates that the certificate will be hosted on the content delivery network.",
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Specify a domain to obtain certificates that include that domain in the SAN field. " +
					"The domain can be a wildcard domain. For example, " +
					"specifying *.example.com will obtain certificates that include img.example.com or www.example.com in the SAN field.",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Specify one or more states to retrieve certificates in those states. " +
					"By default, all certificates in all states are returned. " +
					"You can specify the following states. Multiple states are separated by commas. " +
					"running: Retrieves certificates with a validity period greater than 30 days. " +
					"expired: Retrieves certificates that have already expired. " +
					"expiring_soon: Retrieves certificates with a validity period less than or equal to 30 days but have not yet expired.",
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
			"cert_info": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Specify the location for storing the certificate. " +
								"The parameter can take the following values: " +
								"`volc_cert_center`: indicates that the certificate will be stored in the certificate center." +
								"`cdn_cert_hosting`: indicates that the certificate will be hosted on the content delivery network.",
						},
						"cert_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID indicating the certificate.",
						},
						"cert_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name to which the certificate is issued.",
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Specify one or more states to retrieve certificates in those states. " +
								"By default, all certificates in all states are returned. " +
								"You can specify the following states. Multiple states are separated by commas. " +
								"running: Retrieves certificates with a validity period greater than 30 days. " +
								"expired: Retrieves certificates that have already expired. " +
								"expiring_soon: Retrieves certificates with a validity period less than or equal to 30 days but have not yet expired.",
						},
						"effective_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The issuance time of the certificate is indicated. The unit is Unix timestamp.",
						},
						"expire_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The expiration time of the certificate is indicated. The unit is Unix timestamp.",
						},
						"dns_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain names included in the SAN field of the certificate.",
						},
						"desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remark of the cert.",
						},
						"configured_domain": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The domain name associated with the certificate. " +
								"If the certificate is not yet associated with any domain name, " +
								"the parameter value is null.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCdnCertificatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCdnCertificateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCdnCertificates())
}

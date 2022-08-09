package cr_domain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCrDomains() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCrDomainsRead,
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The CR instance name.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of tag query.",
			},
			"domains": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of repository query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain of image repository.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of domain.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCrDomainsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCrDomainService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCrDomains())
}

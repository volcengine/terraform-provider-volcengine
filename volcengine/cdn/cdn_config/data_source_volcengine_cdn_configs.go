package cdn_config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCdnConfigs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCdnConfigsRead,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name.",
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
			"domain_config": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service type of the domain.",
						},
						"cname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cname of the domain.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The create time of the domain.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name.",
						},
						"lock_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates whether the configuration of this domain name is allowed to be changed.",
						},
						"project": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name.",
						},
						"service_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service region of the domain.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the domain.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of the domain.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCdnConfigsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCdnConfigService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCdnConfigs())
}

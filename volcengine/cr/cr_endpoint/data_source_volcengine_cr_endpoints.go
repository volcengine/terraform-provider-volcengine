package cr_endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCrEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCrEndpointsRead,
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
			"endpoints": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of endpoint query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"registry": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of CR instance.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether public endpoint is enabled.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of public endpoint.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCrEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCrEndpointService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCrEndpoints())
}
package apig_upstream_version

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineApigUpstreamVersions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineApigUpstreamVersionsRead,
		Schema: map[string]*schema.Schema{
			"upstream_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the apig upstream.",
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
			"versions": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of apig upstream version.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of apig upstream version.",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The labels of apig upstream version.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of apig upstream version label.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of apig upstream version label.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineApigUpstreamVersionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewApigUpstreamVersionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineApigUpstreamVersions())
}

package available_resource

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAvailableResources() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAvailableResourcesRead,
		Schema: map[string]*schema.Schema{
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of instance.",
			},
			"bandwith_limit": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The limit of bandwidth.",
			},
			"cloud_disk_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"CloudHDD", "CloudSSD"}, false),
				Description:  "The type of storage. The value can be `CloudHDD` or `CloudSSD`.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of resource query.",
			},
			"regions": {
				Description: "The collection of resource query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of country.",
							Elem:        geoInfo,
						},
						"area": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of area.",
							Elem:        geoInfo,
						},
						"city": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of city.",
							Elem:        geoInfo,
						},
						"isp": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of isp.",
							Elem:        geoInfo,
						},
						"cluster": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The config of cluster.",
							Elem:        geoInfo,
						},
					},
				},
			},
		},
	}
}

var geoInfo = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The id of region.",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name of region.",
		},
		"en_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The english name of region.",
		},
	},
}

func dataSourceVolcengineAvailableResourcesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewResourceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineAvailableResources())
}
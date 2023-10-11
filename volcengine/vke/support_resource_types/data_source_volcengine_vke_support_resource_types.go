package support_resource_types

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVkeVkeSupportResourceTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVkeSupportResourceTypesRead,
		Schema: map[string]*schema.Schema{
			"zone_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list of zone ids. If no parameter value, all available regions is returned.",
			},
			"resource_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list of resource types. Support Ecs or Zone.",
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
			"resources": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_scope": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scope of resource.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of zone.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource.",
						},
						"resource_specifications": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "The resource specifications info.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVkeSupportResourceTypesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVkeVkeSupportResourceTypes())
}

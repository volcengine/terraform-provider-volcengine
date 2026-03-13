package nlb_tag

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNlbTags() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNlbTagsRead,
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the resource.",
			},
			"resource_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of resource ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tag_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the tag. Valid values: `custom`, `system`.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of tag filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the tag.",
						},
						"values": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The values of the tag.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the resource.",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the resource.",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The key of the tag.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the tag.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of tags.",
			},
		},
	}
}

func dataSourceVolcengineNlbTagsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbTagService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNlbTags())
}

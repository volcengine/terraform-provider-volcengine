package iam_tag

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamTags() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamTagsRead,
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the resource. Valid values: User, Role.",
			},
			"resource_names": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The names of the resource.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"next_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The next token of query.",
			},
			"resource_tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the resource.",
						},
						"resource_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the resource.",
						},
						"tag_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The key of the tag.",
						},
						"tag_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the tag.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamTagsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamTagService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamTags())
}

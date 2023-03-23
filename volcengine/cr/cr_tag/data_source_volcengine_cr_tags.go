package cr_tag

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCrTags() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCrTagsRead,
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The CR instance name.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The CR namespace.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository name.",
			},
			"names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of instance names.",
			},
			"types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"Image", "Chart"}, false),
				},
				Set:         schema.HashString,
				Description: "The list of OCI product tag type.",
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
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of repository query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of OCI product tag.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of OCI product tag.",
						},
						"digest": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The digest of OCI product.",
						},
						"push_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last push time of OCI product.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of OCI product.",
						},
						"image_attributes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of image attributes,valid when tag type is Image.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"author": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The image author.",
									},
									"architecture": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The image architecture.",
									},
									"os": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The iamge os.",
									},
									"digest": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The digest of image.",
									},
								},
							},
						},
						"chart_attribute": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The chart attribute,valid when tag type is Chart.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Helm version.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Helm Chart name.",
									},
									"version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Helm Chart version.",
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

func dataSourceVolcengineCrTagsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCrTagService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCrTags())
}
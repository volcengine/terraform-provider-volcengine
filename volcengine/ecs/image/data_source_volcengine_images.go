package image

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineImages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineImagesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Image IDs.",
			},
			"os_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The operating system type of Image.",
			},
			"visibility": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The visibility of Image.",
			},
			"instance_type_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The specification of  Instance.",
			},
			"is_support_cloud_init": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the Image support cloud-init.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Image status.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Image.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Image query.",
			},
			"images": {
				Description: "The collection of Image query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of Image.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of Image.",
						},
						"image_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Image.",
						},
						"image_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of Image.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of Image.",
						},
						"os_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The operating system type of Image.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of Image.",
						},
						"visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The visibility of Image.",
						},
						"architecture": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The architecture of Image.",
						},
						"platform": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The platform of Image.",
						},
						"platform_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The platform version of Image.",
						},
						"os_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of Image operating system.",
						},
						"is_support_cloud_init": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the Image support cloud-init.",
						},
						"share_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The share mode of Image.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size(GiB) of Image.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineImagesRead(d *schema.ResourceData, meta interface{}) error {
	imageService := NewImageService(meta.(*ve.SdkClient))
	return imageService.Dispatcher.Data(imageService, d, DataSourceVolcengineImages())
}

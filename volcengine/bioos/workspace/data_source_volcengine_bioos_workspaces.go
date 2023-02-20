package workspace

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineBioosWorkspaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineBioosWorkspacesRead,
		Schema: map[string]*schema.Schema{
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Name",
					"CreateTime",
				}, false),
				Description: "Sort Field (Name CreateTime).",
			},
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Asc",
					"Desc",
				}, false),
				Description: "The sort order.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Vpc query.",
			},
			"keyword": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the workspace.",
			},
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of workspace ids.",
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of workspaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the workspace.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the workspace.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the workspace.",
						},
						"owner_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the owner of the workspace.",
						},
						"cover_download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the cover.",
						},
						"role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The role of the user.",
						},
						"s3_bucket": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "S3 bucket address.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The creation time of the workspace.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of the workspace.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineBioosWorkspacesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineBioosWorkspaceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineBioosWorkspaces())
}

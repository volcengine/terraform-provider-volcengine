package image_share_permission

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineImageSharePermissions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineImageSharePermissionsRead,
		Schema: map[string]*schema.Schema{
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the image.",
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

			"accounts": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The shared account id of the image.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineImageSharePermissionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewImageSharePermissionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineImageSharePermissions())
}

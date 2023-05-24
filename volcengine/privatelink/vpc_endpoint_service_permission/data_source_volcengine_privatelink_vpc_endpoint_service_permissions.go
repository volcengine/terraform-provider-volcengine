package vpc_endpoint_service_permission

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivatelinkVpcEndpointServicePermissions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivatelinkVpcEndpointServicePermissionRead,
		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Id of service.",
			},
			"permit_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Id of permit account.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Returns the total amount of the data list.",
			},
			"permissions": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permit_account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The permit account id.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcenginePrivatelinkVpcEndpointServicePermissionRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcenginePrivatelinkVpcEndpointServicePermissions())
}

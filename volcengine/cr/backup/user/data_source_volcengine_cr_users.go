package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCrUsers() *schema.Resource {
	return &schema.Resource{
		Read: func(data *schema.ResourceData, meta interface{}) error {
			service := NewCrUserService(meta.(*ve.SdkClient))
			return service.Dispatcher.Data(service, data, DataSourceVolcengineCrUsers())
		},
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cr instance name want to query.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of instance query.",
			},
			"users": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The collection of users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username of cr instance.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of user.",
						},
					},
				},
			},
		},
	}
}

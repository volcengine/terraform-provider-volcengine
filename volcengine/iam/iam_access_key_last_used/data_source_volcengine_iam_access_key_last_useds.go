package iam_access_key_last_used

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamAccessKeyLastUseds() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamAccessKeyLastUsedsRead,
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The access key id.",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user name.",
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
			"access_key_last_useds": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of access key last used.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the last used.",
						},
						"service": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service of the last used.",
						},
						"request_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The request time of the last used.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamAccessKeyLastUsedsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamAccessKeyLastUsedService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamAccessKeyLastUseds())
}

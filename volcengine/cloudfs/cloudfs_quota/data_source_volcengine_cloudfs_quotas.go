package cloudfs_quota

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudfsQuotas() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudfsQuotasRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cloud fs quota query.",
			},
			"quotas": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of account.",
						},
						"fs_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of cloud fs.",
						},
						"fs_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of cloud fs.",
						},
						"quota_enough": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether is enough of cloud fs.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudfsQuotasRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCloudfsQuotas())
}

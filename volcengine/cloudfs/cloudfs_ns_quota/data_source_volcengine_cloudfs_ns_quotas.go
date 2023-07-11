package cloudfs_ns_quota

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudfsNsQuotas() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudfsNsQuotasRead,
		Schema: map[string]*schema.Schema{
			"fs_names": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of fs name.",
			},
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
						"ns_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of cloud fs namespace.",
						},
						"ns_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of cloud fs namespace.",
						},
						"ns_count_per_fs": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "This file stores the number of namespaces under the instance.",
						},
						"ns_quota_per_fs": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "This file stores the total namespace quota under the instance.",
						},
						"fs_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of fs.",
						},
						"quota_enough": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether is enough of cloud fs namespace.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudfsNsQuotasRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCloudfsNsQuotas())
}

package cloudfs_access

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudfsAccesses() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudfsAccessesRead,
		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of file system.",
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
			"accesses": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fs_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of cloud fs.",
						},
						"access_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of access.",
						},
						"access_account_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The account id of access.",
						},
						"access_service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service name of access.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of vpc.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of subnet.",
						},
						"security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of security group.",
						},
						"is_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether is default access.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of access.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudfsAccessesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCloudfsAccesses())
}

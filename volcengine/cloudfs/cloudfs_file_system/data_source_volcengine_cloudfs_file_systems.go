package cloudfs_file_system

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudfsFileSystems() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudfsFileSystemsRead,
		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of file system.",
			},
			"meta_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of file system.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of cloudfs.",
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
			"file_systems": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of file system.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of file system.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of file system.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of region.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of zone.",
						},
						"cache_plan": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The plan of cache.",
						},
						"cache_capacity_tib": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The capacity of cache.",
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
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The mode of file system.",
						},
						"tos_bucket": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The tos bucket.",
						},
						"tos_prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The tos prefix.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time.",
						},
						"mount_point": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The point mount.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudfsFileSystemsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCloudfsFileSystems())
}

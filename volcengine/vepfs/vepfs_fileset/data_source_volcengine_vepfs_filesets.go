package vepfs_fileset

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVepfsFilesets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVepfsFilesetsRead,
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of Vepfs File System.",
			},
			"fileset_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of Vepfs Fileset.",
			},
			"fileset_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of Vepfs Fileset. This field support fuzzy query.",
			},
			"fileset_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The path of Vepfs Fileset. This field support fuzzy query.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The query status list of Vepfs Fileset.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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

			"filesets": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vepfs fileset.",
						},
						"fileset_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vepfs fileset.",
						},
						"fileset_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the vepfs fileset.",
						},
						"fileset_path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The path of the vepfs fileset.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the vepfs fileset.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the vepfs fileset.",
						},
						"iops_qos": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The IOPS Qos of the vepfs fileset.",
						},
						"file_used": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used file number of the vepfs fileset.",
						},
						"file_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Quota for the number of files or directories. A return of 0 indicates that there is no quota limit set for the number of directories after the file.",
						},
						"bandwidth_qos": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth Qos of the vepfs fileset.",
						},
						"capacity_used": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used capacity of the vepfs fileset. Unit: GiB.",
						},
						"capacity_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The capacity limit of the vepfs fileset. Unit: GiB.",
						},
						"max_inode_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max number of inode in the vepfs fileset.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVepfsFilesetsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVepfsFilesetService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVepfsFilesets())
}

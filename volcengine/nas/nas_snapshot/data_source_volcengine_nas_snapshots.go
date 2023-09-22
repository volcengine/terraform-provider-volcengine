package nas_snapshot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNasSnapshots() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNasSnapshotsRead,
		Schema: map[string]*schema.Schema{
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
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Snapshot IDs.",
			},
			"file_system_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of file system.",
			},
			"snapshot_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of snapshot.",
			},
			"snapshot_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of snapshot.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of snapshot.",
			},
			"snapshots": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of snapshot.",
						},
						"snapshot_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of snapshot.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of zone.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of snapshot.",
						},
						"source_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source version info.",
						},
						"source_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of source.",
						},
						"snapshot_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of snapshot.",
						},
						"snapshot_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of snapshot.",
						},
						"retention_days": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The retention days of snapshot.",
						},
						"progress": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The progress of snapshot.",
						},
						"is_encrypt": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether is encrypt.",
						},
						"file_system_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of file system.",
						},
						"file_system_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of file system.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of snapshot.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of snapshot.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNasSnapshotsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNasSnapshots())
}

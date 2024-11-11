package ebs_snapshot_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEbsSnapshotGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEbsSnapshotGroupsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of snapshot group IDs.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance id of snapshot group.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of snapshot group.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of snapshot group.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of snapshot group status. Valid values: `creating`, `available`, `failed`.",
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

			"snapshot_groups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the snapshot group.",
						},
						"snapshot_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the snapshot group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the snapshot group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the snapshot group.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the snapshot group.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance id of the snapshot group.",
						},
						"image_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The image id of the snapshot group.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the snapshot group.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the snapshot group.",
						},
						"tags": ve.TagsSchemaComputed(),
						"snapshots": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The snapshots of the snapshot group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"snapshot_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the snapshot.",
									},
									"snapshot_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the snapshot.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the snapshot.",
									},
									"snapshot_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the snapshot.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the snapshot.",
									},
									"volume_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The volume id of the snapshot.",
									},
									"volume_kind": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The volume kind of the snapshot.",
									},
									"volume_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The volume name of the snapshot.",
									},
									"volume_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The volume size of the snapshot.",
									},
									"volume_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The volume status of the snapshot.",
									},
									"volume_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The volume type of the snapshot.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The zone id of the snapshot.",
									},
									"image_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The image id of the snapshot.",
									},
									"retention_days": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The id of the snapshot.",
									},
									"progress": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The progress of the snapshot.",
									},
									"creation_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The creation time of the snapshot.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the snapshot.",
									},
									"tags": ve.TagsSchemaComputed(),
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEbsSnapshotGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEbsSnapshotGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineEbsSnapshotGroups())
}

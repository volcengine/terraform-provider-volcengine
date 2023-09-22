package nas_file_system

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNasFileSystems() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNasFileSystemsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of nas file system ids.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The status of nas file system.",
			},
			"file_system_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of nas file system. This field supports fuzzy queries.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone id of nas file system.",
			},
			"protocol_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The protocol type of nas file system.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The storage type of nas file system.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The charge type of nas file system.",
			},
			"permission_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The permission group id of nas file system.",
			},
			"mount_point_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The mount point id of nas file system.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of nas file system.",
			},
			"tags": ve.TagsSchema(),

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
			"file_systems": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the nas file system.",
						},
						"file_system_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the nas file system.",
						},
						"file_system_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the nas file system.",
						},
						"file_system_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the nas file system.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the nas file system.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the nas file system.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of the nas file system.",
						},
						"protocol_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol type of the nas file system.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The storage type of the nas file system.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the nas file system.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the nas file system.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id of the nas file system.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id of the nas file system.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone name of the nas file system.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the nas file system.",
						},
						"snapshot_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The snapshot count of the nas file system.",
						},
						"capacity": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The capacity of the nas file system.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The total capacity of the nas file system. Unit: GiB.",
									},
									"used": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The used capacity of the nas file system. Unit: MiB.",
									},
								},
							},
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the nas file system.",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Tags of the nas file system.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Key of Tags.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Value of Tags.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Type of Tags.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNasFileSystemsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNasFileSystemService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNasFileSystems())
}

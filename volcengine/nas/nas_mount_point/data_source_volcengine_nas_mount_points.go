package nas_mount_point

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNasMountPoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNasMountPointsRead,
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the file system.",
			},
			"mount_point_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the mount point.",
			},
			"mount_point_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the mount point.",
			},
			"vpcs_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the vpc.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of nas mount points query.",
			},
			"mount_points": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of mount points.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the mount point.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the mount point.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dns address.",
						},
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The address of the mount point.",
						},
						"mount_point_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the mount point.",
						},
						"mount_point_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the mount point.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the mount point.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the subnet.",
						},
						"subnet_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the subnet.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc.",
						},
						"vpc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the vpc.",
						},
						"permission_group": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The struct of the permission group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"permission_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the permission group.",
									},
									"permission_group_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the permission group.",
									},
									"permission_rule_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of the permission rule.",
									},
									"create_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The creation time of the permission group.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the permission group.",
									},
									"file_system_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of the file system.",
									},
									"file_system_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The file system type of the permission group.",
									},
									"mount_points": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of the mount point.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mount_point_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The id of the mount point.",
												},
												"mount_point_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the mount point.",
												},
												"file_system_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The id of the file system.",
												},
											},
										},
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

func dataSourceVolcengineNasMountPointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineNasMountPointService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNasMountPoints())
}

package nas_permission_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNasPermissionGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNasPermissionGroupsRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter permission groups for specified characteristics.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"PermissionGroupName",
								"PermissionGroupId",
							}, false),
							Description: "Filters permission groups for specified characteristics based on attributes. The parameters that support filtering are as follows: `PermissionGroupName`, `PermissionGroupId`.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the filter item.",
						},
					},
				},
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of nas permission groups query.",
			},
			"permission_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of permissions groups.",
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
						"permission_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of permissions rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"permission_rule_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the permission rule.",
									},
									"cidr_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client IP addresses that are allowed access.",
									},
									"rw_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Permission group read and write rules.",
									},
									"user_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Permission group user permissions.",
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

func dataSourceVolcengineNasPermissionGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineNasPermissionGroupService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNasPermissionGroups())
}

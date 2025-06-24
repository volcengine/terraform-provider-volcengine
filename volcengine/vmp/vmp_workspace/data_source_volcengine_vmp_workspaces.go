package vmp_workspace

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpWorkspaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpWorkspacesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Workspace IDs.",
			},
			"instance_type_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Instance Type IDs.",
			},
			"statuses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Workspace status.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of workspace.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of vmp workspace.",
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The tags of vmp workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"values": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The Value of Tags.",
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
				Description: "The total count of query.",
			},
			"workspaces": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of workspace.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of workspace.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of workspace.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of workspace.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of workspace.",
						},
						"instance_type_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of instance type.",
						},
						"overdue_reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue reclaim time.",
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username of workspace.",
						},
						"prometheus_write_intranet_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The prometheus write intranet endpoint.",
						},
						"prometheus_query_intranet_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The prometheus query intranet endpoint.",
						},
						"delete_protection_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether enable delete protection.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of vmp workspace.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVmpWorkspacesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpWorkspaces())
}

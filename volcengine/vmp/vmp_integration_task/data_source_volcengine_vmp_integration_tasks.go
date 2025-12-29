package vmp_integration_task

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpIntegrationTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpIntegrationTasksRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of integration task IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the integration task.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The status of the integration task. Valid values: `Creating`, `Updating`, `Active`, `Error`, `Deleting`.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The workspace ID.",
			},
			"vke_cluster_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The ID of the VKE cluster.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The deployment environment. Valid values: `Vke` or `Managed`.",
			},
			"integration_tasks": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of integration tasks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the integration task.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the integration task.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the integration task.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the integration task.",
						},
						"environment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deployment environment.",
						},
						"workspace_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The workspace ID.",
						},
						"vke_cluster_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ID of the VKE cluster.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vke_cluster_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information of the VKE cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the VKE cluster.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the VKE cluster.",
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

func dataSourceVolcengineVmpIntegrationTasksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpIntegrationTasks())
}

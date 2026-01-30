package health_check_log_project

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineHealthCheckLogProjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineHealthCheckLogProjectsRead,
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
			"health_check_log_projects": {
				Description: "The collection of health check log projects.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the health check log project.",
						},
						"log_project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the health check log project.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineHealthCheckLogProjectsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewHealthCheckLogProjectService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineHealthCheckLogProjects())
}

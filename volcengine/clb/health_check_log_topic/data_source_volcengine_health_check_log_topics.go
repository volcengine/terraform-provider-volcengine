package health_check_log_topic

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineHealthCheckLogTopics() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineHealthCheckLogTopicsRead,
		Schema: map[string]*schema.Schema{
			"log_topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the log topic.",
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
			"health_check_log_topics": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ID of the CLB instance.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineHealthCheckLogTopicsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewHealthCheckLogTopicService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineHealthCheckLogTopics())
}

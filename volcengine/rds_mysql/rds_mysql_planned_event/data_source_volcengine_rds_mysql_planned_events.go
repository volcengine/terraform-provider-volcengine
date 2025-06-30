package rds_mysql_planned_event

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlPlannedEvents() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlPlannedEventsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the instance.",
			},
			"begin_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of the planned event.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of the planned event.",
			},
			"event_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the planned event.",
			},
			"event_type": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "The type of the planned event.",
			},
			"status": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "The status of the planned event.",
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
			"planned_events": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"business_impact": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business impact of the planned event.",
						},
						"db_engine": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The database engine of the planned event.",
						},
						"event_action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The action of the planned event.",
						},
						"event_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the planned event.",
						},
						"event_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the planned event.",
						},
						"event_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the planned event.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the instance.",
						},
						"max_delay_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The latest postponable time. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"origin_begin_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The initially set start time. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"planned_begin_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the planned execution. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"planned_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the planned execution. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"planned_event_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description information of the operation and maintenance event.",
						},
						"planned_event_reason": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The reasons for the occurrence of the event, " +
								"which are provided to help you understand the reasons for the occurrence of unexpected events.",
						},
						"planned_switch_begin_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the planned switch. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"planned_switch_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the planned switch. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Event status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsMysqlPlannedEventsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlPlannedEventService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlPlannedEvents())
}

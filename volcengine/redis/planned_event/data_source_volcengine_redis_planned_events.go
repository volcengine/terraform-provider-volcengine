package planned_event

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePlannedEvents() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePlannedEventsRead,
		Schema: map[string]*schema.Schema{
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
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of instance.",
			},
			"min_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The earliest execution time of the planned event that needs to be queried." +
					" The format is yyyy-MM-ddTHH:mm:ssZ (UTC).",
			},
			"max_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The latest execution time of the planned events that need to be queried." +
					" The format is yyyy-MM-ddTHH:mm:ssZ (UTC).",
			},
			"planned_events": {
				Description: "The List of planned event information.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Event operation name.",
						},
						"can_cancel": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the current event is allowed to be cancelled for execution.",
						},
						"can_modify_time": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the execution time of the current event can be changed.",
						},
						"event_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Event.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of instance.",
						},
						"max_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The latest execution time at which changes are allowed for the current event.",
						},
						"plan_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The latest execution time of the event plan. The format is yyyy-MM-ddTHH:mm:ssZ (UTC).",
						},
						"plan_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The earliest planned execution time of the event. The format is yyyy-MM-ddTHH:mm:ssZ (UTC).",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of event.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of event.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcenginePlannedEventsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewPlannedEventService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcenginePlannedEvents())
}

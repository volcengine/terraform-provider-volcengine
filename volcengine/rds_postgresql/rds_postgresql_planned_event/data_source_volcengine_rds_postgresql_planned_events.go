package rds_postgresql_planned_event

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlPlannedEvents() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlPlannedEventsRead,
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
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of PostgreSQL instance.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"event_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Event ID.",
			},
			"event_type": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"VersionUpgrade", "HostOffline"}, false),
				},
				Set:         schema.HashString,
				Description: "Event type. Values: VersionUpgrade, HostOffline.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"Canceled", "WaitStart", "WaitSwitch", "Running", "Running_BeforeSwitch", "Running_Switching", "Running_AfterSwitch", "Success", "Failed", "Timeout", "Rollbacking", "RollbackFailed"}, false),
				},
				Set:         schema.HashString,
				Description: "Operation event status. Values: Canceled, WaitStart, WaitSwitch, Running, Running_BeforeSwitch, Running_Switching, Running_AfterSwitch, Success, Failed, Timeout, Rollbacking, RollbackFailed.",
			},
			"planned_begin_time_search_range_start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Time window start for planned execution time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
			},
			"planned_begin_time_search_range_end": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Time window end for planned execution time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
			},
			"planned_switch_time_search_range_start": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Time window start for planned switch time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
			},
			"planned_switch_time_search_range_end": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Time window end for planned switch time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
			},
			"planned_events": {
				Description: "Planned events list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"business_impact": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The impact of operation and maintenance events on the business.",
						},
						"event_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Event ID.",
						},
						"event_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Event type.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance ID.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance name.",
						},
						"max_delay_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Maximum delay time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
						},
						"planned_begin_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Planned execution time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
						},
						"planned_event_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reason for the planned event.",
						},
						"planned_switch_begin_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Planned switch start time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
						},
						"planned_switch_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Planned switch end time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Operation event status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlPlannedEventsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlPlannedEventService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlPlannedEvents())
}

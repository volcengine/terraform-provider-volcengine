package schedule_sql_task

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ScheduleSqlTask can be imported using the id, e.g.
```
$ terraform import volcengine_schedule_sql_task.default resource_id
```

*/

func ResourceVolcengineScheduleSqlTask() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScheduleSqlTaskCreate,
		Read:   resourceVolcengineScheduleSqlTaskRead,
		Update: resourceVolcengineScheduleSqlTaskUpdate,
		Delete: resourceVolcengineScheduleSqlTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"task_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The Name of timed SQL analysis task.",
			},
			"topic_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The log topic ID where the original log to be analyzed for scheduled SQL is located.",
			},
			"dest_region": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The region to which the target log topic belongs. The default is the current region.",
			},
			"dest_topic_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The target log topic ID used for storing the result data of timed SQL analysis.",
			},
			"process_start_time": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeInt,
				Description: "The start time of the scheduled SQL analysis task, " +
					"that is, the time when the first instance is created. " +
					"The format is a timestamp at the second level.",
			},
			"process_end_time": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Schedule the end time of the timed SQL analysis task in the format of a second-level timestamp.",
			},
			"process_time_window": {
				Required: true,
				Type:     schema.TypeString,
				Description: "SQL time window, which refers to the time range for log retrieval and analysis " +
					"when a timed SQL analysis task is running, is in a left-closed and right-open format.",
			},
			"query": {
				Required: true,
				Type:     schema.TypeString,
				Description: "The retrieval and analysis statements for the regular execution of timed SQL analysis " +
					"tasks should conform to the retrieval and analysis syntax of the log service.",
			},
			"request_cycle": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The scheduling cycle of timed SQL analysis tasks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The scheduling cycle or the time point of regular execution (the number of minutes away from 00:00), with a value range of 1 to 1440, and the unit is minutes.",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of Scheduling cycle.",
						},
						"cron_tab": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "Cron expression. The log service specifies the timed execution of alarm tasks through the Cron expression. " +
								"The minimum granularity of Cron expressions is minutes, 24 hours. " +
								"For example, 0 18 * * * indicates that an alarm task is executed exactly at 18:00 every day.",
						},
						"cron_time_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When setting the Type to Cron, the time zone also needs to be set.",
						},
					},
				},
			},
			"status": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Whether to start the scheduled SQL analysis task immediately after completing the task configuration.",
			},
			"process_sql_delay": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "The delay time of each scheduling. The value range is from 0 to 120, and the unit is seconds.",
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A simple description of the timed SQL analysis task.",
			},
		},
	}
	return resource
}

func resourceVolcengineScheduleSqlTaskCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScheduleSqlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineScheduleSqlTask())
	if err != nil {
		return fmt.Errorf("error on creating schedule_sql_task %q, %s", d.Id(), err)
	}
	return resourceVolcengineScheduleSqlTaskRead(d, meta)
}

func resourceVolcengineScheduleSqlTaskRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScheduleSqlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineScheduleSqlTask())
	if err != nil {
		return fmt.Errorf("error on reading schedule_sql_task %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScheduleSqlTaskUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScheduleSqlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineScheduleSqlTask())
	if err != nil {
		return fmt.Errorf("error on updating schedule_sql_task %q, %s", d.Id(), err)
	}
	return resourceVolcengineScheduleSqlTaskRead(d, meta)
}

func resourceVolcengineScheduleSqlTaskDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewScheduleSqlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineScheduleSqlTask())
	if err != nil {
		return fmt.Errorf("error on deleting schedule_sql_task %q, %s", d.Id(), err)
	}
	return err
}

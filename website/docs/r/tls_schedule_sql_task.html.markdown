---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_schedule_sql_task"
sidebar_current: "docs-volcengine-resource-tls_schedule_sql_task"
description: |-
  Provides a resource to manage tls schedule sql task
---
# volcengine_tls_schedule_sql_task
Provides a resource to manage tls schedule sql task
## Example Usage
```hcl
resource "volcengine_tls_schedule_sql_task" "foo" {
  task_name           = "tf-test"
  topic_id            = "8ba48bd7-2493-4300-b1d0-cb760bxxxxxx"
  dest_region         = "cn-beijing"
  dest_topic_id       = "b966e41a-d6a6-4999-bd75-39962xxxxxx"
  process_start_time  = 1751212980
  process_end_time    = 1751295600
  process_time_window = "@m-15m,@m"
  query               = "* | SELECT * limit 10000"
  request_cycle {
    cron_tab       = "0 10 * * *"
    cron_time_zone = "GMT+08:00"
    time           = 1
    type           = "CronTab"
  }
  status            = 1
  process_sql_delay = 60
  description       = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `dest_topic_id` - (Required) The target log topic ID used for storing the result data of timed SQL analysis.
* `process_sql_delay` - (Required) The delay time of each scheduling. The value range is from 0 to 120, and the unit is seconds.
* `process_start_time` - (Required, ForceNew) The start time of the scheduled SQL analysis task, that is, the time when the first instance is created. The format is a timestamp at the second level.
* `process_time_window` - (Required) SQL time window, which refers to the time range for log retrieval and analysis when a timed SQL analysis task is running, is in a left-closed and right-open format.
* `query` - (Required) The retrieval and analysis statements for the regular execution of timed SQL analysis tasks should conform to the retrieval and analysis syntax of the log service.
* `request_cycle` - (Required) The scheduling cycle of timed SQL analysis tasks.
* `status` - (Required) Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
* `task_name` - (Required) The Name of timed SQL analysis task.
* `topic_id` - (Required, ForceNew) The log topic ID where the original log to be analyzed for scheduled SQL is located.
* `description` - (Optional) A simple description of the timed SQL analysis task.
* `dest_region` - (Optional) The region to which the target log topic belongs. The default is the current region.
* `process_end_time` - (Optional, ForceNew) Schedule the end time of the timed SQL analysis task in the format of a second-level timestamp.

The `request_cycle` object supports the following:

* `time` - (Required) The scheduling cycle or the time point of regular execution (the number of minutes away from 00:00), with a value range of 1 to 1440, and the unit is minutes.
* `type` - (Required) The type of Scheduling cycle.
* `cron_tab` - (Optional) Cron expression. The log service specifies the timed execution of alarm tasks through the Cron expression. The minimum granularity of Cron expressions is minutes, 24 hours. For example, 0 18 * * * indicates that an alarm task is executed exactly at 18:00 every day.
* `cron_time_zone` - (Optional) When setting the Type to Cron, the time zone also needs to be set.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ScheduleSqlTask can be imported using the id, e.g.
```
$ terraform import volcengine_schedule_sql_task.default resource_id
```


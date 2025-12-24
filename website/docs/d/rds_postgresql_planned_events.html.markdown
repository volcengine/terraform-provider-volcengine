---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_planned_events"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_planned_events"
description: |-
  Use this data source to query detailed information of rds postgresql planned events
---
# volcengine_rds_postgresql_planned_events
Use this data source to query detailed information of rds postgresql planned events
## Example Usage
```hcl
data "volcengine_rds_postgresql_planned_events" "example" {
  instance_id                            = "postgres-72715e0d9f58"
  instance_name                          = "test-01"
  event_type                             = ["VersionUpgrade"]
  status                                 = ["WaitStart", "Running"]
  planned_switch_time_search_range_start = "2025-12-01T02:06:53.000Z"
  planned_switch_time_search_range_end   = "2025-12-15T17:40:53.000Z"
}
```
## Argument Reference
The following arguments are supported:
* `event_id` - (Optional) Event ID.
* `event_type` - (Optional) Event type. Values: VersionUpgrade, HostOffline.
* `instance_id` - (Optional) The id of the PostgreSQL instance.
* `instance_name` - (Optional) The name of PostgreSQL instance.
* `output_file` - (Optional) File name where to save data source results.
* `planned_begin_time_search_range_end` - (Optional) Time window end for planned execution time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
* `planned_begin_time_search_range_start` - (Optional) Time window start for planned execution time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
* `planned_switch_time_search_range_end` - (Optional) Time window end for planned switch time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
* `planned_switch_time_search_range_start` - (Optional) Time window start for planned switch time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
* `status` - (Optional) Operation event status. Values: Canceled, WaitStart, WaitSwitch, Running, Running_BeforeSwitch, Running_Switching, Running_AfterSwitch, Success, Failed, Timeout, Rollbacking, RollbackFailed.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `planned_events` - Planned events list.
    * `business_impact` - The impact of operation and maintenance events on the business.
    * `event_id` - Event ID.
    * `event_type` - Event type.
    * `instance_id` - Instance ID.
    * `instance_name` - Instance name.
    * `max_delay_time` - Maximum delay time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
    * `planned_begin_time` - Planned execution time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
    * `planned_event_reason` - Reason for the planned event.
    * `planned_switch_begin_time` - Planned switch start time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
    * `planned_switch_end_time` - Planned switch end time. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC).
    * `region` - Region.
    * `status` - Operation event status.
* `total_count` - The total count of query.



---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_planned_events"
sidebar_current: "docs-volcengine-datasource-rds_mysql_planned_events"
description: |-
  Use this data source to query detailed information of rds mysql planned events
---
# volcengine_rds_mysql_planned_events
Use this data source to query detailed information of rds mysql planned events
## Example Usage
```hcl
data "volcengine_rds_mysql_planned_events" "foo" {
  instance_id = "mysql-b51d37110dd1"
}
```
## Argument Reference
The following arguments are supported:
* `begin_time` - (Optional) The start time of the planned event.
* `end_time` - (Optional) The end time of the planned event.
* `event_id` - (Optional) The id of the planned event.
* `event_type` - (Optional) The type of the planned event.
* `instance_id` - (Optional) The id of the instance.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) The status of the planned event.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `planned_events` - The collection of query.
    * `business_impact` - The business impact of the planned event.
    * `db_engine` - The database engine of the planned event.
    * `event_action` - The action of the planned event.
    * `event_id` - The id of the planned event.
    * `event_name` - The name of the planned event.
    * `event_type` - The type of the planned event.
    * `instance_id` - The id of the instance.
    * `instance_name` - The name of the instance.
    * `max_delay_time` - The latest postponable time. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `origin_begin_time` - The initially set start time. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `planned_begin_time` - The start time of the planned execution. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `planned_end_time` - The end time of the planned execution. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `planned_event_description` - Description information of the operation and maintenance event.
    * `planned_event_reason` - The reasons for the occurrence of the event, which are provided to help you understand the reasons for the occurrence of unexpected events.
    * `planned_switch_begin_time` - The start time of the planned switch. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `planned_switch_end_time` - The end time of the planned switch. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `region` - The region.
    * `status` - Event status.
* `total_count` - The total count of query.



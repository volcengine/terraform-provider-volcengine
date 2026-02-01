---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_search_traces"
sidebar_current: "docs-volcengine-datasource-tls_search_traces"
description: |-
  Use this data source to query detailed information of tls search traces
---
# volcengine_tls_search_traces
Use this data source to query detailed information of tls search traces
## Example Usage
```hcl
data "volcengine_tls_seracg_traces" "example" {
  trace_instance_id = "b28b19bd-a539-453a-8919-fda3ef6a22fe"
  # trace_id          = "c415ff6a-7141-4fe9-9e6c-9ddce4e4c189"
}
```
## Argument Reference
The following arguments are supported:
* `trace_instance_id` - (Required) Trace instance ID.
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) Query conditions.

The `attributes` object supports the following:

* `key` - (Required) Attribute key.
* `value` - (Required) Attribute value.

The `query` object supports the following:

* `asc` - (Optional) Whether to sort results in ascending order. true means ascending, false means descending.
* `attributes` - (Optional) Attributes.
* `duration_max` - (Optional) Maximum trace duration in microseconds.
* `duration_min` - (Optional) Minimum trace duration in microseconds.
* `kind` - (Optional) Type of the trace.
* `limit` - (Optional) Maximum number of records to return, used for pagination.
* `offset` - (Optional) Offset for paginated query.
* `operation_name` - (Optional) Operation name, used to filter traces with specific operation.
* `order` - (Optional) Sorting field. Supported fields: Kind, Name, ServiceName, Start, End, Duration, and indexed fields in Attributes.
* `service_name` - (Optional) Service name, used to filter traces from specific service.
* `start_time_max` - (Optional) Maximum start time for searching traces, in microsecond timestamp format.
* `start_time_min` - (Optional) Minimum start time for searching traces, in microsecond timestamp format.
* `status_code` - (Optional) Trace status code, used to filter traces with specific status.
* `trace_id` - (Optional) Trace ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of tls trace query.
* `traces` - The collection of tls trace query.
    * `attributes` - Trace attributes.
    * `duration` - Trace duration in microseconds.
    * `end_time` - Trace end time in microseconds.
    * `operation_name` - Operation name.
    * `service_name` - Service name.
    * `start_time` - Trace start time in microseconds.
    * `status_code` - Trace status code.
    * `trace_id` - Trace ID.



---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_log_searches"
sidebar_current: "docs-volcengine-datasource-tls_log_searches"
description: |-
  Use this data source to query detailed information of tls log searches
---
# volcengine_tls_log_searches
Use this data source to query detailed information of tls log searches
## Example Usage
```hcl

```
## Argument Reference
The following arguments are supported:
* `end_time` - (Required) The end time of the log.
* `query` - (Required) The query of the log.
* `start_time` - (Required) The start time of the log.
* `topic_id` - (Required) The ID of the topic.
* `accurate_query` - (Optional) Whether to use accurate query.
* `context` - (Optional) The context of the log.
* `highlight` - (Optional) Whether to highlight the log.
* `limit` - (Optional) The limit of the logs.
* `output_file` - (Optional) File name where to save data source results.
* `sort` - (Optional) The sort of the log.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `logs` - The collection of query result.
    * `analysis_result` - The analysis result of the query.
    * `analysis` - Whether the result is analysis.
    * `context` - The context of the log.
    * `elapsed_millisecond` - The elapsed time of the query.
    * `highlight` - The highlight of the query.
        * `key` - The key of the highlight.
        * `value` - The value of the highlight.
    * `hit_count` - The count of the logs.
    * `limit` - The limit of the logs.
    * `list_over` - Whether the list is over.
    * `logs` - The list of the logs.
        * `content` - The content of the log.
        * `filename` - The filename of the log.
        * `log_id` - The ID of the log.
        * `source` - The source of the log.
        * `timestamp` - The timestamp of the log.
    * `result_status` - The status of the query.
* `total_count` - The total count of the logs.



---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_log_contexts"
sidebar_current: "docs-volcengine-datasource-tls_log_contexts"
description: |-
  Use this data source to query detailed information of tls log contexts
---
# volcengine_tls_log_contexts
Use this data source to query detailed information of tls log contexts
## Example Usage
```hcl
# Search Logs (Trigger SearchLogs)
data "volcengine_tls_log_searches" "default" {
  topic_id   = "3c57a110-399a-43b3-bc3c-5d60e065239a"
  query      = "*"
  start_time = 1768448896000
  end_time   = 1768450896000
  limit      = 10
}

# 1. Describe Log Context (Trigger DescribeLogContext)
data "volcengine_tls_log_contexts" "default" {
  topic_id       = data.volcengine_tls_log_searches.default.topic_id
  context_flow   = data.volcengine_tls_log_searches.default.logs[0].logs[0].content["__context_flow__"]
  package_offset = tonumber(data.volcengine_tls_log_searches.default.logs[0].logs[0].content["__package_offset__"])
  source         = data.volcengine_tls_log_searches.default.logs[0].logs[0].source
  prev_logs      = 10
  next_logs      = 10
}
```
## Argument Reference
The following arguments are supported:
* `context_flow` - (Required) The context flow of the log.
* `package_offset` - (Required) The package offset of the log.
* `source` - (Required) The source of the log.
* `topic_id` - (Required) The ID of the topic.
* `describe_log_context` - (Optional) Whether to describe log context.
* `next_logs` - (Optional) The number of next logs.
* `output_file` - (Optional) File name where to save data source results.
* `prev_logs` - (Optional) The number of previous logs.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `log_contexts` - The list of log contexts.
    * `log_context_infos` - The infos of context log.
    * `next_over` - Whether the next logs are over.
    * `prev_over` - Whether the previous logs are over.



---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_trace_instances"
sidebar_current: "docs-volcengine-datasource-tls_trace_instances"
description: |-
  Use this data source to query detailed information of tls trace instances
---
# volcengine_tls_trace_instances
Use this data source to query detailed information of tls trace instances
## Example Usage
```hcl
# Example 1: Query by trace instance name
data "volcengine_tls_trace_instances" "by_name" {
  project_id          = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  trace_instance_name = "测试trace"
}

# Example 2: Query by status
data "volcengine_tls_trace_instances" "by_status" {
  project_id = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  status     = "CREATED"
}
```
## Argument Reference
The following arguments are supported:
* `cs_account_channel` - (Optional) CS account channel identifier.
* `iam_project_name` - (Optional) The IAM project name.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) The ID of the project.
* `project_name` - (Optional) The name of the project.
* `status` - (Optional) The status of the trace instance.
* `trace_instance_id` - (Optional) The ID of the trace instance.
* `trace_instance_name` - (Optional) The name of the trace instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of trace instances.
* `trace_instances` - The list of trace instances.
    * `backend_config` - The backend config of the trace instance.
        * `archive_ttl` - Archive storage duration in days.
        * `auto_split` - Whether to enable auto split.
        * `cold_ttl` - Infrequent storage duration in days.
        * `enable_hot_ttl` - Whether to enable tiered storage.
        * `hot_ttl` - Standard storage duration in days.
        * `max_split_partitions` - Max split partitions.
        * `ttl` - Total log retention time in days.
    * `create_time` - The create time of the trace instance.
    * `cs_account_channel` - CS account channel identifier.
    * `dependency_topic_id` - The ID of the dependency topic.
    * `dependency_topic_topic_name` - The name of the dependency topic.
    * `description` - The description of the trace instance.
    * `modify_time` - The update time of the trace instance.
    * `project_id` - The ID of the project.
    * `project_name` - The name of the project.
    * `trace_instance_id` - The ID of the trace instance.
    * `trace_instance_name` - The name of the trace instance.
    * `trace_instance_status` - The status of the trace instance.
    * `trace_topic_id` - The ID of the trace topic.
    * `trace_topic_name` - The name of the trace topic.



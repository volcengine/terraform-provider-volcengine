---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_flow_logs"
sidebar_current: "docs-volcengine-datasource-flow_logs"
description: |-
  Use this data source to query detailed information of flow logs
---
# volcengine_flow_logs
Use this data source to query detailed information of flow logs
## Example Usage
```hcl
data "volcengine_flow_logs" "foo" {
  flow_log_ids = ["fl-13g4fqngluhog3n6nu57o****"]
}
```
## Argument Reference
The following arguments are supported:
* `aggregation_interval` - (Optional) The aggregation interval of flow log. Unit: minute. Valid values: `1`, `5`, `10`.
* `description` - (Optional) The description of flow log.
* `flow_log_ids` - (Optional) A list of flow log IDs.
* `flow_log_name` - (Optional) The name of flow log.
* `log_project_id` - (Optional) The ID of log project.
* `log_topic_id` - (Optional) The ID of log topic.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of flow log.
* `resource_id` - (Optional) The ID of resource.
* `resource_type` - (Optional) The type of resource. Valid values: `vpc`, `subnet`, `eni`.
* `status` - (Optional) The status of flow log. Valid values: `Active`, `Pending`, `Inactive`, `Creating`, `Deleting`.
* `tags` - (Optional) Tags.
* `traffic_type` - (Optional) The type of traffic. Valid values: `All`, `Allow`, `Drop`.
* `vpc_id` - (Optional) The ID of VPC.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `flow_logs` - The collection of query.
    * `aggregation_interval` - The aggregation interval of flow log. Unit: minute. Valid values: `1`, `5`, `10`.
    * `business_status` - The business status of flow log.
    * `created_at` - The created time of flow log.
    * `description` - The description of flow log.
    * `flow_log_id` - The ID of flow log.
    * `flow_log_name` - The name of flow log.
    * `id` - The ID of flow log.
    * `lock_reason` - The reason why flow log is locked.
    * `log_project_id` - The ID of log project.
    * `log_topic_id` - The ID of log topic.
    * `project_name` - The project name of flow log.
    * `resource_id` - The ID of resource.
    * `resource_type` - The type of resource. Valid values: `vpc`, `subnet`, `eni`.
    * `status` - The status of flow log. Valid values: `Active`, `Pending`, `Inactive`, `Creating`, `Deleting`.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `traffic_type` - The type of traffic. Valid values: `All`, `Allow`, `Drop`.
    * `updated_at` - The updated time of flow log.
    * `vpc_id` - The ID of VPC.
* `total_count` - The total count of query.



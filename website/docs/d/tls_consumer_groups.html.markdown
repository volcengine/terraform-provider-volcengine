---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_consumer_groups"
sidebar_current: "docs-volcengine-datasource-tls_consumer_groups"
description: |-
  Use this data source to query detailed information of tls consumer groups
---
# volcengine_tls_consumer_groups
Use this data source to query detailed information of tls consumer groups
## Example Usage
```hcl
data "volcengine_tls_consumer_groups" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `consumer_group_name` - (Optional) The name of the consumer group.
* `iam_project_name` - (Optional) IAM log project name.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) The log project ID to which the consumption group belongs.
* `project_name` - (Optional) The name of the log item to which the consumption group belongs.
* `topic_id` - (Optional) The log topic ID to which the consumer belongs.
* `topic_name` - (Optional) The name of the log topic to which the consumption group belongs.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `consumer_groups` - List of log service consumption groups.
    * `consumer_group_name` - The name of the consumer group.
    * `heartbeat_ttl` - The time of heart rate expiration, measured in seconds, has a value range of 1 to 300.
    * `ordered_consume` - Whether to consume in sequence.
    * `project_id` - The log project ID to which the consumption group belongs.
    * `project_name` - The name of the log item to which the consumption group belongs.
    * `topic_id` - The list of log topic ids to be consumed by the consumer group.
* `total_count` - The total count of query.



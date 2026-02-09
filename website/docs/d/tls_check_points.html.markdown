---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_check_points"
sidebar_current: "docs-volcengine-datasource-tls_check_points"
description: |-
  Use this data source to query detailed information of tls check points
---
# volcengine_tls_check_points
Use this data source to query detailed information of tls check points
## Example Usage
```hcl
data "volcengine_tls_check_points" "default" {
  project_id          = "7a8ac13e-8e3e-4392-ae77-aea8efa49bbf"
  topic_id            = "33124cc3-15c4-4cdc-9a8a-cc64a9d593dd"
  shard_id            = "0"
  consumer_group_name = "tf-consumer-group"
}
```
## Argument Reference
The following arguments are supported:
* `project_id` - (Required) The ID of the project.
* `shard_id` - (Required) The ID of the shard.
* `topic_id` - (Required) The ID of the topic.
* `consumer_group_name` - (Optional) The name of the consumer group.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `check_points` - The list of checkpoints.
    * `checkpoint` - The checkpoint value.
    * `shard_id` - The ID of the shard.



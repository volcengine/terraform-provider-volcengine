---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_groups"
sidebar_current: "docs-volcengine-datasource-rocketmq_groups"
description: |-
  Use this data source to query detailed information of rocketmq groups
---
# volcengine_rocketmq_groups
Use this data source to query detailed information of rocketmq groups
## Example Usage
```hcl
data "volcengine_rocketmq_groups" "foo" {
  instance_id = "rocketmq-cnoeea6b32118fc2"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of rocketmq instance.
* `group_id` - (Optional) The id of rocketmq group. This field support fuzzy query.
* `group_type` - (Optional) The type of rocketmq group. Valid values: `TCP`.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rocketmq_groups` - The collection of query.
    * `consumed_clients` - The consumed topic information of the rocketmq group.
        * `client_address` - The address of the consumed client.
        * `client_id` - The id of the consumed client.
        * `diff` - The amount of message.
        * `language` - The language of the consumed client.
        * `version` - The version of the consumed client.
    * `consumed_topics` - The consumed topic information of the rocketmq group.
        * `queue_num` - The queue number of the rocketmq topic.
        * `sub_string` - The sub string of the rocketmq topic.
        * `topic_name` - The name of the rocketmq topic.
    * `create_time` - The create time of the rocketmq group.
    * `description` - The description of the rocketmq group.
    * `group_id` - The id of the rocketmq group.
    * `group_type` - The type of the rocketmq group.
    * `is_sub_same` - Whether the subscription relationship of consumer instance groups within the group is consistent.
    * `message_delay_time` - The message delay time of the rocketmq group. The unit is milliseconds.
    * `message_model` - The message model of the rocketmq group.
    * `status` - The status of the rocketmq group.
    * `total_consume_rate` - The total consume rate of the rocketmq group. The unit is per second.
    * `total_diff` - The total amount of unconsumed messages.
* `total_count` - The total count of query.



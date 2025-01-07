---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_topics"
sidebar_current: "docs-volcengine-datasource-rocketmq_topics"
description: |-
  Use this data source to query detailed information of rocketmq topics
---
# volcengine_rocketmq_topics
Use this data source to query detailed information of rocketmq topics
## Example Usage
```hcl
data "volcengine_rocketmq_topics" "foo" {
  instance_id = "rocketmq-cnoeea6b32118fc2"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of rocketmq instance.
* `message_type` - (Optional) The type of the rocketmq message. Setting this parameter means filtering the Topic list based on the specified message type. The value explanation is as follows:
0: Regular message
1: Transaction message
2: Partition order message
3: Global sequential message
4: Delay message.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `topic_name` - (Optional) The name of the rocketmq topic. This field support fuzzy query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rocketmq_topics` - The collection of query.
    * `access_policies` - The access policies of the rocketmq topic.
        * `access_key` - The access key of the rocketmq key.
        * `authority` - The authority of the rocketmq key for the current topic.
    * `create_time` - The create time of the rocketmq topic.
    * `description` - The description of the rocketmq topic.
    * `groups` - The groups information of the rocketmq topic.
        * `group_id` - The id of the rocketmq group.
        * `message_model` - The message model of the rocketmq group.
        * `sub_string` - The sub string of the rocketmq group.
    * `instance_id` - The id of rocketmq instance.
    * `message_type` - The type of the rocketmq message.
    * `queue_number` - The number of the rocketmq topic queue.
    * `queues` - The queues information of the rocketmq topic.
        * `end_offset` - The end offset of the rocketmq queue.
        * `last_update_timestamp` - The last update timestamp of the rocketmq queue.
        * `message_count` - The message count of the rocketmq queue.
        * `queue_id` - The id of the rocketmq queue.
        * `start_offset` - The start offset of the rocketmq queue.
    * `topic_name` - The name of the rocketmq topic.
* `total_count` - The total count of query.



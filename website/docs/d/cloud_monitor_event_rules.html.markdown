---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_event_rules"
sidebar_current: "docs-volcengine-datasource-cloud_monitor_event_rules"
description: |-
  Use this data source to query detailed information of cloud monitor event rules
---
# volcengine_cloud_monitor_event_rules
Use this data source to query detailed information of cloud monitor event rules
## Example Usage
```hcl
data "volcengine_cloud_monitor_event_rules" "foo" {
  rule_name = "tftest"
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `rule_name` - (Optional) Rule name, search rules by name using fuzzy search.
* `source` - (Optional) Event source.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of query.
    * `account_id` - The id of the account.
    * `contact_group_ids` - When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.
    * `contact_methods` - List of contact methods.
    * `created_at` - The create time.
    * `description` - The description of the rule.
    * `effect_end_at` - The end time of the rule.
    * `effect_start_at` - The start time of the rule.
    * `endpoint` - When the alarm notification method is alarm callback, it triggers the callback address.
    * `event_bus_name` - The name of the event bus.
    * `event_source` - The source of the event.
    * `event_type` - The event type.
    * `filter_pattern` - Filter mode, also known as event matching rules. Custom matching rules are not currently supported.
        * `source` - Event source corresponding to pattern matching.
        * `type` - The list of corresponding event types in pattern matching, currently set to match any.
    * `id` - The id of the rule.
    * `level` - The level of the rule.
    * `message_queue` - The triggered message queue when the alarm notification method is Kafka message queue.
        * `instance_id` - The kafka instance id.
        * `region` - The region.
        * `topic` - The topic name.
        * `type` - The message queue type, only support kafka now.
        * `vpc_id` - The vpc id.
    * `region` - The name of the region.
    * `rule_id` - The id of the rule.
    * `rule_name` - The name of the rule.
    * `status` - Enable the state of the rule.
    * `tls_target` - The alarm method for log service triggers the configuration of the log service.
        * `project_id` - The project id.
        * `project_name` - The project name.
        * `region_name_cn` - The Chinese region name.
        * `region_name_en` - The English region name.
        * `topic_id` - The topic id.
        * `topic_name` - The topic name.
    * `updated_at` - The updated time.
* `total_count` - The total count of query.



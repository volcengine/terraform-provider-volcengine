---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm_notify_groups"
sidebar_current: "docs-volcengine-datasource-tls_alarm_notify_groups"
description: |-
  Use this data source to query detailed information of tls alarm notify groups
---
# volcengine_tls_alarm_notify_groups
Use this data source to query detailed information of tls alarm notify groups
## Example Usage
```hcl
data "volcengine_tls_alarm_notify_groups" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `alarm_notify_group_id` - (Optional) The id of the alarm notify group.
* `alarm_notify_group_name` - (Optional) The name of the alarm notify group.
* `iam_project_name` - (Optional) The name of the iam project.
* `output_file` - (Optional) File name where to save data source results.
* `receiver_name` - (Optional) The name of the receiver.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `groups` - The list of the notify groups.
    * `alarm_notify_group_id` - The id of the notify group.
    * `alarm_notify_group_name` - Name of the notification group.
    * `create_time` - The create time the notification.
    * `iam_project_name` - The iam project name.
    * `modify_time` - The modification time the notification.
    * `notice_rules` - The list of the notice rules.
        * `has_end_node` - Whether there is an end node behind.
        * `has_next` - Whether to continue to the next level of condition judgment.
        * `receiver_infos` - List of IAM users to receive alerts.
            * `alarm_content_template_id` - The alarm content template id.
            * `alarm_webhook_at_users` - The alarm webhook at users.
            * `alarm_webhook_integration_id` - The alarm webhook integration id.
            * `alarm_webhook_integration_name` - The alarm webhook integration name.
            * `alarm_webhook_is_at_all` - The alarm webhook is at all.
            * `end_time` - The end time.
            * `general_webhook_body` - The webhook body.
            * `general_webhook_headers` - The general webhook headers.
                * `key` - The key of the header.
                * `value` - The value of the header.
            * `general_webhook_method` - The general webhook method.
            * `general_webhook_url` - The webhook url.
            * `general_webhook` - The general webhook.
            * `receiver_channels` - The list of the receiver channels.
            * `receiver_names` - List of the receiver names.
            * `receiver_type` - The receiver type.
            * `start_time` - The start time.
        * `rule_node` - The rule node.
            * `children` - The children of the rule node.
                * `type` - The type of the rule node.
                * `value` - The value of the rule node.
            * `type` - The type of the rule node.
            * `value` - The value of the rule node.
    * `notify_type` - The notify group type.
    * `receivers` - List of IAM users to receive alerts.
        * `alarm_content_template_id` - The alarm content template id.
        * `alarm_webhook_at_users` - The alarm webhook at users.
        * `alarm_webhook_integration_id` - The alarm webhook integration id.
        * `alarm_webhook_integration_name` - The alarm webhook integration name.
        * `alarm_webhook_is_at_all` - The alarm webhook is at all.
        * `end_time` - The end time.
        * `general_webhook_body` - The webhook body.
        * `general_webhook_headers` - The general webhook headers.
            * `key` - The key of the header.
            * `value` - The value of the header.
        * `general_webhook_method` - The general webhook method.
        * `general_webhook_url` - The webhook url.
        * `general_webhook` - The general webhook.
        * `receiver_channels` - The list of the receiver channels.
        * `receiver_names` - List of the receiver names.
        * `receiver_type` - The receiver type.
        * `start_time` - The start time.
* `total_count` - The total count of query.



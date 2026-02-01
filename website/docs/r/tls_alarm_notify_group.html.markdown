---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm_notify_group"
sidebar_current: "docs-volcengine-resource-tls_alarm_notify_group"
description: |-
  Provides a resource to manage tls alarm notify group
---
# volcengine_tls_alarm_notify_group
Provides a resource to manage tls alarm notify group
## Example Usage
```hcl
resource "volcengine_tls_alarm_notify_group" "foo" {
  iam_project_name        = "default"
  alarm_notify_group_name = "tf-test-modify-b"
  notify_type             = ["Recovery"]
  receivers {
    receiver_type        = "User"
    receiver_names       = ["jonny"]
    receiver_channels    = ["Email"]
    start_time           = "23:00:00"
    end_time             = "23:59:59"
    general_webhook_url  = "https://www.volcengine.com/docs/6470/112220?lang=zh"
    general_webhook_body = "test"
    general_webhook_headers {
      key   = "test"
      value = "test"
    }
    general_webhook_method = "POST"
  }

  notice_rules {
    has_next     = false
    has_end_node = true
    rule_node {
      type  = "Operation"
      value = ["OR"]
      children {
        type = "Condition"
        value = [
          "NotifyType",
          "in",
          "[\"1\"]"
        ]
      }
    }
    receiver_infos {
      receiver_type     = "User"
      receiver_names    = ["jonny"]
      receiver_channels = ["Email", "Sms"]
      start_time        = "23:00:00"
      end_time          = "23:59:59"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `alarm_notify_group_name` - (Required) The name of the notify group.
* `iam_project_name` - (Optional) The name of the iam project.
* `notice_rules` - (Optional) The list of the notice rules.
* `notify_type` - (Optional) The notify type.
Trigger: Alarm Trigger
Recovery: Alarm Recovery.
* `receivers` - (Optional) List of IAM users to receive alerts.

The `children` object supports the following:

* `type` - (Optional) The type of the rule node.
* `value` - (Optional) The value of the rule node.

The `general_webhook_headers` object supports the following:

* `key` - (Optional) The key of the header.
* `value` - (Optional) The value of the header.

The `notice_rules` object supports the following:

* `has_end_node` - (Optional) Whether there is an end node behind.
* `has_next` - (Optional) Whether to continue to the next level of condition judgment.
* `receiver_infos` - (Optional) List of IAM users to receive alerts.
* `rule_node` - (Optional) The rule node.

The `receiver_infos` object supports the following:

* `alarm_content_template_id` - (Optional) The alarm content template id.
* `alarm_webhook_at_users` - (Optional) The alarm webhook at users.
* `alarm_webhook_integration_id` - (Optional) The alarm webhook integration id.
* `alarm_webhook_integration_name` - (Optional) The alarm webhook integration name.
* `alarm_webhook_is_at_all` - (Optional) The alarm webhook is at all.
* `end_time` - (Optional) The end time.
* `general_webhook_body` - (Optional) The webhook body.
* `general_webhook_headers` - (Optional) The general webhook headers.
* `general_webhook_method` - (Optional) The general webhook method.
* `general_webhook_url` - (Optional) The webhook url.
* `receiver_channels` - (Optional) The list of the receiver channels.
* `receiver_names` - (Optional) List of the receiver names.
* `receiver_type` - (Optional) The receiver type.
* `start_time` - (Optional) The start time.

The `receivers` object supports the following:

* `end_time` - (Required) The end time.
* `receiver_channels` - (Required) The list of the receiver channels. Currently supported channels: Email, Sms, Phone.
* `receiver_names` - (Required) List of the receiver names.
* `receiver_type` - (Required) The receiver type, Can be set as: `User`(The id of user).
* `start_time` - (Required) The start time.
* `alarm_content_template_id` - (Optional) The alarm content template id.
* `alarm_webhook_at_users` - (Optional) The alarm webhook at users.
* `alarm_webhook_integration_id` - (Optional) The alarm webhook integration id.
* `alarm_webhook_integration_name` - (Optional) The alarm webhook integration name.
* `alarm_webhook_is_at_all` - (Optional) The alarm webhook is at all.
* `general_webhook_body` - (Optional) The webhook body.
* `general_webhook_headers` - (Optional) The general webhook headers.
* `general_webhook_method` - (Optional) The general webhook method.
* `general_webhook_url` - (Optional) The webhook url.

The `rule_node` object supports the following:

* `children` - (Optional) The children of the rule node.
* `type` - (Optional) The type of the rule node.
* `value` - (Optional) The value of the rule node.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `alarm_notify_group_id` - The alarm notification group id.


## Import
tls alarm notify group can be imported using the id, e.g.
```
$ terraform import volcengine_tls_alarm_notify_group.default fa************
```


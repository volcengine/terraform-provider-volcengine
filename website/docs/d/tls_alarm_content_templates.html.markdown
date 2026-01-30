---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm_content_templates"
sidebar_current: "docs-volcengine-datasource-tls_alarm_content_templates"
description: |-
  Use this data source to query detailed information of tls alarm content templates
---
# volcengine_tls_alarm_content_templates
Use this data source to query detailed information of tls alarm content templates
## Example Usage
```hcl
data "volcengine_tls_alarm_content_templates" "foo" {
  # Filter by name (fuzzy matching)
  # alarm_content_template_name = "test-alarm"

  # Filter by specific IDs
  # ids = ["alarm-content-template-123456", "alarm-content-template-789012"]
}
```
## Argument Reference
The following arguments are supported:
* `alarm_content_template_id` - (Optional) The id of the alarm content template.
* `alarm_content_template_name` - (Optional) The name of the alarm content template. Fuzzy matching is supported.
* `asc` - (Optional) Whether to ascend.
* `ids` - (Optional) A list of alarm content template IDs.
* `order_field` - (Optional) The order field.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `templates` - The list of alarm content templates.
    * `alarm_content_template_id` - The ID of the alarm content template.
    * `alarm_content_template_name` - The name of the alarm content template.
    * `content` - The content of the alarm content template.
    * `create_time` - The creation time of the alarm content template.
    * `description` - The description of the alarm content template.
    * `ding_talk` - The ding_talk content of the alarm content template.
        * `content` - The content of the ding_talk content template.
        * `locale` - The locale of the ding_talk content template.
        * `title` - The title of the ding_talk content template.
    * `email` - The email content of the alarm content template.
        * `content` - The content of the email content template.
        * `locale` - The locale of the email content template.
        * `subject` - The subject of the email content template.
    * `is_default` - Whether the alarm content template is default.
    * `lark` - The lark content of the alarm content template.
        * `content` - The content of the lark content template.
        * `locale` - The locale of the lark content template.
        * `title` - The title of the lark content template.
    * `sms` - The sms content of the alarm content template.
        * `content` - The content of the sms content template.
        * `locale` - The locale of the sms content template.
    * `type` - The type of the alarm content template.
    * `update_time` - The update time of the alarm content template.
    * `vms` - The vms content of the alarm content template.
        * `content` - The content of the vms content template.
        * `locale` - The locale of the vms content template.
    * `webhook` - The webhook content of the alarm content template.
        * `content` - The content of the webhook content template.
    * `wechat` - The wechat content of the alarm content template.
        * `content` - The content of the wechat content template.
        * `locale` - The locale of the wechat content template.
* `total_count` - The total count of alarm content templates.



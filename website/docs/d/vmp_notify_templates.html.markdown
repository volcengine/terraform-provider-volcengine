---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_notify_templates"
sidebar_current: "docs-volcengine-datasource-vmp_notify_templates"
description: |-
  Use this data source to query detailed information of vmp notify templates
---
# volcengine_vmp_notify_templates
Use this data source to query detailed information of vmp notify templates
## Example Usage
```hcl
resource "volcengine_vmp_notify_template" "foo" {
  name        = "acc-test-vmp-notify-template"
  description = "acc-test-vmp"
  channel     = "WeComBotWebhook"
  active {
    title   = "acc-test-active-template-title"
    content = "acc-test-active-template-content"
  }
  resolved {
    title   = "acc-test-resolved-template-title"
    content = "acc-test-resolved-template-content"
  }
}

data "volcengine_vmp_notify_templates" "default" {
  ids = [volcengine_vmp_notify_template.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `channel` - (Optional) The channel of notify template. Valid values: `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.
* `ids` - (Optional) A list of IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of notify template. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `notify_templates` - The collection of query.
    * `active` - The active notify template info.
        * `content` - The content of notify template.
        * `title` - The title of notify template.
    * `channel` - The channel of notify template. Valid values: `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.
    * `create_time` - The create time of notify template.
    * `description` - The description of notify template.
    * `id` - The ID of notify template.
    * `name` - The name of notify template.
    * `resolved` - The resolved notify template info.
        * `content` - The content of notify template.
        * `title` - The title of notify template.
    * `update_time` - The update time of notify template.
* `total_count` - The total count of query.



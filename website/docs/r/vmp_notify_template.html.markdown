---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_notify_template"
sidebar_current: "docs-volcengine-resource-vmp_notify_template"
description: |-
  Provides a resource to manage vmp notify template
---
# volcengine_vmp_notify_template
Provides a resource to manage vmp notify template
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
```
## Argument Reference
The following arguments are supported:
* `active` - (Required) The active notify template info.
* `channel` - (Required, ForceNew) The channel of notify template. Valid values: `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.
* `name` - (Required) The name of notify template.
* `resolved` - (Required) The resolved notify template info.
* `description` - (Optional) The description of notify template.

The `active` object supports the following:

* `content` - (Required) The content of notify template.
* `title` - (Required) The title of notify template.

The `resolved` object supports the following:

* `content` - (Required) The content of notify template.
* `title` - (Required) The title of notify template.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of notify template.
* `update_time` - The update time of notify template.


## Import
VmpNotifyTemplate can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_notify_template.default resource_id
```


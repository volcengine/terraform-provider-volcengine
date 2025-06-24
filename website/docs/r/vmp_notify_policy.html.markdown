---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_notify_policy"
sidebar_current: "docs-volcengine-resource-vmp_notify_policy"
description: |-
  Provides a resource to manage vmp notify policy
---
# volcengine_vmp_notify_policy
Provides a resource to manage vmp notify policy
## Example Usage
```hcl
resource "volcengine_vmp_contact" "foo" {
  name  = "acc-test-contact"
  email = "acctest1@tftest.com"

  webhook {
    address = "https://www.acctest1.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest1.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest1.com"
    at_mobiles = ["18046891812"]
  }
  phone_number {
    country_code = "+86"
    number       = "18310101010"
  }
}

resource "volcengine_vmp_contact" "foo1" {
  name  = "acc-test-contact2"
  email = "acctest2@tftest.com"

  webhook {
    address = "https://www.acctest2.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest2.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest2.com"
    at_mobiles = ["18046891813"]
  }
  phone_number {
    country_code = "+86"
    number       = "18310101011"
  }
}

resource "volcengine_vmp_contact_group" "foo" {
  name        = "acc-test"
  contact_ids = [volcengine_vmp_contact.foo.id]
}

resource "volcengine_vmp_contact_group" "foo1" {
  name        = "acc-test-1"
  contact_ids = [volcengine_vmp_contact.foo1.id]
}

resource "volcengine_vmp_notify_policy" "foo" {
  name        = "acc-test-1"
  description = "acc-test-1"
  levels {
    level             = "P1"
    contact_group_ids = [volcengine_vmp_contact_group.foo.id]
    channels          = ["Email", "Webhook"]
  }
  levels {
    level             = "P0"
    contact_group_ids = [volcengine_vmp_contact_group.foo1.id]
    channels          = ["LarkBotWebhook"]
  }
}
```
## Argument Reference
The following arguments are supported:
* `levels` - (Required) The levels of the notify policy.
* `name` - (Required) The name of the notify policy.
* `channel_notify_template_ids` - (Optional) The channel notify template for the alarm notification policy.
* `description` - (Optional) The description of the notify policy.

The `levels` object supports the following:

* `channels` - (Required) The alarm notification method of the alarm notification policy, the optional value can be `Email`, `Webhook`, `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.
* `contact_group_ids` - (Required) The contact group for the alarm notification policy.
* `level` - (Required) The level of the policy, the value can be one of the following: `P0`, `P1`, `P2`.
* `resolved_channels` - (Required) The resolved alarm notification method of the alarm notification policy, the optional value can be `Email`, `Webhook`, `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VMP Notify Policy can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_notify_policy.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```


---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_notify_policies"
sidebar_current: "docs-volcengine-datasource-vmp_notify_policies"
description: |-
  Use this data source to query detailed information of vmp notify policies
---
# volcengine_vmp_notify_policies
Use this data source to query detailed information of vmp notify policies
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

data "volcengine_vmp_notify_policies" "foo" {
  ids = [volcengine_vmp_notify_policy.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `channel_notify_template_ids` - (Optional) The channel notify template for the alarm notification policy.
* `contact_group_ids` - (Optional) The contact group for the alarm notification policy.
* `ids` - (Optional) A list of notify policy ids.
* `name` - (Optional) The name of notify policy.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `notify_policies` - The list of notify policies.
    * `channel_notify_template_ids` - The channel notify template for the alarm notification policy.
    * `create_time` - The create time of notify policy.
    * `description` - The description of notify policy.
    * `id` - The id of the notify policy.
    * `levels` - The levels of the notify policy.
        * `channels` - The alarm notification method of the alarm notification policy.
        * `contact_group_ids` - The contact group for the alarm notification policy.
        * `level` - The level of the policy.
        * `resolved_channels` - The resolved alarm notification method of the alarm notification policy.
    * `name` - The name of notify policy.
* `total_count` - The total count of query.



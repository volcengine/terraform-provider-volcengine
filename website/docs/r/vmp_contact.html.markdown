---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_contact"
sidebar_current: "docs-volcengine-resource-vmp_contact"
description: |-
  Provides a resource to manage vmp contact
---
# volcengine_vmp_contact
Provides a resource to manage vmp contact
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
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of the contact.
* `ding_talk_bot_webhook` - (Optional) The ding talk bot webhook of contact.
* `email` - (Optional) The email of the contact.
* `lark_bot_webhook` - (Optional) The lark bot webhook of contact.
* `phone_number` - (Optional) The phone number of contact.
* `we_com_bot_webhook` - (Optional) The we com bot webhook of contact.
* `webhook` - (Optional) The webhook of contact.

The `ding_talk_bot_webhook` object supports the following:

* `address` - (Required) The address of webhook.
* `at_mobiles` - (Optional) The mobiles of user.
* `at_user_ids` - (Optional) The ids of user.
* `secret_key` - (Optional) The secret key of webhook.

The `lark_bot_webhook` object supports the following:

* `address` - (Required) The address of webhook.
* `secret_key` - (Optional) The secret key of webhook.

The `phone_number` object supports the following:

* `country_code` - (Required) The country code of phone number. The value is `+86`.
* `number` - (Required) The number of phone number.

The `we_com_bot_webhook` object supports the following:

* `address` - (Required) The address of webhook.
* `at_user_ids` - (Optional) The ids of user.

The `webhook` object supports the following:

* `address` - (Required) The address of webhook.
* `token` - (Optional) The token of webhook.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `contact_group_ids` - A list of contact group ids.
* `create_time` - The create time of contact.
* `email_active` - Whether the email of contact active.


## Import
VMP Contact can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_contact.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```


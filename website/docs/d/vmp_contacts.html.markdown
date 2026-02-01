---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_contacts"
sidebar_current: "docs-volcengine-datasource-vmp_contacts"
description: |-
  Use this data source to query detailed information of vmp contacts
---
# volcengine_vmp_contacts
Use this data source to query detailed information of vmp contacts
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

data "volcengine_vmp_contacts" "foo" {
  ids = [volcengine_vmp_contact.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `email` - (Optional) The email of contact.
* `ids` - (Optional) A list of contact ids.
* `name` - (Optional) The name of contact.
* `output_file` - (Optional) File name where to save data source results.
* `sort_by` - (Optional) The sort field of query.
* `sort_order` - (Optional) The sort order of query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `contacts` - The collection of query.
    * `contact_group_ids` - A list of contact group ids.
    * `create_time` - The create time of contact.
    * `ding_talk_bot_webhook` - The ding talk bot webhook of contact.
        * `address` - The address of webhook.
        * `at_mobiles` - The mobiles of user.
        * `at_user_ids` - The ids of user.
        * `secret_key` - The secret key of webhook.
    * `email_active` - Whether the email of contact active.
    * `email` - The email of contact.
    * `id` - The ID of contact.
    * `lark_bot_webhook` - The lark bot webhook of contact.
        * `address` - The address of webhook.
        * `secret_key` - The secret key of webhook.
    * `name` - The name of contact.
    * `phone_number_active` - Whether phone number is active.
    * `phone_number` - The phone number of contact.
        * `country_code` - The country code of phone number.
        * `number` - The number of phone number.
    * `we_com_bot_webhook` - The we com bot webhook of contact.
        * `address` - The address of webhook.
        * `at_user_ids` - The ids of user.
    * `webhook` - The webhook of contact.
        * `address` - The address of webhook.
        * `token` - The token of webhook.
* `total_count` - The total count of query.



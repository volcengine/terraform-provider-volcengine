---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_contact_groups"
sidebar_current: "docs-volcengine-datasource-vmp_contact_groups"
description: |-
  Use this data source to query detailed information of vmp contact groups
---
# volcengine_vmp_contact_groups
Use this data source to query detailed information of vmp contact groups
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
  contact_ids = [volcengine_vmp_contact.foo.id, volcengine_vmp_contact.foo1.id]
}

data "volcengine_vmp_contact_groups" "foo" {
  ids = [volcengine_vmp_contact_group.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of contact group ids.
* `name` - (Optional) The name of contact group.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `contact_groups` - The collection of query.
    * `contact_ids` - A list of contact IDs.
    * `create_time` - The create time of contact group.
    * `id` - The ID of contact group.
    * `name` - The name of contact group.
* `total_count` - The total count of query.



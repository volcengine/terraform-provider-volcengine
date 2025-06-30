---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_contact_group"
sidebar_current: "docs-volcengine-resource-vmp_contact_group"
description: |-
  Provides a resource to manage vmp contact group
---
# volcengine_vmp_contact_group
Provides a resource to manage vmp contact group
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
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of the contact group.
* `contact_ids` - (Optional) A list of contact IDs.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of contact group.


## Import
VMP Contact Group can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_contact_group.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```


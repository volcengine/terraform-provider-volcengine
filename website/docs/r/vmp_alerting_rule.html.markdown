---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_alerting_rule"
sidebar_current: "docs-volcengine-resource-vmp_alerting_rule"
description: |-
  Provides a resource to manage vmp alerting rule
---
# volcengine_vmp_alerting_rule
Provides a resource to manage vmp alerting rule
## Example Usage
```hcl
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "**********"
}

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

resource "volcengine_vmp_notify_group_policy" "foo" {
  name        = "acc-test-1"
  description = "acc-test-1"
  levels {
    level           = "P2"
    group_by        = ["__rule__"]
    group_wait      = "35"
    group_interval  = "30"
    repeat_interval = "30"
  }
  levels {
    level           = "P0"
    group_by        = ["__rule__"]
    group_wait      = "30"
    group_interval  = "30"
    repeat_interval = "30"
  }
  levels {
    level           = "P1"
    group_by        = ["__rule__"]
    group_wait      = "40"
    group_interval  = "45"
    repeat_interval = "30"
  }
}

resource "volcengine_vmp_alerting_rule" "foo" {
  name                   = "acc-test-1"
  description            = "acc-test-1"
  notify_policy_id       = volcengine_vmp_notify_policy.foo.id
  notify_group_policy_id = volcengine_vmp_notify_group_policy.foo.id
  query {
    workspace_id = volcengine_vmp_workspace.foo.id
    prom_ql      = "sum(up)"
  }
  levels {
    level      = "P0"
    for        = "0s"
    comparator = ">="
    threshold  = 2.0
  }
  levels {
    level      = "P1"
    for        = "0s"
    comparator = ">="
    threshold  = 1.0
  }
  levels {
    level      = "P2"
    for        = "0s"
    comparator = ">="
    threshold  = 0.5
  }
}
```
## Argument Reference
The following arguments are supported:
* `levels` - (Required) The alerting levels of the vmp alerting rule.
* `name` - (Required) The name of the vmp alerting rule.
* `notify_group_policy_id` - (Required) The id of the notify group policy.
* `query` - (Required) The alerting query of the vmp alerting rule.
* `description` - (Optional) The description of the vmp alerting rule.
* `notify_policy_id` - (Optional) The id of the notify policy.

The `levels` object supports the following:

* `comparator` - (Required) The comparator of the vmp alerting rule. Valid values: `>`, `>=`, `<`, `<=`, `==`, `!=`.
* `for` - (Required) The duration of the alerting rule. Valid values: `0s`, `1m`, `2m`, `5m`, `10m`.
* `level` - (Required) The level of the vmp alerting rule. Valid values: `P0`, `P1`, `P2`. The value of this field cannot be duplicate.
* `threshold` - (Required) The threshold of the vmp alerting rule.

The `query` object supports the following:

* `prom_ql` - (Required) The prom ql of query.
* `workspace_id` - (Required) The id of the workspace.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the vmp alerting rule.
* `status` - The status of the vmp alerting rule.
* `update_time` - The update time of the vmp alerting rule.


## Import
VmpAlertingRule can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_alerting_rule.default 5bd29e81-2717-4ac8-a1a6-d76da2b1****
```


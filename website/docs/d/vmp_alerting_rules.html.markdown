---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_alerting_rules"
sidebar_current: "docs-volcengine-datasource-vmp_alerting_rules"
description: |-
  Use this data source to query detailed information of vmp alerting rules
---
# volcengine_vmp_alerting_rules
Use this data source to query detailed information of vmp alerting rules
## Example Usage
```hcl
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "***********"
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

data "volcengine_vmp_alerting_rules" "foo" {
  ids = [volcengine_vmp_alerting_rule.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of vmp alerting rule IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of vmp alerting rule. This field support fuzzy query.
* `notify_group_policy_ids` - (Optional) A list of notify group policy IDs.
* `notify_policy_ids` - (Optional) A list of notify policy IDs.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) The status of vmp alerting rule. Valid values: `Running`, `Disabled`.
* `type` - (Optional) The type of vmp alerting rule. Valid values: `vmp/PromQL`.
* `workspace_id` - (Optional) The workspace id of vmp alerting rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `alerting_rules` - The collection of query.
    * `annotations` - The annotations of the vmp alerting rule.
        * `name` - The name of the annotation.
        * `value` - The value of the annotation.
    * `create_time` - The create time of the vmp alerting rule.
    * `description` - The description of the vmp alerting rule.
    * `id` - The id of the vmp alerting rule.
    * `labels` - The labels of the vmp alerting rule.
        * `key` - The name of the label.
        * `value` - The value of the label.
    * `levels` - The alerting levels of the vmp alerting rule.
        * `comparator` - The comparator of the vmp alerting rule.
        * `for` - The duration of the alerting rule.
        * `level` - The level of the vmp alerting rule.
        * `threshold` - The threshold of the vmp alerting rule.
    * `name` - The name of the vmp alerting rule.
    * `notify_group_policy_id` - The notify group policy id of the vmp alerting rule.
    * `notify_policy_id` - The notify policy id of the vmp alerting rule.
    * `query` - The alerting query of the vmp alerting rule.
        * `prom_ql` - The prom ql of query.
        * `workspace_id` - The id of the workspace.
    * `status` - The status of the vmp alerting rule.
    * `type` - The type of the vmp alerting rule.
    * `update_time` - The update time of the vmp alerting rule.
* `total_count` - The total count of query.



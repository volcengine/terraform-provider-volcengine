---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_vpc_firewall_acl_rule"
sidebar_current: "docs-volcengine-resource-cfw_vpc_firewall_acl_rule"
description: |-
  Provides a resource to manage cfw vpc firewall acl rule
---
# volcengine_cfw_vpc_firewall_acl_rule
Provides a resource to manage cfw vpc firewall acl rule
## Example Usage
```hcl
resource "volcengine_cfw_address_book" "foo" {
  group_name   = "acc-test-address-book"
  description  = "acc-test"
  group_type   = "ip"
  address_list = ["192.168.1.1", "192.168.2.2"]
}

resource "volcengine_cfw_vpc_firewall_acl_rule" "foo" {
  vpc_firewall_id   = "vfw-ydmjakzksgf7u99j****"
  action            = "accept"
  destination_type  = "group"
  destination       = volcengine_cfw_address_book.foo.id
  proto             = "TCP"
  source_type       = "net"
  source            = "0.0.0.0/0"
  description       = "acc-test-control-policy"
  dest_port_type    = "port"
  dest_port         = "300"
  repeat_type       = "Weekly"
  repeat_start_time = "01:00"
  repeat_end_time   = "11:00"
  repeat_days       = [2, 5]
  start_time        = 1736092800
  end_time          = 1738339140
  priority          = 1
  status            = true
}
```
## Argument Reference
The following arguments are supported:
* `action` - (Required) The action of the vpc firewall acl rule. Valid values: `accept`, `deny`, `monitor`.
* `destination_type` - (Required) The destination type of the vpc firewall acl rule. Valid values: `net`, `group`, `location`, `domain`.
* `destination` - (Required) The destination of the vpc firewall acl rule.
* `proto` - (Required) The proto of the vpc firewall acl rule. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.
* `source_type` - (Required) The source type of the vpc firewall acl rule. Valid values: `net`, `group`.
* `source` - (Required) The source of the vpc firewall acl rule.
* `vpc_firewall_id` - (Required, ForceNew) The id of the vpc firewall.
* `description` - (Optional) The description of the vpc firewall acl rule.
* `dest_port_type` - (Optional) The dest port type of the vpc firewall acl rule. Valid values: `port`, `group`.
* `dest_port` - (Optional) The dest port of the vpc firewall acl rule.
* `end_time` - (Optional) The end time of the vpc firewall acl rule. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
 When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
* `priority` - (Optional) The priority of the vpc firewall acl rule. Default is 0. This field is only effective when creating a control policy.0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.
* `repeat_days` - (Optional) The repeat days of the vpc firewall acl rule. When the value of repeat_type is one of `Weekly`, `Monthly`, this field is required.
 When the repeat_type is `Weekly`, the valid value range is 0~6.
 When the repeat_type is `Monthly`, the valid value range is 1~31.
* `repeat_end_time` - (Optional) The repeat end time of the vpc firewall acl rule. Accurate to the minute, in the format of hh: mm. For example: 12:00.
 When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.
* `repeat_start_time` - (Optional) The repeat start time of the vpc firewall acl rule. Accurate to the minute, in the format of hh: mm. For example: 12:00.
 When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.
* `repeat_type` - (Optional) The repeat type of the vpc firewall acl rule. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
* `start_time` - (Optional) The start time of the vpc firewall acl rule. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
 When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
* `status` - (Optional) Whether to enable the vpc firewall acl rule. Default is false.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account id of the vpc firewall acl rule.
* `effect_status` - The effect status of the vpc firewall acl rule. 1: Not yet effective, 2: Issued in progress, 3: Effective.
* `hit_cnt` - The hit count of the vpc firewall acl rule.
* `is_effected` - Whether the vpc firewall acl rule is effected.
* `prio` - The priority of the vpc firewall acl rule.
* `rule_id` - The rule id of the vpc firewall acl rule.
* `update_time` - The update time of the vpc firewall acl rule.
* `use_count` - The use count of the vpc firewall acl rule.
* `vpc_firewall_name` - The name of the vpc firewall.


## Import
VpcFirewallAclRule can be imported using the vpc_firewall_id:rule_id, e.g.
```
$ terraform import volcengine_vpc_firewall_acl_rule.default resource_id
```


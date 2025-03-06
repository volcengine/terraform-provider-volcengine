---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_nat_firewall_control_policy"
sidebar_current: "docs-volcengine-resource-cfw_nat_firewall_control_policy"
description: |-
  Provides a resource to manage cfw nat firewall control policy
---
# volcengine_cfw_nat_firewall_control_policy
Provides a resource to manage cfw nat firewall control policy
## Example Usage
```hcl
resource "volcengine_cfw_address_book" "foo" {
  group_name   = "acc-test-address-book"
  description  = "acc-test"
  group_type   = "ip"
  address_list = ["192.168.1.1", "192.168.2.2"]
}

resource "volcengine_cfw_nat_firewall_control_policy" "foo" {
  direction         = "in"
  nat_firewall_id   = "nfw-ydmkayvjsw2vsavx****"
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
  repeat_days       = [2, 5, 4]
  start_time        = 1736092800
  end_time          = 1738339140
  priority          = 1
  status            = true
}
```
## Argument Reference
The following arguments are supported:
* `action` - (Required) The action of the nat firewall control policy. Valid values: `accept`, `deny`, `monitor`.
* `destination_type` - (Required) The destination type of the nat firewall control policy. Valid values: `net`, `group`, `location`, `domain`.
* `destination` - (Required) The destination of the nat firewall control policy.
* `direction` - (Required, ForceNew) The direction of the nat firewall control policy. Valid values: `in`, `out`.
* `nat_firewall_id` - (Required, ForceNew) The id of the nat firewall.
* `proto` - (Required) The proto of the nat firewall control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.
* `source_type` - (Required) The source type of the nat firewall control policy. Valid values: `net`, `group`.
* `source` - (Required) The source of the nat firewall control policy.
* `description` - (Optional) The description of the nat firewall control policy.
* `dest_port_type` - (Optional) The dest port type of the nat firewall control policy. Valid values: `port`, `group`.
* `dest_port` - (Optional) The dest port of the nat firewall control policy.
* `end_time` - (Optional) The end time of the nat firewall control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
 When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
* `priority` - (Optional) The priority of the nat firewall control policy. Default is 0. This field is only effective when creating a control policy.0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.
* `repeat_days` - (Optional) The repeat days of the nat firewall control policy. When the value of repeat_type is one of `Weekly`, `Monthly`, this field is required.
 When the repeat_type is `Weekly`, the valid value range is 0~6.
 When the repeat_type is `Monthly`, the valid value range is 1~31.
* `repeat_end_time` - (Optional) The repeat end time of the nat firewall control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
 When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.
* `repeat_start_time` - (Optional) The repeat start time of the nat firewall control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
 When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.
* `repeat_type` - (Optional) The repeat type of the nat firewall control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
* `start_time` - (Optional) The start time of the nat firewall control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
 When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
* `status` - (Optional) Whether to enable the nat firewall control policy. Default is false.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account id of the nat firewall control policy.
* `effect_status` - The effect status of the nat firewall control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.
* `hit_cnt` - The hit count of the nat firewall control policy.
* `is_effected` - Whether the nat firewall control policy is effected.
* `nat_firewall_name` - The name of the nat firewall.
* `prio` - The priority of the nat firewall control policy.
* `rule_id` - The rule id of the nat firewall control policy.
* `update_time` - The update time of the nat firewall control policy.
* `use_count` - The use count of the nat firewall control policy.


## Import
NatFirewallControlPolicy can be imported using the direction_nat_firewall_id:rule_id, e.g.
```
$ terraform import volcengine_nat_firewall_control_policy.default resource_id
```


---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_nat_firewall_control_policy_priority"
sidebar_current: "docs-volcengine-resource-cfw_nat_firewall_control_policy_priority"
description: |-
  Provides a resource to manage cfw nat firewall control policy priority
---
# volcengine_cfw_nat_firewall_control_policy_priority
Provides a resource to manage cfw nat firewall control policy priority
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

resource "volcengine_cfw_nat_firewall_control_policy_priority" "foo" {
  direction       = "in"
  nat_firewall_id = "nfw-ydmkayvjsw2vsavx****"
  rule_id         = volcengine_cfw_nat_firewall_control_policy.foo.rule_id
  new_prio        = 2
}
```
## Argument Reference
The following arguments are supported:
* `direction` - (Required, ForceNew) The direction of the nat firewall control policy. Valid values: `in`, `out`.
* `nat_firewall_id` - (Required, ForceNew) The id of the nat firewall.
* `rule_id` - (Required, ForceNew) The rule id of the nat firewall control policy.
* `new_prio` - (Optional) The new priority of the nat firewall control policy. The priority increases in order from 1, with lower priority indicating higher priority.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `prio` - The priority of the nat firewall control policy.


## Import
NatFirewallControlPolicyPriority can be imported using the direction_nat_firewall_id:rule_id, e.g.
```
$ terraform import volcengine_nat_firewall_control_policy_priority.default resource_id
```


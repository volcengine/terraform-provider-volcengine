---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_address_book"
sidebar_current: "docs-volcengine-resource-cfw_address_book"
description: |-
  Provides a resource to manage cfw address book
---
# volcengine_cfw_address_book
Provides a resource to manage cfw address book
## Example Usage
```hcl
resource "volcengine_cfw_address_book" "foo" {
  group_name   = "acc-test-address-book-1"
  description  = "acc-test"
  group_type   = "ip"
  address_list = ["192.168.1.1", "192.168.2.2"]
}

#resource "volcengine_cfw_control_policy" "foo" {
#  direction         = "in"
#  action            = "accept"
#  destination_type  = "group"
#  destination       = "${volcengine_cfw_address_book.foo.id}"
#  proto             = "TCP"
#  source_type       = "net"
#  source            = "0.0.0.0/0"
#  description       = "acc-test-control-policy-new-0"
#  dest_port_type    = "port"
#  dest_port         = "300"
#  repeat_type       = "Weekly"
#  repeat_start_time = "01:00"
#  repeat_end_time   = "11:00"
#  repeat_days       = [2, 5, 4]
#  start_time        = 1736092800
#  end_time          = 1738339140
#  priority          = 1
#  status            = true
#}


#resource "volcengine_cfw_vpc_firewall_acl_rule" "foo" {
#  vpc_firewall_id   = "vfw-ydmjakzksgf7u99j6sby"
#  action            = "accept"
#  destination_type  = "group"
#  destination       = volcengine_cfw_address_book.foo.id
#  proto             = "TCP"
#  source_type       = "net"
#  source            = "0.0.0.0/0"
#  description       = "acc-test-control-policy-new"
#  dest_port_type    = "port"
#  dest_port         = "300"
#  repeat_type       = "Weekly"
#  repeat_start_time = "01:00"
#  repeat_end_time   = "11:00"
#  repeat_days       = [2, 5, 3]
#  start_time        = 1736092800
#  end_time          = 1738339140
#  priority          = 0
#  status            = true
#}
#
#resource "volcengine_cfw_vpc_firewall_acl_rule_priority" "foo" {
#  vpc_firewall_id = "vfw-ydmjakzksgf7u99j6sby"
#  rule_id         = volcengine_cfw_vpc_firewall_acl_rule.foo.rule_id
#  new_prio        = 3
#}

resource "volcengine_cfw_nat_firewall_control_policy" "foo" {
  direction         = "in"
  nat_firewall_id   = "nfw-ydmkayvjsw2vsavxj9id"
  action            = "accept"
  destination_type  = "group"
  destination       = volcengine_cfw_address_book.foo.id
  proto             = "TCP"
  source_type       = "net"
  source            = "0.0.0.0/0"
  description       = "acc-test-control-policy-new-0"
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
  nat_firewall_id = "nfw-ydmkayvjsw2vsavxj9id"
  rule_id         = volcengine_cfw_nat_firewall_control_policy.foo.rule_id
  new_prio        = 2
}

data "volcengine_cfw_nat_firewall_control_policies" "foo" {
  direction       = "in"
  nat_firewall_id = "nfw-ydmkayvjsw2vsavxj9id"
  rule_id         = [volcengine_cfw_nat_firewall_control_policy.foo.rule_id]
}

#resource "volcengine_vpc" "foo" {
#  vpc_name   = "acc-test-vpc"
#  cidr_block = "172.16.0.0/16"
#  count      = 2
#}
#
#resource "volcengine_cfw_dns_control_policy" "foo" {
#  description      = "acc-test-dns-control-policy"
#  destination_type = "domain"
#  destination      = "www.test.com"
#  source {
#    vpc_id = volcengine_vpc.foo[0].id
#    region = "cn-guilin-boe"
#  }
#  source {
#    vpc_id = volcengine_vpc.foo[1].id
#    region = "cn-guilin-boe"
#  }
#}

#data "volcengine_cfw_address_books" "foo" {
#  group_name = "acc-test-address-book"
#}
#
#data "volcengine_cfw_control_policies" "foo" {
#  direction = "in"
#  action    = ["deny"]
#}
#
#data "volcengine_cfw_vpc_firewall_acl_rules" "foo" {
#  vpc_firewall_id = "vfw-ydmjakzksgf7u99j6sby"
#  action          = ["deny"]
#}
#
#data "volcengine_cfw_dns_control_policies" "foo" {
#  ids = [volcengine_cfw_dns_control_policy.foo.id]
#}
```
## Argument Reference
The following arguments are supported:
* `address_list` - (Required) The address list of the address book.
 When group_type is `ip`, fill in IPv4/CIDRV4 addresses in the address list.
 When group_type is `port`, fill in the port information in the address list, supporting two formats: 22 and 100/200.
 When group_type is `domain`, fill in the domain name information in the address list.
* `group_name` - (Required) The name of the address book.
* `group_type` - (Required, ForceNew) The type of the address book. Valid values: `ip`, `port`, `domain`.
* `description` - (Optional) The description of the address book.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ref_cnt` - The reference count of the address book.


## Import
AddressBook can be imported using the id, e.g.
```
$ terraform import volcengine_address_book.default resource_id
```


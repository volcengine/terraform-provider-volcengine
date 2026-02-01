---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_traffic_mirror_filter_rule"
sidebar_current: "docs-volcengine-resource-traffic_mirror_filter_rule"
description: |-
  Provides a resource to manage traffic mirror filter rule
---
# volcengine_traffic_mirror_filter_rule
Provides a resource to manage traffic mirror filter rule
## Example Usage
```hcl
resource "volcengine_traffic_mirror_filter" "foo" {
  traffic_mirror_filter_name = "acc-test-traffic-mirror-filter"
  description                = "acc-test"
  project_name               = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_traffic_mirror_filter_rule" "foo-ingress" {
  traffic_mirror_filter_id = volcengine_traffic_mirror_filter.foo.id
  traffic_direction        = "ingress"
  description              = "acc-test"
  priority                 = 1
  policy                   = "reject"
  protocol                 = "all"
  source_cidr_block        = "10.0.1.0/24"
  source_port_range        = "-1/-1"
  destination_cidr_block   = "10.0.0.0/24"
  destination_port_range   = "-1/-1"
}

resource "volcengine_traffic_mirror_filter_rule" "foo-egress" {
  traffic_mirror_filter_id = volcengine_traffic_mirror_filter.foo.id
  traffic_direction        = "egress"
  description              = "acc-test"
  priority                 = 2
  policy                   = "reject"
  protocol                 = "tcp"
  source_cidr_block        = "10.0.1.0/24"
  source_port_range        = "80/80"
  destination_cidr_block   = "10.0.0.0/24"
  destination_port_range   = "88/90"
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Required) The destination cidr block of traffic mirror filter rule.
* `policy` - (Required) The policy of traffic mirror filter rule. Valid values: `accept`, `reject`.
* `protocol` - (Required) The protocol of traffic mirror filter rule. Valid values: `tcp`, `udp`, `icmp`, `all`.
* `source_cidr_block` - (Required) The source cidr block of traffic mirror filter rule.
* `traffic_direction` - (Required) The traffic direction of traffic mirror filter rule. Valid values: `ingress`; `egress`.
* `traffic_mirror_filter_id` - (Required, ForceNew) The ID of traffic mirror filter.
* `description` - (Optional) The description of traffic mirror filter rule.
* `destination_port_range` - (Optional) The destination port range of traffic mirror filter rule. When the protocol is `all` or `icmp`, the value is `-1/-1`. 
When the protocol is `tcp` or `udp`, the value can be `1/200`, `80/80`, which means port 1 to port 200, port 80.
* `priority` - (Optional) The priority of traffic mirror filter rule. Valid values: 1~1000. Default value is 1.
* `source_port_range` - (Optional) The source port range of traffic mirror filter rule. When the protocol is `all` or `icmp`, the value is `-1/-1`. 
When the protocol is `tcp` or `udp`, the value can be `1/200`, `80/80`, which means port 1 to port 200, port 80.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - The create time of traffic mirror filter rule.
* `status` - The status of traffic mirror filter rule.
* `traffic_mirror_filter_rule_id` - The ID of traffic mirror filter rule.
* `updated_at` - The last update time of traffic mirror filter rule.


## Import
TrafficMirrorFilterRule can be imported using the id, e.g.
```
$ terraform import volcengine_traffic_mirror_filter_rule.default resource_id
```


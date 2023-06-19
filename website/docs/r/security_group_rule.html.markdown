---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_security_group_rule"
sidebar_current: "docs-volcengine-resource-security_group_rule"
description: |-
  Provides a resource to manage security group rule
---
# volcengine_security_group_rule
Provides a resource to manage security group rule
## Example Usage
```hcl
resource "volcengine_security_group_rule" "g1test3" {
  direction         = "egress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "10.0.0.0/8"
  description       = "tft1234"
}

resource "volcengine_security_group_rule" "g1test2" {
  direction         = "egress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "10.0.0.0/24"
}

resource "volcengine_security_group_rule" "g1test1" {
  direction         = "egress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "10.0.0.0/24"
  priority          = 2
}


resource "volcengine_security_group_rule" "g1test0" {
  direction         = "ingress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 80
  port_end          = 80
  cidr_ip           = "10.0.0.0/24"
  priority          = 2
  policy            = "drop"
  description       = "tft"
}

resource "volcengine_security_group_rule" "g1test06" {
  direction         = "ingress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  source_group_id   = "sg-3rfe5j4xdnklc5zsk2hcw5c6q"
  priority          = 2
  policy            = "drop"
  description       = "tft"
}
```
## Argument Reference
The following arguments are supported:
* `direction` - (Required, ForceNew) Direction of rule, ingress (inbound) or egress (outbound).
* `port_end` - (Required, ForceNew) Port end of egress/ingress Rule.
* `port_start` - (Required, ForceNew) Port start of egress/ingress Rule.
* `protocol` - (Required, ForceNew) Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
* `security_group_id` - (Required, ForceNew) Id of SecurityGroup.
* `cidr_ip` - (Optional, ForceNew) Cidr ip of egress/ingress Rule.
* `description` - (Optional) description of a egress rule.
* `policy` - (Optional, ForceNew) Access strategy.
* `priority` - (Optional, ForceNew) Priority of a security group rule.
* `source_group_id` - (Optional, ForceNew) ID of the source security group whose access permission you want to set.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - Status of SecurityGroup.


## Import
SecurityGroupRule can be imported using the id, e.g.
```
$ terraform import volcengine_security_group_rule.default ID is a string concatenated with colons(SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp:SourceGroupId:Direction:Policy:Priority)
```


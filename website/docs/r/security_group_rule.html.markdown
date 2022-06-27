---
subcategory: "VPC"
layout: "vestack"
page_title: "Vestack: vestack_security_group_rule"
sidebar_current: "docs-vestack-resource-security_group_rule"
description: |-
  Provides a resource to manage security group rule
---
# vestack_security_group_rule
Provides a resource to manage security group rule
## Example Usage
```hcl
resource "vestack_security_group_rule" "g1test3" {
  direction         = "egress"
  security_group_id = "sg-273ycgql3ig3k7fap8t3dyvqx"
  protocol          = "tcp"
  port_start        = "8000"
  port_end          = "9003"
  cidr_ip           = "10.0.0.0/8"
}
```
## Argument Reference
The following arguments are supported:
* `direction` - (Required, ForceNew) Direction of rule, ingress (inbound) or egress (outbound).
* `port_end` - (Required, ForceNew) Port end of egress/ingress Rule.
* `port_start` - (Required, ForceNew) Port start of egress/ingress Rule.
* `protocol` - (Required, ForceNew) Protocol of the SecurityGroup.
* `security_group_id` - (Required) Id of SecurityGroup.
* `cidr_ip` - (Optional, ForceNew) Cidr ip of egress/ingress Rule.
* `description` - (Optional) description of a egress rule.
* `policy` - (Optional, ForceNew) Access strategy.
* `priority` - (Optional) Priority of a security group rule.
* `source_group_id` - (Optional) ID of the source security group whose access permission you want to set.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - Status of SecurityGroup.


## Import
SecurityGroupRule can be imported using the id, e.g.
```
$ terraform import vestack_security_group_rule.default ID is a string concatenated with colons(SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp)
```


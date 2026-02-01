---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_security_group_rules"
sidebar_current: "docs-volcengine-datasource-security_group_rules"
description: |-
  Use this data source to query detailed information of security group rules
---
# volcengine_security_group_rules
Use this data source to query detailed information of security group rules
## Example Usage
```hcl
data "volcengine_security_group_rules" "default" {
  security_group_id = "sg-13f2nau7x93wg3n6nu3z5sxib"
}
```
## Argument Reference
The following arguments are supported:
* `security_group_id` - (Required) SecurityGroup ID.
* `cidr_ip` - (Optional) Cidr ip of egress/ingress Rule.
* `direction` - (Optional) Direction of rule, ingress (inbound) or egress (outbound).
* `output_file` - (Optional) File name where to save data source results.
* `protocol` - (Optional) Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.
* `source_group_id` - (Optional) ID of the source security group whose access permission you want to set.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `security_group_rules` - The collection of SecurityGroup query.
    * `cidr_ip` - Cidr ip of egress/ingress Rule.
    * `creation_time` - The creation time of security group rule.
    * `description` - description of a group rule.
    * `direction` - Direction of rule, ingress (inbound) or egress (outbound).
    * `policy` - Access strategy.
    * `port_end` - Port end of egress/ingress Rule.
    * `port_start` - Port start of egress/ingress Rule.
    * `priority` - Priority of a security group rule.
    * `protocol` - Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.
    * `security_group_id` - Id of SecurityGroup.
    * `source_group_id` - ID of the source security group whose access permission you want to set.
    * `update_time` - The update time of security group rule.



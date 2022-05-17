---
subcategory: "VPC"
layout: "vestack"
page_title: "Vestack: vestack_security_group"
sidebar_current: "docs-vestack-resource-security_group"
description: |-
  Provides a resource to manage security group
---
# vestack_security_group
Provides a resource to manage security group
## Example Usage
```hcl
resource "vestack_security_group" "g1test1" {
  vpc_id = "sg-273ycgql3ig3k7fap8t3dyvqx"
}
```
## Argument Reference
The following arguments are supported:
* `vpc_id` - (Required, ForceNew) Id of the VPC.
* `description` - (Optional) Description of SecurityGroup.
* `security_group_name` - (Optional) Name of SecurityGroup.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - Creation time of SecurityGroup.
* `status` - Status of SecurityGroup.


## Import
SecurityGroup can be imported using the id, e.g.
```
$ terraform import vestack_security_group.default sg-273ycgql3ig3k7fap8t3dyvqx
```


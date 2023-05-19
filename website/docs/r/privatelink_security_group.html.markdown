---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_security_group"
sidebar_current: "docs-volcengine-resource-privatelink_security_group"
description: |-
  Provides a resource to manage privatelink security group
---
# volcengine_privatelink_security_group
Provides a resource to manage privatelink security group
## Example Usage
```hcl
resource "volcengine_privatelink_security_group" "foo" {
  endpoint_id       = "ep-2byz5npiuu1hc2dx0efkv7ehc"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Required, ForceNew) The id of the endpoint.
* `security_group_id` - (Required, ForceNew) The id of the security group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateLink Security Group Service can be imported using the endpoint id and security group id, e.g.
```
$ terraform import volcengine_privatelink_security_group.default ep-2fe630gurkl37k5gfuy33****:sg-xxxxx
```


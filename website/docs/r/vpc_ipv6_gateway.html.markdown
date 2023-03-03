---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_ipv6_gateway"
sidebar_current: "docs-volcengine-resource-vpc_ipv6_gateway"
description: |-
  Provides a resource to manage vpc ipv6 gateway
---
# volcengine_vpc_ipv6_gateway
Provides a resource to manage vpc ipv6 gateway
## Example Usage
```hcl
resource "volcengine_vpc_ipv6_gateway" "foo" {
  vpc_id      = "vpc-12afxho4sxyio17q7y2kkp8ej"
  name        = "tf-test-1"
  description = "test"
}
```
## Argument Reference
The following arguments are supported:
* `vpc_id` - (Required, ForceNew) The ID of the VPC which the Ipv6Gateway belongs to.
* `description` - (Optional) The description of the Ipv6Gateway.
* `name` - (Optional) The name of the Ipv6Gateway.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Ipv6Gateway can be imported using the id, e.g.
```
$ terraform import volcengine_vpc_ipv6_gateway.default ipv6gw-12bcapllb5ukg17q7y2sd3thx
```


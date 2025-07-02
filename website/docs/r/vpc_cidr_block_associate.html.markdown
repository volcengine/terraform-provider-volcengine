---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_cidr_block_associate"
sidebar_current: "docs-volcengine-resource-vpc_cidr_block_associate"
description: |-
  Provides a resource to manage vpc cidr block associate
---
# volcengine_vpc_cidr_block_associate
Provides a resource to manage vpc cidr block associate
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "192.168.0.0/20"
  project_name = "default"
}

resource "volcengine_vpc_cidr_block_associate" "foo" {
  vpc_id               = volcengine_vpc.foo.id
  secondary_cidr_block = "192.168.16.0/20"
}
```
## Argument Reference
The following arguments are supported:
* `secondary_cidr_block` - (Required, ForceNew) The secondary cidr block of the VPC.
* `vpc_id` - (Required, ForceNew) The id of the VPC.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The VpcCidrBlockAssociate is not support import.


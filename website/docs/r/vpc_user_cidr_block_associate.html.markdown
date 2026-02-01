---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_user_cidr_block_associate"
sidebar_current: "docs-volcengine-resource-vpc_user_cidr_block_associate"
description: |-
  Provides a resource to manage vpc user cidr block associate
---
# volcengine_vpc_user_cidr_block_associate
Provides a resource to manage vpc user cidr block associate
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "10.200.0.0/16"
  project_name = "default"
}

resource "volcengine_vpc_user_cidr_block_associate" "foo1" {
  vpc_id          = volcengine_vpc.foo.id
  user_cidr_block = "10.201.0.0/16"
}

resource "volcengine_vpc_user_cidr_block_associate" "foo2" {
  vpc_id          = volcengine_vpc.foo.id
  user_cidr_block = "10.202.0.0/16"
}

resource "volcengine_vpc_user_cidr_block_associate" "foo3" {
  vpc_id          = volcengine_vpc.foo.id
  user_cidr_block = "10.203.0.0/16"
}
```
## Argument Reference
The following arguments are supported:
* `user_cidr_block` - (Required, ForceNew) The user cidr block of the VPC.
* `vpc_id` - (Required, ForceNew) The id of the VPC.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The VpcCidrBlockAssociate is not support import.


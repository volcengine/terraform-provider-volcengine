---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_vpc"
sidebar_current: "docs-volcengine-resource-veenedge_vpc"
description: |-
  Provides a resource to manage veenedge vpc
---
# volcengine_veenedge_vpc
Provides a resource to manage veenedge vpc
## Example Usage
```hcl
resource "volcengine_veenedge_vpc" "foo" {
  vpc_name     = "tf-test-2"
  cluster_name = "b****t02"
}
```
## Argument Reference
The following arguments are supported:
* `cluster_name` - (Required, ForceNew) The name of the cluster.
* `desc` - (Required) The description of the VPC.
* `vpc_name` - (Required, ForceNew) The name of the VPC.
* `cidr` - (Optional, ForceNew) The cidr info.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VPC can be imported using the id, e.g.
```
$ terraform import volcengine_veenedge_vpc.default vpc-mizl7m1k
```

If you need to create a VPC, you need to apply for permission from the administrator in advance.
You can only delete the vpc from web consul


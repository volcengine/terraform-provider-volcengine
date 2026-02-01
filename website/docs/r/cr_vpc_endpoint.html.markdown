---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_vpc_endpoint"
sidebar_current: "docs-volcengine-resource-cr_vpc_endpoint"
description: |-
  Provides a resource to manage cr vpc endpoint
---
# volcengine_cr_vpc_endpoint
Provides a resource to manage cr vpc endpoint
## Example Usage
```hcl
resource "volcengine_cr_vpc_endpoint" "foo" {
  registry = "enterprise-1"
  vpcs {
    vpc_id     = "vpc-3resbfzl3xgjk5zsk2iuq3vhk"
    account_id = 000000
  }
  vpcs {
    vpc_id    = "vpc-3red9li8dd8g05zsk2iadytvy"
    subnet_id = "subnet-2d62do4697i8058ozfdszxl30"
  }

}
```
## Argument Reference
The following arguments are supported:
* `registry` - (Required, ForceNew) The Cr Registry name.
* `vpcs` - (Required) List of vpc meta. When apply is executed for the first time, the vpcs in the tf file will be added to the existing vpcs, and subsequent apply will overwrite the existing vpcs with the vpcs in the tf file.

The `vpcs` object supports the following:

* `account_id` - (Optional) The id of the account. When you need to expose the Enterprise Edition instance to a VPC under another primary account, you need to specify the ID of the primary account to which the VPC belongs.
* `subnet_id` - (Optional) The id of the subnet. If not specified, the subnet with the most remaining IPs under the VPC will be automatically selected.
* `vpc_id` - (Optional) The id of the vpc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CR Vpc endpoint can be imported using the crVpcEndpoint:registry, e.g.
```
$ terraform import volcengine_cr_vpc_endpoint.default crVpcEndpoint:cr-basic
```


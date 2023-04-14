---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc"
sidebar_current: "docs-volcengine-resource-vpc"
description: |-
  Provides a resource to manage vpc
---
# volcengine_vpc
Provides a resource to manage vpc
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name     = "tf-project-1"
  cidr_block   = "172.16.0.0/16"
  dns_servers  = ["8.8.8.8", "114.114.114.114"]
  project_name = "AS_test"
}
```
## Argument Reference
The following arguments are supported:
* `cidr_block` - (Required, ForceNew) A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
* `description` - (Optional) The description of the VPC.
* `dns_servers` - (Optional) The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
* `project_name` - (Optional) The ProjectName of the VPC.
* `tags` - (Optional) Tags.
* `vpc_name` - (Optional) The name of the VPC.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account ID of VPC.
* `associate_cens` - The associate cen list of VPC.
    * `cen_id` - The ID of CEN.
    * `cen_owner_id` - The owner ID of CEN.
    * `cen_status` - The status of CEN.
* `auxiliary_cidr_blocks` - The auxiliary cidr block list of VPC.
* `creation_time` - Creation time of VPC.
* `nat_gateway_ids` - The nat gateway ID list of VPC.
* `route_table_ids` - The route table ID list of VPC.
* `security_group_ids` - The security group ID list of VPC.
* `status` - Status of VPC.
* `subnet_ids` - The subnet ID list of VPC.
* `update_time` - The update time of VPC.
* `vpc_id` - The ID of VPC.


## Import
VPC can be imported using the id, e.g.
```
$ terraform import volcengine_vpc.default vpc-mizl7m1kqccg5smt1bdpijuj
```


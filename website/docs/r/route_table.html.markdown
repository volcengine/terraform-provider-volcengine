---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_route_table"
sidebar_current: "docs-volcengine-resource-route_table"
description: |-
  Provides a resource to manage route table
---
# volcengine_route_table
Provides a resource to manage route table
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
  project_name = "default"
}

resource "volcengine_route_table" "foo" {
  vpc_id           = volcengine_vpc.foo.id
  route_table_name = "acc-test-route-table"
  description      = "tf-test"
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `vpc_id` - (Required, ForceNew) The id of the VPC.
* `description` - (Optional) The description of the route table.
* `project_name` - (Optional) The ProjectName of the route table.
* `route_table_name` - (Optional) The name of the route table.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Route table can be imported using the id, e.g.
```
$ terraform import volcengine_route_table.default vtb-274e0syt9av407fap8tle16kb
```


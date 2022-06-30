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
resource "volcengine_route_table" "foo" {
  vpc_id           = "vpc-2744hsjr475s07fap8t4jbl1c"
  route_table_name = "tf-test1"
  description      = "tf-test1"
}
```
## Argument Reference
The following arguments are supported:
* `vpc_id` - (Required, ForceNew) The id of the VPC.
* `description` - (Optional) The description of the route table.
* `route_table_name` - (Optional) The name of the route table.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Route table can be imported using the id, e.g.
```
$ terraform import volcengine_route_table.default vtb-274e0syt9av407fap8tle16kb
```


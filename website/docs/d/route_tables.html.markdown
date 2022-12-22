---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_route_tables"
sidebar_current: "docs-volcengine-datasource-route_tables"
description: |-
  Use this data source to query detailed information of route tables
---
# volcengine_route_tables
Use this data source to query detailed information of route tables
## Example Usage
```hcl
data "volcengine_route_tables" "default" {
  ids              = ["vtb-274e19skkuhog7fap8u4i8ird", "vtb-2744hslq5b7r47fap8tjomgnj"]
  route_table_name = "vpc-fast"
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of route table ids.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of the route table.
* `route_table_name` - (Optional) A name of route table.
* `vpc_id` - (Optional) An id of VPC.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `route_tables` - The collection of route tables.
    * `account_id` - The account id of the route table creator.
    * `creation_time` - The create time of the route table.
    * `description` - The description of the route table.
    * `id` - The id of the route table.
    * `project_name` - The ProjectName of the route table.
    * `route_table_id` - The id of the route table.
    * `route_table_name` - The name of the route table.
    * `route_table_type` - The type of the route table.
    * `subnet_ids` - The list of the subnet ids to which the entry table associates.
    * `update_time` - The last update time of the route table.
    * `vpc_id` - The id of the virtual private cloud (VPC) to which the route entry belongs.
    * `vpc_name` - The name of the virtual private cloud (VPC) to which the route entry belongs.
* `total_count` - The total count of route table query.



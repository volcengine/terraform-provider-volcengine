---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_subnets"
sidebar_current: "docs-volcengine-datasource-subnets"
description: |-
  Use this data source to query detailed information of subnets
---
# volcengine_subnets
Use this data source to query detailed information of subnets
## Example Usage
```hcl
data "volcengine_subnets" "default" {
  ids = ["subnet-274zsa5kfmj287fap8soo5e19"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Subnet IDs.
* `name_regex` - (Optional) A Name Regex of Subnet.
* `output_file` - (Optional) File name where to save data source results.
* `route_table_id` - (Optional) The ID of route table which subnet associated with.
* `subnet_name` - (Optional) The subnet name to query.
* `vpc_id` - (Optional) The ID of VPC which subnet belongs to.
* `zone_id` - (Optional) The ID of zone which subnet belongs to.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `subnets` - The collection of Subnet query.
    * `account_id` - The account ID which the subnet belongs to.
    * `available_ip_address_count` - The count of available ip address.
    * `cidr_block` - The cidr block of Subnet.
    * `creation_time` - Creation time of Subnet.
    * `description` - The description of Subnet.
    * `id` - The ID of Subnet.
    * `network_acl_id` - The ID of network acl which this subnet associate with.
    * `route_table_id` - The ID of route table.
    * `route_table_type` - The type of route table.
    * `route_table` - The route table information.
        * `route_table_id` - The route table ID.
        * `route_table_type` - The route table type.
    * `status` - The Status of Subnet.
    * `subnet_name` - The Name of Subnet.
    * `total_ipv4_count` - The Count of ipv4.
    * `update_time` - Update time of Subnet.
    * `vpc_id` - The Vpc ID of Subnet.
    * `zone_id` - The ID of Zone.
* `total_count` - The total count of Subnet query.



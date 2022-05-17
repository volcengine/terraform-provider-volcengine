---
subcategory: "VPC"
layout: "vestack"
page_title: "Vestack: vestack_security_groups"
sidebar_current: "docs-vestack-datasource-security_groups"
description: |-
  Use this data source to query detailed information of security groups
---
# vestack_security_groups
Use this data source to query detailed information of security groups
## Example Usage
```hcl
data "vestack_security_groups" "default" {
  ids = ["sg-273ycgql3ig3k7fap8t3dyvqx"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of SecurityGroup IDs.
* `name_regex` - (Optional) A Name Regex of SecurityGroup.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `security_groups` - The collection of SecurityGroup query.
  * `creation_time` - The creation time of SecurityGroup.
  * `description` - The description of SecurityGroup.
  * `id` - The ID of SecurityGroup.
  * `security_group_id` - The ID of SecurityGroup.
  * `security_group_name` - The Name of SecurityGroup.
  * `status` - The Status of SecurityGroup.
  * `type` - A Name Regex of SecurityGroup.
  * `vpc_id` - The ID of Vpc.
* `total_count` - The total count of SecurityGroup query.



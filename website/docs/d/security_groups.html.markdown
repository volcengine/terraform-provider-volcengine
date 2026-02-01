---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_security_groups"
sidebar_current: "docs-volcengine-datasource-security_groups"
description: |-
  Use this data source to query detailed information of security groups
---
# volcengine_security_groups
Use this data source to query detailed information of security groups
## Example Usage
```hcl
data "volcengine_security_groups" "default" {
  ids = ["sg-273ycgql3ig3k7fap8t3dyvqx"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of SecurityGroup IDs.
* `name_regex` - (Optional) A Name Regex of SecurityGroup.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of SecurityGroup.
* `security_group_names` - (Optional) The list of security group name to query.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The ID of vpc where security group is located.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `security_groups` - The collection of SecurityGroup query.
    * `creation_time` - The creation time of SecurityGroup.
    * `description` - The description of SecurityGroup.
    * `id` - The ID of SecurityGroup.
    * `project_name` - The ProjectName of SecurityGroup.
    * `security_group_id` - The ID of SecurityGroup.
    * `security_group_name` - The Name of SecurityGroup.
    * `status` - The Status of SecurityGroup.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `type` - The type of SecurityGroup.
    * `vpc_id` - The ID of Vpc.
* `total_count` - The total count of SecurityGroup query.



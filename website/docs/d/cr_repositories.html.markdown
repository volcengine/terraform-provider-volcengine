---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_repositories"
sidebar_current: "docs-volcengine-datasource-cr_repositories"
description: |-
  Use this data source to query detailed information of cr repositories
---
# volcengine_cr_repositories
Use this data source to query detailed information of cr repositories
## Example Usage
```hcl
data "volcengine_cr_repositories" "foo" {
  registry = "tf-1"
  # access_levels = ["Private"]
  # namespaces = ["namespace*"]
  names = ["repo*"]
}
```
## Argument Reference
The following arguments are supported:
* `registry` - (Required) The CR instance name.
* `access_levels` - (Optional) The list of instance access level.
* `names` - (Optional) The list of instance names.
* `namespaces` - (Optional) The list of instance namespace.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `repositories` - The collection of repository query.
    * `access_level` - The access level of repository.
    * `create_time` - The creation time of repository.
    * `description` - The description of repository.
    * `name` - The name of repository.
    * `namespace` - The namespace of repository.
    * `update_time` - The last update time of repository.
* `total_count` - The total count of instance query.



---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_registries"
sidebar_current: "docs-volcengine-datasource-cr_registries"
description: |-
  Use this data source to query detailed information of cr registries
---
# volcengine_cr_registries
Use this data source to query detailed information of cr registries
## Example Usage
```hcl
data "volcengine_cr_registries" "foo" {
  # names=["liaoliuqing-prune-test"]
  # types=["Enterprise"]
  statuses {
    phase     = "Running"
    condition = "Ok"
  }
}
```
## Argument Reference
The following arguments are supported:
* `names` - (Optional) The list of registry names to query.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The list of registry statuses.
* `types` - (Optional) The list of registry types to query.

The `statuses` object supports the following:

* `condition` - (Optional) The condition of registry.
* `phase` - (Optional) The phase of status.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `registries` - The collection of registry query.
    * `charge_type` - The charge type of registry.
    * `create_time` - The creation time of registry.
    * `domains` - The domain of registry.
        * `domain` - The domain of registry.
        * `type` - The domain type of registry.
    * `name` - The name of registry.
    * `status` - The status of registry.
        * `conditions` - The condition of registry.
        * `phase` - The phase status of registry.
    * `type` - The type of registry.
    * `user_status` - The status of user.
    * `username` - The username of cr instance.
* `total_count` - The total count of registry query.



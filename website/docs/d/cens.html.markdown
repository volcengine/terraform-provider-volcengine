---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cens"
sidebar_current: "docs-volcengine-datasource-cens"
description: |-
  Use this data source to query detailed information of cens
---
# volcengine_cens
Use this data source to query detailed information of cens
## Example Usage
```hcl
resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 2
}

data "volcengine_cens" "foo" {
  ids = volcengine_cen.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `cen_names` - (Optional) A list of cen names.
* `ids` - (Optional) A list of cen IDs.
* `name_regex` - (Optional) A Name Regex of cen.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `cens` - The collection of cen query.
    * `account_id` - The account ID of the cen.
    * `cen_bandwidth_package_ids` - A list of bandwidth package IDs of the cen.
    * `cen_id` - The ID of the cen.
    * `cen_name` - The name of the cen.
    * `creation_time` - The create time of the cen.
    * `description` - The description of the cen.
    * `id` - The ID of the cen.
    * `project_name` - The ProjectName of the cen instance.
    * `status` - The status of the cen.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the cen.
* `total_count` - The total count of cen query.



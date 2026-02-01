---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_groups"
sidebar_current: "docs-volcengine-datasource-cloud_identity_groups"
description: |-
  Use this data source to query detailed information of cloud identity groups
---
# volcengine_cloud_identity_groups
Use this data source to query detailed information of cloud identity groups
## Example Usage
```hcl
resource "volcengine_cloud_identity_group" "foo" {
  group_name   = "acc-test-group-${count.index}"
  display_name = "tf-test-group-${count.index}"
  join_type    = "Manual"
  description  = "tf"

  count = 2
}

data "volcengine_cloud_identity_groups" "foo" {
  group_name = "acc-test-group"
  join_type  = "Manual"
}
```
## Argument Reference
The following arguments are supported:
* `display_name` - (Optional) The display name of cloud identity group.
* `group_name` - (Optional) The name of cloud identity group.
* `join_type` - (Optional) The join type of cloud identity group. Valid values: `Auto`, `Manual`.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `groups` - The collection of query.
    * `created_time` - The created time of the cloud identity group.
    * `description` - The description of the cloud identity group.
    * `display_name` - The display name of the cloud identity group.
    * `group_id` - The id of the cloud identity group.
    * `group_name` - The name of the cloud identity group.
    * `id` - The id of the cloud identity group.
    * `join_type` - The email of the cloud identity group.
    * `source` - The source of the cloud identity group.
    * `updated_time` - The updated time of the cloud identity group.
* `total_count` - The total count of query.



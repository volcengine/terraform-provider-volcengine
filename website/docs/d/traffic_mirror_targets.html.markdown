---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_traffic_mirror_targets"
sidebar_current: "docs-volcengine-datasource-traffic_mirror_targets"
description: |-
  Use this data source to query detailed information of traffic mirror targets
---
# volcengine_traffic_mirror_targets
Use this data source to query detailed information of traffic mirror targets
## Example Usage
```hcl
data "volcengine_traffic_mirror_targets" "foo" {
  traffic_mirror_target_ids = ["tmt-rry7yljufsw0v0x58w2****"]
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of traffic mirror target.
* `tags` - (Optional) Tags.
* `traffic_mirror_target_ids` - (Optional) A list of traffic mirror target IDs.
* `traffic_mirror_target_name` - (Optional) The name of traffic mirror target.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `traffic_mirror_targets` - The collection of query.
    * `created_at` - The create time of traffic mirror target.
    * `description` - The description of traffic mirror target.
    * `id` - The ID of traffic mirror target.
    * `instance_id` - The instance id of traffic mirror target.
    * `instance_type` - The instance type of traffic mirror target.
    * `project_name` - The project name of traffic mirror target.
    * `status` - The status of traffic mirror target.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `traffic_mirror_target_id` - The ID of traffic mirror target.
    * `traffic_mirror_target_name` - The name of traffic mirror target.
    * `updated_at` - The update time of traffic mirror target.



---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_traffic_mirror_sessions"
sidebar_current: "docs-volcengine-datasource-traffic_mirror_sessions"
description: |-
  Use this data source to query detailed information of traffic mirror sessions
---
# volcengine_traffic_mirror_sessions
Use this data source to query detailed information of traffic mirror sessions
## Example Usage
```hcl
data "volcengine_traffic_mirror_sessions" "foo" {
  traffic_mirror_session_ids = ["tms-mjpcyvp71r0g5smt1ayf****"]
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `network_interface_id` - (Optional) The ID of network interface.
* `output_file` - (Optional) File name where to save data source results.
* `packet_length` - (Optional) The packet length of traffic mirror session.
* `priority` - (Optional) The priority of traffic mirror session.
* `project_name` - (Optional) The project name of traffic mirror session.
* `tags` - (Optional) Tags.
* `traffic_mirror_filter_id` - (Optional) The ID of traffic mirror filter.
* `traffic_mirror_session_ids` - (Optional) A list of traffic mirror session IDs.
* `traffic_mirror_session_names` - (Optional) A list of traffic mirror session names.
* `traffic_mirror_target_id` - (Optional) The ID of traffic mirror target.
* `virtual_network_id` - (Optional) The ID of virtual network.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `traffic_mirror_sessions` - The collection of query.
    * `business_status` - The business status of traffic mirror session.
    * `created_at` - The create time of traffic mirror session.
    * `description` - The description of traffic mirror session.
    * `id` - The ID of traffic mirror session.
    * `lock_reason` - The lock reason of traffic mirror session.
    * `packet_length` - The packet length of traffic mirror session.
    * `priority` - The priority of traffic mirror session.
    * `project_name` - The project name of traffic mirror session.
    * `status` - The status of traffic mirror session.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `traffic_mirror_filter_id` - The ID of traffic mirror filter.
    * `traffic_mirror_session_id` - The ID of traffic mirror session.
    * `traffic_mirror_session_name` - The name of traffic mirror session.
    * `traffic_mirror_source_ids` - The IDs of traffic mirror source.
    * `traffic_mirror_target_id` - The ID of traffic mirror target.
    * `updated_at` - The update time of traffic mirror session.
    * `virtual_network_id` - The ID of virtual network.



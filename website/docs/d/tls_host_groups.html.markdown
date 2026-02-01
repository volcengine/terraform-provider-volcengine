---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_host_groups"
sidebar_current: "docs-volcengine-datasource-tls_host_groups"
description: |-
  Use this data source to query detailed information of tls host groups
---
# volcengine_tls_host_groups
Use this data source to query detailed information of tls host groups
## Example Usage
```hcl
data "volcengine_tls_host_groups" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `auto_update` - (Optional) Whether enable auto update.
* `hidden` - (Optional) Whether to hide host groups in exclusive resources.
* `host_group_id` - (Optional) The id of host group.
* `host_group_name` - (Optional) The name of host group.
* `host_identifier` - (Optional) The identifier of host.
* `iam_project_name` - (Optional) The project name of iam.
* `output_file` - (Optional) File name where to save data source results.
* `service_logging` - (Optional) Whether enable service logging.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `infos` - The collection of query.
    * `host_group_info` - The info of host group.
        * `auto_update` - Whether enable auto update.
        * `create_time` - The create time of host group.
        * `host_count` - The host count of host group.
        * `host_group_id` - The id of host group.
        * `host_group_name` - The name of host group.
        * `host_group_type` - The type of host group.
        * `host_identifier` - The identifier of host.
        * `iam_project_name` - The project name of iam.
        * `modify_time` - The modify time of host group.
        * `rule_count` - The rule count of host.
        * `service_logging` - Whether enable service logging.
        * `update_end_time` - The update end time of log collector.
        * `update_start_time` - The update start time of log collector.
* `total_count` - The total count of query.



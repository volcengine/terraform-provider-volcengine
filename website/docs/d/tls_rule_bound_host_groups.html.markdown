---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_rule_bound_host_groups"
sidebar_current: "docs-volcengine-datasource-tls_rule_bound_host_groups"
description: |-
  Use this data source to query detailed information of tls rule bound host groups
---
# volcengine_tls_rule_bound_host_groups
Use this data source to query detailed information of tls rule bound host groups
## Example Usage
```hcl
data "volcengine_tls_rule_bound_host_groups" "default" {
  rule_id = "048dc010-6bb1-4189-858a-281d654d6686"
}
```
## Argument Reference
The following arguments are supported:
* `rule_id` - (Required) The ID of the rule.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `host_groups` - The collection of Host Group query.
    * `auto_update` - Whether to enable auto update.
    * `create_time` - The creation time of the host group.
    * `host_group_id` - The ID of the host group.
    * `host_group_name` - The name of the host group.
    * `host_group_type` - The type of the host group.
    * `host_identifier` - The identifier of the host.
    * `iam_project_name` - The name of the iam project.
    * `modify_time` - The modification time of the host group.
    * `service_logging` - Whether to enable service logging.
    * `update_end_time` - The end time of auto update.
    * `update_start_time` - The start time of auto update.
* `total_count` - The total count of query.



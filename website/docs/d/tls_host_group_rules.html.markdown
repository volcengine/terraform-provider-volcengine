---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_host_group_rules"
sidebar_current: "docs-volcengine-datasource-tls_host_group_rules"
description: |-
  Use this data source to query detailed information of tls host group rules
---
# volcengine_tls_host_group_rules
Use this data source to query detailed information of tls host group rules
## Example Usage
```hcl
data "volcengine_tls_host_group_rules" "default" {
  host_group_id = "4af86d32-cb9c-4eac-adb1-75f2567789be"
}
```
## Argument Reference
The following arguments are supported:
* `host_group_id` - (Required) The id of host group.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rule_infos` - The collection of rule info.
    * `container_rule` - Container collection rules.
    * `create_time` - The create time of rule.
    * `exclude_paths` - Collect the blacklist list.
    * `extract_rule` - The extract rule.
    * `input_type` - The type of input.
    * `log_sample` - The sample of the log.
    * `log_type` - The type of log.
    * `modify_time` - The modify time of rule.
    * `paths` - The paths of rule.
    * `pause` - The pause status of rule.
    * `rule_id` - The id of rule.
    * `rule_name` - The name of rule.
    * `topic_id` - The id of topic.
    * `topic_name` - The name of topic.
    * `user_define_rule` - User-defined collection rules.
* `total_count` - The total count of query.



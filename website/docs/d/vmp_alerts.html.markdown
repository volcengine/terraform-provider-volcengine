---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_alerts"
sidebar_current: "docs-volcengine-datasource-vmp_alerts"
description: |-
  Use this data source to query detailed information of vmp alerts
---
# volcengine_vmp_alerts
Use this data source to query detailed information of vmp alerts
## Example Usage
```hcl
data "volcengine_vmp_alerts" "default" {
  ids = ["9a4f84-0868efcb795c2ac4-73cefd4b3263****"]
}
```
## Argument Reference
The following arguments are supported:
* `alerting_rule_ids` - (Optional) A list of alerting rule IDs.
* `current_phase` - (Optional) The status of vmp alert. Valid values: `Pending`, `Active`, `Resolved`, `Disabled`.
* `desc` - (Optional) Whether to use descending sorting.
* `ids` - (Optional) A list of vmp alert IDs.
* `level` - (Optional) The level of vmp alert. Valid values: `P0`, `P1`, `P2`.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `alerts` - The collection of query.
    * `alerting_rule_id` - The id of the vmp alerting rule.
    * `alerting_rule_query` - The alerting query of the vmp alerting rule.
        * `prom_ql` - The prom ql of query.
        * `workspace_id` - The id of the workspace.
    * `alerting_rule_type` - The type of the vmp alerting rule.
    * `current_level` - The current level of the vmp alert.
    * `current_phase` - The status of the vmp alert.
    * `id` - The id of the vmp alert.
    * `initial_alert_timestamp` - The start time of the vmp alert. Format: RFC3339.
    * `last_alert_timestamp` - The last time of the vmp alert. Format: RFC3339.
    * `levels` - The alerting levels of the vmp alert.
        * `comparator` - The comparator of the vmp alerting rule.
        * `for` - The duration of the alerting rule.
        * `level` - The level of the vmp alerting rule.
        * `threshold` - The threshold of the vmp alerting rule.
    * `resolve_alert_timestamp` - The end time of the vmp alert. Format: RFC3339.
    * `resource` - The alerting resource of the vmp alert.
        * `labels` - The labels of alerting resource.
            * `key` - The key of the label.
            * `value` - The value of the label.
* `total_count` - The total count of query.



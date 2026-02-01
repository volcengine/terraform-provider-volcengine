---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_traffic_mirror_filter_rules"
sidebar_current: "docs-volcengine-datasource-traffic_mirror_filter_rules"
description: |-
  Use this data source to query detailed information of traffic mirror filter rules
---
# volcengine_traffic_mirror_filter_rules
Use this data source to query detailed information of traffic mirror filter rules
## Example Usage
```hcl
data "volcengine_traffic_mirror_filter_rules" "foo" {
  traffic_mirror_filter_ids = ["tmf-mivro9v5x24g5smt1bsq****"]
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of traffic mirror filter.
* `tags` - (Optional) Tags.
* `traffic_mirror_filter_ids` - (Optional) A list of traffic mirror filter IDs.
* `traffic_mirror_filter_names` - (Optional) A list of traffic mirror filter names.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `traffic_mirror_filter_rules` - The collection of query.
    * `created_at` - The create time of traffic mirror filter rule.
    * `description` - The description of traffic mirror filter rule.
    * `destination_cidr_block` - The destination cidr block of traffic mirror filter rule.
    * `destination_port_range` - The destination port range of traffic mirror filter rule.
    * `id` - The ID of traffic mirror filter rule.
    * `policy` - The policy of traffic mirror filter rule.
    * `priority` - The priority of traffic mirror filter rule.
    * `protocol` - The protocol of traffic mirror filter rule.
    * `source_cidr_block` - The source cidr block of traffic mirror filter rule.
    * `source_port_range` - The source port range of traffic mirror filter rule.
    * `status` - The status of traffic mirror filter rule.
    * `traffic_direction` - The traffic direction of traffic mirror filter rule.
    * `traffic_mirror_filter_id` - The ID of traffic mirror filter.
    * `traffic_mirror_filter_rule_id` - The ID of traffic mirror filter rule.
    * `updated_at` - The last update time of traffic mirror filter rule.



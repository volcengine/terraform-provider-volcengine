---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clb_rules"
sidebar_current: "docs-volcengine-datasource-clb_rules"
description: |-
  Use this data source to query detailed information of clb rules
---
# volcengine_clb_rules
Use this data source to query detailed information of clb rules
## Example Usage
```hcl
data "volcengine_clb_rules" "default" {
  listener_id = "lsn-273ywvnmiu70g7fap8u2xzg9d"
  ids         = ["rule-273z9jo9v3mrk7fap8sq8v5x7"]
}

output "data" {
  value = data.volcengine_clb_rules.default
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required) The Id of listener.
* `ids` - (Optional) A list of Rule IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of Rule query.
  * `description` - The Description of Rule.
  * `domain` - The Domain of Rule.
  * `id` - The Id of Rule.
  * `rule_id` - The Id of Rule.
  * `server_group_id` - The Id of Server Group.
  * `url` - The Url of Rule.



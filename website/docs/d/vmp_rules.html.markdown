---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_rules"
sidebar_current: "docs-volcengine-datasource-vmp_rules"
description: |-
  Use this data source to query detailed information of vmp rules
---
# volcengine_vmp_rules
Use this data source to query detailed information of vmp rules
## Example Usage
```hcl
data "volcengine_vmp_rules" "default" {
  workspace_id = "baa02ffb-6f22-43c4-841b-ecf90ded****"
  kind         = "Recording"
}
```
## Argument Reference
The following arguments are supported:
* `kind` - (Required) The kind of rule.
* `workspace_id` - (Required) The id of workspace.
* `name` - (Optional) The name of rule.
* `output_file` - (Optional) File name where to save data source results.
* `rule_file_names` - (Optional) The name of rule file.
* `rule_group_names` - (Optional) The name of rule group.
* `status` - (Optional) The status of rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of query.
    * `expr` - The expr of rule.
    * `kind` - The kind of rule.
    * `labels` - The labels of rule.
        * `key` - The key of label.
        * `value` - The value of label.
    * `last_evaluation` - The last evaluation of rule.
    * `name` - The name of rule.
    * `reason` - The reason of rule.
    * `rule_file_name` - The name of rule file.
    * `rule_group_name` - The name of rule group.
    * `status` - The status of rule.
* `total_count` - The total count of query.



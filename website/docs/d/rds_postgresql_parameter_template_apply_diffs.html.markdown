---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_parameter_template_apply_diffs"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_parameter_template_apply_diffs"
description: |-
  Use this data source to query detailed information of rds postgresql parameter template apply diffs
---
# volcengine_rds_postgresql_parameter_template_apply_diffs
Use this data source to query detailed information of rds postgresql parameter template apply diffs
## Example Usage
```hcl
data "volcengine_rds_postgresql_parameter_template_apply_diffs" "diffs" {
  instance_id = "postgres-72715e0d9f58"
  template_id = "postgresql-ef66e3807988595a"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the PostgreSQL instance.
* `template_id` - (Required) The id of the template.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `parameters` - Changes in instance parameters after applying the specified parameter template.
    * `name` - The name of the parameter.
    * `new_value` - The running value defined for this parameter in the parameter template.
    * `old_value` - The current running value of this parameter in the instance.
    * `restart` - Indicates whether a restart is required after the parameter is modified.
* `total_count` - The total count of query.



---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_parameter_template"
sidebar_current: "docs-volcengine-resource-rds_mysql_parameter_template"
description: |-
  Provides a resource to manage rds mysql parameter template
---
# volcengine_rds_mysql_parameter_template
Provides a resource to manage rds mysql parameter template
## Example Usage
```hcl
resource "volcengine_rds_mysql_parameter_template" "foo" {
  template_name         = "test"
  template_type         = "Mysql"
  template_type_version = "MySQL_8_0"
  template_params {
    name          = "auto_increment_increment"
    running_value = "1"
  }
  template_desc = "test"
}
```
## Argument Reference
The following arguments are supported:
* `template_name` - (Required) Parameter template name.
* `template_params` - (Required) Parameters contained in the parameter template.
* `template_type_version` - (Required) Database version of parameter template. Value range:
MySQL_5_7: Default value. MySQL 5.7 version.
MySQL_8_0: MySQL 8.0 version.
* `template_type` - (Required, ForceNew) Database type of parameter template. The default value is Mysql.
* `template_desc` - (Optional) Parameter template description.

The `template_params` object supports the following:

* `name` - (Required) Instance parameter name.
Description: When using CreateParameterTemplate and ModifyParameterTemplate as request parameters, only Name and RunningValue need to be passed in.
* `running_value` - (Required) Parameter running value.
Description: When making request parameters in CreateParameterTemplate and ModifyParameterTemplate, only Name and RunningValue need to be passed in.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsMysqlParameterTemplate can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_parameter_template.default resource_id
```


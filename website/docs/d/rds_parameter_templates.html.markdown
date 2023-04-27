---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_parameter_templates"
sidebar_current: "docs-volcengine-datasource-rds_parameter_templates"
description: |-
  Use this data source to query detailed information of rds parameter templates
---
# volcengine_rds_parameter_templates
(Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds parameter templates
## Example Usage
```hcl
data "volcengine_rds_parameter_templates" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of RDS parameter template.
* `output_file` - (Optional) File name where to save data source results.
* `template_category` - (Optional) Parameter template type, range of values:
DBEngine - Engine parameters.
* `template_source` - (Optional) Template source, value range:
System - System
User - the user.
* `template_type_version` - (Optional) Parameter template database version, value range:
MySQL_Community_5_7 - MySQL 5.7
MySQL_8_0 - MySQL 8.0.
* `template_type` - (Optional) Parameter template database type, range of values:
MySQL - MySQL database.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rds_parameter_templates` - The collection of RDS parameter templates query.
    * `create_time` - Creation time.
    * `id` - The ID of the RDS parameter template.
    * `need_restart` - Whether the template contains parameters that need to be restarted.
    * `parameter_num` - The number of parameters the template contains.
    * `template_desc` - The description of the RDS parameter template.
    * `template_id` - The ID of the RDS parameter template.
    * `template_name` - The name of the RDS parameter template.
    * `template_params` - Parameters contained in the template.
        * `default_value` - Parameter default value.
        * `description` - Parameter description.
        * `name` - Parameter name.
        * `restart` - Whether the modified parameters need to be restarted to take effect.
        * `running_value` - Parameter running value.
        * `value_range` - Parameter value range.
    * `template_type_version` - Parameter template database version, value range:
MySQL_Community_5_7 - MySQL 5.7
MySQL_8_0 - MySQL 8.0.
    * `template_type` - Parameter template database type, range of values:
MySQL - MySQL database.
    * `update_time` - Update time.
* `total_count` - The total count of RDS parameter templates query.



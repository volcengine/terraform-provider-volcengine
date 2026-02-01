---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_parameter_templates"
sidebar_current: "docs-volcengine-datasource-rds_mysql_parameter_templates"
description: |-
  Use this data source to query detailed information of rds mysql parameter templates
---
# volcengine_rds_mysql_parameter_templates
Use this data source to query detailed information of rds mysql parameter templates
## Example Usage
```hcl
data "volcengine_rds_mysql_parameter_templates" "foo" {
  template_category = "DBEngine"
  template_source   = "User"
  #template_type = ""
  #template_type_version = ""
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `template_category` - (Optional) Template category, with a value of DBEngine (database engine parameters).
* `template_source` - (Optional) Parameter template source, value range: System. User.
* `template_type_version` - (Optional) Database version of parameter template. Value range:
MySQL_5_7: Default value. MySQL 5.7 version.
MySQL_8_0: MySQL 8.0 version.
* `template_type` - (Optional) Database type of parameter template. The default value is Mysql.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `templates` - The collection of query.
    * `account_id` - The account ID.
    * `create_time` - Creation time.
    * `need_restart` - Does the template contain parameters that require restart.
    * `parameter_num` - The number of parameters contained in the template.
    * `project_name` - The project to which the template belongs.
    * `template_category` - Template category, with a value of DBEngine (database engine parameter).
    * `template_desc` - Parameter template description.
    * `template_id` - Parameter template ID.
    * `template_name` - Parameter template name.
    * `template_params` - Parameters contained in the template.
        * `default_value` - Parameter default value.
        * `description` - Parameter description.
        * `name` - Instance parameter name.
Description: When using CreateParameterTemplate and ModifyParameterTemplate as request parameters, only Name and RunningValue need to be passed in.
        * `restart` - Is it necessary to restart the instance for the changes to take effect.
        * `running_value` - Parameter running value.
Description: When making requests with CreateParameterTemplate and ModifyParameterTemplate as request parameters, only Name and RunningValue need to be passed in.
        * `value_range` - Value range of parameters.
    * `template_source` - The type of parameter template. Values:
System: System template.
User: User template.
    * `template_type_version` - Parameter template database version, value range:
"MySQL_5_7": MySQL 5.7 version.
"MySQL_8_0": MySQL 8.0 version.
    * `template_type` - The database type of the parameter template. The default value is Mysql.
    * `update_time` - Modification time of the template.
* `total_count` - The total count of query.



---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_parameter_templates"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_parameter_templates"
description: |-
  Use this data source to query detailed information of rds postgresql parameter templates
---
# volcengine_rds_postgresql_parameter_templates
Use this data source to query detailed information of rds postgresql parameter templates
## Example Usage
```hcl
data "volcengine_rds_postgresql_parameter_templates" "templates" {
  template_category     = "DBEngine"
  template_type         = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `template_category` - (Optional) Classification of parameter templates. The current value can only be DBEngine.
* `template_source` - (Optional) The source of the parameter template. The current value can only be User.
* `template_type_version` - (Optional) PostgreSQL compatible versions. The current value can only be PostgreSQL_11/12/13/14/15/16/17.
* `template_type` - (Optional) The type of the parameter template. The current value can only be PostgreSQL.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `template_infos` - Parameter template list.
    * `account_id` - Account ID.
    * `create_time` - Creation time of the parameter template. The format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `need_restart` - Indicates whether the parameter template change requires a restart.
    * `parameter_num` - Number of parameters in the parameter template.
    * `template_category` - Classification of parameter templates. The current value can only be DBEngine.
    * `template_desc` - Description information of the parameter template.
    * `template_id` - Parameter template ID.
    * `template_name` - Parameter template name.
    * `template_params` - Parameter configuration of the parameter template.
        * `checking_code` - The value range of the parameter.
        * `default_value` - Parameter default value. Refers to the default value provided in the default template corresponding to this instance.
        * `description_zh` - The description of the parameter in Chinese.
        * `description` - The description of the parameter in English.
        * `force_restart` - Indicates whether a restart is required after the parameter is modified.
        * `name` - The name of the parameter.
        * `type` - The type of the parameter.
        * `value` - The current value of the parameter.
    * `template_source` - The source of the parameter template. The current value can only be User.
    * `template_type_version` - PostgreSQL compatible versions. The current value can only be PostgreSQL_11/12/13/14/15/16/17.
    * `template_type` - The type of the parameter template. The current value can only be PostgreSQL.
    * `update_time` - Update time of the parameter template. The format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
* `total_count` - The total count of query.



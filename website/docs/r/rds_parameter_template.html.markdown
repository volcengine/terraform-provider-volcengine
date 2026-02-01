---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_parameter_template"
sidebar_current: "docs-volcengine-resource-rds_parameter_template"
description: |-
  Provides a resource to manage rds parameter template
---
# volcengine_rds_parameter_template
(Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds parameter template
## Example Usage
```hcl
resource "volcengine_rds_parameter_template" "foo" {
  template_desc         = "created by terraform"
  template_name         = "tf-template"
  template_type         = "MySQL"
  template_type_version = "MySQL_Community_5_7"
  template_params {
    name          = "auto_increment_increment"
    running_value = "2"
  }
  template_params {
    name          = "slow_query_log"
    running_value = "ON"
  }
  template_params {
    name          = "net_retry_count"
    running_value = "33"
  }
}
```
## Argument Reference
The following arguments are supported:
* `template_name` - (Required) Parameter template name.
* `template_params` - (Required) Template parameters. InstanceParam only needs to pass Name and RunningValue.
* `template_desc` - (Optional) Parameter template description.
* `template_type_version` - (Optional, ForceNew) Parameter template database version, value range:
MySQL_Community_5_7 - MySQL 5.7 (default)
MySQL_8_0 - MySQL 8.0.
* `template_type` - (Optional, ForceNew) Parameter template database type, range of values:
MySQL - MySQL database. (Defaults).

The `template_params` object supports the following:

* `name` - (Optional) Parameter name.
* `running_value` - (Optional) Parameter running value.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RDS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_parameter_template.default mysql-sys-80bb93aa14be22d0
```


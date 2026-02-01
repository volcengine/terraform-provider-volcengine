---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_parameter_template"
sidebar_current: "docs-volcengine-resource-rds_postgresql_parameter_template"
description: |-
  Provides a resource to manage rds postgresql parameter template
---
# volcengine_rds_postgresql_parameter_template
Provides a resource to manage rds postgresql parameter template
## Example Usage
```hcl
resource "volcengine_rds_postgresql_parameter_template" "tpl_base" {
  template_name         = "tf-pg-pt-base"
  template_type         = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
  template_desc         = "base template for clone"

  template_params {
    name  = "auto_explain.log_analyze"
    value = "off"
  }
  template_params {
    name  = "auto_explain.log_buffers"
    value = "on"
  }
}

resource "volcengine_rds_postgresql_parameter_template" "tpl_clone" {
  template_name         = "tf-pg-pt-clone"
  src_template_id       = "postgresql-b62f5687df914b1c"
  template_desc         = "cloned by terraform"
  template_type         = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
}

resource "volcengine_rds_postgresql_parameter_template" "tpl_export" {
  template_name         = "tf-pg-pt-export"
  instance_id           = "postgres-72715e0d9f58"
  template_desc         = "exported from instance"
  template_type         = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
}
```
## Argument Reference
The following arguments are supported:
* `template_name` - (Required) Parameter template name.
* `template_type_version` - (Required) The version of PostgreSQL supported by the parameter template. The current value can be PostgreSQL_11/12/13/14/15/16/17.
* `instance_id` - (Optional, ForceNew) The ID of the instance to export the current parameters as a parameter template. If set, the parameter template will be created based on the current parameters of the instance.
* `src_template_id` - (Optional, ForceNew) The ID of the source parameter template to clone. If set, the parameter template will be cloned from the source template.
* `template_desc` - (Optional) The description of the parameter template. The maximum length is 200 characters.
* `template_params` - (Optional) Parameter configuration of the parameter template.
* `template_type` - (Optional, ForceNew) The type of the parameter template. The current value can only be PostgreSQL.

The `template_params` object supports the following:

* `name` - (Required) The name of the parameter.
* `value` - (Required) The value of the parameter.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `template_id` - Parameter template ID.


## Import
RdsPostgresqlParameterTemplate can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_parameter_template.default resource_id
```


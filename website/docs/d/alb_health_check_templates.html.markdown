---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_health_check_templates"
sidebar_current: "docs-volcengine-datasource-alb_health_check_templates"
description: |-
  Use this data source to query detailed information of alb health check templates
---
# volcengine_alb_health_check_templates
Use this data source to query detailed information of alb health check templates
## Example Usage
```hcl
data "volcengine_alb_health_check_templates" "foo" {
  ids = ["hctpl-1iidd1tobnim874adhf708uwf"]
  tags {
    key   = "key1"
    value = "value2"
  }
}
```
## Argument Reference
The following arguments are supported:
* `health_check_template_name` - (Optional) The name of health check template to query.
* `ids` - (Optional) The list of health check templates to query.
* `name_regex` - (Optional) A Name Regex of health check template.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name to query.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `health_check_templates` - The collection of health check template query.
    * `create_time` - The creation time of the health check template.
    * `description` - The description of health check template.
    * `health_check_domain` - The domain name to health check.
    * `health_check_http_code` - The normal HTTP status code for health check, the default is http_2xx, http_3xx, separated by commas.
    * `health_check_http_version` - The HTTP version of health check.
    * `health_check_interval` - The interval for performing health checks, the default value is 2, and the value is 1-300.
    * `health_check_method` - The health check method, support `GET` and `HEAD`.
    * `health_check_port` - The port for health check. 0 means use backend server port for health check, 1-65535 means use the specified port.
    * `health_check_protocol` - The protocol of health check, support HTTP and TCP.
    * `health_check_template_id` - The ID of health check template.
    * `health_check_template_name` - The name of health check template.
    * `health_check_timeout` - The timeout of health check response,the default value is 2, and the value is 1-60.
    * `health_check_uri` - The uri to health check,default is `/`.
    * `healthy_threshold` - The healthy threshold of the health check, the default is 3, the value is 2-10.
    * `id` - The id of the health check template.
    * `project_name` - The project name to which the health check template belongs.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `unhealthy_threshold` - The unhealthy threshold of the health check, the default is 3, the value is 2-10.
    * `update_time` - The last update time of the health check template.
* `total_count` - The total count of health check template query.



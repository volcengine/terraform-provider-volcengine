---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_health_check_template"
sidebar_current: "docs-volcengine-resource-alb_health_check_template"
description: |-
  Provides a resource to manage alb health check template
---
# volcengine_alb_health_check_template
Provides a resource to manage alb health check template
## Example Usage
```hcl
resource "volcengine_alb_health_check_template" "foo" {
  health_check_template_name = "acc-test-template-1"
  description                = "acc-test3"
  health_check_interval      = 8
  health_check_timeout       = 11
  healthy_threshold          = 2
  unhealthy_threshold        = 3
  health_check_method        = "HEAD"
  health_check_domain        = "test.com"
  health_check_uri           = "/"
  health_check_http_code     = "http_2xx"
  health_check_protocol      = "HTTP"
  health_check_http_version  = "HTTP1.1"
}
```
## Argument Reference
The following arguments are supported:
* `health_check_template_name` - (Required) The health check template name.
* `description` - (Optional) The description of health check template.
* `health_check_domain` - (Optional) The domain name to health check.
* `health_check_http_code` - (Optional) The normal HTTP status code for health check, the default is http_2xx, http_3xx, separated by commas.
* `health_check_http_version` - (Optional) The HTTP version of health check.
* `health_check_interval` - (Optional) The interval for performing health checks, the default value is 2, and the value is 1-300.
* `health_check_method` - (Optional) The health check method,default is `GET`, support `GET` and `HEAD`.
* `health_check_protocol` - (Optional) THe protocol of health check,only support HTTP.
* `health_check_timeout` - (Optional) The timeout of health check response,the default value is 2, and the value is 1-60.
* `health_check_uri` - (Optional) The uri to health check,default is `/`.
* `healthy_threshold` - (Optional) The healthy threshold of the health check, the default is 3, the value is 2-10.
* `unhealthy_threshold` - (Optional) The unhealthy threshold of the health check, the default is 3, the value is 2-10.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
AlbHealthCheckTemplate can be imported using the id, e.g.
```
$ terraform import volcengine_alb_health_check_template.default hctpl-123*****432
```


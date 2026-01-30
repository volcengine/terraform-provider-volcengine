---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_health_check_log_topic"
sidebar_current: "docs-volcengine-resource-health_check_log_topic"
description: |-
  Provides a resource to manage health check log topic
---
# volcengine_health_check_log_topic
Provides a resource to manage health check log topic
## Example Usage
```hcl
resource "volcengine_health_check_log_topic" "example" {
  log_topic_id     = "82fddbd8-4140-4527-****-b89d2aae4a61"
  load_balancer_id = "clb-mim12q0soe805smt1be*****"
}
resource "volcengine_health_check_log_topic" "example1" {
  log_topic_id     = "82fddbd8-4140-4527-****-b89d2aae4a61"
  load_balancer_id = "clb-13g5i2cbg6nsw3n6nu5r*****"
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The ID of the CLB instance.
* `log_topic_id` - (Required, ForceNew) The ID of the log topic.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
HealthCheckLogTopic can be imported using the id, e.g.
```
$ terraform import volcengine_health_check_log_topic.default log_topic_id:load_balancer_id
```


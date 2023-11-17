---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_server_group"
sidebar_current: "docs-volcengine-resource-alb_server_group"
description: |-
  Provides a resource to manage alb server group
---
# volcengine_alb_server_group
Provides a resource to manage alb server group
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id            = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group"
  description       = "acc-test"
  server_group_type = "instance"
  scheduler         = "wlc"
  project_name      = "default"
  health_check {
    enabled  = "on"
    interval = 3
    timeout  = 3
    method   = "GET"
  }
  sticky_session_config {
    sticky_session_enabled = "on"
    sticky_session_type    = "insert"
    cookie_timeout         = "1100"
  }
}
```
## Argument Reference
The following arguments are supported:
* `vpc_id` - (Required, ForceNew) The vpc id of the Alb server group.
* `description` - (Optional) The description of the Alb server group.
* `health_check` - (Optional) The health check config of the Alb server group. The enable status of health check function defaults to `on`.
* `project_name` - (Optional) The project name of the Alb server group.
* `scheduler` - (Optional) The scheduling algorithm of the Alb server group. Valid values: `wrr`, `wlc`, `sh`.
* `server_group_name` - (Optional) The name of the Alb server group.
* `server_group_type` - (Optional, ForceNew) The type of the Alb server group. Valid values: `instance`, `ip`. Default is `instance`.
* `sticky_session_config` - (Optional) The sticky session config of the Alb server group. The enable status of sticky session function defaults to `off`.

The `health_check` object supports the following:

* `domain` - (Optional) The domain of health check.
* `enabled` - (Optional) The enable status of health check function. Valid values: `on`, `off`. Default is `on`.
* `healthy_threshold` - (Optional) The healthy threshold of health check. Valid value range in 2~10. Default is 3.
* `http_code` - (Optional) The normal http status code of health check, the value can be `http_2xx` or `http_3xx` or `http_4xx` or `http_5xx`.
* `http_version` - (Optional) The http version of health check. Valid values: `HTTP1.0`, `HTTP1.1`. Default is `HTTP1.0`.
* `interval` - (Optional) The interval executing health check. Unit: second. Valid value range in 1~300. Default is 2.
* `method` - (Optional) The method of health check. Valid values: `GET` or `HEAD`. Default is `HEAD`.
* `timeout` - (Optional) The response timeout of health check. Unit: second. Valid value range in 1~60. Default is 2.
* `unhealthy_threshold` - (Optional) The unhealthy threshold of health check. Valid value range in 2~10. Default is 3.
* `uri` - (Optional) The uri of health check.

The `sticky_session_config` object supports the following:

* `cookie_timeout` - (Optional) The cookie timeout of the sticky session. Unit: second. Valid value range in 1~86400. Default is 1000. This field is required when the value of the `sticky_session_type` is `insert`.
* `cookie` - (Optional) The cookie name of the sticky session. This field is required when the value of the `sticky_session_type` is `server`.
* `sticky_session_enabled` - (Optional) The enable status of sticky session. Valid values: `on`, `off`. Default is `off`.
* `sticky_session_type` - (Optional) The cookie handle type of the sticky session. Valid values: `insert`, `server`. Default is `insert`. This field is required when the value of the `sticky_session_enabled` is `on`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the Alb server group.
* `listeners` - The listener information of the Alb server group.
* `server_count` - The server count of the Alb server group.
* `status` - The status of the Alb server group.
* `update_time` - The update time of the Alb server group.


## Import
AlbServerGroup can be imported using the id, e.g.
```
$ terraform import volcengine_alb_server_group.default resource_id
```


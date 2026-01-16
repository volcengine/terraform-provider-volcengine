---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_server_groups"
sidebar_current: "docs-volcengine-datasource-alb_server_groups"
description: |-
  Use this data source to query detailed information of alb server groups
---
# volcengine_alb_server_groups
Use this data source to query detailed information of alb server groups
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id            = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group-${count.index}"
  description       = "acc-test"
  server_group_type = "instance"
  scheduler         = "sh"
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
  count = 3
}

data "volcengine_alb_server_groups" "foo" {
  ids = volcengine_alb_server_group.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Alb server group IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of Alb server group.
* `server_group_names` - (Optional) A list of Alb server group name.
* `server_group_type` - (Optional) The type of Alb server group. Valid values: `instance`, `ip`.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The vpc id of Alb server group.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `server_groups` - The collection of query.
    * `create_time` - The create time of the Alb server group.
    * `cross_zone_enabled` - Whether to enable cross-zone load balancing for the server group.
    * `description` - The description of the Alb server group.
    * `health_check` - The health check config of the Alb server group.
        * `domain` - The domain of health check.
        * `enabled` - The enable status of health check function.
        * `healthy_threshold` - The healthy threshold of health check.
        * `http_code` - The normal http status code of health check.
        * `http_version` - The http version of health check.
        * `interval` - The interval executing health check.
        * `method` - The method of health check.
        * `port` - The port of health check.
        * `protocol` - The protocol of health check.
        * `unhealthy_threshold` - The unhealthy threshold of health check.
        * `uri` - The uri of health check.
    * `id` - The ID of the Alb server group.
    * `ip_address_type` - The ip address type of the server group.
    * `listeners` - The listener information of the Alb server group.
    * `project_name` - The project name of the Alb server group.
    * `protocol` - The backend protocol of the Alb server group.
    * `scheduler` - The scheduler algorithm of the Alb server group.
    * `server_count` - The server count of the Alb server group.
    * `server_group_id` - The ID of the Alb server group.
    * `server_group_name` - The name of the Alb server group.
    * `server_group_type` - The type of the Alb server group.
    * `servers` - The server information of the Alb server group.
        * `description` - The description of the server group server.
        * `instance_id` - The id of the ecs instance or the network interface.
        * `ip` - The private ip of the server group server.
        * `port` - The port receiving request of the server group server.
        * `remote_enabled` - Whether to enable the remote IP function.
        * `server_id` - The id of the server group server.
        * `type` - The type of the server group server.
        * `weight` - The weight of the server group server.
    * `status` - The status of the Alb server group.
    * `sticky_session_config` - The sticky session config of the Alb server group.
        * `cookie_timeout` - The cookie timeout of the sticky session.
        * `cookie` - The cookie name of the sticky session.
        * `sticky_session_enabled` - The enable status of sticky session.
        * `sticky_session_type` - The cookie handle type of the sticky session.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the Alb server group.
    * `vpc_id` - The vpc id of the Alb server group.
* `total_count` - The total count of query.



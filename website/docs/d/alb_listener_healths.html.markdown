---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_listener_healths"
sidebar_current: "docs-volcengine-datasource-alb_listener_healths"
description: |-
  Use this data source to query detailed information of alb listener healths
---
# volcengine_alb_listener_healths
Use this data source to query detailed information of alb listener healths
## Example Usage
```hcl
data "volcengine_alb_listener_healths" "example" {
  listener_ids    = ["lsn-xoetdjk3dzwg54ov5ewpam7c", "lsn-bdcxfof3fy808dv40ofappua"]
  only_un_healthy = true
  project_name    = "default"
}
```
## Argument Reference
The following arguments are supported:
* `listener_ids` - (Required) A list of Listener IDs.
* `only_un_healthy` - (Optional) Whether to return only backend servers with abnormal health check status.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the listener.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `listeners` - The collection of listener health query.
    * `backend_servers` - The list of backend server health details.
        * `instance_id` - The ID of the ECS instance or ENI.
        * `ip` - The IP address of the backend server.
        * `port` - The port of the backend server.
        * `rule_number` - The number of forwarding rules associated with the backend server.
        * `server_group_id` - The ID of the backend server group.
        * `server_group_name` - The name of the backend server group.
        * `server_id` - The ID of the backend server.
        * `status` - The health status of the backend server. Value: Up, Down.
        * `type` - The type of backend server. Value: ecs, eni.
    * `listener_id` - The ID of the listener.
    * `status` - The status of the listener. Value: Active, Error, NoTarget, Disabled.
    * `total_backend_server_count` - The total count of backend servers under the listener.
    * `un_healthy_count` - The count of backend servers with abnormal health check status.
* `total_count` - The total count of query.



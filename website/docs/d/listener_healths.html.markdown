---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_listener_healths"
sidebar_current: "docs-volcengine-datasource-listener_healths"
description: |-
  Use this data source to query detailed information of listener healths
---
# volcengine_listener_healths
Use this data source to query detailed information of listener healths
## Example Usage
```hcl
data "volcengine_listener_healths" "foo" {
  listener_id = "lsn-mjkyvug6pwxs5smt1b9*****"
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required) The ID of the listener.
* `only_un_healthy` - (Optional) Whether to return only unhealthy backend servers. Valid values: `true`, `false`.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `health_info` - The health info of backend servers.
    * `listener_status` - The health check status of the listener. Valid values: `Active`, `Error`, `Disabled`, `NoTarget`.
    * `results` - The backend server health status results.
        * `instance_id` - The ECS instance or ENI ID.
        * `ip` - The IP address of the backend server.
        * `port` - The port of the backend server.
        * `rule_number` - The number of forwarding rules associated with the backend server. TCP/UDP listeners return 0.
        * `server_group_id` - The server group ID that the backend server belongs to.
        * `server_id` - The backend server ID.
        * `server_type` - The backend server type. Valid values: `ecs`, `eni`.
        * `status` - The health status of the backend server. Valid values: `Up`, `Down`.
        * `updated_at` - The last update time of the backend server.
    * `un_healthy_count` - The count of unhealthy backend servers.
* `total_count` - The total count of backend servers.



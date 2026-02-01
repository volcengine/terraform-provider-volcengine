---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_server_group_servers"
sidebar_current: "docs-volcengine-datasource-alb_server_group_servers"
description: |-
  Use this data source to query detailed information of alb server group servers
---
# volcengine_alb_server_group_servers
Use this data source to query detailed information of alb server group servers
## Example Usage
```hcl
data "volcengine_alb_server_group_servers" "foo" {
  server_group_id = "rsp-1g7317vrcx3pc2zbhq4c3i6a2"
  instance_ids    = ["i-yecutyh340qc6io*****"]
}
```
## Argument Reference
The following arguments are supported:
* `server_group_id` - (Required) The ID of the ServerGroup.
* `instance_ids` - (Optional) A list of instance IDs. When the backend server is ECS, the parameter value is the ID of the ECS. When the backend server is a secondary network interface card, the parameter value is the ID of the secondary network interface card.
* `ips` - (Optional) A list of private IP addresses.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `servers` - The server list of ServerGroup.
    * `description` - The description of the instance.
    * `id` - The server id of instance in ServerGroup.
    * `instance_id` - The ID of ecs instance or the network card bound to ecs instance.
    * `ip` - The private ip of the instance.
    * `port` - The port receiving request.
    * `remote_enabled` - Whether to enable remote IP function. Optional choice contains `on`, `off`.
    * `server_id` - The server id of instance in ServerGroup.
    * `type` - The type of instance. Optional choice contains `ecs`, `eni`, `ip`.
    * `weight` - The weight of the instance.
* `total_count` - The total count of ServerGroupServer query.



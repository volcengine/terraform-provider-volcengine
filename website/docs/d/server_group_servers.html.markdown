---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_server_group_servers"
sidebar_current: "docs-volcengine-datasource-server_group_servers"
description: |-
  Use this data source to query detailed information of server group servers
---
# volcengine_server_group_servers
Use this data source to query detailed information of server group servers
## Example Usage
```hcl
data "volcengine_server_group_servers" "default" {
  ids             = ["rs-273z9pv8mtfcw7fap8sp6ie8k"]
  server_group_id = "rsp-273z9pt9lpdds7fap8sqdvfrf"
}
```
## Argument Reference
The following arguments are supported:
* `server_group_id` - (Required) The ID of the ServerGroup.
* `ids` - (Optional) The list of ServerGroupServer IDs.
* `name_regex` - (Optional) A Name Regex of ServerGroupServer.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `servers` - The server list of ServerGroup.
    * `description` - The description of the instance.
    * `id` - The server id of instance in ServerGroup.
    * `instance_id` - The ID of ecs instance or the network card bound to ecs instance.
    * `ip` - The private ip of the instance.
    * `port` - The port receiving request.
    * `server_id` - The server id of instance in ServerGroup.
    * `type` - The type of instance. Optional choice contains `ecs`, `eni`.
    * `weight` - The weight of the instance.
* `total_count` - The total count of ServerGroupServer query.



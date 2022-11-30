---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_server_groups"
sidebar_current: "docs-volcengine-datasource-server_groups"
description: |-
  Use this data source to query detailed information of server groups
---
# volcengine_server_groups
Use this data source to query detailed information of server groups
## Example Usage
```hcl
data "volcengine_server_groups" "default" {
  ids = ["rsp-273yv0kir1vk07fap8tt9jtwg", "rsp-273yxuqfova4g7fap8tyemn6t", "rsp-273z9pt9lpdds7fap8sqdvfrf"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of ServerGroup IDs.
* `load_balancer_id` - (Optional) The id of the Clb.
* `name_regex` - (Optional) A Name Regex of ServerGroup.
* `output_file` - (Optional) File name where to save data source results.
* `server_group_name` - (Optional) The name of the ServerGroup.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `groups` - The collection of ServerGroup query.
    * `create_time` - The create time of the ServerGroup.
    * `description` - The description of the ServerGroup.
    * `id` - The ID of the ServerGroup.
    * `server_group_id` - The ID of the ServerGroup.
    * `server_group_name` - The name of the ServerGroup.
    * `update_time` - The update time of the ServerGroup.
* `total_count` - The total count of ServerGroup query.



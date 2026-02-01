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
data "volcengine_zones" "foo" {
}
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_clb" "foo" {
  type               = "public"
  subnet_id          = volcengine_subnet.foo.id
  load_balancer_spec = "small_1"
  description        = "acc0Demo"
  load_balancer_name = "acc-test-create"
  eip_billing_config {
    isp              = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth        = 1
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id  = volcengine_clb.foo.id
  server_group_name = "acc-test-create"
  description       = "hello demo11"
}

data "volcengine_server_groups" "foo" {
  ids = [volcengine_server_group.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of ServerGroup IDs.
* `load_balancer_id` - (Optional) The id of the Clb.
* `name_regex` - (Optional) A Name Regex of ServerGroup.
* `output_file` - (Optional) File name where to save data source results.
* `server_group_name` - (Optional) The name of the ServerGroup.
* `tags` - (Optional) Tags.
* `type` - (Optional) The type of ServerGroup. Valid values: `instance`, `ip`.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `groups` - The collection of ServerGroup query.
    * `address_ip_version` - The address IP version of the ServerGroup.
    * `any_port_enabled` - Whether full port forwarding is enabled.
    * `create_time` - The create time of the ServerGroup.
    * `description` - The description of the ServerGroup.
    * `id` - The ID of the ServerGroup.
    * `listeners` - The listeners of the ServerGroup.
    * `load_balancer_id` - The ID of the LoadBalancer.
    * `server_group_id` - The ID of the ServerGroup.
    * `server_group_name` - The name of the ServerGroup.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `type` - The type of the ServerGroup.
    * `update_time` - The update time of the ServerGroup.
* `total_count` - The total count of ServerGroup query.



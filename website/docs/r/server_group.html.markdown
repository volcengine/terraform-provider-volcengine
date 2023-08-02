---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_server_group"
sidebar_current: "docs-volcengine-resource-server_group"
description: |-
  Provides a resource to manage server group
---
# volcengine_server_group
Provides a resource to manage server group
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
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The ID of the Clb.
* `description` - (Optional) The description of ServerGroup.
* `server_group_id` - (Optional) The ID of the ServerGroup.
* `server_group_name` - (Optional) The name of the ServerGroup.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ServerGroup can be imported using the id, e.g.
```
$ terraform import volcengine_server_group.default rsp-273yv0kir1vk07fap8tt9jtwg
```


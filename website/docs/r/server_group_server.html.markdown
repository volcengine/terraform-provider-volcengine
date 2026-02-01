---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_server_group_server"
sidebar_current: "docs-volcengine-resource-server_group_server"
description: |-
  Provides a resource to manage server group server
---
# volcengine_server_group_server
Provides a resource to manage server group server
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
  type              = "instance"
}

resource "volcengine_server_group" "foo_ip" {
  load_balancer_id  = volcengine_clb.foo.id
  server_group_name = "acc-test-create-ip"
  description       = "hello demo ip server group"
  type              = "ip"
}

resource "volcengine_security_group" "foo" {
  vpc_id              = volcengine_vpc.foo.id
  security_group_name = "acc-test-security-group"
}

resource "volcengine_ecs_instance" "foo" {
  image_id             = "image-ycjwwciuzy5pkh54xx8f"
  instance_type        = "ecs.c3i.large"
  instance_name        = "acc-test-ecs-name"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
}

resource "volcengine_server_group_server" "foo" {
  server_group_id = volcengine_server_group.foo.id
  instance_id     = volcengine_ecs_instance.foo.id
  type            = "ecs"
  weight          = 100
  port            = 80
  description     = "This is a acc test server"
}

resource "volcengine_server_group_server" "foo_eni" {
  server_group_id = volcengine_server_group.foo.id
  instance_id     = "eni-btgpz5my7ta85h0b2ur*****"
  type            = "eni"
  weight          = 100
  port            = 8080
  description     = "This is a acc test server use eni"
}

resource "volcengine_server_group_server" "foo_ip" {
  server_group_id = volcengine_server_group.foo_ip.id
  instance_id     = "192.168.*.*"
  ip              = "192.168.*.*"
  type            = "ip"
  weight          = 80
  port            = 400
  description     = "This is a acc test server use ip"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of ecs instance or the network card bound to ecs instance. When the `type` is `ip`, this parameter is an IP address.
* `port` - (Required) The port receiving request.
* `server_group_id` - (Required, ForceNew) The ID of the ServerGroup.
* `type` - (Required, ForceNew) The type of instance. Optional choice contains `ecs`, `eni`, `ip`. When the `type` of `server_group_id` is `ip`, only `ip` is supported.
* `description` - (Optional) The description of the instance.
* `ip` - (Optional, ForceNew) The private ip of the instance.
* `weight` - (Optional) The weight of the instance, range in 0~100.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `server_id` - The server id of instance in ServerGroup.


## Import
ServerGroupServer can be imported using the id, e.g.
```
$ terraform import volcengine_server_group_server.default rsp-274xltv2*****8tlv3q3s:rs-3ciynux6i1x4w****rszh49sj
```


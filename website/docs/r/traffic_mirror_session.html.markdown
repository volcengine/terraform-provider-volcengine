---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_traffic_mirror_session"
sidebar_current: "docs-volcengine-resource-traffic_mirror_session"
description: |-
  Provides a resource to manage traffic mirror session
---
# volcengine_traffic_mirror_session
Provides a resource to manage traffic mirror session
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

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-sg"
  vpc_id              = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  os_type          = "Linux"
  visibility       = "public"
  instance_type_id = "ecs.g3il.large"
}

resource "volcengine_ecs_instance" "foo" {
  instance_name        = "acc-test-ecs"
  description          = "acc-test"
  host_name            = "tf-acc-test"
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_type        = "ecs.g3il.large"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
  project_name         = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_network_interface" "foo" {
  network_interface_name = "acc-test-eni"
  description            = "acc-test"
  subnet_id              = volcengine_subnet.foo.id
  security_group_ids     = [volcengine_security_group.foo.id]
  primary_ip_address     = "172.16.0.253"
  port_security_enabled  = false
  private_ip_address     = ["172.16.0.2"]
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_network_interface" "foo-source" {
  network_interface_name = "acc-test-eni"
  description            = "acc-test"
  subnet_id              = volcengine_subnet.foo.id
  security_group_ids     = [volcengine_security_group.foo.id]
  primary_ip_address     = "172.16.0.252"
  port_security_enabled  = false
  private_ip_address     = ["172.16.0.5"]
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_network_interface_attach" "foo" {
  network_interface_id = volcengine_network_interface.foo.id
  instance_id          = volcengine_ecs_instance.foo.id
}

resource "volcengine_network_interface_attach" "foo-source" {
  network_interface_id = volcengine_network_interface.foo-source.id
  instance_id          = volcengine_ecs_instance.foo.id
}

resource "volcengine_traffic_mirror_target" "foo" {
  instance_type              = "NetworkInterface"
  instance_id                = volcengine_network_interface.foo.id
  traffic_mirror_target_name = "acc-test-traffic-mirror-target"
  description                = "acc-test"
  project_name               = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  depends_on = [volcengine_network_interface_attach.foo]
}

resource "volcengine_traffic_mirror_filter" "foo" {
  traffic_mirror_filter_name = "acc-test-traffic-mirror-filter"
  description                = "acc-test"
  project_name               = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_traffic_mirror_session" "foo" {
  traffic_mirror_filter_id    = volcengine_traffic_mirror_filter.foo.id
  traffic_mirror_target_id    = volcengine_traffic_mirror_target.foo.id
  traffic_mirror_session_name = "acc-test-traffic-mirror-session"
  description                 = "acc-test"
  network_interface_id        = volcengine_network_interface.foo-source.id
  virtual_network_id          = 10
  packet_length               = 1000
  priority                    = 1
  project_name                = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  depends_on = [volcengine_network_interface_attach.foo-source]
}
```
## Argument Reference
The following arguments are supported:
* `network_interface_id` - (Required, ForceNew) The ID of network interface.
* `priority` - (Required) The priority of traffic mirror session. Valid values: 1~32766.
* `traffic_mirror_filter_id` - (Required) The ID of traffic mirror filter.
* `traffic_mirror_target_id` - (Required) The ID of traffic mirror target.
* `description` - (Optional) The description of the traffic mirror session.
* `packet_length` - (Optional) The packet length of traffic mirror session.
* `project_name` - (Optional) The project name of traffic mirror session.
* `tags` - (Optional) Tags.
* `traffic_mirror_session_name` - (Optional) The name of the traffic mirror session.
* `virtual_network_id` - (Optional) The ID of virtual network.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `business_status` - The business status of traffic mirror session.
* `created_at` - The create time of traffic mirror session.
* `lock_reason` - The lock reason of traffic mirror session.
* `status` - The status of traffic mirror session.
* `updated_at` - The update time of traffic mirror session.


## Import
TrafficMirrorSession can be imported using the id, e.g.
```
$ terraform import volcengine_traffic_mirror_session.default resource_id
```


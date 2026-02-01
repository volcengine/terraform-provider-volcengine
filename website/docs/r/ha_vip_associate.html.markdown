---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_ha_vip_associate"
sidebar_current: "docs-volcengine-resource-ha_vip_associate"
description: |-
  Provides a resource to manage ha vip associate
---
# volcengine_ha_vip_associate
Provides a resource to manage ha vip associate
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

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip"
  description = "acc-test"
  subnet_id   = volcengine_subnet.foo.id
  ip_address  = "172.16.0.5"
}

resource "volcengine_ha_vip_associate" "foo" {
  ha_vip_id     = volcengine_ha_vip.foo.id
  instance_type = "NetworkInterface"
  instance_id   = volcengine_network_interface.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `ha_vip_id` - (Required, ForceNew) The id of the Ha Vip.
* `instance_id` - (Required, ForceNew) The id of the associated instance.
* `instance_type` - (Optional, ForceNew) The type of the associated instance. Valid values: `EcsInstance`, `NetworkInterface`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
HaVipAssociate can be imported using the ha_vip_id:instance_id, e.g.
```
$ terraform import volcengine_ha_vip_associate.default havip-2byzv8icq1b7k2dx0eegb****:eni-2d5wv84h7onpc58ozfeeu****
```


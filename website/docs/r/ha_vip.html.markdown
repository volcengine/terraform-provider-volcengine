---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_ha_vip"
sidebar_current: "docs-volcengine-resource-ha_vip"
description: |-
  Provides a resource to manage ha vip
---
# volcengine_ha_vip
Provides a resource to manage ha vip
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

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip"
  description = "acc-test"
  subnet_id   = volcengine_subnet.foo.id
  #  ip_address = "172.16.0.5"
}

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByTraffic"
}

resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id   = volcengine_ha_vip.foo.id
  instance_type = "HaVip"
}
```
## Argument Reference
The following arguments are supported:
* `subnet_id` - (Required, ForceNew) The subnet id of the Ha Vip.
* `description` - (Optional) The description of the Ha Vip.
* `ha_vip_name` - (Optional) The name of the Ha Vip.
* `ip_address` - (Optional, ForceNew) The ip address of the Ha Vip.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `associated_eip_address` - The associated eip address of the Ha Vip.
* `associated_eip_id` - The associated eip id of the Ha Vip.
* `associated_instance_ids` - The associated instance ids of the Ha Vip.
* `associated_instance_type` - The associated instance type of the Ha Vip.
* `created_at` - The create time of the Ha Vip.
* `master_instance_id` - The master instance id of the Ha Vip.
* `project_name` - The project name of the Ha Vip.
* `status` - The status of the Ha Vip.
* `updated_at` - The update time of the Ha Vip.
* `vpc_id` - The vpc id of the Ha Vip.


## Import
HaVip can be imported using the id, e.g.
```
$ terraform import volcengine_ha_vip.default havip-2byzv8icq1b7k2dx0eegb****
```


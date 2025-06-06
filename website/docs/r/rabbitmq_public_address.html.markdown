---
subcategory: "RABBITMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rabbitmq_public_address"
sidebar_current: "docs-volcengine-resource-rabbitmq_public_address"
description: |-
  Provides a resource to manage rabbitmq public address
---
# volcengine_rabbitmq_public_address
Provides a resource to manage rabbitmq public address
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

resource "volcengine_rabbitmq_instance" "foo" {
  zone_ids             = [data.volcengine_zones.foo.zones[0].id, data.volcengine_zones.foo.zones[1].id, data.volcengine_zones.foo.zones[2].id]
  subnet_id            = volcengine_subnet.foo.id
  version              = "3.8.18"
  user_name            = "acc-test-user"
  user_password        = "93f0cb0614Aab12"
  compute_spec         = "rabbitmq.n3.x2.small"
  storage_space        = 300
  instance_name        = "acc-test-rabbitmq"
  instance_description = "acc-test"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "BGP"
  name         = "acc-test-eip"
  description  = "acc-test"
  project_name = "default"
}

resource "volcengine_rabbitmq_public_address" "foo" {
  instance_id = volcengine_rabbitmq_instance.foo.id
  eip_id      = volcengine_eip_address.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Required, ForceNew) The id of the eip.
* `instance_id` - (Required, ForceNew) The id of rabbitmq instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RabbitmqPublicAddress can be imported using the instance_id:eip_id, e.g.
```
$ terraform import volcengine_rabbitmq_public_address.default resource_id
```


---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_public_address"
sidebar_current: "docs-volcengine-resource-rocketmq_public_address"
description: |-
  Provides a resource to manage rocketmq public address
---
# volcengine_rocketmq_public_address
Provides a resource to manage rocketmq public address
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

resource "volcengine_rocketmq_instance" "foo" {
  zone_ids             = [data.volcengine_zones.foo.zones[0].id]
  subnet_id            = volcengine_subnet.foo.id
  version              = "4.8"
  compute_spec         = "rocketmq.n1.x2.micro"
  storage_space        = 300
  auto_scale_queue     = true
  file_reserved_time   = 10
  instance_name        = "acc-test-rocketmq"
  instance_description = "acc-test"
  project_name         = "default"
  charge_info {
    charge_type = "PostPaid"
  }
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

resource "volcengine_rocketmq_public_address" "foo" {
  instance_id = volcengine_rocketmq_instance.foo.id
  eip_id      = volcengine_eip_address.foo.id
  ssl_mode    = "permissive"
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Required, ForceNew) The id of the eip.
* `instance_id` - (Required, ForceNew) The id of rocketmq instance.
* `ssl_mode` - (Optional, ForceNew) The ssl mode of the rocketmq instance. Valid values: `enforcing`, `permissive`. Default is `permissive`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RocketmqPublicAddress can be imported using the instance_id:eip_id, e.g.
```
$ terraform import volcengine_rocketmq_public_address.default resource_id
```


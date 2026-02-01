---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_allow_list_associate"
sidebar_current: "docs-volcengine-resource-rocketmq_allow_list_associate"
description: |-
  Provides a resource to manage rocketmq allow list associate
---
# volcengine_rocketmq_allow_list_associate
Provides a resource to manage rocketmq allow list associate
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

resource "volcengine_rocketmq_allow_list" "foo" {
  allow_list_name = "acc-test-allow-list"
  allow_list_desc = "acc-test"
  allow_list      = ["192.168.0.0/24", "192.168.2.0/24"]
}

resource "volcengine_rocketmq_allow_list_associate" "foo" {
  instance_id   = volcengine_rocketmq_instance.foo.id
  allow_list_id = volcengine_rocketmq_allow_list.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_id` - (Required, ForceNew) The id of the rocketmq allow list.
* `instance_id` - (Required, ForceNew) The id of the rocketmq instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RocketmqAllowListAssociate can be imported using the instance_id:allow_list_id, e.g.
```
$ terraform import volcengine_rocketmq_allow_list_associate.default resource_id
```


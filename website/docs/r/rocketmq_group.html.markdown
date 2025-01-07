---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_group"
sidebar_current: "docs-volcengine-resource-rocketmq_group"
description: |-
  Provides a resource to manage rocketmq group
---
# volcengine_rocketmq_group
Provides a resource to manage rocketmq group
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

resource "volcengine_rocketmq_group" "foo" {
  instance_id = volcengine_rocketmq_instance.foo.id
  group_id    = "acc-test-rocketmq-group"
  description = "acc-test"
}
```
## Argument Reference
The following arguments are supported:
* `group_id` - (Required, ForceNew) The id of rocketmq group.
* `instance_id` - (Required, ForceNew) The id of rocketmq instance.
* `description` - (Optional, ForceNew) The description of rocketmq group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the rocketmq group.
* `group_type` - The type of the rocketmq group.
* `is_sub_same` - Whether the subscription relationship of consumer instance groups within the group is consistent.
* `message_delay_time` - The message delay time of the rocketmq group. The unit is milliseconds.
* `message_model` - The message model of the rocketmq group.
* `status` - The status of the rocketmq group.
* `total_consume_rate` - The total consume rate of the rocketmq group. The unit is per second.
* `total_diff` - The total amount of unconsumed messages.


## Import
RocketmqGroup can be imported using the instance_id:group_id, e.g.
```
$ terraform import volcengine_rocketmq_group.default resource_id
```


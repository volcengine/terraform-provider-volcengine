---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_topic"
sidebar_current: "docs-volcengine-resource-rocketmq_topic"
description: |-
  Provides a resource to manage rocketmq topic
---
# volcengine_rocketmq_topic
Provides a resource to manage rocketmq topic
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

resource "volcengine_rocketmq_access_key" "foo" {
  instance_id   = volcengine_rocketmq_instance.foo.id
  description   = "acc-test-key"
  all_authority = "SUB"
}

resource "volcengine_rocketmq_topic" "foo" {
  instance_id  = volcengine_rocketmq_instance.foo.id
  topic_name   = "acc-test-rocketmq-topic"
  description  = "acc-test"
  queue_number = 2
  message_type = 1
  access_policies {
    access_key = volcengine_rocketmq_access_key.foo.access_key
    authority  = "PUB"
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of rocketmq instance.
* `message_type` - (Required, ForceNew) The type of the message. Valid values: `0`: Regular message, `1`: Transaction message, `2`: Partition order message, `3`: Global sequential message, `4`: Delay message.
* `queue_number` - (Required, ForceNew) The maximum number of queues for the current topic, which cannot exceed the remaining available queues for the current rocketmq instance.
* `topic_name` - (Required, ForceNew) The name of the rocketmq topic.
* `access_policies` - (Optional) The access policies of the rocketmq topic. This field can only be added or modified. Deleting this field is invalid.
* `description` - (Optional, ForceNew) The description of the rocketmq topic.

The `access_policies` object supports the following:

* `access_key` - (Required) The access key of the rocketmq key.
* `authority` - (Required) The authority of the rocketmq key for the current topic. Valid values: `ALL`, `PUB`, `SUB`, `DENY`. Default is `DENY`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `groups` - The groups information of the rocketmq topic.
    * `group_id` - The id of the rocketmq group.
    * `message_model` - The message model of the rocketmq group.
    * `sub_string` - The sub string of the rocketmq group.
* `queues` - The queues information of the rocketmq topic.
    * `end_offset` - The end offset of the rocketmq queue.
    * `last_update_timestamp` - The last update timestamp of the rocketmq queue.
    * `message_count` - The message count of the rocketmq queue.
    * `queue_id` - The id of the rocketmq queue.
    * `start_offset` - The start offset of the rocketmq queue.


## Import
RocketmqTopic can be imported using the instance_id:topic_name, e.g.
```
$ terraform import volcengine_rocketmq_topic.default resource_id
```


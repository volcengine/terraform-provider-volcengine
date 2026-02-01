---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_instance"
sidebar_current: "docs-volcengine-resource-rocketmq_instance"
description: |-
  Provides a resource to manage rocketmq instance
---
# volcengine_rocketmq_instance
Provides a resource to manage rocketmq instance
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
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required) The charge information of the rocketmq instance.
* `compute_spec` - (Required) The compute spec of the rocketmq instance.
* `file_reserved_time` - (Required, ForceNew) The reserved time of messages on the RocketMQ server of the message queue. Messages that exceed the reserved time will be cleared after expiration. The unit is in hours. Valid value range is 1~72.
* `storage_space` - (Required) The storage space of the rocketmq instance.
* `subnet_id` - (Required, ForceNew) The subnet id of the rocketmq instance.
* `version` - (Required, ForceNew) The version of the rocketmq instance. Valid values: `4.8`.
* `zone_ids` - (Required, ForceNew) The zone id of the rocketmq instance. Support specifying multiple availability zones.
* `auto_scale_queue` - (Optional) Whether to create queue automatically when the spec of the instance is changed. This field is effective only when modifying `compute_field` and `storage_space`.
* `instance_description` - (Optional) The instance description of the rocketmq instance.
* `instance_name` - (Optional) The instance name of the rocketmq instance.
* `project_name` - (Optional) The project name of the rocketmq instance.
* `tags` - (Optional) Tags.

The `charge_info` object supports the following:

* `charge_type` - (Required) The charge type of the rocketmq instance. Valid values: `PostPaid`, `PrePaid`.
* `auto_renew` - (Optional) Whether to automatically renew in prepaid scenarios. Default is false.
* `period_unit` - (Optional) The purchase cycle in the prepaid scenario. Valid values: `Monthly`, `Yearly`. Default is `Monthly`.
* `period` - (Optional) Purchase duration in prepaid scenarios. When PeriodUnit is specified as `Monthly`, the value range is 1-9. When PeriodUnit is specified as `Yearly`, the value range is 1-3. Default is 1.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account id of the rocketmq instance.
* `apply_private_dns_to_public` - Whether the private dns to public function is enabled for the rocketmq instance.
* `available_queue_number` - The available queue number of the rocketmq instance.
* `connection_info` - The connection information of the rocketmq.
    * `endpoint_address_ip` - The endpoint address ip of the rocketmq.
    * `endpoint_type` - The endpoint type of the rocketmq.
    * `internal_endpoint` - The internal endpoint of the rocketmq.
    * `network_type` - The network type of the rocketmq.
    * `public_endpoint` - The public endpoint of the rocketmq.
* `create_time` - The create time of the rocketmq instance.
* `eip_id` - The eip id of the rocketmq instance.
* `enable_ssl` - Whether the ssl authentication is enabled for the rocketmq instance.
* `instance_status` - The status of the rocketmq instance.
* `region_id` - The region id of the rocketmq instance.
* `ssl_mode` - The ssl mode of the rocketmq instance.
* `used_group_number` - The used group number of the rocketmq instance.
* `used_queue_number` - The used queue number of the rocketmq instance.
* `used_storage_space` - The used storage space of the rocketmq instance.
* `used_topic_number` - The used topic number of the rocketmq instance.
* `vpc_id` - The vpc id of the rocketmq instance.


## Import
RocketmqInstance can be imported using the id, e.g.
```
$ terraform import volcengine_rocketmq_instance.default resource_id
```


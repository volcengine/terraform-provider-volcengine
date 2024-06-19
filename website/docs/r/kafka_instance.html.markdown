---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_instance"
sidebar_current: "docs-volcengine-resource-kafka_instance"
description: |-
  Provides a resource to manage kafka instance
---
# volcengine_kafka_instance
Provides a resource to manage kafka instance
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
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

resource "volcengine_kafka_instance" "foo" {
  instance_name        = "acc-test-kafka"
  instance_description = "tf-test"
  version              = "2.2.2"
  compute_spec         = "kafka.20xrate.hw"
  subnet_id            = volcengine_subnet.foo.id
  user_name            = "tf-user"
  user_password        = "tf-pass!@q1"
  charge_type          = "PostPaid"
  storage_space        = 300
  partition_number     = 350
  project_name         = "default"
  tags {
    key   = "k1"
    value = "v1"
  }

  parameters {
    parameter_name  = "MessageMaxByte"
    parameter_value = "12"
  }
  parameters {
    parameter_name  = "LogRetentionHours"
    parameter_value = "70"
  }
  parameters {
    parameter_name  = "MessageTimestampType"
    parameter_value = "CreateTime"
  }
  parameters {
    parameter_name  = "OffsetRetentionMinutes"
    parameter_value = "10080"
  }
  parameters {
    parameter_name  = "AutoDeleteGroup"
    parameter_value = "false"
  }
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Required) The charge type of instance, the value can be `PrePaid` or `PostPaid`.
* `compute_spec` - (Required) The compute spec of instance.
* `subnet_id` - (Required, ForceNew) The subnet id of instance.
* `user_name` - (Required, ForceNew) The user name of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `user_password` - (Required, ForceNew) The user password of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `version` - (Required, ForceNew) The version of instance, the value can be `2.2.2` or `2.8.2`.
* `auto_renew` - (Optional) The auto renew flag of instance. Only effective when instance_charge_type is PrePaid. Default is false.
* `instance_description` - (Optional) The description of instance.
* `instance_name` - (Optional) The name of instance.
* `need_rebalance` - (Optional) Whether enable rebalance. Only effected in modify when compute_spec field is changed.
* `parameters` - (Optional) Parameter of the instance.
* `partition_number` - (Optional) The partition number of instance.
* `period` - (Optional) The period of instance. Only effective when instance_charge_type is PrePaid. Unit is Month.
* `project_name` - (Optional) The project name of instance.
* `rebalance_time` - (Optional) The rebalance time.
* `storage_space` - (Optional) The storage space of instance.
* `storage_type` - (Optional, ForceNew) The storage type of instance. The value can be ESSD_FlexPL or ESSD_PL0.
* `tags` - (Optional) The tags of instance.

The `parameters` object supports the following:

* `parameter_name` - (Required) Parameter name.
* `parameter_value` - (Required) Parameter value.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
KafkaInstance can be imported using the id, e.g.
```
$ terraform import volcengine_kafka_instance.default kafka-insbjwbbwb
```


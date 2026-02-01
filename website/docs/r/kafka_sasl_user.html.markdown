---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_sasl_user"
sidebar_current: "docs-volcengine-resource-kafka_sasl_user"
description: |-
  Provides a resource to manage kafka sasl user
---
# volcengine_kafka_sasl_user
Provides a resource to manage kafka sasl user
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
}

resource "volcengine_kafka_sasl_user" "foo" {
  user_name     = "acc-test-user"
  instance_id   = volcengine_kafka_instance.foo.id
  user_password = "suqsnis123!"
  description   = "tf-test"
  all_authority = true
  password_type = "Scram"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of instance.
* `user_name` - (Required, ForceNew) The name of user.
* `user_password` - (Required, ForceNew) The password of user.
* `all_authority` - (Optional) Whether this user has read and write permissions for all topics. Default is true.
* `description` - (Optional, ForceNew) The description of user.
* `password_type` - (Optional, ForceNew) The type of password. Valid values are `Scram` and `Plain`. Default is `Plain`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
KafkaSaslUser can be imported using the kafka_id:username, e.g.
```
$ terraform import volcengine_kafka_sasl_user.default kafka-cnngbnntswg1****:tfuser
```


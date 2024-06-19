---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_topic"
sidebar_current: "docs-volcengine-resource-kafka_topic"
description: |-
  Provides a resource to manage kafka topic
---
# volcengine_kafka_topic
Provides a resource to manage kafka topic
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

resource "volcengine_kafka_topic" "foo" {
  topic_name       = "acc-test-topic"
  instance_id      = volcengine_kafka_instance.foo.id
  description      = "tf-test"
  partition_number = 15
  replica_number   = 3

  parameters {
    min_insync_replica_number = 2
    message_max_byte          = 10
    log_retention_hours       = 96
  }

  all_authority = false
  access_policies {
    user_name     = volcengine_kafka_sasl_user.foo.user_name
    access_policy = "Pub"
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The instance id of the kafka topic.
* `partition_number` - (Required) The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.
* `topic_name` - (Required, ForceNew) The name of the kafka topic.
* `access_policies` - (Optional) The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.
* `all_authority` - (Optional) Whether the kafka topic is configured to be accessible by all users. Default: true.
* `description` - (Optional) The description of the kafka topic.
* `parameters` - (Optional) The parameters of the kafka topic.
* `replica_number` - (Optional, ForceNew) The number of replica in kafka topic. The value can be 2 or 3. Default is 3.

The `access_policies` object supports the following:

* `access_policy` - (Required) The access policy of SASL user. Valid values: `PubSub`, `Pub`, `Sub`.
* `user_name` - (Required) The name of SASL user.

The `parameters` object supports the following:

* `log_retention_hours` - (Optional) The retention hours of log. Unit: hour. Valid values: 0-2160. Default is 72.
* `message_max_byte` - (Optional) The max byte of message. Unit: MB. Valid values: 1-12. Default is 10.
* `min_insync_replica_number` - (Optional) The min number of sync replica. The default value is the replica number minus 1.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
KafkaTopic can be imported using the instance_id:topic_name, e.g.
```
$ terraform import volcengine_kafka_topic.default kafka-cnoeeapetf4s****:topic
```


---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_topics"
sidebar_current: "docs-volcengine-datasource-kafka_topics"
description: |-
  Use this data source to query detailed information of kafka topics
---
# volcengine_kafka_topics
Use this data source to query detailed information of kafka topics
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

data "volcengine_kafka_topics" "default" {
  instance_id = volcengine_kafka_topic.foo.instance_id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of kafka instance.
* `name_regex` - (Optional) A Name Regex of kafka topic.
* `output_file` - (Optional) File name where to save data source results.
* `partition_number` - (Optional) The number of partition in kafka topic.
* `replica_number` - (Optional) The number of replica in kafka topic.
* `topic_name` - (Optional) The name of kafka topic. This field supports fuzzy query.
* `user_name` - (Optional) When a user name is specified, only the access policy of the specified user for this Topic will be returned.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `topics` - The collection of query.
    * `access_policies` - The access policies info of the kafka topic.
        * `access_policy` - The access policy of SASL user.
        * `user_name` - The name of SASL user.
    * `all_authority` - Whether the kafka topic is configured to be accessible by all users.
    * `create_time` - The create time of the kafka topic.
    * `description` - The description of the kafka topic.
    * `parameters` - The parameters of the kafka topic.
        * `log_retention_hours` - The retention hours of log.
        * `message_max_byte` - The max byte of message.
        * `min_insync_replica_number` - The min number of sync replica.
    * `partition_number` - The number of partition in the kafka topic.
    * `replica_number` - The number of replica in the kafka topic.
    * `status` - The status of the kafka topic.
    * `topic_name` - The name of the kafka topic.
* `total_count` - The total count of query.



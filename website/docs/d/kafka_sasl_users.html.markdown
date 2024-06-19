---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_sasl_users"
sidebar_current: "docs-volcengine-datasource-kafka_sasl_users"
description: |-
  Use this data source to query detailed information of kafka sasl users
---
# volcengine_kafka_sasl_users
Use this data source to query detailed information of kafka sasl users
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

data "volcengine_kafka_sasl_users" "default" {
  instance_id = volcengine_kafka_instance.foo.id
  user_name   = volcengine_kafka_sasl_user.foo.user_name
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of instance.
* `output_file` - (Optional) File name where to save data source results.
* `user_name` - (Optional) The user name, support fuzzy matching.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `users` - The collection of query.
    * `all_authority` - Whether this user has read and write permissions for all topics.
    * `create_time` - The create time.
    * `description` - The description of user.
    * `password_type` - The type of password.
    * `user_name` - The name of user.



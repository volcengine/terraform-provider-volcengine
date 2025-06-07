# query available zones in current region
data "volcengine_zones" "foo" {
}

# create vpc
resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
  dns_servers  = ["8.8.8.8", "114.114.114.114"]
  project_name = "default"
}

# create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

# create kafka instance
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

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "BGP"
  name         = "acc-test-eip"
  description  = "tf-test"
  project_name = "default"
}

resource "volcengine_kafka_public_address" "foo" {
  instance_id = volcengine_kafka_instance.foo.id
  eip_id      = volcengine_eip_address.foo.id
}

resource "volcengine_kafka_group" "foo" {
  instance_id = volcengine_kafka_instance.foo.id
  group_id    = "acc-test-group"
  description = "tf-test"
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
}

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

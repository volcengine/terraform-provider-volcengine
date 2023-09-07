resource "volcengine_redis_allow_list" "foo" {
  allow_list = ["192.168.0.0/24"]
  allow_list_name = "acc-test-allowlist"
}

data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_redis_instance" "foo"{
  zone_ids = [data.volcengine_zones.foo.zones[0].id]
  instance_name = "acc-test-tf-redis"
  sharded_cluster = 1
  password = "1qaz!QAZ12"
  node_number = 2
  shard_capacity = 1024
  shard_number = 2
  engine_version = "5.0"
  subnet_id = volcengine_subnet.foo.id
  deletion_protection = "disabled"
  vpc_auth_mode = "close"
  charge_type = "PostPaid"
  port = 6381
  project_name = "default"
}

resource "volcengine_redis_allow_list_associate" "foo" {
  allow_list_id = volcengine_redis_allow_list.foo.id
  instance_id = volcengine_redis_instance.foo.id
}
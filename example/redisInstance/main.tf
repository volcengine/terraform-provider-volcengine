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

# create redis instance
resource "volcengine_redis_instance" "foo" {
  instance_name       = "tf-test2"
  sharded_cluster     = 1
  password            = "1qaz!QAZ12"
  node_number         = 4
  shard_capacity      = 1024
  shard_number        = 2
  engine_version      = "5.0"
  subnet_id           = volcengine_subnet.foo.id
  deletion_protection = "disabled"
  vpc_auth_mode       = "close"
  charge_type         = "PostPaid"
  port                = 6381
  project_name        = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  tags {
    key   = "k3"
    value = "v3"
  }

  param_values {
    name  = "active-defrag-cycle-min"
    value = "5"
  }
  param_values {
    name  = "active-defrag-cycle-max"
    value = "28"
  }

  backup_period = [1, 2, 3]
  backup_hour   = 6
  backup_active = true

  create_backup     = false
  apply_immediately = true

  multi_az = "enabled"
  configure_nodes {
    az = "cn-beijing-a"
  }
  configure_nodes {
    az = "cn-beijing-b"
  }
  configure_nodes {
    az = "cn-beijing-c"
  }
  configure_nodes {
    az = "cn-beijing-b"
  }
}
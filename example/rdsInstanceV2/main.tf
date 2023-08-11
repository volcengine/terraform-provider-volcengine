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

resource "volcengine_rds_instance_v2" "foo" {
  db_engine_version = "MySQL_5_7"
  node_info {
    node_type = "Primary"
    node_spec = "rds.mysql.2c4g"
    zone_id   = data.volcengine_zones.foo.zones[0].id
  }
  node_info {
    node_type = "Secondary"
    node_spec = "rds.mysql.2c4g"
    zone_id   = data.volcengine_zones.foo.zones[0].id
  }
  storage_type           = "LocalSSD"
  storage_space          = 100
  vpc_id                 = volcengine_vpc.foo.id
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "tf-test-v2"
  lower_case_table_names = "1"
  charge_info {
    charge_type = "PostPaid"
  }
}
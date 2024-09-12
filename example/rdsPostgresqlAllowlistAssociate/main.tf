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


resource "volcengine_rds_postgresql_instance" "foo" {
  db_engine_version = "PostgreSQL_12"
  node_spec         = "rds.postgres.1c2g"
  primary_zone_id   = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id = data.volcengine_zones.foo.zones[0].id
  storage_space     = 40
  subnet_id         = volcengine_subnet.foo.id
  instance_name     = "acc-test-postgresql"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
  tags {
    key   = "tfk1"
    value = "tfv1"
  }
  parameters {
    name  = "auto_explain.log_analyze"
    value = "off"
  }
  parameters {
    name  = "auto_explain.log_format"
    value = "text"
  }
}

resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24"]
}

resource "volcengine_rds_postgresql_allowlist_associate" "foo" {
  instance_id   = volcengine_rds_postgresql_instance.foo.id
  allow_list_id = volcengine_rds_postgresql_allowlist.foo.id
}

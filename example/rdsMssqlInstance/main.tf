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

resource "volcengine_rds_mssql_instance" "foo" {
  db_engine_version = "SQLServer_2019_Std"
  instance_type = "HA"
  node_spec = "rds.mssql.se.ha.d2.2c4g"
  storage_space = 20
  subnet_id = [volcengine_subnet.foo.id]
  super_account_password = "Tftest110"
  instance_name = "acc-test-mssql"
  project_name = "default"
  charge_info {
    charge_type = "PostPaid"
  }
  tags {
    key = "k1"
    value = "v1"
  }

  backup_time = "18:00Z-19:00Z"
  full_backup_period = ["Monday", "Tuesday"]
  backup_retention_period = 14
}

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

# create mysql instance
resource "volcengine_rds_mysql_instance" "foo" {
  db_engine_version      = "MySQL_5_7"
  node_spec              = "rds.mysql.1c2g"
  primary_zone_id        = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
  storage_space          = 80
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "acc-test-mysql-instance"
  lower_case_table_names = "1"
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }

  charge_info {
    charge_type = "PostPaid"
  }

  parameters {
    parameter_name  = "auto_increment_increment"
    parameter_value = "2"
  }
  parameters {
    parameter_name  = "auto_increment_offset"
    parameter_value = "5"
  }
}

# create mysql instance readonly node
resource "volcengine_rds_mysql_instance_readonly_node" "foo" {
  instance_id = volcengine_rds_mysql_instance.foo.id
  node_spec   = "rds.mysql.2c4g"
  zone_id     = data.volcengine_zones.foo.zones[0].id
}

# create mysql allow list
resource "volcengine_rds_mysql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24"]
}

# associate mysql allow list to mysql instance
resource "volcengine_rds_mysql_allowlist_associate" "foo" {
  allow_list_id = volcengine_rds_mysql_allowlist.foo.id
  instance_id   = volcengine_rds_mysql_instance.foo.id
}

# create mysql database
resource "volcengine_rds_mysql_database" "foo" {
  db_name     = "acc-test-database"
  instance_id = volcengine_rds_mysql_instance.foo.id
}

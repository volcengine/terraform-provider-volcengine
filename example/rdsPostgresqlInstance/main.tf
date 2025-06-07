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

# create postgresql instance
resource "volcengine_rds_postgresql_instance" "foo" {
  db_engine_version = "PostgreSQL_12"
  node_spec         = "rds.postgres.1c2g"
  primary_zone_id   = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id = data.volcengine_zones.foo.zones[0].id
  storage_space     = 40
  subnet_id         = volcengine_subnet.foo.id
  instance_name     = "acc-test-postgresql-instance"
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

# create postgresql instance readonly node
resource "volcengine_rds_postgresql_instance_readonly_node" "foo" {
  instance_id = volcengine_rds_postgresql_instance.foo.id
  node_spec   = "rds.postgres.1c2g"
  zone_id     = data.volcengine_zones.foo.zones[0].id
}

# create postgresql allow list
resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24"]
}

# associate postgresql allow list to postgresql instance
resource "volcengine_rds_postgresql_allowlist_associate" "foo" {
  instance_id   = volcengine_rds_postgresql_instance.foo.id
  allow_list_id = volcengine_rds_postgresql_allowlist.foo.id
}

# create postgresql database
resource "volcengine_rds_postgresql_database" "foo" {
  db_name     = "acc-test-database"
  instance_id = volcengine_rds_postgresql_instance.foo.id
  c_type      = "C"
  collate     = "zh_CN.utf8"
}

# create postgresql account
resource "volcengine_rds_postgresql_account" "foo" {
  account_name       = "acc-test-account"
  account_password   = "9wc@********12"
  account_type       = "Normal"
  instance_id        = volcengine_rds_postgresql_instance.foo.id
  account_privileges = "Inherit,Login,CreateRole,CreateDB"
}

# create postgresql schema
resource "volcengine_rds_postgresql_schema" "foo" {
  db_name     = volcengine_rds_postgresql_database.foo.db_name
  instance_id = volcengine_rds_postgresql_instance.foo.id
  owner       = volcengine_rds_postgresql_account.foo.account_name
  schema_name = "acc-test-schema"
}

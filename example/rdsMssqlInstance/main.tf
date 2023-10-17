resource "volcengine_rds_mssql_instance" "foo" {
  db_engine_version = "SQLServer_2019_Std"
  instance_type = "HA"
  node_spec = "rds.mssql.se.ha.d2.2c4g"
  storage_space = 20
  subnet_id = "subnet-2bzxbx25dmygw2dx0eg14e3yn"
  super_account_password = "Tftest110"
  db_time_zone = "China Standard Time"
  instance_name = "tf-test"
  server_collation = "Chinese_PRC_CI_AS"
  project_name = "default"
  charge_info {
    charge_type = "PostPaid"
  }
  tags {
    key = "tf-key1"
    value = "tf-value1"
  }
  backup_time = "18:00Z-19:00Z"
  full_backup_period = ["Monday", "Tuesday"]
  backup_retention_period = 14
}
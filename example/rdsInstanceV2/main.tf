resource "volcengine_rds_instance_v2" "foo" {
  db_engine_version = "MySQL_5_7"
  instance_type = "HA"
  node_info {
    node_type = "Primary"
    node_spec = "rds.mysql.1c2g"
    zone_id = "cn-guilin-a"
  }
  node_info {
    node_type = "Secondary"
    node_spec = "rds.mysql.1c2g"
    zone_id = "cn-guilin-a"
  }
  node_info {
    node_type = "ReadOnly"
    node_spec = "rds.mysql.1c2g"
    zone_id = "cn-guilin-a"
  }
  storage_type = "LocalSSD"
  storage_space = 100
  vpc_id = "vpc-2d6ym9l9mqlfk58ozfd64sej3"
  subnet_id = "subnet-2d6yma8y0394w58ozfemu5vmi"
  instance_name = "tf-test-v2"
  lower_case_table_names = "1"
  charge_info {
    charge_type = "PostPaid"
  }
}
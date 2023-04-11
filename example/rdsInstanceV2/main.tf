resource "volcengine_rds_instance_v2" "foo" {
  db_engine_version = "MySQL_5_7"
  instance_type = "HA"
  node_info {
    node_type = "Primary"
    node_spec = "rds.mysql.2c4g"
    zone_id = "cn-beijing-a"
  }
  node_info {
    node_type = "Secondary"
    node_spec = "rds.mysql.2c4g"
    zone_id = "cn-beijing-a"
  }
  node_info {
    node_type = "ReadOnly"
    node_spec = "rds.mysql.1c2g"
    zone_id = "cn-beijing-a"
  }
  storage_type = "LocalSSD"
  storage_space = 100
  vpc_id = "vpc-13fawddpwi41s3n6nu4g2y8bt"
  subnet_id = "subnet-mj92ij84m5fk5smt1arvwrtw"
  instance_name = "tf-test-v2"
  lower_case_table_names = "1"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "yyy"
}
resource "volcengine_rds_instance" "foo" {
  region = "cn-north-3"
  db_engine = "MySQL"
  db_engine_version = "MySQL_Community_8_0"
  vpc_id = "vpc-2740cxyk9im0w7fap8u013dfe"
  instance_type = "HA"
  charge_type = "PostPaid"
  zone = "cn-north-3-a"
  storage_type = "LocalSSD"
  storage_space_gb = 100
  instance_spec_name = "rds.mysql.1c2g"
}
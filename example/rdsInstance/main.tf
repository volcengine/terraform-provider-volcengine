resource "volcengine_rds_instance" "foo" {
  region = "cn-north-4"
  zone = "cn-langfang-b"
  instance_name = "tf-test"
  db_engine = "MySQL"
  db_engine_version = "MySQL_Community_5_7"
  vpc_id = "vpc-3cj17x7u9bzeo6c6rrtzfpaeb"
  instance_type = "HA"
  charge_type = "PostPaid"
  storage_type = "LocalSSD"
  storage_space_gb = 100
  instance_spec_name = "rds.mysql.1c2g"
  subnet_id = "subnet-1g0d4fkh1nabk8ibuxx1jtvss"
}
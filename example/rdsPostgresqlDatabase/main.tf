resource "volcengine_rds_postgresql_database" "foo" {
  db_name     = "acc-test"
  instance_id = "postgres-95*******233"
  c_type      = "C"
  collate     = "zh_CN.utf8"
}
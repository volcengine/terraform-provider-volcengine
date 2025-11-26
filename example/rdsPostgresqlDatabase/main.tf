resource "volcengine_rds_postgresql_database" "foo" {
  db_name     = "acc-test"
  instance_id = "postgres-95*******233"
  c_type      = "C"
  collate     = "zh_CN.utf8"
  owner       = "super"
}
resource "volcengine_rds_postgresql_database" "clone_example" {
  db_name     = "clone-test"
  source_db_name = "acc-test"
  instance_id = "postgres-95*******233"
  data_option = "Metadata"
}
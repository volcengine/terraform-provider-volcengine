data "volcengine_rds_postgresql_instance_parameters" "example" {
  instance_id    = "postgres-72715e0d9f58"
  parameter_name = "wal_level"
}
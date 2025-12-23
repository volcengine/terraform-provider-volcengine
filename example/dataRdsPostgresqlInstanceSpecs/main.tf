data "volcengine_rds_postgresql_instance_specs" "example" {
  zone_id           = "cn-chongqing-a"
  db_engine_version = "PostgreSQL_12"
  spec_code         = "rds.postgres.32c128g"
  storage_type      = "LocalSSD"
}
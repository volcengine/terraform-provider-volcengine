data "volcengine_rds_postgresql_database_endpoints" "example" {
  instance_id = "postgres-72715e0d9f58"
  name_regex  = "默认.*"
}
data "volcengine_rds_postgresql_instance_parameter_logs" "example" {
  instance_id = "postgres-72715e0d9f58"
  start_time  = "2025-12-01T00:00:00.000Z"
  end_time    = "2025-12-15T23:59:59.999Z"
}
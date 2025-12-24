data "volcengine_rds_postgresql_instance_failover_logs" "example" {
  instance_id      = "postgres-72******9f58"
  query_start_time = "2025-12-10T16:00:00Z"
  query_end_time   = "2025-12-12T17:00:00Z"
  limit            = 1000
}
data "volcengine_rds_postgresql_planned_events" "example" {
  instance_id = "postgres-72715e0d9f58"
  instance_name = "test-01"
  event_type  = ["VersionUpgrade"]
  status      = ["WaitStart", "Running"]
  planned_switch_time_search_range_start = "2025-12-01T02:06:53.000Z"
  planned_switch_time_search_range_end   = "2025-12-15T17:40:53.000Z"
}

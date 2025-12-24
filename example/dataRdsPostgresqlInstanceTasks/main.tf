data "volcengine_rds_postgresql_instance_tasks" "example" {
  # Choose one of TaskId or time window
  # task_id = "202512121649255DCB10D567104F714DDE-660239"

  # Or filter by time window (≤ 7 days)
  instance_id          = "postgres-72715e0d9f58"
  creation_start_time  = "2025-12-10T21:30:00Z"
  creation_end_time    = "2025-12-15T23:40:00Z"

  # Optional filters
  task_action          = "ModifyDBEndpointReadWriteFlag"
  task_status          = ["Running", "Success"]
  project_name         = "default"
}

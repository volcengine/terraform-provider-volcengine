data "volcengine_rds_postgresql_instance_backup_detacheds" "example" {
  project_name     = "default"
  backup_status    = "Success"
  backup_type      = "Full"
  backup_start_time = "2025-12-01T00:00:00.000Z"
  backup_end_time   = "2025-12-15T23:59:59.999Z"
}
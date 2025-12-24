data "volcengine_rds_postgresql_backup_downloads" "example" {
  instance_id = "postgres-72715e0d9f58"
  backup_id   = "20251214-200431-0698LD"
}
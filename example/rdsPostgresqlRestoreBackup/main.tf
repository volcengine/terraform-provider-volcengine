resource "volcengine_rds_postgresql_restore_backup" "example" {
  backup_id               = "20251214-200431-0698LD"
  source_db_instance_id   = "postgres-72715e0d9f58"
  target_db_instance_id   = "postgres-72715e0d9f58"
  target_db_instance_account = "super"

  databases {
    db_name     = "test"
    new_db_name = "test_restored"
  }
}
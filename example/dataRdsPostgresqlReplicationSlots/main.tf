data "volcengine_rds_postgresql_replication_slots" "example" {
  instance_id = "postgres-72715e0d9f58"
  slot_name   = "my_standby_slot1"
  slot_status = "INACTIVE"
  slot_type   = "physical"
}
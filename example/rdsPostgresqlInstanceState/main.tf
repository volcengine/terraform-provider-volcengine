resource "volcengine_rds_postgresql_instance_state" "example" {
  instance_id = "postgres-72715e0d9f58"
  action      = "Restart"
}
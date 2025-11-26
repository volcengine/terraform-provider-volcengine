resource "volcengine_rds_postgresql_instance_ssl" "example" {
  instance_id           = "postgres-72715e0d9f58"
  ssl_enable            = true
  force_encryption      = true
}
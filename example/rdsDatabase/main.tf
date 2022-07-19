resource "volcengine_rds_database" "foo" {
  instance_id = "mysql-0fdd3bab2e7c"
  db_name = "foo"
  character_set_name = "utf8mb4"
}
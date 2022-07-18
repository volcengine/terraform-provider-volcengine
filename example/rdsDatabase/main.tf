resource "volcengine_rds_database" "foo" {
  instance_id = "mysql-42b38c769c4b"
  db_name = "merge_requests"
  character_set_name = "utf8mb4"
}
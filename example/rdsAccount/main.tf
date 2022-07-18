resource "volcengine_rds_account" "foo" {
  instance_id = "mysql-42b38c769c4b"
  account_name = "test"
  account_password = "Aatest123"
  account_type = "Normal"
}
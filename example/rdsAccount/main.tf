resource "volcengine_rds_account" "foo" {
  instance_id = "mysql-0fdd3bab2e7c"
  account_name = "test"
  account_password = "Aatest123"
  account_type = "Normal"
}
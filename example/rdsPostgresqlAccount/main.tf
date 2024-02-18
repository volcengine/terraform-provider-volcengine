resource "volcengine_rds_postgresql_account" "foo" {
  account_name       = "acc-test-account"
  account_password   = "93c@*****!ab12"
  account_type       = "Super"
  instance_id        = "postgres-954*****7233"
}

resource "volcengine_rds_postgresql_account" "foo1" {
  account_name       = "acc-test-account1"
  account_password   = "9wc@****b12"
  account_type       = "Normal"
  instance_id        = "postgres-95*****7233"
  account_privileges = "Inherit,Login,CreateRole,CreateDB"
}
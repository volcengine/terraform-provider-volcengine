resource "volcengine_rds_account" "app_name" {
  instance_id = "mysql-0fdd3bab2e7c"
  account_name = "terraform-test-app"
  account_password = "Aatest123"
  account_type = "Normal"
}

resource "volcengine_rds_account_privilege" "foo" {
  instance_id = "mysql-0fdd3bab2e7c"
  account_name = volcengine_rds_account.app_name.account_name

  db_privileges {
    db_name = "foo"
    account_privilege = "Custom"
    account_privilege_str = "ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES"
  }

  db_privileges {
    db_name = "bar"
    account_privilege = "DDLOnly"
  }

  db_privileges {
    db_name = "demo"
    account_privilege = "ReadWrite"
  }
}
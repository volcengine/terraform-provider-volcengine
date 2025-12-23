resource "volcengine_rds_postgresql_data_backup" "example" {
  instance_id        = "postgres-72715e0d9f58"
  backup_scope       = "Instance"
  backup_method      = "Physical"
  backup_type        = "Full"
  backup_description = "tf demo full backup2"
}

resource "volcengine_rds_postgresql_data_backup" "example1" {
  instance_id        = "postgres-72715e0d9f58"
  backup_scope       = "Instance"
  backup_method      = "Logical"
  backup_description = "tf demo logical backup"
}

resource "volcengine_rds_postgresql_data_backup" "example2" {
  instance_id        = "postgres-72715e0d9f58"
  backup_scope       = "Database"
  backup_method      = "Logical"
  backup_description = "tf demo database full backup"
  backup_meta {
    db_name = "test"
  }
  backup_meta {
    db_name = "test-1"
  }
}
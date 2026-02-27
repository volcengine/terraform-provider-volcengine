resource "volcengine_kms_secret_backup" "example" {
  secret_name = "Test-1"
}

resource "volcengine_kms_secret_restore" "default" {
  secret_data_key = volcengine_kms_secret_backup.example.secret_data_key
  backup_data     = volcengine_kms_secret_backup.example.backup_data
  signature       = volcengine_kms_secret_backup.example.signature
}
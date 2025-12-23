resource "volcengine_tos_bucket_logging" "foo" {
  bucket_name   = "tflyb7"
  logging_enabled {
    target_bucket = "tflyb78"
    target_prefix = "logs1/"
    role = "ServiceRoleforTOSLogging"
  }
}
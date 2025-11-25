resource "volcengine_tos_bucket_object_lock_configuration" "foo" {
  bucket_name         = "tflyb7"
  
  rule {
    default_retention {
      mode = "COMPLIANCE"
      days = 31
    }
  }
}

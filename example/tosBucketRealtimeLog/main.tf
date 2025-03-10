// When deleting this resource, the tls related resources such as project and topic will not be automatically deleted
resource "volcengine_tos_bucket_realtime_log" "foo" {
  bucket_name = "terraform-demo"
  role        = "TOSLogArchiveTLSRole"
  access_log_configuration {
    ttl = 6
  }
}

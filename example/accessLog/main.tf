# Enable CLB Access Log (TOS Bucket)
resource "volcengine_access_log" "tos_example" {
  load_balancer_id = "clb-13g5i2cbg6nsw3n6nu5r*****"
  bucket_name      = "tos-bucket"
}

# Enable CLB Access Log (TLS)
resource "volcengine_access_log" "tls_example" {
  load_balancer_id = "clb-13g5i2cbg6nsw3n6nu5r*****"
  delivery_type    = "tls"
  tls_project_id   = "d8c6e4c2-8d22-****-****-9811f2067580"
  tls_topic_id     = "081aa4ff-991b-****-****-5d573dcf4ba4"
}
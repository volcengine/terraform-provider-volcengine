# Example: Create a TLS trace instance
resource "volcengine_tls_trace_instance" "foo" {
  project_id          = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  trace_instance_name = "tf-trace-instance-nn"
  description         = "This is an instance-modify"
  backend_config  {
        ttl             = 90
        enable_hot_ttl  = true
        hot_ttl         = 60
        cold_ttl        = 30
        archive_ttl     = 0
        auto_split      = true
        max_split_partitions = 10
       }
}

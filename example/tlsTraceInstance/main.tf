# Example: Create a TLS trace instance
resource "volcengine_tls_trace_instance" "foo" {
  project_id          = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  trace_instance_name = "tf-trace-instance"
  description         = "This is an example trace instance"
  backend_config  {
      ttl             = 60
      enable_hot_ttl  = true
      hot_ttl         = 30
      cold_ttl        = 30
      archive_ttl     = 0
      auto_split      = true
      max_split_partitions = 10
     }
}

output "tls_trace_instance_id" {
  value = volcengine_tls_trace_instance.foo.id
}

output "tls_trace_instance_name" {
  value = volcengine_tls_trace_instance.foo.trace_instance_name
}

output "tls_trace_instance_description" {
  value = volcengine_tls_trace_instance.foo.description
}

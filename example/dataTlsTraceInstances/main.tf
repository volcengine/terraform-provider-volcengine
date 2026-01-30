# Example 1: Query by trace instance name
data "volcengine_tls_trace_instances" "by_name" {
  project_id          = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  trace_instance_name = "测试trace"
}

# Example 2: Query by status
data "volcengine_tls_trace_instances" "by_status" {
  project_id = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  status     = "CREATED"
}



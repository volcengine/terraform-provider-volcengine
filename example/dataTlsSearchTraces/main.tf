data "volcengine_tls_search_traces" "default" {
  trace_instance_id = "e7985388-7b6a-4a15-8013-23556598f0d3"
  query {
      limit = 10
  }
}

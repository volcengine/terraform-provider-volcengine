data "volcengine_tls_search_traces" "default" {
  trace_instance_id = "ac368174-2353-4e5d-859d-84c8bd255590"
  query {
      limit = 10
  }
}

# Basic example - query tags for specific resources
data "volcengine_tls_tags" "basic" {
  resource_type = "project"
  resource_ids  = ["6e6ea17f-ee1d-494f-83f7-c3ecc5c351ea"]
  max_results   = 10
}

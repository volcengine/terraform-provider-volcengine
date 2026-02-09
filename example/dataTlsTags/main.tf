
# Basic example - query tags for specific resources
data "volcengine_tls_tags" "basic" {
  resource_type = "project"
  resource_ids  = ["b01a99c0-cf7b-482f-b317-6563865111c6"]
  max_results   = 10
}

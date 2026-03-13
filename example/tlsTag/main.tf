# Example: Add tags to a TLS topic
resource "volcengine_tls_tag" "foo" {
  resource_id   = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  resource_type = "project"
  tags {
      key   = "environment"
      value = "production"
    }
  tags {
    key = "key1"
    value = "value2"
  }
}

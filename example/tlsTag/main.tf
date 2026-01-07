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


output "tls_tag_id" {
  value = volcengine_tls_tag.foo.id
}

output "tls_tag_resource_id" {
  value = volcengine_tls_tag.foo.resource_id
}

output "tls_tag_resource_type" {
  value = volcengine_tls_tag.foo.resource_type
}

output "tls_tag_tags" {
  value = volcengine_tls_tag.foo.tags
}

output "tls_tag_resource_new_id" {
  value = volcengine_tls_tag.foo.id
}

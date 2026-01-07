
# Basic example - query tags for specific resources
data "volcengine_tls_tag_resources" "basic" {
  resource_type = "project"
  resource_ids  = ["6e6ea17f-ee1d-494f-83f7-c3ecc5c351ea"]
  max_results   = 10
}

# Example with tag filters and max_results
data "volcengine_tls_tag_resources" "with_filters" {
  resource_type = "project"
  resource_ids  = ["project-123456", "project-789012"]
  max_results   = 50
  tag_filters {
    key    = "environment"
    values = ["production", "development"]
  }
  tag_filters {
    key    = "department"
    values = ["devops"]
  }
}


# Example with pagination using max_results
data "volcengine_tls_tag_resources" "first_page" {
  resource_type = "topic"
  resource_ids  = ["topic-123456"]
  max_results   = 20
}


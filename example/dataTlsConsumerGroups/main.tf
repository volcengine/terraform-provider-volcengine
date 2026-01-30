
#Basic example - query all consumer groups

data "volcengine_tls_consumer_groups" "all" {
}

output "all_consumer_groups" {
  value = data.volcengine_tls_consumer_groups.all.consumer_groups
}

# Example with consumer group name filter

data "volcengine_tls_consumer_groups" "by_name" {
  consumer_group_name = "test-consumer-group"
}

output "by_name_consumer_groups" {
  value = data.volcengine_tls_consumer_groups.by_name.consumer_groups
}

# Example with project filter

data "volcengine_tls_consumer_groups" "by_project" {
  project_id = "project-123456"
}

output "by_project_consumer_groups" {
  value = data.volcengine_tls_consumer_groups.by_project.consumer_groups
}

# Example with topic filter

data "volcengine_tls_consumer_groups" "by_topic" {
  topic_id = "topic-123456"
}

output "by_topic_consumer_groups" {
  value = data.volcengine_tls_consumer_groups.by_topic.consumer_groups
}

# Example with multiple filters

data "volcengine_tls_consumer_groups" "with_multiple_filters" {
  project_id          = "project-123456"
  consumer_group_name = "test"
}

output "multiple_filters_consumer_groups" {
  value = data.volcengine_tls_consumer_groups.with_multiple_filters.consumer_groups
}
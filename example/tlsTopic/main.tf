resource "volcengine_tls_topic" "foo" {
  project_id = "e020c978-4f05-40e1-9167-0113d3ef****"
  topic_name = "tf-test-topic"
  description = "test"
  ttl = 10
  shard_count = 2
  auto_split = true
  max_split_shard = 10
  enable_tracking = true
  time_key = "request_time"
  time_format = "%Y-%m-%dT%H:%M:%S,%f"
  tags {
    key = "k1"
    value = "v1"
  }
}
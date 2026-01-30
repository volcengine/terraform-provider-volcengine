resource "volcengine_tls_topic" "foo" {
  project_id      = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  topic_name      = "tf-topic-5"
  description     = "test"
  ttl             = 60
  shard_count     = 2
  auto_split      = true
  max_split_shard = 10
  enable_tracking = true
  time_key        = "request_time"
  time_format     = "%Y-%m-%dT%H:%M:%S,%f"
  tags {
    key   = "k1"
    value = "v1"
  }
  log_public_ip  = true
  enable_hot_ttl = true
  hot_ttl        = 30
  cold_ttl       = 30
  archive_ttl    = 0
  encrypt_conf {
    enable       = true
    encrypt_type = "default"
  }
}

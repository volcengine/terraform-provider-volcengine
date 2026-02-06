data "volcengine_tls_log_cursors" "default" {
  topic_id = "e101b8c8-77e7-4ae3-91c1-2532ee480e7d"
  shard_id = 0
  from     = "begin"
}

output "cursor" {
  value = data.volcengine_tls_log_cursors.default.log_cursors[0].cursor
}

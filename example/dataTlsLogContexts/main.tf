# Search Logs (Trigger SearchLogs)
data "volcengine_tls_log_searches" "default" {
  topic_id    = "3c57a110-399a-43b3-bc3c-5d60e065239a"
  query      = "*"
  start_time = 1773017877000
  end_time   = 1773067877000
  limit      = 100
}

# 1. Describe Log Context (Trigger DescribeLogContext)
data "volcengine_tls_log_contexts" "default" {
  topic_id       = data.volcengine_tls_log_searches.default.topic_id
  context_flow   = data.volcengine_tls_log_searches.default.logs[0].logs[0].content["__context_flow__"]
  package_offset = tonumber(data.volcengine_tls_log_searches.default.logs[0].logs[0].content["__package_offset__"])
  source         = data.volcengine_tls_log_searches.default.logs[0].logs[0].source
  prev_logs      = 10
  next_logs      = 10
}
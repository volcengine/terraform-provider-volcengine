# 1. Describe Log Context (Trigger DescribeLogContext)
data "volcengine_tls_log_contexts" "default" {
  topic_id       = data.volcengine_tls_search_logs.default.topic_id
  context_flow   = data.volcengine_tls_search_logs.default.logs[0].logs[0].content["__context_flow__"]
  package_offset = tonumber(data.volcengine_tls_search_logs.default.logs[0].logs[0].content["__package_offset__"])
  source         = data.volcengine_tls_search_logs.default.logs[0].logs[0].source
  prev_logs      = 10
  next_logs      = 10
}
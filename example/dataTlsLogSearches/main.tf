# Search Logs (Trigger SearchLogs)
data "volcengine_tls_log_searches" "default" {
  topic_id    = "3c57a110-399a-43b3-bc3c-5d60e065239a"
  query      = "*"
  start_time = 1773017877000
  end_time   = 1773067877000
  limit      = 100
}
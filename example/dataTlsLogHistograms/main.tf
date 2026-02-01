data "volcengine_tls_log_histograms" "default" {
  topic_id    = "3c57a110-399a-43b3-bc3c-5d60e065239a"
  query      = "*"
  start_time = 1768448896000
  end_time   = 1768450896000
  interval   = 60000
}
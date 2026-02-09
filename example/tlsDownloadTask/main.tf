resource "volcengine_tls_download_task" "foo" {
  topic_id         = "3c57a110-399a-43b3-bc3c-5d60e065239a"
  task_name        = "tf-test-download"
  query            = "*"
  start_time       = 1768448896
  end_time         = 1768450896
  compression      = "gzip"
  data_format      = "json"
  limit            = 10000000
  sort             = "asc"
  allow_incomplete = false
  task_type        = 1
  log_context_infos {
    source = "your ip"
    context_flow = "1768450893021#4258909d8fc97e7d-286d6d5f6966623c-6943"
    package_offset = "4833728523"
  }
}

output "tls_download_task_id" {
  value = volcengine_tls_download_task.foo.task_id
}